package unicode

import "unicode"

var _ = unicode.Is

// [unicode.Is]
func Is(ranges *[]Range, r rune) bool {
	r32 := uint32(r)
	for _, rang := range *ranges {
		if r32 < rang.Lo {
			return false
		}
		if r32 <= rang.Hi {
			return rang.Stride == 1 || (r32-rang.Lo)%rang.Stride == 0
		}
	}
	return false
}
