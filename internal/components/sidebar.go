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

// navCursor cycles a 1-cell "activity" bar beside the active menu item.
var navCursor = []string{"▌", "▍", "▎", "▏", "▎", "▍"}

// NavRail renders a fixed-width vertical menu with key hints at the bottom.
// The active item carries an animated cursor so the whole app shows life even
// while the content body is idle. When navFocused is true the active item is a
// solid highlighted bar; otherwise it is a muted selection to show the content
// zone currently has focus.
func NavRail(items []SidebarItem, selectedIndex int, width, frame int, navFocused bool) string {
	if width <= 0 {
		width = navRailWidth
	}

	lines := make([]string, 0, len(items)+5)
	for i, it := range items {
		if it.Name == "" {
			continue
		}
		if i == selectedIndex {
			if navFocused {
				marker := navCursor[frame%len(navCursor)]
				label := styles.NavActiveStyle.Bold(true).Render(marker + " " + it.Name)
				pad := width - lipgloss.Width(label)
				if pad < 0 {
					pad = 0
				}
				bar := styles.NavActiveStyle.Bold(true).Render(strings.Repeat(" ", pad))
				lines = append(lines, label+bar)
			} else {
				label := styles.NavSelectedInactive.Render("  " + it.Name)
				lines = append(lines, label)
			}
		} else {
			label := styles.NavItemStyle.Render("  " + it.Name + " ")
			lines = append(lines, label)
		}
	}

	lines = append(lines, "")
	if navFocused {
		lines = append(lines,
			styles.NavZoneHighlight.Render("◤ NAV"),
			styles.MutedStyle.Render("↑↓ screen"),
			styles.MutedStyle.Render("→ focus"),
		)
	} else {
		lines = append(lines,
			styles.NavZoneHighlight.Render("◤ SCREEN"),
			styles.MutedStyle.Render("↑↓ scroll"),
			styles.MutedStyle.Render("←/Esc to nav"),
		)
	}

	block := lipgloss.JoinVertical(lipgloss.Left, lines...)
	return FitBlock(block, width)
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
// Both columns start at the top so the vertical rule stays a continuous line
// next to the menu and the NAV hints, instead of floating mid-screen.
func JoinShell(rail, content string, height int) string {
	content = PadRailHeight(content, height)
	rail = PadRailHeight(rail, height)
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
	return NavRail(items, selectedIndex, navRailWidth, 0, true)
}
