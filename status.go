package goose

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

)

// ErrorEncoder defines a function type for encoding errors into HTTP responses.
// Implementations should write the error to the provided http.ResponseWriter.
type ErrorEncoder func(ctx context.Context, err error, response http.ResponseWriter)

type ErrorDecoder func(ctx context.Context, response *http.Response, factory ErrorFactory) (error, bool)

type ErrorFactory func() error

type defaultError struct {
	statusCode int
	headers    http.Header
	body       any
}

func NewError(statusCode int, body any, headers ...string) error {
	err := &defaultError{
		statusCode: statusCode,
		body:       body,
		headers:    http.Header{},
	}
	if len(headers)/2 != 0 {
		panic("goose: headers length must be even")
	}
	for i := 0; i < len(headers); i += 2 {
		err.headers.Add(headers[i], headers[i+1])
	}
	return err
}

func (e *defaultError) Error() string {
	return fmt.Sprintf("goose: http error, status code: %d, body: %s", e.statusCode, e.body)
}

func (e *defaultError) StatusCode() int {
	return e.statusCode
}

func (e *defaultError) SetStatusCode(code int) {
	e.statusCode = code
}

func (e *defaultError) Headers() http.Header {
	return e.headers
}

func (e *defaultError) SetHeaders(h http.Header) {
	e.headers = h
}

func (e *defaultError) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.body)
}

func (e *defaultError) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.body)
}

type StatusCodeGetter interface {
	StatusCode() int
}

type HeaderGetter interface {
	Headers() http.Header
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
func DefaultEncodeError(ctx context.Context, respErr error, response http.ResponseWriter) {
	// Default to 500 status code unless error provides specific status code
	code := http.StatusInternalServerError
	if statusCodeGetter, ok := respErr.(StatusCodeGetter); ok {
		code = statusCodeGetter.StatusCode()
	}

	// Default to plain text content type and error message as body
	contentType, body := PlainContentType, []byte(respErr.Error())
	// If the error implements json.Marshaler, try to marshal it as JSON
	if marshaler, ok := respErr.(json.Marshaler); ok {
		if jsonBody, err := marshaler.MarshalJSON(); err != nil {
			log.Println("goose: body marshal error: ", err)
		} else {
			contentType, body = JsonContentType, jsonBody
		}
	}

	header := response.Header()
	// Set response content type header
	header.Set(ContentTypeKey, contentType)
	// If error provides custom headers, add them to the response
	keys := make([]string, 0)
	if headerGetter, ok := respErr.(HeaderGetter); ok {
		for key, values := range headerGetter.Headers() {
			for _, v := range values {
				header.Add(key, v)
				keys = append(keys, key)
			}
		}
	}
	keysJson, _ := json.Marshal(keys)
	header.Add(ErrorKey, string(keysJson))

	// Write HTTP status code and response body
	response.WriteHeader(code)
	_, respErr = response.Write(body)
	if respErr != nil {
		log.Println("goose: DefaultEncodeError, response write error: ", respErr)
	}
}

type StatusCodeSetter interface {
	SetStatusCode(code int)
}

type HeaderSetter interface {
	SetHeaders(h http.Header)
}

func DefaultErrorFactory() error {
	return &defaultError{}
}

func DefaultDecodeError(ctx context.Context, response *http.Response, factory ErrorFactory) (error, bool) {
	keysJson := response.Header.Get(ErrorKey)
	if keysJson == "" {
		return nil, false
	}
	respErr := factory()

	if statusCodeGetter, ok := respErr.(StatusCodeSetter); ok {
		statusCodeGetter.SetStatusCode(response.StatusCode)
	}

	if headerSetter, ok := respErr.(HeaderSetter); ok {
		keys := make([]string, 0)
		err := json.Unmarshal([]byte(keysJson), &keys)
		if err != nil {
			log.Println("goose: header key unmarshal error: ", err)
		} else {
			headers := make(http.Header, len(keys))
			for _, key := range keys {
				for _, value := range response.Header.Values(key) {
					headers.Add(key, value)
				}
			}
			headerSetter.SetHeaders(headers)
		}
	}

	body, _ := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if unmarshaler, ok := respErr.(json.Unmarshaler); ok {
		if err := unmarshaler.UnmarshalJSON(body); err != nil {
			log.Println("goose: body unmarshal error: ", err)
		}
	}
	return respErr, true
}
