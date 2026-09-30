# ==============================================================================
# Stage 1: Builder
# ==============================================================================
FROM golang:1.24-alpine AS builder

WORKDIR /src

# Install build essentials and security certs
RUN apk add --no-cache git ca-certificates tzdata

# Install Templ CLI
RUN go install github.com/a-h/templ/cmd/templ@latest

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate Templ components
RUN templ generate

# Compile fully static, stripped Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -trimpath \
    -ldflags="-s -w -extldflags '-static'" \
    -o /app/server cmd/server/main.go

# ==============================================================================
# Stage 2: Minimal Distroless Runtime (<20MB)
# ==============================================================================
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /

# Copy static binary from builder
COPY --from=builder --chown=nonroot:nonroot /app/server /server

# Expose HTTP port
EXPOSE 8080

# Run with non-root security context
USER nonroot:nonroot

# Environment variables
ENV PORT=8080 \
    APP_ENV=production

ENTRYPOINT ["/server"]
