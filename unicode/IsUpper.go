package unicode

import "unicode"

var _ = unicode.IsUpper

// [unicode.IsUpper]
func IsUpper(r rune) bool { return Is(Upper, r) }
