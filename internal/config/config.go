package config

import (
	"os"
	"strconv"
)

// Config holds all application configuration parameters.
type Config struct {
	Port        int
	Environment string
	Version     string
	AppName     string
}

// Load loads configuration from environment variables with sensible defaults.
func Load() *Config {
	port := 8080
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "production"
	}

	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "1.0.0"
	}

	return &Config{
		Port:        port,
		Environment: env,
		Version:     version,
		AppName:     "Cloud Resume Engine",
	}
}
