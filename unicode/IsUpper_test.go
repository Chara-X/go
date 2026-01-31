package unicode_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode"
)

func ExampleIsUpper() {
	fmt.Printf("%t\n", unicode.IsUpper('A'))
	fmt.Printf("%t\n", unicode.IsUpper('a'))
	// Output:
	// true
	// false
}
