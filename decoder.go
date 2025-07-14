package gonic

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// OnValidationErrCallback defines a callback function type for validation errors
// This callback will be invoked when ValidateRequest encounters a validation error
type OnValidationErrCallback func(ctx context.Context, err error)

// CustomDecodeRequest provides custom request decoding functionality
// Parameters:
//   - ctx: Context object
//   - r: HTTP request object
//   - req: Proto.Message to be decoded
//
// Returns:
//   - bool: Indicates whether custom decoding was performed
//   - error: Any error that occurred during decoding
//
// Behavior:
//  1. Checks if req implements UnmarshalRequest method
//  2. If implemented, invokes the method for decoding
//  3. If not implemented, returns false indicating no custom decoding was done
func CustomDecodeRequest(ctx context.Context, r *http.Request, req proto.Message) (bool, error) {
	unmarshaler, ok := req.(interface {
		UnmarshalRequest(context.Context, *http.Request) error
	})
	if ok {
		return true, unmarshaler.UnmarshalRequest(ctx, r)
	}
	return false, nil
}

// ValidateRequest validates the request parameters
// Parameters:
//   - ctx: Context object
//   - req: Proto.Message to validate
//   - fast: Whether to perform fast validation (skip deep validation)
//   - callback: Callback function for validation errors
//
// Returns:
//   - error: Validation error if any
//
// Behavior:
//
//	Based on fast parameter:
//	- fast=true: Attempts to call Validate() or Validate(false)
//	- fast=false: Attempts to call ValidateAll() or Validate(true) or Validate()
//	If validation fails and callback is provided, invokes the callback
func ValidateRequest(ctx context.Context, req proto.Message, fast bool, callback OnValidationErrCallback) (err error) {
	if fast {
		switch v := req.(type) {
		case interface{ Validate() error }:
			err = v.Validate()
		case interface{ Validate(all bool) error }:
			err = v.Validate(false)
		}
	} else {
		switch v := req.(type) {
		case interface{ ValidateAll() error }:
			err = v.ValidateAll()
		case interface{ Validate(all bool) error }:
			err = v.Validate(true)
		case interface{ Validate() error }:
			err = v.Validate()
		}
	}

	if err == nil {
		return nil
	}

	if callback != nil {
		callback(ctx, err)
	}
	return err
}

// DecodeRequest decodes HTTP request body into a proto.Message
// Parameters:
//   - ctx: Context object
//   - r: HTTP request object
//   - req: Target proto.Message
//   - unmarshalOptions: protojson unmarshal options
//
// Returns:
//   - error: Decoding error if any
//
// Behavior:
//  1. Reads the request body
//  2. Unmarshals the data into target proto.Message using protojson
func DecodeRequest(ctx context.Context, r *http.Request, req proto.Message, unmarshalOptions protojson.UnmarshalOptions) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := unmarshalOptions.Unmarshal(data, req); err != nil {
		return err
	}
	return nil
}

// DecodeHttpBody decodes HTTP request body into HttpBody object
// Parameters:
//   - ctx: Context object
//   - r: HTTP request object
//   - body: Target HttpBody object
//
// Returns:
//   - error: Decoding error if any
//
// Behavior:
//  1. Reads the request body data
//  2. Sets HttpBody's Data and ContentType fields
func DecodeHttpBody(ctx context.Context, r *http.Request, body *httpbody.HttpBody) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	body.Data = data
	body.ContentType = r.Header.Get(ContentTypeKey)
	return nil
}

// DecodeHttpRequest decodes HTTP request into HttpRequest object
// Parameters:
//   - ctx: Context object
//   - r: HTTP request object
//   - request: Target HttpRequest object
//
// Returns:
//   - error: Decoding error if any
//
// Behavior:
//  1. Reads the request body data
//  2. Sets method, URI, headers and body fields
func DecodeHttpRequest(ctx context.Context, r *http.Request, request *rpchttp.HttpRequest) error {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	request.Method = r.Method
	request.Uri = r.URL.String()
	for key, values := range r.Header {
		for _, value := range values {
			request.Headers = append(request.Headers, &rpchttp.HttpHeader{Key: key, Value: value})
		}
	}
	request.Body = data
	return nil
}

// FormGetter defines a generic function type for form data retrieval
// Parameters:
//   - form: Form data
//   - key: Form field key
//
// Returns:
//   - T: Retrieved value
//   - error: Retrieval error if any
type FormGetter[T any] func(form url.Values, key string) (T, error)

// DecodeForm decodes form data
// Parameters:
//   - pre: Pre-existing error (if any, will be returned immediately)
//   - form: Form data
//   - key: Form field key
//   - f: Form data getter function
//
// Returns:
//   - T: Decoded value
//   - error: Decoding error if any
//
// Behavior:
//  1. If pre is not nil, returns pre error immediately
//  2. Otherwise invokes f to get form value
func DecodeForm[T any](pre error, form url.Values, key string, f FormGetter[T]) (T, error) {
	return breakOnError[T](pre)(func() (T, error) { return f(form, key) })
}

// breakOnError provides error interception functionality
// Parameters:
//   - pre: Pre-existing error
//
// Returns:
//
//	A function that will:
//	1. Return pre error immediately if pre is not nil
//	2. Otherwise execute the provided function f
func breakOnError[T any](pre error) func(f func() (T, error)) (T, error) {
	return func(f func() (T, error)) (T, error) {
		if pre != nil {
			var v T
			return v, pre
		}
		return f()
	}
}
