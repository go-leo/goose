package accesslog

import (
	"context"
	"net/http"
	"reflect"
	"time"
	"unsafe"

	"github.com/go-leo/goose"
	"golang.org/x/exp/slog"
)

type LoggerFactory func(ctx context.Context) *slog.Logger

type options struct {
	loggerFactory LoggerFactory
	level         slog.Level
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

func WithLoggerFactory(loggerFactory LoggerFactory) Option {
	return func(o *options) {
		o.loggerFactory = loggerFactory
	}
}

func WithLevel(level slog.Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func Middleware(opts ...Option) goose.MiddlewareFunc {
	opt := defaultOptions().apply(opts...)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if opt.loggerFactory == nil {
				next.ServeHTTP(w, r)
				return
			}
			startTime := time.Now()
			sw := &statusCodeResponseWriter{ResponseWriter: w}
			next.ServeHTTP(sw, r)
			ctx := r.Context()
			logger := opt.loggerFactory(ctx)
			route := getRoute(r)
			builder := new(builder).
				System().
				StartTime(startTime).
				Deadline(ctx).
				Method(r.Method).
				URI(r.RequestURI).
				Proto(r.Proto).
				Host(r.Host).
				RemoteAddress(r.RemoteAddr).
				Status(sw.statusCode).
				Latency(time.Since(startTime))
			logger.LogAttrs(ctx, opt.level, route, builder.Build()...)
		})
	}
}

type statusCodeResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusCodeResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func getRoute(r *http.Request) string {
	defer func() {
		_ = recover()
	}()
	strVal := reflect.ValueOf(r).Elem().FieldByName("pat").Elem().FieldByName("str")
	return reflect.NewAt(strVal.Type(), unsafe.Pointer(strVal.UnsafeAddr())).Elem().String()
}
