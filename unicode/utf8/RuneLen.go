package utf8

import (
	"unicode"
	"unicode/utf8"
)

var _ = utf8.RuneLen

const (
	rune1Max     = 1<<7 - 1
	rune2Max     = 1<<11 - 1
	surrogateMin = 0xD800
	surrogateMax = 0xDFFF
	rune3Max     = 1<<16 - 1
	rune4Max     = unicode.MaxRune
)

// [utf8.RuneLen]
func RuneLen(r rune) int {
	switch {
	case r <= rune1Max:
		return 1
	case r <= rune2Max:
		return 2
	case r >= surrogateMin && r <= surrogateMax:
		return -1
	case r <= rune3Max:
		return 3
	case r <= rune4Max:
		return 4
	}
	return -1
}
