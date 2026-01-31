package utf8

import "unicode/utf8"

var _ = utf8.EncodeRune

// [utf8.EncodeRune]
func EncodeRune(p []byte, r rune) int {
	switch {
	case r <= rune1Max:
		p[0] = byte(r)
		return 1
	case r <= rune2Max:
		p[0], p[1] = 0xC0|byte(r>>6), 0x80|byte(r)&0x3F
		return 2
	case r >= surrogateMin && r <= surrogateMax:
		return 0
	case r <= rune3Max:
		p[0], p[1], p[2] = 0xE0|byte(r>>12), 0x80|byte(r>>6)&0x3F, 0x80|byte(r)&0x3F
		return 3
	case r <= rune4Max:
		p[0], p[1], p[2], p[3] = 0xF0|byte(r>>18), 0x80|byte(r>>12)&0x3F, 0x80|byte(r>>6)&0x3F, 0x80|byte(r)&0x3F
		return 4
	}
	return 0
}
