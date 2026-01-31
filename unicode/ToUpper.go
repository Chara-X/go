package unicode

import "unicode"

var _ = unicode.ToUpper

// [unicode.ToUpper]
func ToUpper(r rune) rune { return To(UpperCase, r) }
