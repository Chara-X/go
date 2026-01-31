package utf8_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode/utf8"
)

func ExampleRuneLen() {
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
	// Output:
	// 1
	// 3
}
