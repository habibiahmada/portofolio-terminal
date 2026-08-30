package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderCertificatesContent() string {
	cw := m.contentWidth()
	lines := []string{
		styles.SectionTitleStyle.Render("▸ Licenses & Certifications"),
		styles.MutedStyle.Render(fmt.Sprintf("%d verified credentials across cloud, fullstack engineering, and AI systems.", len(m.certificates))),
		"",
	}

	// Featured / Pinned Awards
	var pinned []data.Certificate
	var rest []data.Certificate
	for _, c := range m.certificates {
		if c.Pinned {
			pinned = append(pinned, c)
		} else {
			rest = append(rest, c)
		}
	}

	if len(pinned) > 0 {
		lines = append(lines, styles.PrimaryText.Render("★ Featured Honors & Awards"))
		for _, p := range pinned {
			badge := styles.BadgeAccentStyle.Render("[Award]")
			meta := styles.MutedStyle.Render(p.Issuer + " · " + p.Date)
			wrapped := components.WrapText(p.Name, cw-12)
			if len(wrapped) == 0 {
				wrapped = []string{p.Name}
			}
			lines = append(lines,
				fmt.Sprintf("  ★ %s %s", badge, styles.NormalStyle.Render(wrapped[0])),
			)
			for _, wl := range wrapped[1:] {
				lines = append(lines, fmt.Sprintf("          %s", styles.NormalStyle.Render(wl)))
			}
			lines = append(lines, fmt.Sprintf("          %s", meta))
		}
		lines = append(lines, "")
	}

	// All Certifications list
	lines = append(lines,
		styles.SectionTitleStyle.Render("▸ Professional Certifications"),
		styles.MutedStyle.Render("Verified courses, technical assessments, and accreditations:"),
		"",
	)

	for i, c := range rest {
		issuerTag := styles.BadgeStyle.Render(fmt.Sprintf("[%-10s]", c.Issuer))
		date := styles.MutedStyle.Render("[" + c.Date + "]")

		name := c.Name
		if idx := strings.Index(name, ": "); idx != -1 && strings.EqualFold(name[:idx], c.Issuer) {
			name = name[idx+2:]
		}

		wrapW := cw - 26
		if wrapW < 20 {
			wrapW = 20
		}
		wrapped := components.WrapText(name, wrapW)
		if len(wrapped) == 0 {
			wrapped = []string{name}
		}

		firstLine := fmt.Sprintf("  %2d. %s %s %s", i+1, issuerTag, date, styles.NormalStyle.Render(wrapped[0]))
		lines = append(lines, firstLine)
		for _, extra := range wrapped[1:] {
			lines = append(lines, fmt.Sprintf("                      %s", styles.NormalStyle.Render(extra)))
		}
	}

	return strings.Join(lines, "\n")
}
