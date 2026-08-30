package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderContactContent() string {
	lines := []string{
		styles.SectionTitleStyle.Render("Contact"),
		styles.SuccessStyle.Render("● " + m.profile.Availability),
		"",
		styles.LinkStyle.Render(m.profile.Email),
		styles.MutedStyle.Render("Remote (WIB) · reply within 48h"),
	}

	for _, s := range m.socials {
		lines = append(lines, fmt.Sprintf("%s  %s",
			styles.MutedStyle.Render(s.Name),
			styles.LinkStyle.Render(s.URL),
		))
	}
	return strings.Join(lines, "\n")
}
