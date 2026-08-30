package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderServicesContent() string {
	cw := m.contentWidth()
	lines := []string{
		styles.SectionTitleStyle.Render("▸ Services & Solutions"),
		styles.MutedStyle.Render("End-to-end web engineering for products, startups, and teams - from discovery through production deployment."),
		"",
	}

	for _, sv := range m.services {
		numBadge := styles.BadgeAccentStyle.Render("[" + sv.Number + "]")
		catBadge := styles.BadgeStyle.Render("[" + sv.Category + "]")
		title := styles.PrimaryText.Render(sv.Title)

		header := fmt.Sprintf("  %s %s %s", numBadge, catBadge, title)
		for _, hl := range components.WrapText(header, cw) {
			lines = append(lines, hl)
		}

		descLines := components.WrapText(sv.Description, cw-6)
		for _, dl := range descLines {
			lines = append(lines, "     "+styles.NormalStyle.Render(dl))
		}
		lines = append(lines, "")
	}

	// 4-step operating process
	lines = append(lines,
		styles.SectionTitleStyle.Render("▸ How I Ship (4-Step Process)"),
		styles.MutedStyle.Render("The repeatable workflow I apply across client work and personal projects:"),
		"",
	)

	for _, st := range m.process {
		num := styles.BadgeAccentStyle.Render("[" + st.Number + "]")
		title := styles.PrimaryText.Render(st.Title)
		descLines := components.WrapText(st.Description, cw-6)
		desc := styles.NormalStyle.Render(strings.Join(descLines, "\n     "))
		lines = append(lines, fmt.Sprintf("  %s %s\n     %s", num, title, desc), "")
	}

	// CTA line
	cta := styles.PromptStyle.Render("  Inquiries: "+m.profile.Email) + "  ·  " + styles.LinkStyle.Render("https://"+m.profile.Website)
	for _, wl := range components.WrapText(cta, cw) {
		lines = append(lines, wl)
	}

	return strings.Join(lines, "\n")
}
