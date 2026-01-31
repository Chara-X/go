package unicode_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode"
)

func ExampleToUpper() {
	fmt.Printf("%#U\n", unicode.ToUpper('g'))
	// Output: U+0047 'G'
}
