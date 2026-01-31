package slog

import (
	"fmt"
	"log/slog"
	"sync/atomic"
)

var _ slog.Leveler

// [slog.Leveler]
type Leveler interface{ Level() Level }

// [slog.Level]
type Level int

// [slog.Level.Level]
func (l Level) Level() Level { return l }

// [slog.Level.String]
func (l Level) String() string {
	switch {
	case l < LevelInfo:
		return fmt.Sprintf("DEBUG%+d", l-LevelDebug)
	case l < LevelWarn:
		return fmt.Sprintf("INFO%+d", l-LevelInfo)
	case l < LevelError:
		return fmt.Sprintf("WARN%+d", l-LevelWarn)
	default:
		return fmt.Sprintf("ERROR%+d", l-LevelError)
	}
}

// [slog.LevelVar]
type LevelVar struct{ val atomic.Int64 }

// [slog.LevelVar.Level]
func (v *LevelVar) Level() Level { return Level(v.val.Load()) }

// [slog.LevelVar.Set]
func (v *LevelVar) Set(l Level) { v.val.Store(int64(l)) }
