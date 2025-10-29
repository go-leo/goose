package resolver

import (
	"fmt"
	"net/url"
)

type ResolverError struct {
	target *url.URL
}

func (e *ResolverError) Error() string {
	return fmt.Sprintf("resolver: scheme %s is not supported", e.target.Scheme)
}

func (e *ResolverError) Target() *url.URL {
	return e.target
}
