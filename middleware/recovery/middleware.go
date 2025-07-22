package recovery

import (
	"runtime"

	"github.com/gin-gonic/gin"
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

type HandlerFunc func(ctx *gin.Context, p any)

// RecoveryHandler customizes the function for recovering from a panic.
func RecoveryHandler(f HandlerFunc) Option {
	return func(o *options) {
		o.handler = f
	}
}

func Middleware(opts ...Option) gin.HandlerFunc {
	opt := defaultOptions().apply(opts...)
	return func(c *gin.Context) {
		defer func() {
			p := recover()
			if p == nil {
				return
			}
			opt.handler(c, p)
		}()
		c.Next()
	}
}

func defaultHandler(ctx *gin.Context, p any) {
	stack := make([]byte, 64<<10)
	stack = stack[:runtime.Stack(stack, false)]
	slog.ErrorContext(ctx, "panic caught", "panic", p, "stack", string(stack))
}
