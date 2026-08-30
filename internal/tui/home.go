package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderHomeContent renders the home screen: a scrollable page of sections
// matching the website (hero, companies, featured, services, press, process,
// CTA).
func (m *App) renderHomeContent() string {
	contentWidth := m.width - sidebarWidth - 2

	sections := make([]string, 0, 8)
	sections = append(sections, m.renderHeroSection(contentWidth))

	if m.width >= 60 {
		sections = append(sections, m.renderCompaniesSection(contentWidth))
	}
	sections = append(sections, m.renderFeaturedSection(contentWidth))
	if m.width >= 60 {
		sections = append(sections, m.renderServicesPreviewSection(contentWidth))
	}
	sections = append(sections, m.renderPressSection(contentWidth))
	sections = append(sections, m.renderProcessSection(contentWidth))
	sections = append(sections, m.renderHomeCTASection(contentWidth))

	sep := "\n" + styles.RuleStyle.Render(strings.Repeat("─", max(contentWidth-2, 8))) + "\n"
	return styles.ContentStyle.Render(strings.Join(sections, sep))
}

// renderHeroSection renders the home hero with signature, H1, subtitle, and
// availability badge.
func (m *App) renderHeroSection(width int) string {
	sig := components.Signature(width)
	nameBlock := styles.HeroTitleStyle.Render(strings.ToUpper(m.profile.BoostName))
	if m.height >= 24 && width >= 90 {
		nameBlock = m.renderHeroName(width)
	}

	badge := styles.SuccessStyle.Render("● " + m.profile.Availability)

	h1 := styles.SectionTitleStyle.Render(
		"Building digital experiences that " + styles.PrimaryText.Render("actually matter"),
	)
	sub := styles.NormalStyle.Render(
		"Frontend-leaning full-stack developer. I craft clear interfaces and the APIs behind them, with a bias for performance you can measure, not just claim.",
	)

	hints := []string{
		"[ ↑↓ ] Explore",
		"[ P ] Projects",
		"[ C ] Contact",
		"[ V ] CV",
	}
	hintLine := components.Hints(hints, width)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		sig,
		"",
		badge,
		"",
		nameBlock,
		h1,
		sub,
		"",
		hintLine,
	)
}

// renderCompaniesSection renders the "Collaborations & Trusted By" marquee.
func (m *App) renderCompaniesSection(width int) string {
	label := styles.LabelStyle.Render("// Companies")
	heading := styles.SectionTitleStyle.Render("Collaborations & Trusted By")

	names := make([]string, 0, len(m.companies))
	for _, c := range m.companies {
		names = append(names, styles.SubtitleStyle.Render("● "+c.Name))
	}
	line := strings.Join(names, "   ")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		heading,
		line,
	)
}

// renderFeaturedSection renders the featured projects cards.
func (m *App) renderFeaturedSection(width int) string {
	label := styles.LabelStyle.Render("// Featured Work")
	heading := lipgloss.JoinHorizontal(lipgloss.Left,
		styles.SectionTitleStyle.Render("Featured Projects"),
		"   ",
		styles.LinkStyle.Render("All Projects →"),
	)

	featured := make([]string, 0)
	for _, p := range m.projects {
		if !p.Featured {
			continue
		}
		featured = append(featured, renderProjectCard(p, m.width, false))
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		heading,
		"",
		components.Cards(featured...),
	)
}

// renderServicesPreviewSection renders the abbreviated service cards.
func (m *App) renderServicesPreviewSection(width int) string {
	label := styles.LabelStyle.Render("// Services")
	heading := styles.SectionTitleStyle.Render("What I can help with")

	lines := make([]string, 0, len(m.services))
	for _, sv := range m.services {
		head := fmt.Sprintf("%s / %s  %s", styles.SelectedStyle.Render(sv.Number), styles.MutedStyle.Render(sv.Category), styles.NormalStyle.Render(sv.Title))
		lines = append(lines, head)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		heading,
		"",
		strings.Join(lines, "\n"),
		"",
		styles.MutedStyle.Render("All services → press Enter on Services in the sidebar"),
	)
}

// renderPressSection renders the two spotlight stories.
func (m *App) renderPressSection(width int) string {
	label := styles.LabelStyle.Render("// Press")
	heading := styles.SectionTitleStyle.Render("Spotlight")

	cards := make([]string, 0, len(m.press))
	for _, pr := range m.press {
		card := components.Card(
			styles.NormalStyle.Render(pr.Title),
			"",
			styles.MutedStyle.Render(pr.Body),
			"",
			styles.LinkStyle.Render("→ "+pr.CTALabel),
		)
		cards = append(cards, card)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		heading,
		"",
		components.Cards(cards...),
	)
}

// renderProcessSection renders the four-step "How I ship" sequence.
func (m *App) renderProcessSection(width int) string {
	label := styles.LabelStyle.Render("// How I ship")
	heading := styles.SectionTitleStyle.Render("Process over performance theater")
	sub := styles.MutedStyle.Render("A short operating manual, the same sequence I use on school systems, internships, and competition deadlines.")

	steps := make([]string, 0, len(m.process))
	for _, ps := range m.process {
		head := fmt.Sprintf("%s %s", styles.SelectedStyle.Render(ps.Number+"."), styles.NormalStyle.Render(ps.Title))
		steps = append(steps, head, styles.MutedStyle.Render("   "+ps.Description), "")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		heading,
		sub,
		"",
		strings.Join(steps, "\n"),
	)
}

// renderHomeCTASection renders the closing contact CTA.
func (m *App) renderHomeCTASection(width int) string {
	label := styles.LabelStyle.Render("// Contact")
	heading := styles.SectionTitleStyle.Render("Need a web product that actually ships in the next 90 days?")
	body := styles.NormalStyle.Render(
		"Open to freelance and full-time. Remote (WIB). Write to " +
			styles.LinkStyle.Render(m.profile.Email) +
			". I usually reply within 48 hours.",
	)
	cta := fmt.Sprintf("%s   %s", styles.SelectedStyle.Render("[ C ] Let's talk"), styles.MutedStyle.Render("· Browse projects with [ P ]"))

	return styles.PrimaryCardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			label,
			heading,
			"",
			body,
			"",
			cta,
		),
	)
}

// renderHeroName renders the hero name: first token as large FIGlet text when
// it fits, remainder as normal hero text. Falls back to plain hero text.
func (m *App) renderHeroName(width int) string {
	fields := strings.Fields(strings.ToUpper(m.profile.BoostName))
	if len(fields) == 0 {
		return styles.HeroTitleStyle.Render("")
	}
	fig := components.Figlet(fields[0], width)
	if len(fields) == 1 {
		return fig
	}
	return fig + "\n" + styles.HeroTitleStyle.Render(strings.Join(fields[1:], " "))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
