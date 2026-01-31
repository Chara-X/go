package bufio

import (
	"bufio"
	"io"
)

var _ bufio.Reader

// [bufio.Reader]
type Reader struct {
	r    io.Reader
	buf  []byte
	i, j int
}

// [bufio.NewReader]
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r, buf: make([]byte, 4096)}
}

// [bufio.Reader.Read]
func (b *Reader) Read(p []byte) (n int, err error) {
	if b.i == b.j {
		b.i = 0
		b.j, _ = b.r.Read(b.buf)
		if b.j == 0 {
			return 0, io.EOF
		}
	}
	n = copy(p, b.buf[b.i:b.j])
	b.i += n
	return n, nil
}

// [bufio.Reader.Peek]
func (b *Reader) Peek(n int) ([]byte, error) {
	for b.j-b.i < min(n, len(b.buf)) {
		copy(b.buf, b.buf[b.i:b.j])
		b.i = 0
		b.j -= b.i
		var n, _ = b.r.Read(b.buf[b.j:])
		if n == 0 {
			return b.buf[:b.j], io.ErrUnexpectedEOF
		}
		b.j += n
	}
	return b.buf[b.i : b.i+n], nil
}
