package goose

import (
	"net/http"
	"net/url"
)

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
