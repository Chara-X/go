package expvar

import (
	"expvar"
	"net/http"
)

var _ = expvar.Handler

// [expvar.Handler]
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte(vars.String()))
	})
}
