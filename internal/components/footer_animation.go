package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// wave is a single-line equalizer cycle for the footer brand animation.
var wave = []string{
	"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁",
}

// FooterBar renders a single footer row: brand on the left, keyboard hints on
// the right. Animation is minimal to keep CPU/RAM usage low.
func FooterBar(frame, width int, hints []FooterHint) string {
	if width < 20 {
		width = 20
	}

	left := footerBrand(frame, width)

	parts := make([]string, 0, len(hints))
	for _, h := range hints {
		parts = append(parts, h.Key+" "+h.Label)
	}
	right := styles.FooterStyle.Render(strings.Join(parts, "  ·  "))

	leftW := lipgloss.Width(left)
	rightW := lipgloss.Width(right)
	gap := width - leftW - rightW
	if gap < 2 {
		// Stack on tiny widths: brand then hints.
		return lipgloss.JoinVertical(lipgloss.Left, left, right)
	}

	spacer := strings.Repeat(" ", gap)
	row := left + spacer + right
	return styles.FooterStyle.Width(width).Render(row)
}

// footerBrand renders the left-side brand mark with optional mini animation.
func footerBrand(frame, width int) string {
	mark := styles.FooterWordmark.Render("habibiahmada") + styles.HeaderDot.Render(".")

	if width < 40 {
		return styles.FooterArtStyle.Render(mark)
	}

	anim := ""
	// Equalizer only on wide terminals — skip on smaller to save render cost.
	if width >= 90 {
		anim = " " + miniEqualizer(frame)
	}

	return styles.FooterArtStyle.Render(
		styles.PromptStyle.Render(">_") + " " + mark + anim,
	)
}

func miniEqualizer(frame int) string {
	n := 5
	if n > len(wave) {
		n = len(wave)
	}
	start := frame % len(wave)
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		parts = append(parts, styles.FooterBarStyle.Render(wave[(start+i)%len(wave)]))
	}
	return strings.Join(parts, "")
}

// FooterArtline is deprecated — use FooterBar. Kept for tests during migration.
func FooterArtline(frame int, width int, website string) string {
	_ = website
	return footerBrand(frame, width)
}

// FooterWithHint is deprecated — use FooterBar.
func FooterWithHint(artline string, hints []FooterHint, width int) string {
	_ = artline
	parts := make([]string, 0, len(hints))
	for _, h := range hints {
		parts = append(parts, h.Key+" "+h.Label)
	}
	return styles.FooterStyle.Render(strings.Join(parts, "  ·  "))
}
