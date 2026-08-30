package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Card renders a bordered block containing the given lines stacked vertically.
// A blank line is inserted between cards so a list of cards reads cleanly.
func Card(lines ...string) string {
	content := strings.Join(lines, "\n")
	return styles.CardStyle.Render(content)
}

// Cards renders multiple cards separated by a blank line.
func Cards(cards ...string) string {
	filtered := make([]string, 0, len(cards))
	for _, c := range cards {
		if c != "" {
			filtered = append(filtered, c)
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, filtered...)
}
