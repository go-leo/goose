// Package resolver provides URL resolution functionality for the goose client
package resolver

import (
	"context"
	"net/url"
	"strings"
)

// Ensure HttpResolver implements the Resolver interface
var _ Resolver = (*HttpResolver)(nil)

// init registers the HttpResolver when the package is initialized
func init() {
	RegisterResolver(&HttpResolver{})
}

// HttpResolver is a resolver that handles URLs with "http" scheme
type HttpResolver struct{}

// Resolve resolves an HTTP target URL by copying all its components
// Parameters:
//   - ctx: Context for the resolution operation
//   - target: Target URL with "http" scheme to resolve
//
// Returns:
//   - *url.URL: Resolved URL with all components copied from target
//   - error: Error if the target scheme is not "http" (case-insensitive)
func (r HttpResolver) Resolve(ctx context.Context, target *url.URL) (*url.URL, error) {
	// Check if the target URL scheme matches this resolver's scheme (case-insensitive)
	if !strings.EqualFold(target.Scheme, r.Scheme()) {
		return nil, &ResolverError{target: target}
	}

	// Create a new URL and copy all components from the target
	resolved := &url.URL{}
	resolved.Scheme = target.Scheme
	resolved.Opaque = target.Opaque
	resolved.User = target.User
	resolved.Host = target.Host
	resolved.Path = target.Path
	resolved.RawPath = target.RawPath
	resolved.OmitHost = target.OmitHost
	resolved.ForceQuery = target.ForceQuery
	resolved.RawQuery = target.RawQuery
	resolved.Fragment = target.Fragment
	resolved.RawFragment = target.RawFragment

	return resolved, nil
}

// Scheme returns the scheme that this resolver handles
// Returns:
//   - string: "http", indicating this resolver handles URLs with "http" scheme
func (r HttpResolver) Scheme() string {
	return "http"
}
