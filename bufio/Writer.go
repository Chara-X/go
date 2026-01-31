package bufio

import (
	"bufio"
	"io"
)

var _ bufio.Writer

// [bufio.Writer]
type Writer struct {
	w   io.Writer
	buf []byte
	n   int
}

// [bufio.NewWriter]
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w, buf: make([]byte, 4096)}
}

// [bufio.Writer.Write]
func (b *Writer) Write(p []byte) (n int, err error) {
	n = len(p)
	for len(p) > 0 {
		var n = copy(b.buf[b.n:], p)
		b.n += n
		p = p[n:]
		if b.n == len(b.buf) {
			b.Flush()
		}
	}
	return n, nil
}

// [bufio.Writer.Flush]
func (b *Writer) Flush() error {
	b.w.Write(b.buf[:b.n])
	b.n = 0
	return nil
}
