package utf8_test

import (
	"fmt"

	"github.com/Chara-X/go/unicode/utf8"
)

func ExampleRuneCount() {
	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))
	// Output:
	// bytes = 13
	// runes = 9
}
