package clog

import (
	"context"
	"log/slog"
)

func Debug(ctx context.Context, msg string) {
	cl.logWithLevel(ctx, msg, nil, slog.LevelDebug)
}

func DebugWithError(ctx context.Context, err error) {
	cl.logWithLevel(ctx, "", err, slog.LevelDebug)
}

func Info(ctx context.Context, msg string) {
	cl.logWithLevel(ctx, msg, nil, slog.LevelInfo)
}

func InfoWithError(ctx context.Context, err error) {
	cl.logWithLevel(ctx, "", err, slog.LevelInfo)
}

func Warn(ctx context.Context, msg string) {
	cl.logWithLevel(ctx, msg, nil, slog.LevelWarn)
}

func WarnWithError(ctx context.Context, err error) {
	cl.logWithLevel(ctx, "", err, slog.LevelWarn)
}

func Error(ctx context.Context, err error) {
	cl.logWithLevel(ctx, "", err, slog.LevelError)
}

func Fatal(ctx context.Context, err error) {
	cl.logWithLevel(ctx, "", err, levelFatal)
}
