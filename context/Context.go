package context

import (
	"context"
	"time"
)

var _ context.Context

// [context.Context]
type Context interface {
	Value(key any) any
	Done() <-chan struct{}
}
type backgroundCtx struct{}

// [context.Background]
func Background() Context                   { return &backgroundCtx{} }
func (backgroundCtx) Done() <-chan struct{} { return nil }
func (backgroundCtx) Value(key any) any     { return nil }

type valueCtx struct {
	Context
	key, val any
}

// [context.WithValue]
func WithValue(parent Context, key, val any) Context {
	return &valueCtx{parent, key, val}
}
func (c *valueCtx) Value(key any) any {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}

type timeoutCtx struct {
	Context
	done chan struct{}
}

// [context.WithTimeout]
func WithTimeout(parent Context, timeout time.Duration) (ctx Context, cancel func()) {
	var c = &timeoutCtx{parent, make(chan struct{})}
	go func() {
		select {
		case <-parent.Done():
			close(c.done)
		case <-time.After(timeout):
			close(c.done)
		case <-c.done:
		}
	}()
	return c, func() { close(c.done) }
}
func (c *timeoutCtx) Done() <-chan struct{} { return c.done }
