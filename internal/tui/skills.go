package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/sanitize"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderSkillsContent() string {
	cw := m.contentWidth()
	lines := []string{
		styles.SectionTitleStyle.Render("▸ Tools & Technologies"),
		styles.MutedStyle.Render("Languages, frameworks, and infrastructure I use to design, build, and ship production-ready web applications."),
		"",
	}

	categories := sanitize.SkillCategories(data.GetSkillCategories())
	for _, cat := range categories {
		icon := styles.BadgeAccentStyle.Render("[" + cat.Icon + "]")
		title := styles.PrimaryText.Render(cat.Name)
		header := fmt.Sprintf("  %s %s", icon, title)

		chips := make([]string, 0, len(cat.Skills))
		for _, sk := range cat.Skills {
			chips = append(chips, components.SkillChip(sk))
		}
		rawChipLine := strings.Join(chips, "  ")

		var chipBlock string
		if cw > 0 && lipgloss.Width(rawChipLine) > cw-4 {
			grid := components.SkillGrid(cat.Skills, cw-4)
			var indented []string
			for _, gl := range strings.Split(grid, "\n") {
				indented = append(indented, "   "+gl)
			}
			chipBlock = strings.Join(indented, "\n")
		} else {
			chipBlock = "   " + rawChipLine
		}

		descLines := components.WrapText(cat.Description, cw-6)
		var indentedDesc []string
		for _, dl := range descLines {
			indentedDesc = append(indentedDesc, "   "+styles.MutedStyle.Render(dl))
		}
		descBlock := strings.Join(indentedDesc, "\n")

		lines = append(lines, header, chipBlock, descBlock, "")
	}

	overview := styles.MutedStyle.Render("● 16+ core technologies across the stack - from responsive UI through APIs, databases, and cloud deployment.")
	lines = append(lines, overview)

	return strings.Join(lines, "\n")
}
