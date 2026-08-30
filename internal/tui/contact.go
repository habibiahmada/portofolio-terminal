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
		"or just want to say hi — I'm always open to connecting. " +
		"Feel free to reach out through any of the channels below."

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
			icon:  "✉",
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
		rows = append(rows, fmt.Sprintf("%s %s  %s", icon, label, value))
	}

	// Response time note
	note := styles.MutedStyle.Render("Typically reply within 24–48 hours · Remote (WIB timezone)")

	return strings.Join(append([]string{heading}, rows...), "\n") + "\n\n" + note
}

// renderContactSocials — social media links in a structured list.
func (m *App) renderContactSocials(cw int) string {
	heading := styles.SectionTitleStyle.Render("▸ Socials")

	socialIcons := map[string]string{
		"GitHub":   "◈",
		"LinkedIn": "◆",
		"Instagram":"◦",
		"Email":    "✉",
	}

	var rows []string
	for _, s := range m.socials {
		icon := socialIcons[s.Name]
		if icon == "" {
			icon = "·"
		}
		label := styles.NormalStyle.Render(s.Name)
		link := styles.LinkStyle.Render(s.URL)
		rows = append(rows, fmt.Sprintf("%s  %s  %s",
			styles.MutedStyle.Render(icon),
			label,
			link,
		))
	}
	return strings.Join(append([]string{heading}, rows...), "\n")
}

// renderContactCTA — a closing call to action.
func (m *App) renderContactCTA(cw int) string {
	cta := styles.NormalStyle.Render(
		"Open to freelance projects, full-time opportunities, " +
			"and interesting collaborations. Let's build something great together.",
	)
	prompt := styles.PromptStyle.Render(">_ " + m.profile.Website)
	return cta + "\n" + prompt
}
