// Package resolver provides URL resolution functionality for the goose client
package resolver

import (
	"context"
	"net/url"
)

// DefaultHttpScheme is the default HTTP scheme used by the DefaultResolver
var DefaultHttpScheme = "http"

// init registers the DefaultResolver when the package is initialized
func init() {
	RegisterResolver(&DefaultResolver{})
}

// DefaultResolver is a resolver that handles URLs with empty schemes
type DefaultResolver struct {
	HttpScheme string // Custom HTTP scheme to use instead of default
}

// Resolve resolves a target URL by copying its components and setting the scheme
// Parameters:
//   - ctx: Context for the resolution operation
//   - target: Target URL to resolve
//
// Returns:
//   - *url.URL: Resolved URL with scheme set
//   - error: Error if the target scheme doesn't match this resolver's scheme
func (r DefaultResolver) Resolve(ctx context.Context, target *url.URL) (*url.URL, error) {
	// Check if the target URL scheme matches this resolver's scheme
	if target.Scheme != r.Scheme() {
		return nil, &ResolverError{target: target}
	}

	// Create a new URL with the resolved scheme
	resolved := &url.URL{}

	// Set the scheme based on HttpScheme field or default scheme
	if r.HttpScheme != "" {
		resolved.Scheme = r.HttpScheme
	} else {
		resolved.Scheme = DefaultHttpScheme
	}

	// Copy all other URL components from the target
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

// Scheme returns the scheme that this resolver handles (empty string for default resolver)
// Returns:
//   - string: Empty string, indicating this resolver handles URLs with no scheme
func (r DefaultResolver) Scheme() string {
	return ""
}
