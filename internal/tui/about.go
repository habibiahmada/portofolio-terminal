package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderAboutContent renders the about screen: hero, stats, intro, and CTAs.
func (m *App) renderAboutContent() string {
	contentWidth := m.width - sidebarWidth - 2

	// Hero.
	h1 := lipgloss.JoinHorizontal(lipgloss.Left,
		styles.PrimaryText.Render("Habibi"),
		"  ",
		styles.SectionTitleStyle.Render("Ahmad Aziz"),
	)

	oneLiner := styles.NormalStyle.Render(
		"Full-Stack Web Developer experienced in building responsive apps and CMS products. Skilled in crafting end-to-end features using Next.js, React, Laravel, Node.js, and WordPress to deliver production-ready solutions.",
	)

	ctaLine := fmt.Sprintf("%s   %s",
		styles.SelectedStyle.Render("[ Enter ] Let's Collaborate"),
		styles.MutedStyle.Render("[ C ] Contact · [ P ] Projects"),
	)

	// Stats bar.
	stats := make([]string, 0, len(m.profile.Stats))
	for i, s := range m.profile.Stats {
		block := lipgloss.JoinVertical(
			lipgloss.Left,
			styles.PrimaryText.Render(s.Value),
			styles.MutedStyle.Render(s.Label),
		)
		stats = append(stats, block)
		if i < len(m.profile.Stats)-1 {
			stats = append(stats, "   ")
		}
	}

	// Intro with micro-illustration side by side (when wide enough).
	label := styles.LabelStyle.Render("// About")
	h2 := styles.SectionTitleStyle.Render("A Glimpse Into" + "\n" + "Who I Am")

	para1 := styles.NormalStyle.Render(
		"As a Software Engineering graduate from SMKN 1 Karawang, I currently work as a Web Developer at PT Webekspres Teknologi Indonesia. My expertise lies in developing tailored client websites, architecting CMS platforms, and deploying scalable full-stack features for production environments.",
	)
	para2 := styles.NormalStyle.Render(
		"Driven by a deep passion for software architecture and modern web technologies, I thrive on solving complex technical challenges. My development philosophy centers on writing clean, scalable code and crafting intuitive digital experiences that deliver tangible impact.",
	)

	introText := lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		h2,
		"",
		para1,
		"",
		para2,
		"",
		styles.MutedStyle.Render("View my journey → navigate to Experience"),
	)

	var body string
	art := components.AboutTerminal()
	if art != "" && contentWidth >= 90 {
		body = lipgloss.JoinHorizontal(lipgloss.Top, introText, "   ", art)
	} else {
		body = introText
	}

	return styles.ContentStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			h1,
			"",
			oneLiner,
			"",
			ctaLine,
			"",
			"",
			lipgloss.JoinHorizontal(lipgloss.Left, strings.Join(stats, "")),
			"",
			"",
			"",
			body,
		),
	)
}
