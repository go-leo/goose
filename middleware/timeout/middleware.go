package timeout

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const key = "X-Leo-Timeout"

func Middleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := duration
		value := c.GetHeader(key)
		if value != "" {
			switch {
			case strings.HasSuffix(value, "n"):
				value = value + "s"
			case strings.HasSuffix(value, "u"):
				value = value + "s"
			case strings.HasSuffix(value, "m"):
				value = value + "s"
			case strings.HasSuffix(value, "S"):
				value = strings.Replace(value, "S", "s", 1)
			case strings.HasSuffix(value, "M"):
				value = strings.Replace(value, "M", "m", 1)
			case strings.HasSuffix(value, "H"):
				value = strings.Replace(value, "H", "h", 1)
			}
			incomingDuration, err := time.ParseDuration(value)
			if err != nil {
				slog.Error("timeout parse error", slog.String("timeout", value), slog.String("error", err.Error()))
			} else {
				timeout = min(incomingDuration, duration)
			}
		}
		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
