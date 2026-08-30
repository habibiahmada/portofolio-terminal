package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderExperienceContent() string {
	cw := m.contentWidth()
	lines := []string{
		styles.SectionTitleStyle.Render("▸ Work Experience"),
		styles.MutedStyle.Render("Roles and engineering programs that shaped my end-to-end development skills."),
		"",
	}

	for _, w := range m.work {
		badge := ""
		if w.Badge != "" {
			badge = "  " + styles.BadgeAccentStyle.Render("["+w.Badge+"]")
		}

		role := styles.PrimaryText.Render(w.Role)
		company := styles.NormalStyle.Render(w.Company)
		period := styles.MutedStyle.Render("[" + w.Period + "]")

		header := fmt.Sprintf("● %s  %s · %s%s", period, role, company, badge)

		locLine := ""
		if w.Location != "" {
			locLine = "  " + styles.MutedStyle.Render("📍 "+w.Location)
		}

		var detailLines []string
		for _, d := range w.Details {
			wrapped := components.WrapText(d, cw-6)
			if len(wrapped) > 0 {
				detailLines = append(detailLines, "  ▸ "+styles.NormalStyle.Render(wrapped[0]))
				for _, dl := range wrapped[1:] {
					detailLines = append(detailLines, "    "+styles.NormalStyle.Render(dl))
				}
			}
		}

		lines = append(lines, header)
		if locLine != "" {
			lines = append(lines, locLine)
		}
		if len(detailLines) > 0 {
			lines = append(lines, detailLines...)
		}
		lines = append(lines, "")
	}

	lines = append(lines,
		styles.SectionTitleStyle.Render("▸ Education & Foundations"),
		styles.MutedStyle.Render("Academic background and formal foundations in software engineering."),
		"",
	)

	for _, e := range m.education {
		title := styles.PrimaryText.Render(e.Title)
		school := styles.NormalStyle.Render(e.School)
		period := styles.MutedStyle.Render("[" + e.Period + "]")

		header := fmt.Sprintf("● %s  %s · %s", period, title, school)

		var descLines []string
		if e.Description != "" {
			wrapped := components.WrapText(e.Description, cw-4)
			for _, dl := range wrapped {
				descLines = append(descLines, "  "+styles.MutedStyle.Render(dl))
			}
		}

		lines = append(lines, header)
		if len(descLines) > 0 {
			lines = append(lines, descLines...)
		}
		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}
