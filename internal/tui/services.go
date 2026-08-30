package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderServicesContent renders the services screen: 5 numbered service cards,
// the process section, and a contact CTA.
func (m *App) renderServicesContent() string {
	label := styles.LabelStyle.Render("// My Services")
	title := styles.SectionTitleStyle.Render("Comprehensive Solutions")
	sub := styles.MutedStyle.Render(
		"From wireframe concepts to fully animated frontends and scalable servers. I build performant products that stand out.",
	)

	cards := make([]string, 0, len(m.services))
	for _, sv := range m.services {
		head := lipgloss.JoinHorizontal(lipgloss.Left,
			styles.SelectedStyle.Render(sv.Number+" / "),
			styles.MutedStyle.Render(sv.Category),
			"  ",
			styles.SectionTitleStyle.Render(sv.Title),
		)
		card := components.Card(
			head,
			"",
			styles.NormalStyle.Render(sv.Description),
		)
		cards = append(cards, card)
	}

	// Process section.
	processLabel := styles.LabelStyle.Render("// How I ship")
	processTitle := styles.SectionTitleStyle.Render("Process over performance theater")
	steps := make([]string, 0, len(m.process))
	for _, ps := range m.process {
		steps = append(steps,
			fmt.Sprintf("%s %s", styles.SelectedStyle.Render(ps.Number+"."), styles.NormalStyle.Render(ps.Title)),
			styles.MutedStyle.Render("   "+ps.Description),
			"",
		)
	}

	// CTA.
	cta := fmt.Sprintf("%s   %s",
		styles.SelectedStyle.Render("[ C ] Let's talk"),
		styles.MutedStyle.Render("→ Contact"),
	)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		title,
		sub,
		"",
		components.Cards(cards...),
		"",
		processLabel,
		processTitle,
		"",
		strings.Join(steps, "\n"),
		"",
		styles.PrimaryCardStyle.Render(cta),
	)

	return styles.ContentStyle.Render(content)
}
