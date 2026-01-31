package unicode_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode"
)

func ExampleIsLower() {
	fmt.Printf("%t\n", unicode.IsLower('a'))
	fmt.Printf("%t\n", unicode.IsLower('A'))
	// Output:
	// true
	// false
}
