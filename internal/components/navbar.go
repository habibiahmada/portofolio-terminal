package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// NavItem is one entry in the horizontal navigation bar.
type NavItem struct {
	Key  string
	Name string
}

// NavBar renders a compact horizontal menu centered within width.
func NavBar(items []NavItem, selectedIndex int, width int) string {
	if len(items) == 0 {
		return ""
	}

	parts := make([]string, 0, len(items))
	for i, it := range items {
		label := it.Name
		if i == selectedIndex {
			label = styles.NavActiveStyle.Render("[" + it.Name + "]")
		} else {
			label = styles.NavItemStyle.Render(it.Name)
		}
		parts = append(parts, label)
	}

	line := strings.Join(parts, styles.MutedStyle.Render(" · "))
	return lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center, line)
}

// Masthead renders a single centered wordmark line.
func Masthead(wordmark, title string, width int) string {
	mark := styles.HeaderWordmark.Render(wordmark) + styles.HeaderDot.Render(".")
	meta := styles.MutedStyle.Render(" · " + title)
	line := mark + meta
	return lipgloss.Place(width, 1, lipgloss.Center, lipgloss.Center, line)
}
