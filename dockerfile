# Development Dockerfile with hot-reload support
FROM golang:1.25.5-alpine

WORKDIR /usr/src/app

# Install development dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Install Air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Copy go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code (for initial build)
COPY . .

# Set environment variables for development
ENV GIN_MODE=debug \
  PORT=3000 \
  TZ=America/Guatemala

# Expose port
EXPOSE 3000

# Use Air for hot-reloading in development
CMD ["air", "-c", ".air.toml"]
