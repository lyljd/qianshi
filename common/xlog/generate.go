package xlog

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/trace"
	"runtime"
	"time"
)

func generateStandardLog(level Level, content string) {
	now := time.Now()

	ch <- map[string]any{
		"service_name": sn,
		"content":      content,
		"level":        level,
		"date":         now.Format(time.DateTime),
		"timestamp":    now.Unix(),
		"caller":       getCaller(),
	}
}

func generateWithTraceLog(ctx context.Context, level Level, content string) {
	now := time.Now()

	ch <- map[string]any{
		"service_name": sn,
		"content":      content,
		"level":        level,
		"date":         now.Format(time.DateTime),
		"timestamp":    now.Unix(),
		"caller":       getCaller(),
		"trace_id":     getTraceID(ctx),
		"span_id":      getSpanID(ctx),
	}
}

func levelConvert(level Level) string {
	switch level {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return "unknown"
	}
}

func getCaller() string {
	_, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf("%s:%d", file, line)
}

func getTraceID(ctx context.Context) string {
	return trace.TraceIDFromContext(ctx)
}

func getSpanID(ctx context.Context) string {
	return trace.SpanIDFromContext(ctx)
}
