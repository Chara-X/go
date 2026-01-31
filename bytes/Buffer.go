package bytes

import (
	"bytes"
	"io"
)

var _ bytes.Buffer

// [bytes.Buffer]
type Buffer struct {
	buf []byte
	off int
}

// [bytes.NewBuffer]
func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }

// [bytes.Buffer.Len]
func (b *Buffer) Len() int { return len(b.buf) - b.off }

// [bytes.Buffer.Read]
func (b *Buffer) Read(p []byte) (n int, err error) {
	if b.off >= len(b.buf) {
		b.buf, b.off = b.buf[:0], 0
		return 0, io.EOF
	}
	n = copy(p, b.buf[b.off:])
	b.off += n
	return n, nil
}

// [bytes.Buffer.Write]
func (b *Buffer) Write(p []byte) (n int, err error) {
	if len(b.buf)+len(p) > cap(b.buf) {
		copy(b.buf, b.buf[b.off:])
		b.buf = b.buf[:len(b.buf)-b.off]
	}
	b.buf = append(b.buf, p...)
	return len(p), nil
}

// [bytes.Buffer.Bytes]
func (b *Buffer) Bytes() []byte { return b.buf[b.off:] }
