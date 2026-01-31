package context_test

import (
	"fmt"
	"time"

	"github.com/Chara-X/go/context"
)

func ExampleWithValue() {
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, favContextKey("color"))
	// Output:
	// found value: Go
	// key not found: color
}
func ExampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("ready")
		cancel()
	case <-ctx.Done():
		fmt.Println("printed context deadline exceeded")
	}
	// Output:
	// printed context deadline exceeded
}
