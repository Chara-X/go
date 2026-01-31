package expvar_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/Chara-X/go/expvar"
)

func ExampleHandler() {
	reqs := expvar.NewInt("requests")
	reqs.Add(10)
	load := expvar.NewFloat("load")
	load.Set(0.75)
	status := expvar.NewString("status")
	status.Set("ok")
	counters := expvar.NewMap("counters")
	counters.Add("hits", 3)
	expvar.Publish("uptime_seconds", expvar.Func(func() any { return 123 }))
	w := httptest.NewRecorder()
	expvar.Handler().ServeHTTP(w, nil)
	var vars map[string]any
	if err := json.NewDecoder(w.Body).Decode(&vars); err != nil {
		panic(err)
	}
	fmt.Println("requests:", vars["requests"])
	fmt.Println("load:", vars["load"])
	fmt.Println("status:", vars["status"])
	fmt.Println("counters:", vars["counters"])
	fmt.Println("uptime_seconds:", vars["uptime_seconds"])
	// Output:
	// requests: 10
	// load: 0.75
	// status: ok
	// counters: map[hits:3]
	// uptime_seconds: 123
}
