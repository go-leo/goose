package goose

import (
	"net/http"
	"net/url"
)

// FormGetter defines a generic function type for form data retrieval
// Parameters:
//   - form: Form data
//   - key: Form field key
//
// Returns:
//   - T: Retrieved value
//   - error: Retrieval error if any
type FormGetter[T any] func(form url.Values, key string) (T, error)

// GetForm decodes form data
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
func GetForm[T any](pre error, form url.Values, key string, f FormGetter[T]) (T, error) {
	return BreakOnError[T](pre)(func() (T, error) { return f(form, key) })
}

// FormFromPath extracts specified key-value pairs from HTTP request path parameters
// and constructs them into url.Values format.
//
// Parameters:
//
//	r: HTTP request object used to retrieve path parameters
//	keys: list of path parameter key names to extract
//
// Returns:
//
//	url.Values: form data containing specified path parameter key-value pairs,
//	            returns nil if keys is nil
func FormFromPath(r *http.Request, keys ...string) url.Values {
	if keys == nil {
		return nil
	}
	form := url.Values{}
	for _, key := range keys {
		form.Add(key, r.PathValue(key))
	}
	return form
}

// FormFromMap converts a map[string]string to url.Values format
// Parameters:
//   - m: String key-value pair mapping to be converted
//
// Returns:
//   - url.Values: Converted form data, returns nil if input is nil
//
// Behavior:
//  1. If m is nil, returns nil immediately
//  2. Creates a new url.Values instance
//  3. Iterates through the input map, adding each key-value pair to the form
func FormFromMap(m map[string]string) url.Values {
	if m == nil {
		return nil
	}
	form := url.Values{}
	for key, value := range m {
		form.Add(key, value)
	}
	return form
}
