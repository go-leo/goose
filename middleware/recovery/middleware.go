package recovery

import (
	"net/http"
	"runtime"

	"github.com/go-leo/goose/server"
	"golang.org/x/exp/slog"
)

type options struct {
	handler HandlerFunc
}
type Option func(*options)

func defaultOptions() *options {
	return &options{
		handler: defaultHandler,
	}
}

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request, p any)

// RecoveryHandler customizes the function for recovering from a panic.
func RecoveryHandler(f HandlerFunc) Option {
	return func(o *options) {
		o.handler = f
	}
}

func Middleware(opts ...Option) server.Middleware {
	opt := defaultOptions().apply(opts...)
	return func(response http.ResponseWriter, request *http.Request, invoker http.HandlerFunc) {
		defer func() {
			p := recover()
			if p == nil {
				return
			}
			opt.handler(response, request, p)
		}()
		invoker(response, request)
	}
}

func defaultHandler(response http.ResponseWriter, request *http.Request, p any) {
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]
	slog.ErrorContext(request.Context(), "panic caught", "panic", p, "stack", string(stack))
}
