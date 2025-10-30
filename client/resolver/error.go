// Package resolver provides URL resolution functionality for the goose client
package resolver

import (
	"fmt"
	"net/url"
)

// ResolverError represents an error that occurs when a URL scheme is not supported by any registered resolver
type ResolverError struct {
	target *url.URL // The target URL that could not be resolved
}

// Error returns a formatted error message indicating which scheme is not supported
// Returns:
//   - string: Error message in the format "resolver: scheme <scheme> is not supported"
func (e *ResolverError) Error() string {
	return fmt.Sprintf("resolver: scheme %s is not supported", e.target.Scheme)
}

// Target returns the target URL that caused the resolver error
// Returns:
//   - *url.URL: The URL that could not be resolved due to unsupported scheme
func (e *ResolverError) Target() *url.URL {
	return e.target
}
