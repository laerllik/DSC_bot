CMD_PATH=./cmd

DSC_BOT=dsc-bot
DSC_BOT_PATH=$(CMD_PATH)/$(DSC_BOT)/main.go

build-dsc-bot:
	@echo "Building $(DSC_BOT)..."
	go build -o bin/$(DSC_BOT) $(DSC_BOT_PATH)


lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

clean:
	$(go clean)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*
	rm -f *.exe
	rm -f coverage.*
	rm -f coverage.html

deps:
	go mod download
	go mod tidy

install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

setup: install-tools deps
	@echo "Development environment ready!"

fmt:
	go fmt ./...

help:
	@echo "  setup             - Prepare workspace: install tools, deps"
	@echo "  build-dsc-bot     - Build dsc-bot"
	@echo "  build-collector   - Build collector"
	@echo "  test              - Run tests"
	@echo "  lint              - Run linter"
	@echo "  fmt               - Format code"
	@echo "  clean             - Clean build artifacts"
	@echo "  help              - Show this help"
