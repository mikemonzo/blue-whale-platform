# Build stage
FROM golang:1.24-alpine AS builder

# Install build dependencies
RUN apk add --no-cache \
    git \
    make \
    gcc \
    g++

# Set working directory
WORKDIR /build

# Copy go.mod and go.sum files
COPY go.work* .
COPY services/idp/go.* ./services/idp/
COPY shared/common/go.* ./shared/common/
COPY shared/pkg/go.* ./shared/pkg/

# Download dependencies
RUN cd services/idp && go mod download

# Copy source code
COPY . .

# Build the application
RUN cd services/idp && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o idp-service ./cmd/main.go

# Final stage
FROM alpine:3.19

# Install runtime dependencies
RUN apk add --no-cache \
    ca-certificates \
    tzdata \
    postgresql-client

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/services/idp/idp-service .
COPY services/idp/config/config.yaml ./config/

# Create non-root user
RUN adduser -D appuser && \
    chown -R appuser:appuser /app

USER appuser

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./idp-service"]