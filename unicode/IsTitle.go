package unicode

import "unicode"

var _ = unicode.IsTitle

// [unicode.IsTitle]
func IsTitle(r rune) bool { return Is(Title, r) }
