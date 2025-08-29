package jwtauth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-leo/goose"
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

func Middleware(keyFunc jwt.Keyfunc, opts ...Option) goose.MiddlewareFunc {
	opt := defaultOptions().apply(opts...)
	realm := "Basic realm=" + strconv.Quote(opt.realm)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, found := parseAuthorization(r.Header.Get("Authorization"))
			if !found {
				w.Header().Set("WWW-Authenticate", realm)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			token, err := jwt.Parse(tokenString, keyFunc)
			if err != nil {
				w.Header().Set("WWW-Authenticate", realm)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				w.Header().Set("WWW-Authenticate", realm)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), ctxKey{}, token))
			next.ServeHTTP(w, r)
		})
	}
}

func parseAuthorization(authorization string) (string, bool) {
	if !strings.HasPrefix(authorization, "Bearer ") {
		return "", false
	}
	return authorization[len("Bearer "):], true
}
