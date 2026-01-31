package testing

import (
	"os"
	"path/filepath"
	"testing"
)

// [testing.T]
type T struct {
	t      *testing.T
	signal chan struct{}
	failed bool
}

func (c *T) Name() string     { return c.t.Name() }
func (c *T) TempDir() string  { return c.t.TempDir() }
func (c *T) Log(args ...any)  { c.t.Log(args...) }
func (c *T) Fatal()           { c.t.Fatal() }
func (c *T) Cleanup(f func()) { c.t.Cleanup(f) }

// [testing.T.Chdir]
func (t *T) Chdir(dir string) {
	oldwd, _ := os.Open(".")
	os.Chdir(dir)
	if !filepath.IsAbs(dir) {
		dir, _ = os.Getwd()
	}
	t.Setenv("PWD", dir)
	t.Cleanup(func() {
		oldwd.Chdir()
		oldwd.Close()
	})
}

// [testing.T.Setenv]
func (t *T) Setenv(key, value string) {
	prevValue, ok := os.LookupEnv(key)
	os.Setenv(key, value)
	if ok {
		t.Cleanup(func() {
			os.Setenv(key, prevValue)
		})
	} else {
		t.Cleanup(func() {
			os.Unsetenv(key)
		})
	}
}

// [testing.T.Run]
func (t *T) Run(name string, f func(t *T)) bool {
	go func() {
		defer func() {
			// t.runCleanup()
			err := recover()
			defer func() {
				t.signal <- struct{}{}
			}()
			if err != nil {
				t.t.Fail()
				// t.report()
				panic(err)
			}
			// t.report()
		}()
		f(t)
	}()
	t = &T{}
	<-t.signal
	return !t.failed
}
