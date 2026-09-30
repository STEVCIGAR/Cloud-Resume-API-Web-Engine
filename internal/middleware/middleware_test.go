package middleware_test

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloud-resume/api/internal/middleware"
)

func TestSecurityHeadersMiddleware(t *testing.T) {
	handler := middleware.SecurityHeadersMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	headers := map[string]string{
		"X-Frame-Options":        "DENY",
		"X-Content-Type-Options": "nosniff",
		"Access-Control-Allow-Origin": "*",
	}

	for k, expected := range headers {
		got := rec.Header().Get(k)
		if got != expected {
			t.Errorf("header %s: expected %q, got %q", k, expected, got)
		}
	}

	if csp := rec.Header().Get("Content-Security-Policy"); csp == "" {
		t.Error("expected Content-Security-Policy to be set")
	}
}

func TestLoggerMiddleware(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))

	handler := middleware.LoggerMiddleware(logger)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test response"))
	}))

	req := httptest.NewRequest(http.MethodGet, "/test-path", nil)
	req.Header.Set("X-Request-ID", "custom-trace-12345")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	if reqID := rec.Header().Get(middleware.HeaderReqID); reqID != "custom-trace-12345" {
		t.Errorf("expected X-Request-ID %q, got %q", "custom-trace-12345", reqID)
	}

	if respTime := rec.Header().Get(middleware.HeaderRespTime); respTime == "" {
		t.Error("expected X-Response-Time to be present")
	}

	logOutput := buf.String()
	if !bytes.Contains(buf.Bytes(), []byte("custom-trace-12345")) {
		t.Errorf("log output missing request_id: %s", logOutput)
	}
}
