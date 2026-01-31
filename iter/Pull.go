package iter

import (
	"iter"
)

var _ = iter.Pull[any]

// [iter.Pull]
func Pull[V any](seq func(yield func(V) bool)) (next func() (V, bool), stop func()) {
	var values, stopped = make(chan V, 1), make(chan struct{})
	go func() {
		defer close(values)
		seq(func(v V) bool {
			select {
			case <-stopped:
				return false
			case values <- v:
				return true
			}
		})
	}()
	return func() (v V, ok bool) {
			v, ok = <-values
			return
		}, func() {
			close(stopped)
		}
}
