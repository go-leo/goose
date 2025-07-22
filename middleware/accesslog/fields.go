package accesslog

import (
	"context"
	"time"

	"golang.org/x/exp/slog"
)

type builder struct {
	fields []slog.Attr
}

func (f *builder) System() *builder {
	f.fields = append(f.fields, slog.String("system", "server"))
	return f
}

func (f *builder) StartTime(startTime time.Time) *builder {
	f.fields = append(f.fields, slog.String("timestamp", startTime.Format(time.RFC3339Nano)))
	return f
}

func (f *builder) Deadline(ctx context.Context) *builder {
	if d, ok := ctx.Deadline(); ok {
		f.fields = append(f.fields, slog.String("deadline", d.Format(time.RFC3339Nano)))
	}
	return f
}

func (f *builder) Latency(duration time.Duration) *builder {
	f.fields = append(f.fields, slog.String("latency", duration.String()))
	return f
}

func (f *builder) Method(method string) *builder {
	f.fields = append(f.fields, slog.String("method", method))
	return f
}

func (f *builder) URI(uri string) *builder {
	f.fields = append(f.fields, slog.String("uri", uri))
	return f
}

func (f *builder) Proto(proto string) *builder {
	f.fields = append(f.fields, slog.String("proto", proto))
	return f
}

func (f *builder) Host(host string) *builder {
	f.fields = append(f.fields, slog.String("host", host))
	return f
}

func (f *builder) RemoteAddress(address string) *builder {
	f.fields = append(f.fields, slog.String("remote_address", address))
	return f
}

func (f *builder) Status(status int) *builder {
	f.fields = append(f.fields, slog.Int("response_status", status))
	return f
}

func (f *builder) Build() []slog.Attr {
	return f.fields
}
