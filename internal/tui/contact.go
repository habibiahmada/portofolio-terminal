package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderContactContent renders the contact screen.
func (m *App) renderContactContent() string {
	title := styles.TitleStyle.Render("Contact")

	contactLines := []string{
		styles.NormalStyle.Render(fmt.Sprintf("Email:    %s", styles.LinkStyle.Render(m.profile.Email))),
		styles.NormalStyle.Render(fmt.Sprintf("GitHub:   %s", styles.LinkStyle.Render(m.profile.GitHub))),
		styles.NormalStyle.Render(fmt.Sprintf("LinkedIn: %s", styles.LinkStyle.Render(m.profile.LinkedIn))),
		styles.NormalStyle.Render(fmt.Sprintf("Website:  %s", styles.LinkStyle.Render(m.profile.Website))),
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(contactLines, "\n"),
		"",
		styles.MutedStyle.Render("Feel free to reach out!"),
	)

	return styles.ContentStyle.Render(content)
}
