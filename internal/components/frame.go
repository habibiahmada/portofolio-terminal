package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// cornerCycle is the animated accent set used for the frame corners and the
// header casing. Each glyph is one cell wide so widths stay stable.
var cornerCycle = []string{"◆", "◈", "◇", "◈"}

// edgeDot cycles for the decorative top/bottom frame rule accents.
var edgeDot = []string{"·", "·", "•", "·", "·", "•"}

// HeaderBar renders a fixed top bar: "wordmark - title" on the left, ambient
// art on the right, with a separator rule below. The casing glyph animates.
func HeaderBar(wordmark, title, location string, width, frame int) string {
	if width <= 0 {
		width = 1
	}

	brand := styles.HeaderWordmark.Render(wordmark) + styles.HeaderDot.Render(".")
	role := styles.MutedStyle.Render(" - " + title)
	loc := styles.HeaderMeta.Render("  " + location)

	left := lipgloss.JoinHorizontal(lipgloss.Left, brand, role, loc)

	ambient := styles.ArtBrandSecondary.Render(edgeDot[frame%len(edgeDot)]) +
		strings.Repeat(" ", 2) +
		styles.FooterBarStyle.Render(cornerCycle[frame%len(cornerCycle)]) +
		strings.Repeat(" ", 1) +
		styles.ArtBrandSecondary.Render(edgeDot[(frame+2)%len(edgeDot)])

	leftW := lipgloss.Width(left)
	ambientW := lipgloss.Width(ambient)
	gap := width - leftW - ambientW
	if gap < 2 {
		ambient = ""
		gap = width - leftW - 1
	}
	if gap < 1 {
		gap = 1
	}
	row := lipgloss.JoinHorizontal(lipgloss.Left, left, strings.Repeat(" ", gap), ambient)

	node := styles.FrameAccentStyle.Render(cornerCycle[frame%len(cornerCycle)])
	span := width - 1
	if span < 1 {
		span = 1
	}
	// A slow, subtle "scan" travels along the header rule: one bright dash moves
	// left-to-right (wrapping) while the rest of the rule stays muted. Width is
	// fixed by span so the layout never shifts.
	scanPos := frame % span
	rule := ""
	for i := 0; i < span-1; i++ {
		if i == scanPos {
			rule += styles.HeaderScanStyle.Render("─")
		} else {
			rule += styles.FrameBorderStyle.Render("─")
		}
	}
	rule += node

	return styles.HeaderBarStyle.Width(width).Render(row) + "\n" +
		styles.HeaderBarStyle.Width(width).Render(rule)
}
