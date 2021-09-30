package request

import (
	"context"
)

type key int

const (
	clientIPKey key = iota
	requestIDKey
)

func WithClientIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, clientIPKey, ip)
}

func ClientIPFrom(ctx context.Context) string {
	return ctx.Value(clientIPKey).(string)
}

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestIDFrom(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}
