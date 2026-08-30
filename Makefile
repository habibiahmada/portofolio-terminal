.PHONY: build build-ssh clean dev test vet lint fmt

# Default target: build all binaries.
all: build

# Cross-compile for all platforms.
build:
	@bash scripts/build.sh

# Build SSH server binary.
build-ssh:
	@bash scripts/build-ssh.sh

# Run the TUI locally for development.
dev:
	go run ./cmd/portfolio

# Run the SSH server locally for development.
dev-ssh:
	go run ./cmd/ssh

# Run all tests.
test:
	go test ./...

# Run go vet.
vet:
	go vet ./...

# Format all Go files.
fmt:
	gofmt -w -s .

# Clean build artifacts.
clean:
	rm -rf dist/
	rm -f habibiahmada-ssh

# Verify code quality.
lint: vet fmt test
