package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// ShimmerText renders text with a slow wave of emphasis sweeping across it.
// One "brightened" character travels left→right (wrapping), and the neighbours
// gently trail behind it. It is a subtle, calm effect — characters never move,
// only their weight/brightness changes — so it does not distract or jitter the
// layout. Used for a headline on the Home screen.
func ShimmerText(text string, frame int, base, highlight lipgloss.Style) string {
	if text == "" {
		return ""
	}
	n := len([]rune(text))
	if n == 0 {
		return ""
	}
	// a three-cell wave: bright, mid, dim leading the travel.
	period := n + 2
	head := frame % period
	runes := []rune(text)

	var b strings.Builder
	for i, r := range runes {
		dist := (head - i + period) % period
		var s lipgloss.Style
		switch {
		case dist == 0:
			s = highlight
		case dist == 1:
			s = base.Copy().Bold(true)
		case dist == 2:
			s = base.Copy().Bold(false).Foreground(styles.ColorSecondary)
		default:
			s = base
		}
		b.WriteString(s.Render(string(r)))
	}
	return b.String()
}
