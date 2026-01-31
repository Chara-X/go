package unicode

import "unicode"

var _ = unicode.ToTitle

// [unicode.ToTitle]
func ToTitle(r rune) rune { return To(TitleCase, r) }
