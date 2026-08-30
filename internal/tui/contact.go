package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderContactContent builds the Contact page with availability status,
// contact methods, social links, and a call to action.
func (m *App) renderContactContent() string {
	cw := m.contentWidth()
	parts := []string{
		m.renderContactHero(cw),
		"",
		m.renderContactMethods(cw),
		"",
		m.renderContactSocials(cw),
		"",
		m.renderContactCTA(cw),
	}
	return strings.Join(parts, "\n")
}

// renderContactHero — title, availability badge, and a brief intro.
func (m *App) renderContactHero(cw int) string {
	title := styles.HeroTitleStyle.Render("Let's Talk")
	badge := styles.SuccessStyle.Render("● " + m.profile.Availability)

	intro := "Whether you have a project idea, need a developer on your team, " +
		"or just want to say hi, I am always open to connecting. " +
		"Pick any channel below and I will get back to you as soon as I can."

	lines := []string{title, badge, ""}
	lines = append(lines, components.WrapText(intro, cw-2)...)
	return strings.Join(lines, "\n")
}

// renderContactMethods — direct contact channels with labels and details.
func (m *App) renderContactMethods(cw int) string {
	heading := styles.SectionTitleStyle.Render("▸ Contact Methods")

	methods := []struct {
		icon  string
		label string
		value string
		link  string
	}{
		{
			icon:  "@",
			label: "Email",
			value: m.profile.Email,
			link:  "mailto:" + m.profile.Email,
		},
		{
			icon:  "◉",
			label: "Website",
			value: m.profile.Website,
			link:  "https://" + m.profile.Website,
		},
	}

	var rows []string
	for _, mt := range methods {
		icon := styles.MutedStyle.Render(mt.icon)
		label := styles.NormalStyle.Render(mt.label)
		value := styles.LinkStyle.Render(mt.value)
		line := fmt.Sprintf("%s %s  %s", icon, label, value)
		for _, wl := range components.WrapText(line, cw) {
			rows = append(rows, wl)
		}
	}

	note := styles.MutedStyle.Render("I typically reply within 24-48 hours. Based in Indonesia (WIB timezone), open to remote collaboration.")
	noteLines := components.WrapText(note, cw)

	return strings.Join(append([]string{heading}, rows...), "\n") + "\n\n" + strings.Join(noteLines, "\n")
}

// renderContactSocials — social media links in a structured list.
func (m *App) renderContactSocials(cw int) string {
	heading := styles.SectionTitleStyle.Render("▸ Socials")

	socialIcons := map[string]string{
		"GitHub":   "◈",
		"LinkedIn": "◆",
		"Instagram":"◦",
		"Email":    "@",
	}

	var rows []string
	for _, s := range m.socials {
		icon := socialIcons[s.Name]
		if icon == "" {
			icon = "·"
		}
		label := styles.NormalStyle.Render(s.Name)
		link := styles.LinkStyle.Render(s.URL)
		line := fmt.Sprintf("%s  %s  %s",
			styles.MutedStyle.Render(icon),
			label,
			link,
		)
		for _, wl := range components.WrapText(line, cw) {
			rows = append(rows, wl)
		}
	}
	return strings.Join(append([]string{heading}, rows...), "\n")
}

// renderContactCTA — a closing call to action.
func (m *App) renderContactCTA(cw int) string {
	cta := styles.NormalStyle.Render(
		"I am open to freelance projects, full-time roles, and interesting collaborations. " +
			"If you have an idea worth building, I would love to hear about it.",
	)
	ctaLines := components.WrapText(cta, cw)
	prompt := styles.PromptStyle.Render(">_ " + m.profile.Website)
	promptLines := components.WrapText(prompt, cw)
	return strings.Join(ctaLines, "\n") + "\n" + strings.Join(promptLines, "\n")
}
