package middleware

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type contextKey string

const (
	RequestIDKey contextKey = "request_id"
	HeaderReqID  string     = "X-Request-ID"
	HeaderRespTime string   = "X-Response-Time"
)

// responseWriter is a wrapper around http.ResponseWriter to capture status code and size.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	bytesWritten int64
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += int64(n)
	return n, err
}

// LoggerMiddleware logs every HTTP request with latency, client IP, Request ID, and status code.
func LoggerMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Extract or generate Request ID
			reqID := r.Header.Get(HeaderReqID)
			if reqID == "" {
				reqID = uuid.New().String()
			}

			// Add Request ID to Response Header and Request Context
			w.Header().Set(HeaderReqID, reqID)
			ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
			r = r.WithContext(ctx)

			// Extract Client IP
			clientIP := getClientIP(r)

			// Wrap ResponseWriter
			wrappedWriter := newResponseWriter(w)

			// Process Request
			next.ServeHTTP(wrappedWriter, r)

			// Calculate Duration
			duration := time.Since(start)
			wrappedWriter.Header().Set(HeaderRespTime, duration.String())

			// Log structured request info
			logLevel := slog.LevelInfo
			if wrappedWriter.statusCode >= 500 {
				logLevel = slog.LevelError
			} else if wrappedWriter.statusCode >= 400 {
				logLevel = slog.LevelWarn
			}

			logger.Log(r.Context(), logLevel, "HTTP request handled",
				slog.String("request_id", reqID),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("query", r.URL.RawQuery),
				slog.Int("status", wrappedWriter.statusCode),
				slog.Duration("latency", duration),
				slog.String("latency_human", duration.String()),
				slog.String("client_ip", clientIP),
				slog.Int64("bytes", wrappedWriter.bytesWritten),
				slog.String("user_agent", r.UserAgent()),
			)
		})
	}
}

// getClientIP resolves the actual client IP considering reverse proxies.
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return strings.TrimSpace(xrip)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	return r.RemoteAddr
}
