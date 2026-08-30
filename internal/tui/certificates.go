package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderCertificatesContent renders the certificates screen: 3 pinned/featured
// entries plus a flat grid of the remaining certificates.
func (m *App) renderCertificatesContent() string {
	label := styles.LabelStyle.Render("// Certificates")
	title := styles.SectionTitleStyle.Render("Licenses & Certifications")
	sub := styles.MutedStyle.Render(
		"Professional certifications and awards that validate my expertise in software development, cloud computing, and technology innovation.",
	)

	// Pinned (featured).
	pinned := make([]string, 0, 3)
	for _, c := range m.certificates {
		if !c.Pinned {
			continue
		}
		item := styles.SelectedStyle.Render("★ "+c.Name) + "  " + styles.MutedStyle.Render(c.Issuer+" · "+c.Date)
		pinned = append(pinned, item)
	}

	// Remaining as a tag grid.
	names := make([]string, 0, len(m.certificates))
	for _, c := range m.certificates {
		if c.Pinned {
			continue
		}
		names = append(names, c.Name+" · "+c.Issuer)
	}

	contentWidth := m.width - sidebarWidth - 2
	grid := components.TagGrid(names, contentWidth)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		title,
		sub,
		"",
		styles.SectionTitleStyle.Render("Pinned"),
		strings.Join(pinned, "\n"),
		"",
		styles.SectionTitleStyle.Render("All"),
		grid,
	)

	return styles.ContentStyle.Render(content)
}
