#!/usr/bin/env bash
set -euo pipefail

# Cross-compile Go binaries for all supported platforms.
# Binary names follow the npm wrapper convention used by npm/index.js:
#   habibiahmada-<os>-<arch>[.exe]  with os in {linux, darwin, win}
# Output: dist/<binary-name>[.exe]

MODULE="github.com/habibiahmada/habibiahmada-terminal"
OUTPUT_DIR="dist"
BINARY_NAME="habibiahmada"

# Build targets: <goos>:<goarch>:<output-suffix>
TARGETS=(
  "linux:amd64:linux-x64"
  "linux:arm64:linux-arm64"
  "darwin:amd64:darwin-x64"
  "darwin:arm64:darwin-arm64"
  "windows:amd64:win-x64"
)

echo "Building $BINARY_NAME for ${#TARGETS[@]} targets..."

# Clean previous build.
rm -rf "$OUTPUT_DIR"
mkdir -p "$OUTPUT_DIR"

for target in "${TARGETS[@]}"; do
  IFS=':' read -r goos goarch suffix <<< "$target"

  ext=""
  if [ "$goos" = "windows" ]; then
    ext=".exe"
  fi

  output="$OUTPUT_DIR/$BINARY_NAME-$suffix$ext"

  echo "  Building $suffix ($goos/$goarch) -> $output"

  GOOS="$goos" GOARCH="$goarch" CGO_ENABLED=0 \
    go build -trimpath -ldflags="-s -w" -o "$output" "./cmd/portfolio"
done

echo ""
echo "Build complete. Binaries in $OUTPUT_DIR/:"
ls -lh "$OUTPUT_DIR/"

echo ""
echo "Generating npm/checksums.json"
CHECKSUM_FILE="npm/checksums.json"
{
  echo "{"
  first=true
  for f in "$OUTPUT_DIR"/*; do
    name=$(basename "$f")
    hash=$(sha256sum "$f" | awk '{print $1}')
    if [ "$first" = true ]; then
      first=false
    else
      echo ","
    fi
    printf '  "%s": "%s"' "$name" "$hash"
  done
  echo ""
  echo "}"
} > "$CHECKSUM_FILE"
echo "Checksums written to $CHECKSUM_FILE"