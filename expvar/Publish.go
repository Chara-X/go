package expvar

import "expvar"

var _ = expvar.Publish
var vars Map

// [expvar.Publish]
func Publish(name string, v Var) {
	if _, dup := vars.m.LoadOrStore(name, v); dup {
		panic("Reuse of exported var name")
	}
}
