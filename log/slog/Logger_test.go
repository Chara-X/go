package slog_test

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/Chara-X/go/log/slog"
)

func ExampleLogger() {
	r, _ := http.NewRequest("GET", "localhost", nil)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, slog.LevelInfo))
	logger.Log(context.Background(), slog.LevelInfo, "finished", slog.Attr{
		"req", []slog.Attr{
			{"method", r.Method},
			{"url", r.URL.String()},
			{"status", http.StatusOK},
			{"duration", time.Microsecond},
		}})
	// Output:
	// {"time":"2025-12-19T23:24:12.304987963+08:00","level":"INFO+0","msg":"finished","req":{"method":"GET","url":"localhost","status":200,"duration":1000}}
}
