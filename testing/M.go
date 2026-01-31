package testing

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

// [testing.M]
type M struct{ examples []testing.InternalExample }

// [testing.M.Run]
func (m *M) Run() (code int) {
	flag.Parse()
	var timer = time.AfterFunc(*timeout, func() {
		panic("timed out")
	})
	code = 0
	for _, eg := range m.examples {
		if !run(eg) {
			code = 1
		}
	}
	timer.Stop()
	return
}
func run(eg testing.InternalExample) (pass bool) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	go func() {
		var buf strings.Builder
		io.Copy(&buf, r)
		r.Close()
		outC <- buf.String()
	}()
	defer func() {
		w.Close()
		os.Stdout = stdout
		if eg.Output != <-outC {
			pass = false
		}
	}()
	eg.F()
	return
}
