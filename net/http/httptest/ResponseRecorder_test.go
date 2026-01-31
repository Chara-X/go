package httptest_test

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Chara-X/go/net/http/httptest"
)

func ExampleResponseRecorder() {
	handler := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}
