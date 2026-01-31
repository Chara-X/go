package utf8

import "unicode/utf8"

var _ = utf8.DecodeRune

// [utf8.DecodeRune]
func DecodeRune(p []byte) (r rune, size int) {
	switch {
	case p[0]&0x80 == 0:
		return rune(p[0]), 1
	case p[0]&0xE0 == 0xC0:
		return rune(p[0]&0x1F)<<6 | rune(p[1]&0x3F), 2
	case p[0]&0xF0 == 0xE0:
		return rune(p[0]&0x0F)<<12 | rune(p[1]&0x3F)<<6 | rune(p[2]&0x3F), 3
	case p[0]&0xF8 == 0xF0:
		return rune(p[0]&0x07)<<18 | rune(p[1]&0x3F)<<12 | rune(p[2]&0x3F)<<6 | rune(p[3]&0x3F), 4
	}
	return 0, 0
}
