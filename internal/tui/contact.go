package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderContactContent renders the contact screen: CTA copy, email, social
// links, and availability badge.
func (m *App) renderContactContent() string {
	label := styles.LabelStyle.Render("// Contact")

	badge := styles.SuccessStyle.Render("● " + m.profile.Availability)

	h2 := styles.SectionTitleStyle.Render("Need a web product that actually ships in the next 90 days?")

	body := styles.NormalStyle.Render(
		"Open to freelance and full-time. Remote (WIB). Write to " +
			styles.LinkStyle.Render(m.profile.Email) +
			". I usually reply within 48 hours.",
	)

	cta := styles.SelectedStyle.Render("Let's talk → " + m.profile.Email)

	// Social list.
	lines := make([]string, 0, len(m.socials)+1)
	lines = append(lines, "", styles.SectionTitleStyle.Render("Social Profiles"))
	for _, s := range m.socials {
		lines = append(lines, fmt.Sprintf("[%s] %s  %s",
			styles.SelectedStyle.Render(s.Icon),
			styles.NormalStyle.Render(s.Name+":"),
			styles.LinkStyle.Render(s.URL),
		))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		badge,
		"",
		h2,
		"",
		body,
		"",
		styles.PrimaryCardStyle.Render(cta),
		strings.Join(lines, "\n"),
	)

	return styles.ContentStyle.Render(content)
}
