package resolver

import (
	"context"
	"net/url"
	"sync"
)

type Resolver interface {
	Resolve(ctx context.Context, target *url.URL) (*url.URL, error)
	Scheme() string
}

var registered = sync.Map{}

func RegisterResolver(resolver Resolver) {
	registered.Store(resolver.Scheme(), resolver)
}

func Resolve(ctx context.Context, resolver Resolver, targetStr string) (*url.URL, error) {
	target, err := url.Parse(targetStr)
	if err != nil {
		return nil, err
	}
	if resolver != nil {
		return resolver.Resolve(ctx, target)
	}
	if resolver, ok := registered.Load(target.Scheme); ok {
		return resolver.(Resolver).Resolve(ctx, target)
	}
	return nil, &ResolverError{target: target}
}
