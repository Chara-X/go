package testing

import (
	"flag"
	"testing"
	"time"
)

var _ = testing.Init
var timeout *time.Duration

// [testing.Init]
func init() {
	timeout = flag.Duration("test.timeout", 0, "panic test binary after duration `d` (default 0, timeout disabled)")
}
