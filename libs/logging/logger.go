package logging

import (
	"context"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
)

var (
	global *zap.SugaredLogger
)

func InitGlobal(opts ...zap.Option) error {
	logger, err := zap.NewProduction(opts...)
	if err != nil {
		return err
	}

	global = logger.Sugar()
	return nil
}

func WithCtx(ctx context.Context) *zap.SugaredLogger {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return global
	}

	if sc, ok := span.Context().(jaeger.SpanContext); ok {
		return global.Desugar().With(
			zap.Stringer("trace_id", sc.TraceID()),
			zap.Stringer("span_id", sc.SpanID()),
		).Sugar()
	}

	return global
}

func Infof(msg string, args ...any) {
	if global != nil {
		global.Infof(msg, args...)
	} else {
		log.Println("INFO", fmt.Sprintf(msg, args...))
	}
}

func Warnf(msg string, args ...any) {
	if global != nil {
		global.Warnf(msg, args...)
	} else {
		log.Println("WARN", fmt.Sprintf(msg, args...))
	}
}

func Errorf(msg string, args ...any) {
	if global != nil {
		global.Errorf(msg, args...)
	} else {
		log.Println("ERROR", fmt.Sprintf(msg, args...))
	}
}
