// Package resolver provides URL resolution functionality for the goose client
package resolver

import (
	"context"
	"net/url"
	"sync"
)

// Resolver is an interface for resolving target URLs to actual URLs
// Implementations handle different URL schemes and resolve them appropriately
type Resolver interface {
	// Resolve takes a target URL and returns a resolved URL
	// Parameters:
	//   - ctx: Context for the resolution operation
	//   - target: Target URL to resolve
	// Returns:
	//   - *url.URL: Resolved URL
	//   - error: Error if resolution fails
	Resolve(ctx context.Context, target *url.URL) (*url.URL, error)

	// Scheme returns the URL scheme this resolver handles
	// Returns:
	//   - string: URL scheme (e.g., "http", "https", "")
	Scheme() string
}

// registered is a thread-safe map storing registered resolvers by their scheme
var registered = sync.Map{}

// RegisterResolver registers a resolver for its scheme
// This allows the resolver to be automatically used when resolving URLs with matching schemes
// Parameters:
//   - resolver: Resolver to register
func RegisterResolver(resolver Resolver) {
	registered.Store(resolver.Scheme(), resolver)
}

// Resolve resolves a target URL string using the appropriate resolver
// It first tries to use the provided resolver if its scheme matches the target,
// otherwise it looks up a registered resolver by the target's scheme
// Parameters:
//   - ctx: Context for the resolution operation
//   - resolver: Optional resolver to try first (can be nil)
//   - targetStr: Target URL string to resolve
//
// Returns:
//   - *url.URL: Resolved URL
//   - error: Error if parsing or resolution fails
func Resolve(ctx context.Context, resolver Resolver, targetStr string) (*url.URL, error) {
	// Parse the target string into a URL
	target, err := url.Parse(targetStr)
	if err != nil {
		return nil, err
	}

	// If a resolver is provided and its scheme matches the target's scheme, use it
	if resolver != nil && resolver.Scheme() == target.Scheme {
		return resolver.Resolve(ctx, target)
	}

	// Look up a registered resolver by the target's scheme
	if resolver, ok := registered.Load(target.Scheme); ok {
		return resolver.(Resolver).Resolve(ctx, target)
	}

	// No resolver found for the scheme, return an error
	return nil, &ResolverError{target: target}
}
