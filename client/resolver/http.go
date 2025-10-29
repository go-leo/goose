package resolver

import (
	"context"
	"net/url"
	"strings"
)

var _ Resolver = (*HttpResolver)(nil)

func init() {
	RegisterResolver(&HttpResolver{})
}

type HttpResolver struct{}

func (r HttpResolver) Resolve(ctx context.Context, target *url.URL) (*url.URL, error) {
	if !strings.EqualFold(target.Scheme, r.Scheme()) {
		return nil, &ResolverError{target: target}
	}
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

func (r HttpResolver) Scheme() string {
	return "http"
}
