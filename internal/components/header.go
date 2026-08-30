// Package components provides reusable, presentation-only UI pieces for the
// TUI. Components are pure functions: they take data and return a rendered
// string. No business logic, no model state — screens keep their own state and
// call these to render output.
package components

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Header renders the profile masthead (name + title) inside a bordered bar.
func Header(name, title string, width int) string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.TitleStyle.Render(name),
		styles.SubtitleStyle.Render(title),
	)
	return styles.BorderStyle.Width(width - 2).Render(content)
}
