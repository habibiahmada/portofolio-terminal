package components

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// SidebarItem is a single navigation entry shown in the sidebar.
type SidebarItem struct {
	Key  string // stable identifier, e.g. "projects"
	Name string // display name, e.g. "Projects"
}

// Sidebar renders the navigation column. The item at selectedIndex is marked
// with ">", and the item matching activeKey is marked with "▸". Pass -1 for
// selectedIndex on non-menu screens. Empty items are ignored.
func Sidebar(items []SidebarItem, selectedIndex int, activeKey string, height int) string {
	lines := make([]string, 0, len(items))
	for i, it := range items {
		if it.Key == "" || it.Name == "" {
			continue
		}
		switch {
		case i == selectedIndex:
			lines = append(lines, styles.SidebarItemSelectedStyle.Render("> "+it.Name))
		case it.Key == activeKey:
			lines = append(lines, styles.SidebarItemSelectedStyle.Render("▸ "+it.Name))
		default:
			lines = append(lines, styles.SidebarItemStyle.Render("  "+it.Name))
		}
	}

	sb := lipgloss.JoinVertical(lipgloss.Left, lines...)
	return styles.BorderStyle.Width(20).Height(height).Render(sb)
}
