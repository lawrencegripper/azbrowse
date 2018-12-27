package tracing

import (
	"context"
	opentracing "github.com/opentracing/opentracing-go"
)

// SetTag is a shortcut to create tags easily
func SetTag(key string, value interface{}) opentracing.Tag {
	return opentracing.Tag{Key: key, Value: value}
}

// StartSpanFromContext create a span for the context
func StartSpanFromContext(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, operationName, opts...)
	return span, ctx
}
