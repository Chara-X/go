package slog

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"slices"
	"sync"
)

var _ slog.Handler

// [slog.Handler]
type Handler interface {
	WithGroup(name string) Handler
	WithAttrs(attrs []Attr) Handler
	Enabled(context.Context, Level) bool
	Handle(context.Context, Record) error
}

// [slog.JSONHandler]
type JSONHandler struct {
	w      io.Writer
	mu     *sync.Mutex
	level  Leveler
	groups []string
	attrs  []Attr
}

// [slog.NewJSONHandler]
func NewJSONHandler(w io.Writer, level Leveler) *JSONHandler {
	return &JSONHandler{w, &sync.Mutex{}, level, nil, nil}
}

// [slog.JSONHandler.WithGroup]
func (h *JSONHandler) WithGroup(name string) Handler {
	return &JSONHandler{h.w, h.mu, h.level, append(slices.Clone(h.groups), name), h.attrs}
}

// [slog.JSONHandler.WithAttrs]
func (h *JSONHandler) WithAttrs(attrs []Attr) Handler {
	return &JSONHandler{h.w, h.mu, h.level, h.groups, append(h.attrs, slices.Clone(attrs)...)}
}

// [slog.JSONHandler.Enabled]
func (h *JSONHandler) Enabled(_ context.Context, level Level) bool { return level >= h.level.Level() }

// [slog.JSONHandler.Handle]
func (h *JSONHandler) Handle(_ context.Context, r Record) error {
	out := map[string]any{}
	out[TimeKey], out[LevelKey], out[MessageKey] = r.Time, r.Level.String(), r.Message
	obj := out
	for _, group := range h.groups {
		o := map[string]any{}
		obj[group] = o
		obj = o
	}
	for _, attr := range append(h.attrs, r.attrs...) {
		switch v := attr.Value.(type) {
		case []Attr:
			o := map[string]any{}
			for _, attr := range v {
				o[attr.Key] = attr.Value
			}
			obj[attr.Key] = o
		default:
			obj[attr.Key] = v
		}
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	return json.NewEncoder(h.w).Encode(out)
}
