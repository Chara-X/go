package unicode_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode"
)

func ExampleToTitle() {
	fmt.Printf("%#U\n", unicode.ToTitle('g'))
	// Output: U+0047 'G'
}
