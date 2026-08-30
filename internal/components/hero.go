package components

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Hero builds the home welcome block: signature art (if any), the large name,
// the title, and a call-to-action line. Content is centered within width.
func Hero(signature, name, title, cta string, width int) string {
	blocks := make([]string, 0, 5)
	if signature != "" {
		blocks = append(blocks, signature, "")
	}
	blocks = append(blocks,
		styles.HeroTitleStyle.Render(name),
		styles.SubtitleStyle.Render(title),
		"",
		styles.HeroCTAStyle.Render(cta),
	)

	content := lipgloss.JoinVertical(lipgloss.Center, blocks...)
	return lipgloss.Place(width, 0, lipgloss.Center, lipgloss.Top, content)
}
