package middleware

import (
	"log/slog"
	"net/http"
	"os"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LoggingMiddleware(next http.Handler) http.Handler {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		logger.Info("Request received", "method", r.Method, "url", r.URL.Path, "status", wrapped.statusCode)
	})
}
