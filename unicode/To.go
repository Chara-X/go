package unicode

import "unicode"

var _ = unicode.To

// [unicode.To]
func To(cas int, r rune) rune {
	for _, caseRange := range CaseRanges {
		if rune(caseRange.Lo) <= r && r <= rune(caseRange.Hi) {
			return r + caseRange.Delta[cas]
		}
	}
	return r
}
