package unicode

import "unicode"

var _ = unicode.IsLower

// [unicode.IsLower]
func IsLower(r rune) bool { return Is(Lower, r) }
