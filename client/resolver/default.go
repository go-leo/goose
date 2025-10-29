package resolver

import (
	"context"
	"net/url"
)

var DefaultScheme = "http"

func init() {
	RegisterResolver(&DefaultResolver{})
}

type DefaultResolver struct {
	HttpScheme string
}

func (r DefaultResolver) Resolve(ctx context.Context, target *url.URL) (*url.URL, error) {
	if target.Scheme != r.Scheme() {
		return nil, &ResolverError{target: target}
	}
	resolved := &url.URL{}
	if r.HttpScheme != "" {
		resolved.Scheme = r.HttpScheme
	} else {
		resolved.Scheme = DefaultScheme
	}
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

func (r DefaultResolver) Scheme() string {
	return ""
}
