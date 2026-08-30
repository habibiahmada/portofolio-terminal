package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderProjectDetailContent renders a single project detail view.
func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail
	title := styles.TitleStyle.Render(p.Name)
	desc := styles.NormalStyle.Render(p.Description)

	stack := components.TagList(p.Stack)

	links := make([]string, 0, 2)
	if p.GitHub != "" {
		links = append(links, styles.LinkStyle.Render("GitHub: "+p.GitHub))
	}
	if p.Live != "" {
		links = append(links, styles.LinkStyle.Render("Live: "+p.Live))
	}

	closeHint := styles.MutedStyle.Render("← back to projects")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		desc,
		"",
		stack,
		"",
		strings.Join(links, "\n"),
		"",
		closeHint,
	)

	return styles.ContentStyle.Render(content)
}
