package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

const (
	MascotStateNormal = 0
	MascotStateBlink  = 1
	MascotStateAngry  = 2

	mascotWidth = 14
	mascotInner = 12
	mascotSide  = "║"
)

func mascotTopLine() string    { return "╔" + strings.Repeat("═", mascotInner) + "╗" }
func mascotBottomLine() string { return "╚" + strings.Repeat("═", mascotInner) + "╝" }

// MascotFigure returns a compact, dense front-facing CRT monitor mascot with
// expressive interactive states (normal, blinking/happy, and angry with ╬ vein mark).
func MascotFigure(frame int, state ...int) string {
	st := MascotStateNormal
	if len(state) > 0 {
		st = state[0]
	}

	rows := []string{
		mascotAntennaRow(frame, st),
		mascotTopBorderRow(st),
		mascotEyesRow(frame, st),
		mascotMouthRow(frame, st),
		mascotBottomBorderRow(st),
		mascotFeetRow(st),
	}

	for i, r := range rows {
		rows[i] = padMascotLine(r)
	}

	return strings.Join(rows, "\n")
}

// mascotAntennaRow renders the top antenna with blinking LED or anger mark (╬ / ⑊).
func mascotAntennaRow(frame int, state int) string {
	if state == MascotStateAngry {
		// Anger marks in bold red: ╬, ⑊, ⁑ (anime vein pulse without emojis).
		angerMarks := []string{"╬", "⑊", "⁑", "╬"}
		mark := styles.PrimaryText.Render(angerMarks[frame%len(angerMarks)])
		dome := styles.MascotEye.Render("╭─!─╮")
		// Dome stays in exact center (cols 4..8), anger mark at col 10.
		// 4 spaces + 5 (dome) + 1 space + 1 (mark) + 3 spaces = 14 visible cells.
		return "    " + dome + " " + mark + "   "
	}

	if state == MascotStateBlink {
		star := styles.MascotBright.Render("★")
		dome := styles.MascotThin.Render("╭─") + star + styles.MascotThin.Render("─╮")
		// 4 spaces + 5 (dome) + 5 spaces = 14 visible cells.
		return "    " + dome + "     "
	}

	glyphs := []string{"·", "○", "●", "○"}
	stylesList := []lipgloss.Style{styles.MascotDim, styles.MascotMid, styles.MascotBright, styles.MascotMid}
	glyph := stylesList[frame%4].Render(glyphs[frame%4])
	dome := styles.MascotThin.Render("╭─") + glyph + styles.MascotThin.Render("─╮")
	// 4 spaces + 5 (dome) + 5 spaces = 14 visible cells.
	return "    " + dome + "     "
}

func mascotTopBorderRow(state int) string {
	if state == MascotStateAngry {
		return styles.MascotEye.Render(mascotTopLine())
	}
	return styles.MascotBold.Render(mascotTopLine())
}

func mascotBottomBorderRow(state int) string {
	if state == MascotStateAngry {
		return styles.MascotEye.Render(mascotBottomLine())
	}
	return styles.MascotBold.Render(mascotBottomLine())
}

func mascotEyesRow(frame int, state int) string {
	inner := mascotEyesInner(frame, state)
	return mascotSideRow(inner, state)
}

func mascotMouthRow(frame int, state int) string {
	inner := mascotMouthInner(frame, state)
	return mascotSideRow(inner, state)
}

func mascotSideRow(inner string, state int) string {
	inner = padInner(inner)
	sideStyle := styles.MascotBold
	if state == MascotStateAngry {
		sideStyle = styles.MascotEye
	}
	return sideStyle.Render(mascotSide) + inner + sideStyle.Render(mascotSide)
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

func mascotEyesInner(frame int, state int) string {
	if state == MascotStateAngry {
		// Sharp furious angry eyes (alternating ◣ ◢ and \ /).
		if frame%2 == 0 {
			eyeL := styles.MascotEye.Render("◣")
			eyeR := styles.MascotEye.Render("◢")
			return eyeL + strings.Repeat(" ", 6) + eyeR
		}
		eyeL := styles.MascotEye.Render("\\")
		eyeR := styles.MascotEye.Render("/")
		return eyeL + strings.Repeat(" ", 6) + eyeR
	}

	if state == MascotStateBlink {
		// Interactive blinking / happy expressions on click.
		switch frame % 4 {
		case 0:
			// Blink: ─ ─
			return styles.MascotBright.Render("─" + strings.Repeat(" ", 6) + "─")
		case 1:
			// Happy squint: ^ ^
			return styles.MascotBright.Render("^" + strings.Repeat(" ", 6) + "^")
		case 2:
			// Wink: ^ ─
			return styles.MascotBright.Render("^") + strings.Repeat(" ", 6) + styles.MascotBright.Render("─")
		default:
			// Sparkle / wide eyes: ★ ★
			return styles.MascotBright.Render("★" + strings.Repeat(" ", 6) + "★")
		}
	}

	// Normal idle state: natural blink on frame 9.
	if frame%10 == 9 {
		return styles.MascotThin.Render("─" + strings.Repeat(" ", 6) + "─")
	}

	// Normal centered eyes.
	face := styles.MascotEye.Render("●") +
		strings.Repeat(" ", 6) +
		styles.MascotEye.Render("●")
	return face
}

func mascotMouthInner(frame int, state int) string {
	if state == MascotStateAngry {
		// Grumpy frown or gnashing teeth.
		if frame%2 == 0 {
			return styles.MascotEye.Render("╭────╮")
		}
		return styles.MascotEye.Render("╭─┴┴─╮")
	}

	if state == MascotStateBlink {
		// Happy open smile or cute mouth.
		if frame%2 == 0 {
			return styles.MascotBright.Render("╰▽╯")
		}
		return styles.MascotBright.Render("╰──╯")
	}

	mouths := []string{"╰──╯", "╰────╯", "╰──╯", "╰────╯"}
	mouth := styles.MascotMid.Render(mouths[frame%len(mouths)])
	return mouth
}

func mascotFeetRow(state int) string {
	if state == MascotStateAngry {
		return styles.MascotEye.Render("  ╰─┯────┯─╯  ")
	}
	return styles.MascotThin.Render("  ╰─┴────┴─╯  ")
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
// figure right-aligned in a given width.
func Mascot(frame, width int) string {
	if width < 60 {
		return ""
	}
	figure := MascotFigure(frame)
	return lipgloss.NewStyle().Width(width - 4).Align(lipgloss.Right).MarginRight(4).Render(figure)
}
