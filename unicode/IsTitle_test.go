package unicode_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode"
)

func ExampleIsTitle() {
	fmt.Printf("%t\n", unicode.IsTitle('ǅ'))
	fmt.Printf("%t\n", unicode.IsTitle('a'))
	// Output:
	// true
	// false
}
