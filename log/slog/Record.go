package slog

import (
	"log/slog"
	"slices"
	"time"
)

var _ slog.Record

// [slog.Record]
type Record struct {
	Time    time.Time
	Message string
	Level   Level
	attrs   []Attr
}

// [slog.NewRecord]
func NewRecord(t time.Time, level Level, msg string) Record { return Record{t, msg, level, nil} }

// [slog.Record.AddAttrs]
func (r *Record) AddAttrs(attrs ...Attr) { r.attrs = append(r.attrs, attrs...) }

// [slog.Record.Clone]
func (r Record) Clone() Record { return Record{r.Time, r.Message, r.Level, slices.Clone(r.attrs)} }
