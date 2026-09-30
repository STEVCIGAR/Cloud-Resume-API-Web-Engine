package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloud-resume/api/internal/assets"
	"github.com/cloud-resume/api/internal/config"
	"github.com/cloud-resume/api/internal/handlers"
	"github.com/cloud-resume/api/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

func main() {
	startTime := time.Now()

	// 1. Initialize structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// 2. Load Configuration
	cfg := config.Load()
	logger.Info("Starting Cloud Resume Engine",
		slog.String("app", cfg.AppName),
		slog.String("version", cfg.Version),
		slog.String("env", cfg.Environment),
		slog.Int("port", cfg.Port),
	)

	// 3. Setup Router
	r := chi.NewRouter()

	// Core Middleware Pipeline
	r.Use(chimw.Recoverer)
	r.Use(middleware.SecurityHeadersMiddleware)
	r.Use(middleware.LoggerMiddleware(logger))

	// Static Assets (Embedded in binary)
	staticFS, err := assets.GetStaticFS()
	if err != nil {
		logger.Error("Failed to load embedded static filesystem", slog.String("error", err.Error()))
		os.Exit(1)
	}
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// Web UI & Server-Driven HTMX Endpoints
	r.Get("/", handlers.HomeHandler())
	r.Get("/components/skills", handlers.SkillsComponentHandler())

	// REST API Endpoints
	r.Route("/api/v1", func(api chi.Router) {
		api.Get("/health", handlers.HealthHandler(startTime, cfg.Version))
		api.Get("/cv", handlers.CVHandler())
	})

	// 4. HTTP Server with timeouts
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 5. Run server in background Goroutine
	serverErr := make(chan error, 1)
	go func() {
		logger.Info("Server listening for HTTP requests", slog.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// 6. Graceful Shutdown listener
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		logger.Error("Server failed to start", slog.String("error", err.Error()))
		os.Exit(1)
	case sig := <-shutdown:
		logger.Info("Shutdown signal received", slog.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Forced shutdown error", slog.String("error", err.Error()))
			_ = server.Close()
		}
		logger.Info("Server gracefully stopped. Clean exit.")
	}
}
