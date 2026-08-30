package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderProjectsContent renders the projects list.
func (m *App) renderProjectsContent() string {
	title := styles.TitleStyle.Render("Projects")
	hint := styles.MutedStyle.Render("↑↓ to browse • Enter to view details")

	cards := make([]string, 0, len(m.projects))
	for i, p := range m.projects {
		stack := strings.Join(p.Stack, " • ")
		nameStyle := styles.NormalStyle
		if i == m.selectedProject {
			nameStyle = styles.SelectedStyle
		}
		card := styles.CardStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				nameStyle.Render(p.Name),
				styles.NormalStyle.Render(p.Description),
				styles.MutedStyle.Render(stack),
			),
		)
		cards = append(cards, card)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		hint,
		"",
		strings.Join(cards, "\n"),
	)

	return styles.ContentStyle.Render(content)
}
