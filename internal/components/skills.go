package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// skillIcons maps tool names to compact terminal glyphs (universal, no emoji).
var skillIcons = map[string]string{
	"React":        "⚛",
	"Next.js":      "▲",
	"Node.js":      "⬡",
	"TypeScript":   "◈",
	"PostgreSQL":   "◉",
	"Tailwind CSS": "◆",
	"PHP":          "◦",
	"Laravel":      "◈",
	"WordPress":    "◆",
	"Elementor":    "◦",
	"Astra":        "◦",
	"Git":          "◉",
	"GitHub":       "◉",
	"Bootstrap":    "◆",
	"Vercel":       "▲",
	"JavaScript":   "◈",
}

func skillIcon(name string) string {
	if icon, ok := skillIcons[name]; ok {
		return icon
	}
	return "·"
}

// SkillChip renders one skill with an icon prefix — no box borders.
func SkillChip(name string) string {
	icon := styles.MutedStyle.Render(skillIcon(name))
	label := styles.NormalStyle.Render(name)
	return icon + " " + label
}

// SkillGrid lays out skills in horizontal rows that wrap at width.
func SkillGrid(names []string, width int) string {
	if width < 20 {
		width = 20
	}
	sep := "  "
	var rows []string
	current := ""
	for _, name := range names {
		chip := SkillChip(name)
		if current == "" {
			current = chip
			continue
		}
		candidate := current + sep + chip
		if lipgloss.Width(candidate) > width {
			rows = append(rows, current)
			current = chip
		} else {
			current = candidate
		}
	}
	if current != "" {
		rows = append(rows, current)
	}
	return strings.Join(rows, "\n")
}
