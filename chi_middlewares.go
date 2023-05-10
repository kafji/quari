package quari

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

var debug = false

type SlogFormatter struct{}

func NewSlogFormatter() SlogFormatter {
	return SlogFormatter{}
}

func (s SlogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	logger := slog.With(
		"tag", "http/handler",
		"method", r.Method,
		"path", r.URL.Path,
		"peer", r.RemoteAddr,
		"protocol", r.Proto,
	)

	if xs, ok := r.Header[http.CanonicalHeaderKey("x-forwarded-for")]; ok && len(xs) > 0 {
		logger = logger.With("origin", xs[0])
	}

	if xs, ok := r.Header[http.CanonicalHeaderKey("user-agent")]; ok && len(xs) > 0 {
		logger = logger.With("user_agent", xs[0])
	}

	return slogEntry{logger}
}

type slogEntry struct {
	logger *slog.Logger
}

func (s slogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	s.logger.Info("request",
		"status", status,
		"latency_ms", elapsed/time.Millisecond)
}

func (s slogEntry) Panic(v interface{}, stack []byte) {
	s.logger.Error("panic", "err", v)
	if debug {
		middleware.PrintPrettyStack(v)
	}
}

var _ middleware.LogFormatter = NewSlogFormatter()
