package utf8_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode/utf8"
)

func ExampleEncodeRune() {
	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '世')
	fmt.Println(buf)
	fmt.Println(n)
	// Output:
	// [228 184 150]
	// 3
}
