package httptest

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
)

var _ httptest.ResponseRecorder

// [httptest.ResponseRecorder]
type ResponseRecorder struct {
	Code      int
	HeaderMap http.Header
	Body      *bytes.Buffer
}

// [httptest.NewRecorder]
func NewRecorder() *ResponseRecorder {
	return &ResponseRecorder{Code: 200, HeaderMap: make(http.Header), Body: new(bytes.Buffer)}
}

// [httptest.ResponseRecorder.Header]
func (rw *ResponseRecorder) Header() http.Header { return rw.HeaderMap }

// [httptest.ResponseRecorder.WriteHeader]
func (rw *ResponseRecorder) WriteHeader(code int) { rw.Code = code }

// [httptest.ResponseRecorder.Write]
func (rw *ResponseRecorder) Write(buf []byte) (int, error) { return rw.Body.Write(buf) }

// [httptest.ResponseRecorder.Result]
func (rw *ResponseRecorder) Result() *http.Response {
	res := &http.Response{
		Proto:         "HTTP/1.1",
		StatusCode:    rw.Code,
		Status:        strconv.Itoa(rw.Code) + " " + http.StatusText(rw.Code),
		Header:        rw.HeaderMap,
		ContentLength: int64(rw.Body.Len()),
		Body:          io.NopCloser(bytes.NewReader(rw.Body.Bytes())),
	}
	if trailers, ok := rw.HeaderMap["Trailer"]; ok {
		res.Trailer = make(http.Header, len(trailers))
		for _, trailer := range trailers {
			for key := range strings.SplitSeq(trailer, ",") {
				res.Trailer[key] = rw.HeaderMap[key]
			}
		}
	}
	return res
}
