package tar

import (
	"archive/tar"
	"bytes"
	"io"
	"strconv"
)

// [tar.Reader]
type Reader struct {
	r    io.Reader
	size int64
}

// [tar.NewReader]
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// [tar.Reader.Next]
func (tr *Reader) Next() (*tar.Header, error) {
	var blk = make([]byte, 512)
	io.ReadFull(tr.r, blk)
	if bytes.Equal(blk[:], make([]byte, 512)) {
		return nil, io.EOF
	}
	var hdr = &tar.Header{}
	hdr.Name = string(blk[:bytes.IndexByte(blk, 0)])
	hdr.Mode, _ = strconv.ParseInt(string(blk[108:][:bytes.IndexByte(blk[108:], 0)]), 8, 64)
	hdr.Size, _ = strconv.ParseInt(string(blk[124:][:bytes.IndexByte(blk[124:], 0)]), 8, 64)
	tr.size = hdr.Size
	return hdr, nil
}

// [tar.Reader.Read]
func (tr *Reader) Read(b []byte) (int, error) {
	return tr.r.Read(b)
}
