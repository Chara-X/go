package bytes

import "fmt"

func ExampleBuffer() {
	var b = NewBuffer([]byte{})
	b.Write([]byte("123"))
	b.Write([]byte("456"))
	var p = make([]byte, 3)
	b.Read(p)
	fmt.Println(string(p))
	fmt.Println(string(b.Bytes()))
	// Output:
	// 123
	// 456
}
