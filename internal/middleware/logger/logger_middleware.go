package logger

import (
	"net/http"
	"time"
)

// LoggerMiddleware mencatat method, path, dan durasi setiap request
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Milliseconds()

		logMsg := r.Method + " " + r.URL.Path
		Info("%s completed in %dms", logMsg, duration)
	})
}
