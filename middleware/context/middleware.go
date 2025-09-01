package context

import (
	"context"
	"net/http"

	"github.com/go-leo/goose/server"
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

func Middleware(opts ...Option) server.Middleware {
	opt := defaultOptions().apply(opts...)
	return func(response http.ResponseWriter, request *http.Request, invoker http.HandlerFunc) {
		ctx := request.Context()
		ctx = opt.contextFunc(ctx)
		request = request.WithContext(ctx)
		invoker(response, request)
	}
}
