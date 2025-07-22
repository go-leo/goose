package accesslog

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type LoggerFactory func(ctx *gin.Context) *slog.Logger

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

func Middleware(opts ...Option) gin.HandlerFunc {
	opt := defaultOptions().apply(opts...)
	return func(c *gin.Context) {
		if opt.loggerFactory == nil {
			c.Next()
			return
		}
		startTime := time.Now()
		c.Next()
		logger := opt.loggerFactory(c)
		r := c.Request
		builder := new(builder).
			System().
			StartTime(startTime).
			Deadline(c).
			Method(r.Method).
			URI(r.RequestURI).
			Proto(r.Proto).
			Host(r.Host).
			RemoteAddress(r.RemoteAddr).
			Status(c.Writer.Status()).
			Latency(time.Since(startTime))
		logger.LogAttrs(c, opt.level, c.FullPath(), builder.Build()...)
	}
}
