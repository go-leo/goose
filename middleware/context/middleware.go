package context

import (
	"context"
	"net/http"

	"github.com/go-leo/goose"
)

type ContextFunc func(ctx context.Context) context.Context

type options struct {
	contextFunc ContextFunc
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(o *options)

func defaultOptions() *options {
	return &options{
		contextFunc: func(ctx context.Context) context.Context { return ctx },
	}
}

func WithContextFunc(contextFunc ContextFunc) Option {
	return func(o *options) {
		o.contextFunc = contextFunc
	}
}

func Middleware(opts ...Option) goose.MiddlewareFunc {
	opt := defaultOptions().apply(opts...)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = opt.contextFunc(ctx)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
