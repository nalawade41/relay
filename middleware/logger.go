package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// RequestLogger wraps a handler and logs every request after it finishes.
func RequestLogger(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info(
			"request",
			"mehtod", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start).String(),
		)
	})
}
