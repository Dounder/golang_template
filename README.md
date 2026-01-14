# Go REST API Template

A production-ready Go REST API template with MongoDB, structured logging, and comprehensive tooling.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)
- [Development Tools](#development-tools)
- [Configuration](#configuration)
- [Docker Setup](#docker-setup)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- **Gin Web Framework** - High-performance HTTP web framework
- **MongoDB Integration** - Document database with connection pooling
- **Structured Logging** - slog-based logging with proper context
- **Environment Configuration** - Flexible config management with validation
- **Health Checks** - Database and application health endpoints
- **Docker Support** - Containerized MongoDB setup
- **Code Quality** - golangci-lint with comprehensive rules
- **Pre-commit Hooks** - lefthook for automated code quality checks
- **Hot Reload** - Development with air for live reloading

## Prerequisites

- Go 1.25.5 or later
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/Dounder/golang_template.git
   cd golang_template
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up environment variables**

   ```bash
   cp .env.example .env  # Create from example if available
   ```

   Configure your `.env` file:

   ```env
   # Server Configuration
   SERVER_PORT=3000
   SERVER_STAGE=dev

   # Database Configuration
   MONGO_USER=admin
   MONGO_PASSWORD=password
   MONGODB_URI=mongodb://admin:password@localhost:27017
   MONGO_DATABASE=template_db

   # Application Configuration
   LOG_LEVEL=info
   GIN_MODE=debug
   ```

4. **Start MongoDB**
   ```bash
   make docker-run
   # or
   docker-compose up -d
   ```

## Running the Application

### Development with Hot Reload

```bash
make dev
# or
go install github.com/air-verse/air@latest
air
```

### Production Build

```bash
make build
make run
# or
go build -o bin/glasdou_template ./main.go
./bin/glasdou_template
```

### Using Docker

```bash
# Build and run with Docker
make docker-build
make docker-run
```

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-cover
```

## API Endpoints

### Health Checks

- `GET /api/v1/health` - Application health status
- `GET /api/v1/health/db` - Database connection health

### Response Format

```json
{
  "status": "ok",
  "timestamp": "2024-01-14T10:00:00Z",
  "version": "1.0.0"
}
```

## Project Structure

```
├── main.go                 # Application entry point
├── config/
│   ├── database.go         # MongoDB connection setup
│   └── envs.go             # Configuration management
├── modules/
│   ├── routes.go           # Route registration
│   ├── common/
│   │   └── types/          # Shared types and utilities
│   └── health/             # Health check module
│       ├── health.controller.go
│       ├── health.routes.go
│       └── health.service.go
├── compose.yml             # Docker Compose configuration
├── mongo-init.js          # MongoDB initialization script
├── makefile               # Build and development commands
├── lefthook.yml           # Pre-commit hooks configuration
├── .golangci.yml          # Linting configuration
└── go.mod                 # Go module dependencies
```

## Development Tools

### Code Quality

```bash
# Format code
make fmt

# Run linter
golangci-lint run

# Fix linting issues automatically
golangci-lint run --fix
```

### Pre-commit Hooks

The project uses [lefthook](https://github.com/evilmartians/lefthook) for pre-commit hooks:

- **Format Go files** - Runs `go fmt`
- **Vet Go code** - Runs `go vet`
- **Run golangci-lint** - Comprehensive linting with auto-fix

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Build and run the application
make dev           # Run with hot reload
make test          # Run tests
make test-cover    # Run tests with coverage
make clean         # Remove build artifacts
make fmt           # Format code
make docker-build  # Build Docker image
make docker-run    # Run with Docker Compose
make docker-stop   # Stop Docker containers
make docker-logs   # Show Docker logs
```

## Configuration

The application uses environment-based configuration with validation:

### Server Configuration

- `SERVER_PORT` - Server port (default: 3000)
- `SERVER_STAGE` - Environment stage: dev, prod, staging

### Database Configuration

- `MONGODB_URI` - MongoDB connection URI
- `MONGO_DATABASE` - Database name
- `MONGO_USER` - Database username
- `MONGO_PASSWORD` - Database password

### Application Configuration

- `LOG_LEVEL` - Logging level: debug, info, warn, error
- `GIN_MODE` - Gin framework mode: debug, release, test

## Docker Setup

### MongoDB Container

The project includes a Docker Compose setup for MongoDB:

```yaml
services:
  go_template_mongo:
    image: mongo:7.0
    restart: always
    environment:
      MONGO_INITDB_ROOT_USERNAME: ${MONGO_USER}
      MONGO_INITDB_ROOT_PASSWORD: ${MONGO_PASSWORD}
    ports:
      - '27017:27017'
```

### Database Initialization

The `mongo-init.js` script runs on container startup to set up initial database state.

## Dependencies

### Core Dependencies

- [gin-gonic/gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [go.mongodb.org/mongo-driver/v2](https://github.com/mongodb/mongo-go-driver) - MongoDB driver
- [github.com/joho/godotenv](https://github.com/joho/godotenv) - Environment loading
- [github.com/go-playground/validator/v10](https://github.com/go-playground/validator) - Validation

### Development Dependencies

- [github.com/air-verse/air](https://github.com/air-verse/air) - Live reloading
- [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) - Linting
- [github.com/evilmartians/lefthook](https://github.com/evilmartians/lefthook) - Git hooks

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and conventions
- Write tests for new features
- Ensure all linting checks pass
- Update documentation as needed

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/) for the excellent HTTP framework
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) for database connectivity
- [golangci-lint](https://github.com/golangci/golangci-lint) for code quality tools
- [lefthook](https://github.com/evilmartians/lefthook) for git hooks management

---

Made with ❤️ using Go</content>
<parameter name="filePath">/Users/daramirez/dev/go/template/README.md
