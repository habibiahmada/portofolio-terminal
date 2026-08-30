package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

const navRailWidth = 16

// SidebarItem is a single navigation entry shown in the side rail.
type SidebarItem struct {
	Key  string
	Name string
}

// NavRail renders a fixed-width vertical menu with key hints at the bottom.
func NavRail(items []SidebarItem, selectedIndex int, width int) string {
	if width <= 0 {
		width = navRailWidth
	}

	lines := make([]string, 0, len(items)+5)
	for i, it := range items {
		if it.Name == "" {
			continue
		}
		label := it.Name
		if i == selectedIndex {
			label = styles.NavActiveStyle.Render("[" + it.Name + "]")
		} else {
			label = styles.NavItemStyle.Render(" " + it.Name)
		}
		lines = append(lines, label)
	}

	lines = append(lines, "")
	lines = append(lines,
		styles.MutedStyle.Render("↑↓ screen"),
		styles.MutedStyle.Render("j/k scroll"),
		styles.MutedStyle.Render("Enter open"),
	)

	block := lipgloss.JoinVertical(lipgloss.Left, lines...)
	return lipgloss.NewStyle().Width(width).Render(block)
}

// VerticalRule renders a column divider for the shell height.
func VerticalRule(height int) string {
	if height < 1 {
		height = 1
	}
	line := lipgloss.NewStyle().Foreground(styles.ColorBorder).Render("│")
	lines := make([]string, height)
	for i := range lines {
		lines[i] = line
	}
	return strings.Join(lines, "\n")
}

// JoinShell combines nav rail, divider, and content into one top-aligned block.
// The nav rail is vertically centered within the shell height while content is
// kept top-aligned (it is scrollable).
func JoinShell(rail, content string, height int) string {
	content = PadRailHeight(content, height)
	rail = VAlignBlock(rail, height)
	rule := VerticalRule(height)
	return lipgloss.JoinHorizontal(lipgloss.Top, rail, rule, " ", content)
}

// VAlignBlock vertically centers a block within exactly targetLines so the
// navigation menu sits in the middle of the available height.
func VAlignBlock(block string, targetLines int) string {
	lines := strings.Split(block, "\n")
	if len(lines) > targetLines {
		return strings.Join(lines[:targetLines], "\n")
	}
	if len(lines) == 0 {
		lines = []string{""}
	}
	pad := targetLines - len(lines)
	top := pad / 2
	bottom := pad - top
	out := make([]string, 0, targetLines)
	for i := 0; i < top; i++ {
		out = append(out, "")
	}
	out = append(out, lines...)
	for i := 0; i < bottom; i++ {
		out = append(out, "")
	}
	return strings.Join(out, "\n")
}

// NavRailWidth returns the default side rail column width.
func NavRailWidth() int {
	return navRailWidth
}

// PadRailHeight pads or trims a block to exactly targetLines.
func PadRailHeight(block string, targetLines int) string {
	lines := strings.Split(block, "\n")
	if len(lines) > targetLines {
		return strings.Join(lines[:targetLines], "\n")
	}
	for len(lines) < targetLines {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}

// Sidebar is deprecated — use NavRail.
func Sidebar(items []SidebarItem, selectedIndex int, activeKey string, height int) string {
	_ = activeKey
	_ = height
	return NavRail(items, selectedIndex, navRailWidth)
}
