package util

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

func JegerTrace(msg string) (opentracing.Span, context.Context) {
	trace, ctx := opentracing.StartSpanFromContext(context.Background(), msg)
	return trace, ctx
}
