package basicauth

import (
	"github.com/gin-gonic/gin"
)

type options struct {
	realm string
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Option func(o *options)

func defaultOptions() *options {
	return &options{}
}

func Realm(realm string) Option {
	return func(o *options) {
		o.realm = realm
	}
}

func Middleware(accounts gin.Accounts, opts ...Option) gin.HandlerFunc {
	opt := defaultOptions().apply(opts...)
	return gin.BasicAuthForRealm(accounts, opt.realm)
}
