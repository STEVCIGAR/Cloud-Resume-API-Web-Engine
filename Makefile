# ==============================================================================
# Cloud Resume API & Web - Automation Makefile
# ==============================================================================

APP_NAME := cloud-resume-api
VERSION := 1.0.0
PORT := 8080
BIN_DIR := bin
DOCKER_IMAGE := $(APP_NAME):latest

.PHONY: all build run templ-generate test clean docker-build docker-run lint help

all: templ-generate build

## templ-generate: Compiles .templ files into type-safe Go components
templ-generate:
	@echo "==> Generating Templ Go components..."
	templ generate

## build: Compiles the Go server binary
build: templ-generate
	@echo "==> Building binary..."
	go build -ldflags="-s -w" -o $(BIN_DIR)/server cmd/server/main.go

## run: Runs the application locally
run: templ-generate
	@echo "==> Starting local server on http://localhost:$(PORT)..."
	go run cmd/server/main.go

## test: Runs all unit and integration tests
test: templ-generate
	@echo "==> Running tests..."
	go test -v -race ./...

## docker-build: Builds the minimal multi-stage Distroless Docker image
docker-build:
	@echo "==> Building Docker image $(DOCKER_IMAGE)..."
	docker build -t $(DOCKER_IMAGE) .

## docker-run: Runs the containerized application
docker-run:
	@echo "==> Running Docker container on http://localhost:$(PORT)..."
	docker run --rm -p $(PORT):8080 $(DOCKER_IMAGE)

## clean: Removes build artifacts
clean:
	@echo "==> Cleaning build artifacts..."
	rm -rf $(BIN_DIR)

## help: Shows this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed -e 's/## //'
