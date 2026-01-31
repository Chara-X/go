package unicode_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode"
)

func ExampleToLower() {
	fmt.Printf("%#U\n", unicode.ToLower('G'))
	// Output: U+0067 'g'
}
