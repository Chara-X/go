package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"strings"
)

// [tar.Writer]
type Writer struct {
	w    io.Writer
	size int64
}

// [tar.NewWriter]
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// [tar.Writer.WriteHeader]
func (tw *Writer) WriteHeader(hdr *tar.Header) error {
	tw.Flush()
	var blk = make([]byte, 512)
	copy(blk[:100], hdr.Name)
	copy(blk[100:108], fmt.Sprintf("%07o", hdr.Mode))
	copy(blk[124:136], fmt.Sprintf("%011o", hdr.Size))
	copy(blk[148:156], strings.Repeat(" ", 8))
	var sum int64
	for _, v := range blk {
		sum += int64(v)
	}
	copy(blk[148:156], fmt.Sprintf("%06o", sum))
	tw.w.Write(blk[:])
	tw.size = hdr.Size
	return nil
}

// [tar.Writer.Write]
func (tw *Writer) Write(b []byte) (n int, err error) {
	return tw.w.Write(b)
}

// [tar.Writer.Flush]
func (tw *Writer) Flush() error {
	if n := tw.size % 512; n > 0 {
		tw.w.Write(make([]byte, 512-n))
	}
	return nil
}

// [tar.Writer.Close]
func (tw *Writer) Close() error {
	tw.Flush()
	tw.w.Write(make([]byte, 1024))
	return nil
}
