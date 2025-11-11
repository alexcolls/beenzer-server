.PHONY: help build run dev test clean swagger docker install

# Default target
help:
	@echo "Beenzer Server v3 - Makefile Commands"
	@echo ""
	@echo "  make install    - Install Go dependencies"
	@echo "  make build      - Build the server binary"
	@echo "  make run        - Build and run the server"
	@echo "  make dev        - Run in development mode (with hot reload if air is installed)"
	@echo "  make test       - Run tests"
	@echo "  make test-cover - Run tests with coverage"
	@echo "  make clean      - Remove build artifacts"
	@echo "  make swagger    - Generate Swagger documentation"
	@echo "  make fmt        - Format code"
	@echo "  make lint       - Run linter (requires golangci-lint)"
	@echo "  make tidy       - Tidy and verify go.mod"
	@echo "  make docker     - Build Docker image"
	@echo ""

# Install dependencies
install:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy

# Build binary
build:
	@echo "🔨 Building server..."
	go build -o bin/beenzer-server cmd/server/main.go
	@echo "✅ Build complete: bin/beenzer-server"

# Run the server
run: build
	@echo "🚀 Starting server..."
	./bin/beenzer-server

# Development mode
dev:
	@echo "🔥 Starting development server..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		@echo "⚠️  'air' not found. Installing..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Run tests with coverage
test-cover:
	@echo "🧪 Running tests with coverage..."
	go test -v -cover -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report: coverage.html"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -rf docs/swagger.json docs/swagger.yaml
	@echo "✅ Cleaned"

# Generate Swagger docs
swagger:
	@echo "📚 Generating Swagger documentation..."
	@if command -v swag > /dev/null; then \
		swag init -g cmd/server/main.go -o docs; \
		echo "✅ Swagger docs generated in docs/"; \
	else \
		echo "⚠️  'swag' not found. Installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init -g cmd/server/main.go -o docs; \
		echo "✅ Swagger docs generated in docs/"; \
	fi

# Format code
fmt:
	@echo "✨ Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted"

# Run linter
lint:
	@echo "🔍 Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "⚠️  golangci-lint not found. Please install it."; \
		echo "    https://golangci-lint.run/usage/install/"; \
	fi

# Tidy go.mod
tidy:
	@echo "🧹 Tidying go.mod..."
	go mod tidy
	go mod verify
	@echo "✅ go.mod tidied"

# Build Docker image
docker:
	@echo "🐳 Building Docker image..."
	docker build -t beenzer-server:v3 .
	@echo "✅ Docker image built: beenzer-server:v3"

# Run with race detector
race:
	@echo "🏃 Running with race detector..."
	go run -race cmd/server/main.go

# Quick development workflow
quick: fmt tidy test build
	@echo "✅ Quick workflow complete!"
