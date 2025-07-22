package gincontext

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ContextFunc func(ctx context.Context) context.Context

type options struct {
	contextFunc ContextFunc
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
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

func Middleware(opts ...Option) gin.HandlerFunc {
	o := defaultOptions()
	o.apply(opts...)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = o.contextFunc(ctx)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
