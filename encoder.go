package goose

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/genproto/googleapis/api/httpbody"
	rpchttp "google.golang.org/genproto/googleapis/rpc/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ErrorEncoder defines a function type for encoding errors into HTTP responses.
// Implementations should write the error to the provided http.ResponseWriter.
type ErrorEncoder func(ctx context.Context, err error, w http.ResponseWriter)

// ResponseTransformer defines a function type for transforming protobuf responses
// before encoding. Can be used to modify or wrap responses.
type ResponseTransformer func(ctx context.Context, resp proto.Message) proto.Message

// DefaultTransformResponse is the default response transformer that returns
// the response unchanged.
//
// Parameters:
//
//	ctx - context.Context for the request
//	resp - proto.Message to transform
//
// Returns:
//
//	proto.Message - the same response unchanged
func DefaultTransformResponse(ctx context.Context, resp proto.Message) proto.Message {
	return resp
}

// DefaultEncodeError encodes errors into HTTP responses with appropriate
// status codes and content type. Handles several error types:
// - json.Marshaler: encodes error as JSON if implemented
// - Headers() http.Header: adds headers to response if implemented
// - StatusCode() int: uses custom status code if implemented
//
// Parameters:
//
//	ctx - context.Context for the request
//	err - error to encode
//	w - http.ResponseWriter to write the error response
func DefaultEncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	// Default to plain text content type and error message as body
	contentType, body := PlainContentType, []byte(err.Error())

	// If the error implements json.Marshaler, try to marshal it as JSON
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			contentType, body = JsonContentType, jsonBody
		}
	}

	// Set response content type header
	w.Header().Set(ContentTypeKey, contentType)

	// If error provides custom headers, add them to the response
	if headerGetter, ok := err.(interface{ Headers() http.Header }); ok {
		for k, values := range headerGetter.Headers() {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}

	// Default to 500 status code unless error provides specific status code
	code := http.StatusInternalServerError
	if sc, ok := err.(interface{ StatusCode() int }); ok {
		code = sc.StatusCode()
	}

	// Write HTTP status code and response body
	w.WriteHeader(code)
	_, err = w.Write(body)
	if err != nil {
		log.Default().Println("goose: response write error: ", err)
	}
}

// EncodeResponse encodes a protobuf message as JSON into an HTTP response.
// Sets Content-Type to application/json and status code to 200 OK.
//
// Parameters:
//
//	ctx - context.Context for the request
//	w - http.ResponseWriter to write the response
//	resp - proto.Message to encode
//	marshalOptions - protojson.MarshalOptions for JSON encoding
//
// Returns:
//
//	error - if encoding or writing fails
func EncodeResponse(ctx context.Context, w http.ResponseWriter, resp proto.Message, marshalOptions protojson.MarshalOptions) error {
	// Set response headers for JSON content and HTTP 200 status
	w.Header().Set(ContentTypeKey, JsonContentType)
	w.WriteHeader(http.StatusOK)

	// Marshal the protocol buffer message into JSON
	data, err := marshalOptions.Marshal(resp)
	if err != nil {
		return err
	}

	// Write the JSON data to the response body
	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

// EncodeHttpBody encodes an httpbody.HttpBody into an HTTP response.
// Sets Content-Type from the HttpBody and status code to 200 OK.
//
// Parameters:
//
//	ctx - context.Context for the request
//	w - http.ResponseWriter to write the response
//	resp - *httpbody.HttpBody to encode
//
// Returns:
//
//	error - if writing fails
func EncodeHttpBody(ctx context.Context, w http.ResponseWriter, resp *httpbody.HttpBody) error {
	// Set response headers
	w.Header().Set(ContentTypeKey, resp.GetContentType())
	w.WriteHeader(http.StatusOK)

	// Write response data
	if _, err := w.Write(resp.GetData()); err != nil {
		return err
	}
	return nil
}

// EncodeHttpResponse encodes an rpchttp.HttpResponse into an HTTP response.
// Sets headers, status code and body from the HttpResponse.
//
// Parameters:
//
//	ctx - context.Context for the request
//	w - http.ResponseWriter to write the response
//	resp - *rpchttp.HttpResponse to encode
//
// Returns:
//
//	error - if writing fails
func EncodeHttpResponse(ctx context.Context, w http.ResponseWriter, resp *rpchttp.HttpResponse) error {
	// Set all headers from the RPC response
	for _, header := range resp.GetHeaders() {
		w.Header().Add(header.GetKey(), header.GetValue())
	}

	// Write HTTP status code before body
	w.WriteHeader(int(resp.GetStatus()))

	// Write response body and return any write errors
	if _, err := w.Write(resp.GetBody()); err != nil {
		return err
	}
	return nil
}
