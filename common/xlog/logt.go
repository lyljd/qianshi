package xlog

import (
	"context"
)

func DebugT(ctx context.Context, content string) {
	generateWithTraceLog(ctx, LevelDebug, content)
}

func InfoT(ctx context.Context, content string) {
	generateWithTraceLog(ctx, LevelInfo, content)
}

func WarnT(ctx context.Context, content string) {
	generateWithTraceLog(ctx, LevelWarn, content)
}

func ErrorT(ctx context.Context, content string) {
	generateWithTraceLog(ctx, LevelError, content)
}

func FatalT(ctx context.Context, content string) {
	generateWithTraceLog(ctx, LevelFatal, content)
}
