package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderAboutContent builds a focused personal profile: developer identity,
// background narrative, key career milestones, and engineering philosophy.
func (m *App) renderAboutContent() string {
	cw := m.contentWidth()
	parts := []string{
		m.renderAboutHero(cw),
		"",
		m.renderAboutStats(cw),
		"",
		m.renderAboutPhilosophy(cw),
		"",
		m.renderAboutGuide(cw),
	}
	return strings.Join(parts, "\n")
}

// renderAboutHero — name, title, location, availability badge, and full bio.
func (m *App) renderAboutHero(cw int) string {
	name := styles.HeroTitleStyle.Render(m.profile.Name)
	title := styles.SubtitleStyle.Render(m.profile.Title + " · " + m.profile.Location)
	badge := styles.SuccessStyle.Render("● " + m.profile.Availability)

	bio := "I am a full-stack web developer passionate about building scalable, " +
		"user-centric applications and solid architectural foundations. Having graduated " +
		"in Software Engineering from SMKN 1 Karawang, I currently build web applications " +
		"at PT Webekspres Teknologi Indonesia. My experience spans crafting responsive client " +
		"interfaces, architecting custom CMS platforms, and deploying production-ready full-stack features."

	lines := []string{name, title, "", badge, ""}
	lines = append(lines, components.WrapText(bio, cw-2)...)
	return strings.Join(lines, "\n")
}

// renderAboutStats — headline metrics and key career milestones.
func (m *App) renderAboutStats(cw int) string {
	heading := styles.SectionTitleStyle.Render("▸ Core Metrics & Accolades")

	// Stats pills
	stats := make([]string, 0, len(m.profile.Stats))
	for _, s := range m.profile.Stats {
		stats = append(stats,
			styles.PrimaryText.Render(s.Value)+" "+styles.MutedStyle.Render(s.Label),
		)
	}
	statsLine := "  " + strings.Join(stats, "   ·   ")

	// Key Awards / Milestones
	milestones := []struct {
		badge string
		title string
		desc  string
	}{
		{
			badge: "[Top 15 Capstone]",
			title: "Best Capstone Project — Coding Camp powered by DBS Foundation (2025)",
			desc:  "CultureConnect platform selected among top 15 capstone teams nationally.",
		},
		{
			badge: "[Country Winner]",
			title: "Intel AI for Youth — National Award Winner (2025)",
			desc:  "Smartfarm AI (Agrify) recognized for real-world impact in agricultural intelligence.",
		},
	}

	var rows []string
	for _, ml := range milestones {
		tag := styles.BadgeAccentStyle.Render(ml.badge)
		wrappedTitle := components.WrapText(ml.title, cw-len(ml.badge)-6)
		if len(wrappedTitle) == 0 {
			wrappedTitle = []string{ml.title}
		}

		var blockLines []string
		blockLines = append(blockLines, fmt.Sprintf("  %s %s", tag, styles.NormalStyle.Render(wrappedTitle[0])))
		for _, tl := range wrappedTitle[1:] {
			blockLines = append(blockLines, "     "+styles.NormalStyle.Render(tl))
		}

		wrappedDesc := components.WrapText(ml.desc, cw-6)
		for _, dl := range wrappedDesc {
			blockLines = append(blockLines, "     "+styles.MutedStyle.Render(dl))
		}

		rows = append(rows, strings.Join(blockLines, "\n"))
	}

	lines := []string{heading, statsLine, ""}
	lines = append(lines, rows...)
	return strings.Join(lines, "\n")
}

// renderAboutPhilosophy — core engineering values and software development approach.
func (m *App) renderAboutPhilosophy(cw int) string {
	heading := styles.SectionTitleStyle.Render("▸ Engineering Philosophy")

	principles := []struct {
		title string
		desc  string
	}{
		{
			title: "1. Clean Architecture & Predictability",
			desc:  "Write clear, modular, and maintainable code with honest boundaries and typed contracts.",
		},
		{
			title: "2. Performance & User Experience First",
			desc:  "Optimize for snappy load times, accessible interfaces, and fluid interactions across devices.",
		},
		{
			title: "3. End-to-End Ownership",
			desc:  "Take features from requirement scoping and data schema design all the way to CI/CD and production deploys.",
		},
	}

	var rows []string
	for _, p := range principles {
		title := styles.PrimaryText.Render(p.title)
		desc := styles.NormalStyle.Render(strings.Join(components.WrapText(p.desc, cw-6), "\n     "))
		rows = append(rows, fmt.Sprintf("  %s\n     %s", title, desc))
	}

	return strings.Join(append([]string{heading}, rows...), "\n\n")
}

// renderAboutGuide — hints directing to detailed sections.
func (m *App) renderAboutGuide(cw int) string {
	guideLines := components.WrapText("▸ Explore Experience for career timeline, Skills for full tech stack, and Projects for in-depth case studies.", cw-2)
	var rendered []string
	for _, gl := range guideLines {
		rendered = append(rendered, styles.MutedStyle.Render(gl))
	}
	return strings.Join(rendered, "\n")
}
