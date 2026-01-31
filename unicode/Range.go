package unicode

import "unicode"

var _ unicode.Range32

// [unicode.Range32]
type Range struct {
	Lo     uint32
	Hi     uint32
	Stride uint32
}
