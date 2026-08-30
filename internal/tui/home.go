package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderHomeContent — compact landing: hero (text | art), partners, top projects.
func (m *App) renderHomeContent() string {
	w := m.contentWidth()
	parts := []string{
		m.renderHeroCompact(w),
		m.renderCompaniesInline(w),
		m.renderFeaturedCompact(w),
	}
	return strings.Join(parts, "\n")
}

func (m *App) renderHeroCompact(width int) string {
	badge := styles.SuccessStyle.Render("● " + m.profile.Availability)

	text := lipgloss.JoinVertical(
		lipgloss.Left,
		badge,
		styles.SectionTitleStyle.Render(m.profile.Name),
		styles.MutedStyle.Render(m.profile.Title+" · "+m.profile.Location),
		styles.NormalStyle.Render(
			"Full-stack developer — clear UI, solid APIs, measurable performance.",
		),
	)

	artW := width / 3
	if artW < 22 {
		artW = 22
	}
	art := components.Signature(artW)
	if art == "" {
		art = components.AboutTerminal()
	}

	textW := width - artW - 2
	if textW < 24 {
		return lipgloss.JoinVertical(lipgloss.Left, text, "", art)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.NewStyle().Width(textW).Render(text),
		"  ",
		art,
	)
}

func (m *App) renderCompaniesInline(width int) string {
	names := make([]string, 0, len(m.companies))
	for _, c := range m.companies {
		names = append(names, c.Name)
	}
	line := strings.Join(names, " · ")
	return "\n" + styles.MutedStyle.Render("Trusted · "+components.WrapTextLine(line, width))
}

func (m *App) renderFeaturedCompact(width int) string {
	var lines []string
	n := 0
	for _, p := range m.projects {
		if !p.Featured {
			continue
		}
		if n >= 3 {
			break
		}
		tag := p.Year
		if len(p.Tags) > 0 {
			tag += " · " + p.Tags[0]
		}
		lines = append(lines, fmt.Sprintf("▸ %s  %s",
			styles.NormalStyle.Render(p.Name),
			styles.MutedStyle.Render(tag),
		))
		n++
	}
	if len(lines) == 0 {
		return ""
	}
	hint := styles.MutedStyle.Render("↑↓ screens · Enter projects")
	return "\n" + strings.Join(lines, "\n") + "\n" + hint
}
