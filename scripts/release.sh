#!/usr/bin/env bash
set -euo pipefail

# Build platform binaries and stage them into the npm package for publish.
#
# Usage:
#   scripts/release.sh [version]
#
#   version       npm version to write into npm/package.json (default: derive
#                 from git describe --tags, falling back to the current value).

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
NPM_DIR="$ROOT_DIR/npm"
DIST_DIR="$ROOT_DIR/dist"

VERSION="${1:-}"
if [ -z "$VERSION" ]; then
  VERSION="$(git -C "$ROOT_DIR" describe --tags --abbrev=0 2>/dev/null || true)"
fi

cd "$ROOT_DIR"

echo "==> Building binaries for all platforms"
bash scripts/build.sh

echo "==> Staging binaries into npm/bin/"
mkdir -p "$NPM_DIR/bin"
for binary in "$DIST_DIR"/habibiahmada-*; do
  [ -f "$binary" ] || continue
  cp "$binary" "$NPM_DIR/bin/"
  chmod +x "$NPM_DIR/bin/$(basename "$binary")"
done

echo "==> npm package contents:"
ls -lh "$NPM_DIR/bin/"

if [ -n "$VERSION" ]; then
  echo "==> Setting npm version to $VERSION"
  node -e "const p=require('$NPM_DIR/package.json'); p.version='$VERSION'; require('fs').writeFileSync('$NPM_DIR/package.json', JSON.stringify(p, null, 2)+'\n');"
fi

echo ""
echo "Release ready. Publish with: cd npm && npm publish --access public"