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

// Hints renders a row of bracketed keyboard hint pills, wrapping to fit width.
func Hints(hints []string, width int) string {
	pills := make([]string, 0, len(hints))
	for _, h := range hints {
		pills = append(pills, styles.SelectedStyle.Render(h))
	}
	joined := strings.Join(pills, "  ")
	if width > 0 && lipgloss.Width(joined) > width {
		return lipgloss.JoinVertical(lipgloss.Left, pills...)
	}
	return joined
}

// TagGrid renders a flat list of tag pills, wrapping into rows that fit the
// given width. Unstyled names are styled as tech tags.
func TagGrid(tags []string, width int) string {
	rows := make([]string, 0)
	current := ""
	for _, t := range tags {
		pill := styles.TagStyle.Render(t)
		if current == "" {
			current = pill
			continue
		}
		if lipgloss.Width(current)+lipgloss.Width(pill) > width {
			rows = append(rows, current)
			current = pill
		} else {
			current = current + " " + pill
		}
	}
	if current != "" {
		rows = append(rows, current)
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

// TagColumn renders each tag as a full-width bordered row so a long list
// (certificates) has a straight right edge instead of ragged pill widths.
func TagColumn(tags []string, width int) string {
	if width < 8 {
		width = 8
	}
	inner := width - 4 // rounded border (2) + horizontal padding (2)
	if inner < 4 {
		inner = 4
	}
	box := styles.TagStyle.Copy().
		Width(width - 2).
		MarginRight(0)
	rows := make([]string, 0, len(tags))
	for _, t := range tags {
		rows = append(rows, box.Render(Truncate(t, inner)))
	}
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
