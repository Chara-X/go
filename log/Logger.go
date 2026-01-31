package log

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

var _ log.Logger

// [log.Logger]
type Logger struct {
	out    io.Writer
	prefix string
}

// [log.New]
func New(out io.Writer, prefix string) *Logger {
	return &Logger{out: out, prefix: prefix}
}

// [log.Logger.SetOutput]
func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
}

// [log.Logger.SetPrefix]
func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

// [log.Logger.Output]
func (l *Logger) Output(calldepth int, s string) error {
	var year, month, day = time.Now().Date()
	var _, file, line, _ = runtime.Caller(calldepth)
	l.out.Write(fmt.Appendf(nil, "%s%04d/%02d/%02d %s:%d: %s\n", l.prefix, year, month, day, file, line, s))
	return nil
}
