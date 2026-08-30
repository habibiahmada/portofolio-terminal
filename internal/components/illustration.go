package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/assets"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Variant describes which illustration fidelity fits the current terminal width.
type Variant int

const (
	VariantWide        Variant = iota // >= 100 cols — full signature art
	VariantComfortable                // 80–99 — compact signature
	VariantNarrow                     // 60–79 — inline mini signature
	VariantHidden                     // < 60 — no art, content first
)

// IllustrationVariant picks an illustration variant for a given width.
func IllustrationVariant(width int) Variant {
	switch {
	case width >= 100:
		return VariantWide
	case width >= 80:
		return VariantComfortable
	case width >= 60:
		return VariantNarrow
	default:
		return VariantHidden
	}
}

// signatureArt returns the raw art file matching the variant for the width.
func signatureArt(width int) string {
	switch IllustrationVariant(width) {
	case VariantWide:
		return assets.Art("signature_wide.txt")
	case VariantComfortable:
		return assets.Art("signature_compact.txt")
	case VariantNarrow:
		return assets.Art("signature_mini.txt")
	default:
		return ""
	}
}

// renderArt styles each art line according to its content.
func renderArt(art string) string {
	lines := strings.Split(art, "\n")
	styled := make([]string, 0, len(lines))
	for _, line := range lines {
		styled = append(styled, styleArtLine(line))
	}
	return lipgloss.JoinVertical(lipgloss.Left, styled...)
}

// Signature renders the "Habibi Terminal" signature art, styled per line.
// It returns an empty string when the terminal is too narrow for art.
func Signature(width int) string {
	return renderArt(signatureArt(width))
}

// SignatureBlink renders the signature with a blinking cursor. When showCursor
// is false the ">_" prompt is rendered without its underscore.
func SignatureBlink(width int, showCursor bool) string {
	art := signatureArt(width)
	if art == "" {
		return ""
	}
	if !showCursor {
		art = strings.ReplaceAll(art, ">_", "> ")
	}
	return renderArt(art)
}

// AboutTerminal renders the small terminal illustration for the About screen.
func AboutTerminal() string {
	return renderArt(assets.Art("about_terminal.txt"))
}

// styleArtLine applies the palette to a single art line based on its content.
func styleArtLine(line string) string {
	switch {
	case strings.Contains(line, ">_"):
		return styles.PromptStyle.Render(line)
	case strings.Contains(line, "HABIBI"):
		return styles.ArtBrandPrimary.Render(line)
	case strings.Contains(line, "TERMINAL"):
		return styles.ArtBrandSecondary.Render(line)
	case strings.ContainsAny(line, "╭╮╰╯│"):
		return styles.ArtBorderStyle.Render(line)
	default:
		return line
	}
}
