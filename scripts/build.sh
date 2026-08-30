#!/usr/bin/env bash
set -euo pipefail

# Cross-compile Go binaries for all supported platforms.
# Output: dist/<os>-<arch>[.exe]

MODULE="github.com/habibiahmada/habibiahmada-terminal"
OUTPUT_DIR="dist"
BINARY_NAME="habibiahmada"

# Build targets: OS-ARCH
TARGETS=(
  "linux-amd64"
  "linux-arm64"
  "darwin-amd64"
  "darwin-arm64"
  "windows-amd64"
)

echo "Building $BINARY_NAME for ${#TARGETS[@]} targets..."

# Clean previous build.
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

for target in "${TARGETS[@]}"; do
  IFS='-' read -r os arch <<< "$target"

  # Map target naming to Go GOOS/GOARCH.
  goos="$os"
  goarch="$arch"
  ext=""

  if [ "$os" = "windows" ]; then
    ext=".exe"
  fi

  output="$OUTPUT_DIR/$BINARY_NAME-$target$ext"

  echo "  Building $target -> $output"

  GOOS="$goos" GOARCH="$goarch" CGO_ENABLED=0 \
    go build -trimpath -ldflags="-s -w" -o "$output" "./cmd/portfolio"
done

echo ""
echo "Build complete. Binaries in $OUTPUT_DIR/:"
ls -lh "$OUTPUT_DIR/"
