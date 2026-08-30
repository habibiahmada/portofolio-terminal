package components

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// FooterHint is a single keyboard hint shown in the footer.
type FooterHint struct {
	Key   string // e.g. "↑↓"
	Label string // e.g. "Navigate"
}

// Footer renders the bottom navigation hints separated by dots.
func Footer(hints []FooterHint) string {
	parts := make([]string, 0, len(hints))
	for _, h := range hints {
		parts = append(parts, h.Key+" "+h.Label)
	}
	return styles.FooterStyle.Render(strings.Join(parts, "  •  "))
}
