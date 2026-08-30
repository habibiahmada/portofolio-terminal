package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderExperienceContent renders the experience screen as a vertical timeline
// (4 work entries) plus an education "Foundations" section and a companies
// marquee.
func (m *App) renderExperienceContent() string {
	label := styles.LabelStyle.Render("// Experience")
	title := styles.SectionTitleStyle.Render("Path so far")
	sub := styles.MutedStyle.Render("Roles and programs that taught me to scope, ship, and explain the trade-offs.")

	work := make([]string, 0, len(m.work)+1)
	work = append(work, label, title, sub, "")
	for _, w := range m.work {
		work = append(work, m.renderWorkNode(w))
	}

	// Education.
	edu := make([]string, 0, len(m.education)+3)
	edu = append(edu, "", styles.LabelStyle.Render("// Foundations"), styles.SectionTitleStyle.Render("Education"), "")
	for _, e := range m.education {
		edu = append(edu,
			styles.MutedStyle.Render(e.Period),
			styles.NormalStyle.Render(e.Title),
			styles.SubtitleStyle.Render("   "+e.School),
			styles.MutedStyle.Render("   "+e.Description),
			"",
		)
	}

	// Companies marquee.
	companyNames := make([]string, 0, len(m.companies))
	for _, c := range m.companies {
		companyNames = append(companyNames, styles.SubtitleStyle.Render("● "+c.Name))
	}
	companies := append([]string{""}, strings.Join(companyNames, "   "))

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		append(append(work, edu...), companies...)...,
	)

	return styles.ContentStyle.Render(content)
}

// renderWorkNode renders one timeline entry with connectors and a badge.
func (m *App) renderWorkNode(w data.ExperienceWork) string {
	head := styles.MutedStyle.Render(w.Period+"  ") + styles.NormalStyle.Render(w.Role)
	company := styles.SubtitleStyle.Render("   "+w.Company) + "  " + styles.MutedStyle.Render(w.Location)

	lines := []string{head, company}

	if w.Badge != "" {
		badge := renderBadge(w.Badge)
		lines = append(lines, "   "+badge)
	}

	for _, d := range w.Details {
		lines = append(lines, "   • "+styles.NormalStyle.Render(d))
	}
	lines = append(lines, "")

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

// renderBadge returns a styled status badge for a work entry.
func renderBadge(badge string) string {
	switch badge {
	case "Current":
		return styles.BadgeStyle.Render(badge)
	case "Top 15 Capstone":
		return styles.BadgeAccentStyle.Render(badge)
	default:
		return styles.BadgeNeutralStyle.Render(badge)
	}
}
