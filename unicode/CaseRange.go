package unicode

import "unicode"

var _ unicode.CaseRange

// [unicode.CaseRange]
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta [3]rune
}
