package middleware

import (
	"log/slog"
	"net/http"
	"os"
)

func LoggingMiddleware(next http.Handler) http.HandlerFunc {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received", "method", r.Method, "url", r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
