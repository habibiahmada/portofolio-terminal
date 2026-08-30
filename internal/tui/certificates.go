package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderCertificatesContent renders the certificates screen.
func (m *App) renderCertificatesContent() string {
	title := styles.TitleStyle.Render("Certificates")

	items := make([]string, 0, len(m.certificates))
	for _, c := range m.certificates {
		item := styles.CardStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				styles.SelectedStyle.Render(c.Name),
				styles.MutedStyle.Render(fmt.Sprintf("%s • %s", c.Issuer, c.Date)),
			),
		)
		items = append(items, item)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(items, "\n"),
	)

	return styles.ContentStyle.Render(content)
}
