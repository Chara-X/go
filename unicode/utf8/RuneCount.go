package utf8

import "unicode/utf8"

var _ = utf8.RuneCount

// [utf8.RuneCount]
func RuneCount(p []byte) (n int) {
	for range string(p) {
		n++
	}
	return n
}
