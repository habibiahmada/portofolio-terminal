package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// List renders vertical items and highlights the selected index.
func List(items []string, selected int) string {
	lines := make([]string, 0, len(items))
	for i, item := range items {
		if i == selected {
			lines = append(lines, styles.ListSelectedStyle.Render("▸ "+item))
		} else {
			lines = append(lines, styles.SidebarItemStyle.Render("  "+item))
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

// TagList renders a horizontal row of tag pills joined by a space.
func TagList(tags []string) string {
	pills := make([]string, 0, len(tags))
	for _, t := range tags {
		pills = append(pills, styles.TagStyle.Render(t))
	}
	return strings.Join(pills, " ")
}

// Section renders a section title followed by its body, with spacing.
func Section(title string, body ...string) string {
	parts := make([]string, 0, len(body)+2)
	parts = append(parts, title)
	if len(body) > 0 {
		parts = append(parts, "")
		parts = append(parts, body...)
	}
	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}
