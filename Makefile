# Makefile for next-starter CLI

BINARY_NAME=next-starter
MAIN_FILE=main.go

.PHONY: build clean install test help

# Default target
all: build

# Build the binary
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: ./$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html
	go clean
	@echo "✅ Clean complete"

# Install to GOPATH/bin
install:
	@echo "📦 Installing $(BINARY_NAME)..."
	go install
	@echo "✅ Installed to GOPATH/bin"

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./... -v

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test ./... -v -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

# Show help
help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  clean    - Clean build artifacts"
	@echo "  install  - Install to GOPATH/bin"
	@echo "  test     - Run tests"
	@echo "  help     - Show this help"

# Example usage
example:
	@echo "📚 Example usage:"
	@echo "  make build"
	@echo "  ./$(BINARY_NAME) my-awesome-app"