package unicode

import "unicode"

var _ = unicode.MaxRune

const (
	MaxRune         = '\U0010FFFF' // [unicode.MaxRune]
	ReplacementChar = '\uFFFD'     // [unicode.ReplacementChar]
)
const (
	UpperCase = iota // [unicode.UpperCase]
	LowerCase        // [unicode.LowerCase]
	TitleCase        // [unicode.TitleCase]
)
