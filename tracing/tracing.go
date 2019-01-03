package tracing

import (
	"context"
	opentracing "github.com/opentracing/opentracing-go"
)

var isDebug bool

// IsDebug checks if tracing is configured for debuggging
func IsDebug() bool {
	return isDebug
}

// EnableDebug enables debugging traces for the session
func EnableDebug() {
	isDebug = true
}

// SetTag is a shortcut to create tags easily
func SetTag(key string, value interface{}) opentracing.Tag {
	return opentracing.Tag{Key: key, Value: value}
}

// StartSpanFromContext create a span for the context
func StartSpanFromContext(ctx context.Context, operationName string, opts ...opentracing.StartSpanOption) (opentracing.Span, context.Context) {
	// without a tracing specified this call will fail
	if isDebug {
		span, ctx := opentracing.StartSpanFromContext(ctx, operationName, opts...)
		return span, ctx
	}
	// so we fallback to this one
	span := opentracing.StartSpan(operationName, opts...)
	return span, ctx
}
