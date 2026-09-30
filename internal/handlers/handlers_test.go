package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/cloud-resume/api/internal/data"
	"github.com/cloud-resume/api/internal/handlers"
)

func TestHealthHandler(t *testing.T) {
	startTime := time.Now().Add(-10 * time.Minute)
	handler := handlers.HealthHandler(startTime, "1.0.0")

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp handlers.HealthResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode JSON response: %v", err)
	}

	if resp.Status != "UP" {
		t.Errorf("expected status 'UP', got %q", resp.Status)
	}
	if resp.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got %q", resp.Version)
	}
}

func TestCVHandler(t *testing.T) {
	handler := handlers.CVHandler()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/cv", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resume data.Resume
	if err := json.Unmarshal(rec.Body.Bytes(), &resume); err != nil {
		t.Fatalf("failed to decode JSON resume: %v", err)
	}

	if resume.Profile.Name == "" {
		t.Error("expected non-empty profile name")
	}
	if len(resume.Skills) == 0 {
		t.Error("expected skills list to be non-empty")
	}
}

func TestSkillsComponentHandler(t *testing.T) {
	handler := handlers.SkillsComponentHandler()

	req := httptest.NewRequest(http.MethodGet, "/components/skills?cat=backend", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "Golang (Go)") {
		t.Errorf("expected Golang skill in backend category, got %s", body)
	}
}

func TestHomeHandler(t *testing.T) {
	handler := handlers.HomeHandler()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "Cloud Resume Engine") && !strings.Contains(body, "Alexander Vance") {
		t.Errorf("home page missing key titles, got %s", body)
	}
}
