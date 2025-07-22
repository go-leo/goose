package accesslog

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type LoggerFactory func(ctx context.Context) *slog.Logger

type options struct {
	loggerFactory LoggerFactory
	level         slog.Level
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
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

func Middleware(opts ...Option) gin.HandlerFunc {
	o := defaultOptions()
	o.apply(opts...)
	return func(c *gin.Context) {
		if o.loggerFactory == nil {
			c.Next()
			return
		}
		startTime := time.Now()
		c.Next()
		ctx := c.Request.Context()
		logger := o.loggerFactory(ctx)
		r := c.Request
		builder := new(builder).
			System().
			StartTime(startTime).
			Deadline(ctx).
			Method(r.Method).
			URI(r.RequestURI).
			Proto(r.Proto).
			Host(r.Host).
			RemoteAddress(r.RemoteAddr).
			Status(c.Writer.Status()).
			Latency(time.Since(startTime))
		logger.Log(ctx, o.level, c.FullPath(), builder.Build()...)
	}
}
