package timeout

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-leo/goose/server"
	"golang.org/x/exp/slog"
)

const key = "X-Leo-Timeout"

func Middleware(duration time.Duration) server.Middleware {
	return func(response http.ResponseWriter, request *http.Request, invoker http.HandlerFunc) {
		timeout := duration
		value := request.Header.Get(key)
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
		ctx := request.Context()
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		request = request.WithContext(ctx)
		invoker(response, request)
	}
}
