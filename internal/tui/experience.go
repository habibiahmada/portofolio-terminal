package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderExperienceContent renders the experience screen as a vertical timeline
// (data-driven; flattens to a simple list when there is a single entry).
func (m *App) renderExperienceContent() string {
	title := styles.TitleStyle.Render("Experience")

	if len(m.experiences) == 0 {
		return styles.ContentStyle.Render(title)
	}

	lines := make([]string, 0, len(m.experiences)*5)
	for i, e := range m.experiences {
		lines = append(lines, renderTimelineNode(i, len(m.experiences), e))
	}

	view := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(lines, "\n"),
	)

	return styles.ContentStyle.Render(view)
}

// renderTimelineNode renders one experience entry with timeline connectors.
func renderTimelineNode(i, total int, e data.Experience) string {
	isLast := i == total-1
	hasNext := !isLast

	head := fmt.Sprintf("● %s — %s", e.Period, e.Role)
	headLine := styles.SubtitleStyle.Render(head)

	company := styles.NormalStyle.Render("  " + e.Company)
	meta := styles.MutedStyle.Render("  " + e.Location)

	prevConnector := ""
	if hasNext {
		prevConnector = "│"
	}

	lines := []string{
		headLine,
		company,
		meta,
	}

	for j, d := range e.Details {
		connector := "│"
		if isLast && j == len(e.Details)-1 {
			connector = "  "
		}
		if !isLast && j == len(e.Details)-1 {
			connector = "├"
			prevConnector = "│"
		}
		lines = append(lines, fmt.Sprintf("  %s  • %s", connector, styles.NormalStyle.Render(d)))
	}

	if hasNext {
		lines = append(lines, "  "+prevConnector)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
