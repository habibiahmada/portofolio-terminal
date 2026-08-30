package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderSkillsContent renders the skills screen: a header plus the 16 tools as
// a responsive flat grid of tag pills.
func (m *App) renderSkillsContent() string {
	label := styles.LabelStyle.Render("// Tech Stack")
	title := styles.SectionTitleStyle.Render("Tools & Technologies")
	sub := styles.MutedStyle.Render(
		"The technologies I use daily to turn ideas into functional, high-performing digital reality.",
	)

	contentWidth := m.width - sidebarWidth - 2
	names := make([]string, 0, len(m.skills))
	for _, s := range m.skills {
		names = append(names, s.Name)
	}
	grid := components.TagGrid(names, contentWidth)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		title,
		sub,
		"",
		grid,
	)

	return styles.ContentStyle.Render(content)
}
