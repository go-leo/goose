package context

import (
	"github.com/gin-gonic/gin"
)

type ContextFunc func(ctx *gin.Context) *gin.Context

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
		contextFunc: func(ctx *gin.Context) *gin.Context { return ctx },
	}
}

func WithContextFunc(contextFunc ContextFunc) Option {
	return func(o *options) {
		o.contextFunc = contextFunc
	}
}

func Middleware(opts ...Option) gin.HandlerFunc {
	opt := defaultOptions().apply(opts...)
	return func(c *gin.Context) {
		c = opt.contextFunc(c)
		c.Request = c.Request.WithContext(c)
		c.Next()
	}
}
