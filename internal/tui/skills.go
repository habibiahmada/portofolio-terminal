package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderSkillsContent renders the skills screen grouped by category.
func (m *App) renderSkillsContent() string {
	title := styles.TitleStyle.Render("Skills")

	categories := make(map[string][]data.Skill)
	for _, s := range m.skills {
		categories[s.Category] = append(categories[s.Category], s)
	}

	lines := make([]string, 0)
	for _, cat := range data.GetSkillCategories() {
		lines = append(lines, styles.SubtitleStyle.Render(cat))
		for _, s := range categories[cat] {
			lines = append(lines, fmt.Sprintf("  %s %s", styles.NormalStyle.Render(s.Name), renderSkillBar(s.Level)))
		}
		lines = append(lines, "")
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(lines, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderSkillBar renders a visual skill level bar with a percentage label.
func renderSkillBar(level int) string {
	filled := strings.Repeat("█", level)
	empty := strings.Repeat("░", 5-level)
	percent := level * 20
	return styles.SuccessStyle.Render(filled) +
		styles.MutedStyle.Render(empty) +
		fmt.Sprintf("  %d%%", percent)
}
