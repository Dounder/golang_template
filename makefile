# Go API Makefile
# Simplified commands for running the project

.PHONY: help build run dev test clean fmt lint mod docker-build docker-run docker-stop docker-logs

# Variables
BINARY_NAME=glasdou_template
MAIN_PATH=./main.go
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
REGISTRY=drglasdou/go_template

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'

# Development Commands
build: ## Build the application binary
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

run: build ## Build and run the application
	@echo "Running..."
	@./bin/$(BINARY_NAME)

dev: ## Run the application with hot reload (requires air: go install github.com/air-verse/air@latest)
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "air is not installed. Install it with: go install github.com/air-verse/air@latest"; \
		echo "Running without hot reload..."; \
		go run $(MAIN_PATH); \
	fi

test: ## Run all tests
	@echo "Running tests..."
	@go test -v ./...

test-cover: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

clean: ## Remove build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run linter (requires golangci-lint)
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed. Install it from: https://golangci-lint.run/usage/install/"; \
	fi

mod: ## Tidy go modules
	@echo "Tidying modules..."
	@go mod tidy
	@go mod verify

# Docker Commands
docker-build: ## Build Docker production image (version from git tags)
	@echo "Building Docker image version: $(VERSION)"
	@VERSION=$(VERSION) docker compose -f compose.build.yml build

docker-run: ## Run the application with Docker Compose
	@docker compose up -d

docker-stop: ## Stop all Docker containers
	@docker compose down

docker-logs: ## Show Docker container logs
	@docker compose logs -f

docker-clean: ## Stop containers and remove volumes
	@docker compose down -v

docker-push: ## Push Docker image to registry
	@echo "Pushing version: $(VERSION)"
	@docker push $(REGISTRY):$(VERSION)
	@docker push $(REGISTRY):latest

version: ## Show current version from git tags
	@echo $(VERSION)
