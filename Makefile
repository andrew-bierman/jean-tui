# jean-tui Makefile

.PHONY: build test coverage clean install help

# Default target
all: build

# Build the default binary
build:
	go build -o jean

# Build with custom branding (example: make build-custom NAME=myapp PREFIX=myapp-)
build-custom:
	@if [ -z "$(NAME)" ]; then echo "Usage: make build-custom NAME=myapp"; exit 1; fi
	go build -ldflags "\
		-X github.com/coollabsio/jean-tui/internal/branding.CLIName=$(NAME) \
		-X github.com/coollabsio/jean-tui/internal/branding.SessionPrefix=$(or $(PREFIX),$(NAME)-) \
		-X github.com/coollabsio/jean-tui/internal/branding.ConfigDirName=$(or $(CONFIG),$(NAME)) \
		-X github.com/coollabsio/jean-tui/internal/branding.EnvVarPrefix=$(or $(ENVPREFIX),$(shell echo $(NAME) | tr '[:lower:]-' '[:upper:]_'))" \
		-o $(NAME)

# Build with blank terminal (no agent, worktree-only mode)
build-blank:
	go build -ldflags "\
		-X github.com/coollabsio/jean-tui/internal/branding.AgentCommand= \
		-X github.com/coollabsio/jean-tui/internal/branding.AgentWindowName=shell" \
		-o jean

# Run tests
test:
	go test ./...

# Run tests with verbose output
test-v:
	go test ./... -v

# Run tests with coverage
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out | tail -10

# Generate HTML coverage report
coverage-html: coverage
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Clean build artifacts
clean:
	rm -f jean coverage.out coverage.html
	go clean

# Install to /usr/local/bin
install: build
	sudo cp jean /usr/local/bin/

# Run locally
run:
	go run main.go

# Initialize dependencies
deps:
	go mod tidy
	go mod download

# Show help
help:
	@echo "jean-tui Makefile targets:"
	@echo ""
	@echo "  build          Build default binary (jean)"
	@echo "  build-custom   Build with custom branding (NAME=myapp)"
	@echo "  build-blank    Build with blank terminal (no Claude)"
	@echo "  test           Run tests"
	@echo "  test-v         Run tests with verbose output"
	@echo "  coverage       Run tests with coverage report"
	@echo "  coverage-html  Generate HTML coverage report"
	@echo "  clean          Clean build artifacts"
	@echo "  install        Install to /usr/local/bin"
	@echo "  run            Run locally"
	@echo "  deps           Update dependencies"
	@echo ""
	@echo "Custom build examples:"
	@echo "  make build-custom NAME=ralph-tui"
	@echo "  make build-custom NAME=opencode PREFIX=oc- CONFIG=opencode ENVPREFIX=OPENCODE"
