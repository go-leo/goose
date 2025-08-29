package recovery

import (
	"net/http"
	"runtime"

	"github.com/go-leo/goose"
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

func Middleware(opts ...Option) goose.MiddlewareFunc {
	opt := defaultOptions().apply(opts...)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				p := recover()
				if p == nil {
					return
				}
				opt.handler(w, r, p)
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request, p any) {
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]
	slog.ErrorContext(r.Context(), "panic caught", "panic", p, "stack", string(stack))
}
