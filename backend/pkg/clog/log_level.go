package clog

import "log/slog"

const (
	levelInfo  = slog.LevelInfo
	levelWarn  = slog.LevelWarn
	levelError = slog.LevelError
	levelFatal = slog.Level(10)
)
