package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/cloud-resume/api/internal/data"
)

// HealthResponse represents system health and runtime diagnostic metrics.
type HealthResponse struct {
	Status      string            `json:"status"`
	Version     string            `json:"version"`
	Uptime      string            `json:"uptime"`
	ServerTime  string            `json:"server_time"`
	GoVersion   string            `json:"go_version"`
	Diagnostics SystemDiagnostics `json:"diagnostics"`
}

// SystemDiagnostics captures memory and concurrency stats.
type SystemDiagnostics struct {
	Goroutines int    `json:"goroutines"`
	Allocated  string `json:"allocated_memory"`
	SystemMem  string `json:"system_memory"`
	NumGC      uint32 `json:"num_gc"`
}

// HealthHandler returns real-time server health and runtime memory statistics.
func HealthHandler(startTime time.Time, version string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		resp := HealthResponse{
			Status:     "UP",
			Version:    version,
			Uptime:     time.Since(startTime).Round(time.Second).String(),
			ServerTime: time.Now().UTC().Format(time.RFC3339),
			GoVersion:  runtime.Version(),
			Diagnostics: SystemDiagnostics{
				Goroutines: runtime.NumGoroutine(),
				Allocated:  formatBytes(mem.Alloc),
				SystemMem:  formatBytes(mem.Sys),
				NumGC:      mem.NumGC,
			},
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(resp)
	}
}

// CVHandler returns the complete structured resume data as a JSON payload.
func CVHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(data.DefaultResume)
	}
}

func formatBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	suffixes := []string{"KB", "MB", "GB", "TB"}
	return fmt.Sprintf("%.2f %s", float64(b)/float64(div), suffixes[exp])
}
