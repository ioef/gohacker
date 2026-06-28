# Backend Dockerfile - Debian-based for Go 1.26+
FROM golang:1.26-bookworm AS builder

WORKDIR /app

# Install build dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags="-w -s -X main.version=1.0.0" \
    -o gohacker main.go

# Final stage - Use golang image for runtime (needed for code execution)
FROM golang:1.26-bookworm

# Install runtime dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    wget \
    curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/gohacker .

# Create data directory
RUN mkdir -p /app/data

# Create non-root user for security
RUN useradd -m -u 1000 gohacker && \
    chown -R gohacker:gohacker /app

USER gohacker

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/api/health || exit 1

# Run
CMD ["./gohacker"]
