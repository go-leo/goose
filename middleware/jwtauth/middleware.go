package jwtauth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-leo/goose/server"
	"github.com/golang-jwt/jwt/v5"
)

type ctxKey struct{}

func FromContext(ctx context.Context) (*jwt.Token, bool) {
	v, ok := ctx.Value(ctxKey{}).(*jwt.Token)
	return v, ok
}

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
	return &options{
		realm: "Authorization Required",
	}
}

func Realm(realm string) Option {
	return func(o *options) {
		o.realm = realm
	}
}

func Middleware(keyFunc jwt.Keyfunc, opts ...Option) server.Middleware {
	opt := defaultOptions().apply(opts...)
	realm := "Basic realm=" + strconv.Quote(opt.realm)
	return func(response http.ResponseWriter, request *http.Request, invoker http.HandlerFunc) {
		tokenString, found := parseAuthorization(request.Header.Get("Authorization"))
		if !found {
			response.Header().Set("WWW-Authenticate", realm)
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(tokenString, keyFunc)
		if err != nil {
			response.Header().Set("WWW-Authenticate", realm)
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			response.Header().Set("WWW-Authenticate", realm)
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
		request = request.WithContext(context.WithValue(request.Context(), ctxKey{}, token))
		invoker(response, request)
	}
}

func parseAuthorization(authorization string) (string, bool) {
	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", false
	}
	return authorization[len("Bearer "):], true
}
