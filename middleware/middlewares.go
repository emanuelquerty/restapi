package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type LoggingResponseWriter struct {
	w          http.ResponseWriter
	StatusCode int
	BytesCount int
}

func (lrw *LoggingResponseWriter) Header() http.Header {
	return lrw.w.Header()
}

func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	bytesWritten, err := lrw.w.Write(b)
	lrw.BytesCount += bytesWritten
	return bytesWritten, err
}

func (lrw *LoggingResponseWriter) WriteHeader(statusCode int) {
	lrw.StatusCode = statusCode
	lrw.w.WriteHeader(statusCode)
}

func AccessLogger(logger *slog.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &LoggingResponseWriter{w: w}

		start := time.Now()

		handler.ServeHTTP(lrw, r)

		duration := time.Since(start).Milliseconds()

		responseAtrr := slog.Group("response",
			slog.Int("status", lrw.StatusCode),
			slog.Int("size", lrw.BytesCount),
			slog.String("duration", fmt.Sprintf("%d ms", duration)),
		)

		endpoint := fmt.Sprintf("%s %s %s", r.Method, r.URL.Path, r.Proto)

		totalRequestCounter.WithLabelValues(endpoint).Inc()
		requestDurationObserver.WithLabelValues(endpoint).Observe(float64(duration))

		logger.Info(endpoint, responseAtrr)
	})
}

func SecurityHeaders(logger *slog.Logger, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(w, r)
	})
}
