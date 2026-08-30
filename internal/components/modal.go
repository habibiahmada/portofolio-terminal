package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Modal renders a centered overlay box with the given title and lines on top
// of the current screen. Used for confirmations, help, and non-modal overlays.
func Modal(title string, lines []string, width, height int) string {
	content := []string{"  " + title + "  ", ""}
	content = append(content, lines...)
	box := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(styles.ColorPrimary).
		Padding(1, 2).
		Width(min(width-4, 50)).
		Render(strings.Join(content, "\n"))

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, box)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
