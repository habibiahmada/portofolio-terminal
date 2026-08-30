#!/usr/bin/env bash
set -euo pipefail

# Build the SSH server binary for EC2 deployment.
# Output: habibiahmada-ssh

OUTPUT="habibiahmada-ssh"

echo "Building SSH server binary..."

CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o "$OUTPUT" "./cmd/ssh"

echo "Build complete: $OUTPUT"
