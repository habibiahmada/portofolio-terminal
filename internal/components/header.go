// Package components provides reusable, presentation-only UI pieces for the
// TUI. Components are pure functions: they take data and return a rendered
// string. No business logic, no model state — screens keep their own state and
// call these to render output.
package components

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Header renders the brand masthead: the wordmark "habibiahmada" with a brand
// red dot, plus the role title, inside a bordered bar that spans the width.
func Header(wordmark, title string, width int) string {
	mark := styles.HeaderWordmark.Render(wordmark)
	dot := styles.HeaderDot.Render(".")
	right := styles.HeaderMeta.Render(title)

	row := lipgloss.JoinHorizontal(lipgloss.Center, mark, dot, "   "+right)
	return styles.BorderStyle.Width(width - 2).Render(row)
}
