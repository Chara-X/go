package slog

import "log/slog"

var _ = slog.LevelDebug

// [slog.Level]
const (
	// [slog.LevelDebug]
	LevelDebug Level = -4
	// [slog.LevelInfo]
	LevelInfo Level = 0
	// [slog.LevelWarn]
	LevelWarn Level = 4
	// [slog.LevelError]
	LevelError Level = 8
)
const (
	// [slog.TimeKey]
	TimeKey = "time"
	// [slog.LevelKey]
	LevelKey = "level"
	// [slog.MessageKey]
	MessageKey = "msg"
)
