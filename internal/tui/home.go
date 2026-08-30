package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderHomeContent renders the home/welcome screen with a responsive hero.
func (m *App) renderHomeContent() string {
	contentWidth := m.width - sidebarWidth - 2

	sig := components.Signature(contentWidth)
	nameBlock := styles.HeroTitleStyle.Render(strings.ToUpper(m.profile.Name))
	if m.height >= 24 && contentWidth >= 90 {
		nameBlock = m.renderHeroName(contentWidth)
	}
	cta := "[ ENTER ] Explore Portfolio"

	hero := components.Hero(sig, nameBlock, m.profile.Title, cta, contentWidth)

	// Social links stay below the hero for quick access.
	socialLines := make([]string, 0, len(m.socials))
	for _, s := range m.socials {
		socialLines = append(socialLines, styles.LinkStyle.Render(fmt.Sprintf("→ %s: %s", s.Name, s.URL)))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		hero,
		"",
		strings.Join(socialLines, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderHeroName renders the hero name: first token as large FIGlet text when
// it fits, remainder as normal hero text. Falls back to plain hero text.
func (m *App) renderHeroName(width int) string {
	fields := strings.Fields(strings.ToUpper(m.profile.Name))
	if len(fields) == 0 {
		return styles.HeroTitleStyle.Render("")
	}

	fig := components.Figlet(fields[0], width)
	if len(fields) == 1 {
		return fig
	}
	return fig + "\n" + styles.HeroTitleStyle.Render(strings.Join(fields[1:], " "))
}
