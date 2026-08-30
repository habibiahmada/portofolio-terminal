// Package assets bundles static ASCII art for the TUI using go:embed.
// Keep art files under internal/assets/art/ so they end up inside the binary.
package assets

import (
	"embed"
	"strings"
)

//go:embed art/*
var ArtFS embed.FS

// Art returns the trimmed content of an art file under internal/assets/art/.
// Returns an empty string if the file is missing.
func Art(name string) string {
	b, err := ArtFS.ReadFile("art/" + name)
	if err != nil {
		return ""
	}
	return strings.TrimRight(string(b), "\n")
}
