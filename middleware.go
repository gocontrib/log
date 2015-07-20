package log

import (
	"bufio"
	"errors"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
)

// Middleware creates logger middleware.
func Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return &middleware{h}
	}
}

// Logger middleware.
type middleware struct {
	h http.Handler
}

// ServeHTTP implementation.
func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res := &wrapper{w, 0, 200}
	Info("%s %s", r.Method, r.RequestURI)
	m.h.ServeHTTP(res, r)
	size := humanize.Bytes(uint64(res.written))
	if res.status >= 500 {
		Error("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	} else {
		Info("<< %s %s %d (%s) in %s", r.Method, r.RequestURI, res.status, size, time.Since(start))
	}
}

// Response wrapper to capture status.
type wrapper struct {
	http.ResponseWriter
	written int
	status  int
}

// capture status.
func (w *wrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// capture written bytes.
func (w *wrapper) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.written += n
	return n, err
}

// Hijack implements Hijacker interface
func (w *wrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	var h, ok = w.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("http.Hijacker is not supported")
	}
	return h.Hijack()
}

func (w *wrapper) Flush() {
	var f, ok = w.ResponseWriter.(http.Flusher)
	if !ok {
		panic("http.Flusher is not supported")
	}
	f.Flush()
}

func (w *wrapper) CloseNotify() <-chan bool {
	var c, ok = w.ResponseWriter.(http.CloseNotifier)
	if !ok {
		panic("http.CloseNotifier is not supported")
	}
	return c.CloseNotify()
}

func (w *wrapper) ReadFrom(r io.Reader) (int64, error) {
	rf, ok := w.ResponseWriter.(io.ReaderFrom)
	if !ok {
		return 0, errors.New("io.ReaderFrom is not supported")
	}
	return rf.ReadFrom(r)
}
