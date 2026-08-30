package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

const (
	mascotWidth = 14
	mascotInner = 12
	mascotSide  = "║"
)

func mascotTopLine() string    { return "╔" + strings.Repeat("═", mascotInner) + "╗" }
func mascotBottomLine() string { return "╚" + strings.Repeat("═", mascotInner) + "╝" }

// MascotFigure returns a front-facing monitor-face icon with layered thin bezels,
// bold outer frame, and saturation steps so it reads clearly in the margin.
func MascotFigure(frame int) string {
	rows := []string{
		mascotAntennaRow(frame),
		mascotThinEdge("╭", "─", "╮"),
		styles.MascotBold.Render(mascotTopLine()),
		mascotSideRow(mascotEyesInner(frame)),
		mascotSideRow(styles.MascotThin.Render(strings.Repeat("─", mascotInner))),
		mascotSideRow(mascotMouthInner(frame)),
		styles.MascotBold.Render(mascotBottomLine()),
		mascotThinEdge("╰", "─", "╯"),
	}

	for i, r := range rows {
		rows[i] = padMascotLine(r)
	}

	return strings.Join(rows, "\n")
}

func mascotAntennaRow(frame int) string {
	glyphs := []string{"·", "○", "●", "○"}
	stylesList := []lipgloss.Style{styles.MascotDim, styles.MascotMid, styles.MascotBright, styles.MascotMid}
	glyph := stylesList[frame%4].Render(glyphs[frame%4])
	return centerMascot(glyph)
}

func mascotThinEdge(left, mid, right string) string {
	edge := styles.MascotThin.Render(left + strings.Repeat(mid, mascotInner) + right)
	return centerMascot(edge)
}

func mascotSideRow(inner string) string {
	inner = padInner(inner)
	return styles.MascotBold.Render(mascotSide) + inner + styles.MascotBold.Render(mascotSide)
}

func padInner(content string) string {
	w := lipgloss.Width(content)
	if w >= mascotInner {
		return content
	}
	left := (mascotInner - w) / 2
	right := mascotInner - w - left
	return strings.Repeat(" ", left) + content + strings.Repeat(" ", right)
}

func mascotEyesInner(frame int) string {
	if frame%10 == 9 {
		return styles.MascotThin.Render(strings.Repeat("─", mascotInner))
	}

	// Fixed spacing keeps both eyes centered and facing forward.
	face := styles.MascotEye.Render("●") +
		strings.Repeat(" ", 6) +
		styles.MascotEye.Render("●")
	return padInner(face)
}

func mascotMouthInner(frame int) string {
	mouths := []string{"╰─╯", "╰──╯", "╰─╯", "╰──╯"}
	mouth := styles.MascotMid.Render(mouths[frame%len(mouths)])
	return padInner(mouth)
}

func centerMascot(segment string) string {
	pad := mascotWidth - lipgloss.Width(segment)
	if pad < 0 {
		return segment
	}
	left := pad / 2
	return strings.Repeat(" ", left) + segment + strings.Repeat(" ", pad-left)
}

func padMascotLine(line string) string {
	if w := lipgloss.Width(line); w < mascotWidth {
		return line + strings.Repeat(" ", mascotWidth-w)
	}
	return line
}

// Mascot animates a small terminal monitor face. Kept for callers that want the
// figure right-aligned in a given width; overlay placement lives in the TUI so it
// can sit in the body margin without covering content.
func Mascot(frame, width int) string {
	if width < 60 {
		return ""
	}
	figure := MascotFigure(frame)
	return lipgloss.NewStyle().Width(width - 4).Align(lipgloss.Right).MarginRight(4).Render(figure)
}
