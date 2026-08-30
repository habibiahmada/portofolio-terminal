package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderExperienceContent() string {
	lines := []string{styles.SectionTitleStyle.Render("Experience")}

	for _, w := range m.work {
		badge := ""
		if w.Badge != "" {
			badge = styles.BadgeAccentStyle.Render(w.Badge)
		}
		lines = append(lines, renderExpRow(w.Period, w.Role, w.Company, badge, m.contentWidth()))
	}

	lines = append(lines, "", styles.MutedStyle.Render("Education"))
	for _, e := range m.education {
		lines = append(lines, renderExpRow(e.Period, e.Title, e.School, "", m.contentWidth()))
	}
	return strings.Join(lines, "\n")
}

// renderExpRow lays out one experience/education entry with a fixed-width period
// column, the role+company on the same row, and the optional badge kept inline
// when it fits, otherwise wrapped onto an indented continuation line.
func renderExpRow(period, primary, secondary, badge string, contentWidth int) string {
	const periodCol = 18

	namePart := primary + " · " + secondary

	if badge == "" {
		return fmt.Sprintf("%-*s%s",
			periodCol, styles.MutedStyle.Render(period),
			styles.NormalStyle.Render(namePart),
		)
	}

	// Try to fit role+company and badge on one line.
	baseWidth := periodCol + lipgloss.Width(namePart)
	badgeWidth := lipgloss.Width(badge)
	if baseWidth+2+badgeWidth <= contentWidth {
		return fmt.Sprintf("%-*s%s  %s",
			periodCol, styles.MutedStyle.Render(period),
			styles.NormalStyle.Render(namePart),
			badge,
		)
	}

	// Wrap: keep the badge on its own aligned line.
	return fmt.Sprintf("%-*s%s\n%s%s",
		periodCol, styles.MutedStyle.Render(period),
		styles.NormalStyle.Render(namePart),
		strings.Repeat(" ", periodCol),
		badge,
	)
}
