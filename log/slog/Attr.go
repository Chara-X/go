package slog

import (
	"log/slog"
)

var _ slog.Attr

// [slog.Attr]
type Attr struct {
	Key   string
	Value any
}
