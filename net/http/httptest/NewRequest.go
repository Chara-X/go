package httptest

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = httptest.NewRequest

// [httptest.NewRequest]
func NewRequest(method, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	req.Close = false
	req.RemoteAddr = "192.0.2.1:1234"
	if strings.HasPrefix(url, "https://") {
		req.TLS = &tls.ConnectionState{
			Version:           tls.VersionTLS12,
			HandshakeComplete: true,
			ServerName:        req.Host,
		}
	}
	return req
}
