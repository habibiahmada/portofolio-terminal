package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderAboutContent renders the about screen with a micro-illustration.
func (m *App) renderAboutContent() string {
	title := styles.TitleStyle.Render("About Me")

	art := components.AboutTerminal()

	bio := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.NormalStyle.Render(fmt.Sprintf("Name: %s", m.profile.Name)),
		styles.NormalStyle.Render(fmt.Sprintf("Title: %s", m.profile.Title)),
		styles.NormalStyle.Render(fmt.Sprintf("Location: %s", m.profile.Location)),
		styles.LinkStyle.Render(fmt.Sprintf("Email: %s", m.profile.Email)),
	)

	body := lipgloss.JoinHorizontal(lipgloss.Top, art, "   ", bio)
	if art == "" || m.width < 70 {
		body = lipgloss.JoinVertical(lipgloss.Left, art, "", bio)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		body,
	)

	return styles.ContentStyle.Render(content)
}
