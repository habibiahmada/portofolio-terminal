package tui

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderAboutContent() string {
	stats := make([]string, 0, len(m.profile.Stats))
	for _, s := range m.profile.Stats {
		stats = append(stats, s.Value+" "+s.Label)
	}

	bio := "Software Engineering graduate (SMKN 1 Karawang). Web Developer at PT Webekspres — client sites, CMS, full-stack features in production."

	lines := []string{
		styles.SectionTitleStyle.Render("Habibi Ahmad Aziz"),
		styles.MutedStyle.Render(m.profile.Title + " · " + m.profile.Location),
		styles.NormalStyle.Render(strings.Join(stats, " · ")),
		"",
		styles.NormalStyle.Render(componentsWrap(bio, m.contentWidth())),
		"",
		styles.MutedStyle.Render(m.profile.Employer + " · " + m.profile.School),
	}
	return strings.Join(lines, "\n")
}

func componentsWrap(text string, width int) string {
	return strings.Join(wrapWords(text, width), "\n")
}

func wrapWords(text string, width int) []string {
	if width < 20 {
		width = 20
	}
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}
	var lines []string
	line := words[0]
	for _, w := range words[1:] {
		if len(line)+1+len(w) > width {
			lines = append(lines, line)
			line = w
		} else {
			line += " " + w
		}
	}
	lines = append(lines, line)
	return lines
}
