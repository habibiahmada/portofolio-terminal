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
	name := components.ShimmerText(
		m.profile.Name,
		m.footerFrame,
		styles.SectionTitleStyle.Copy().UnsetMarginBottom(),
		styles.PrimaryText,
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
	stack := textW < 28
	wrapW := width
	if !stack {
		wrapW = textW
	}

	badge := styles.SuccessStyle.Render(
		components.WrapTextLine("● "+m.profile.Availability, wrapW),
	)
	meta := styles.MutedStyle.Render(
		components.WrapTextLine(m.profile.Title+" · "+m.profile.Location, wrapW),
	)
	bio := styles.NormalStyle.Render(
		components.WrapTextLine(
			"Full-stack developer — clear UI, solid APIs, measurable performance.",
			wrapW,
		),
	)
	text := lipgloss.JoinVertical(lipgloss.Left, badge, name, meta, bio)
	if stack {
		return lipgloss.JoinVertical(lipgloss.Left, text, "", art)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, text, "  ", art)
}

func (m *App) renderCompaniesInline(width int) string {
	names := make([]string, 0, len(m.companies))
	for _, c := range m.companies {
		names = append(names, c.Name)
	}
	line := strings.Join(names, " · ")
	return "\n" + styles.MutedStyle.Render(components.WrapTextLine("Trusted · "+line, width))
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
