package middleware

import (
	"log"
	"net/http"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{w, http.StatusOK}
}

func (lrw *logResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		lrw := NewLogResponseWriter(w)
		next.ServeHTTP(lrw, r)
		log.Printf("Finished %s %s with StatusCode: %d", r.Method, r.URL.Path, lrw.statusCode)
	})
}
