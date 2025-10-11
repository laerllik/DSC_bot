# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
GOMOD=$(GOCMD) mod
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

# Project structure
CMD_PATH=./cmd
MAIN_FILE=$(CMD_PATH)/main.go

# Default target
all: build

# Build the project
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

# Run the application
run:
	$(GOCMD) run $(MAIN_FILE)


# Test
test:
	$(GOTEST) -v ./...

test-coverage:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Linting and code quality
lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

# Clean
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*
	rm -f *.exe
	rm -f coverage.*
	rm -f coverage.html

# Dependency management
deps:
	$(GOMOD) download
	$(GOMOD) tidy

deps-update:
	$(GOMOD) download -x
	$(GOMOD) tidy

# Install tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Development setup
setup: install-tools deps
	@echo "Development environment ready!"

fmt:
	$(GOFMT) ./...

# Help
help:
	@echo "  deps         - Download dependencies"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  lint         - Run linter"
	@echo "  fmt          - Format code"
	@echo "  clean        - Clean build artifacts"
	@echo "  setup        - Setup development environment"
	@echo "  help         - Show this help"
