package tui

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderSkillsContent() string {
	names := make([]string, 0, len(m.skills))
	for _, s := range m.skills {
		names = append(names, s.Name)
	}
	return strings.Join([]string{
		styles.SectionTitleStyle.Render("Tools & Technologies"),
		components.SkillGrid(names, m.contentWidth()),
	}, "\n")
}
