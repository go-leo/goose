package timeout

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-leo/goose"
	"golang.org/x/exp/slog"
)

const key = "X-Leo-Timeout"

func Middleware(duration time.Duration) goose.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeout := duration
			value := r.Header.Get(key)
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
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
