package slog

import (
	"context"
	"log/slog"
	"time"
)

var _ slog.Logger

// [slog.Logger]
type Logger struct{ Handler }

// [slog.New]
func New(h Handler) *Logger { return &Logger{h} }

// [slog.Logger.WithGroup]
func (l *Logger) WithGroup(name string) *Logger { return &Logger{l.Handler.WithGroup(name)} }

// [slog.Logger.With]
func (l *Logger) WithAttrs(attrs []Attr) *Logger { return &Logger{l.Handler.WithAttrs(attrs)} }

// [slog.Logger.Log]
func (l *Logger) Log(ctx context.Context, level Level, msg string, attrs ...Attr) {
	if l.Enabled(ctx, level) {
		r := NewRecord(time.Now(), level, msg)
		r.AddAttrs(attrs...)
		l.Handle(ctx, r)
	}
}
