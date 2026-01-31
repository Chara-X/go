package unicode

import "unicode"

var _ = unicode.ToLower

// [unicode.ToLower]
func ToLower(r rune) rune { return To(LowerCase, r) }
