// Package resolver provides URL resolution functionality for the goose client
package resolver

import (
	"context"
	"net/url"
	"strings"
)

// Ensure HttpsResolver implements the Resolver interface
var _ Resolver = (*HttpsResolver)(nil)

// init registers the HttpsResolver when the package is initialized
func init() {
	RegisterResolver(&HttpsResolver{})
}

// HttpsResolver is a resolver that handles URLs with "https" scheme
type HttpsResolver struct{}

// Resolve resolves an HTTPS target URL by copying all its components
// Parameters:
//   - ctx: Context for the resolution operation
//   - target: Target URL with "https" scheme to resolve
//
// Returns:
//   - *url.URL: Resolved URL with all components copied from target
//   - error: Error if the target scheme is not "https" (case-insensitive)
func (r HttpsResolver) Resolve(ctx context.Context, target *url.URL) (*url.URL, error) {
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
//   - string: "https", indicating this resolver handles URLs with "https" scheme
func (r HttpsResolver) Scheme() string {
	return "https"
}
