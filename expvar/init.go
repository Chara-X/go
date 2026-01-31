package expvar

import "net/http"

func init() { http.Handle("GET /vars", Handler()) }
