package tui

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
)

const maxContentWidth = 84

// Cells reserved beside the inner content column so the assembled shell
// (nav + rule + focus rail + content + scrollbar) never exceeds the terminal.
const (
	focusRailCells  = 2 // "· " / "▎ "
	minBodyMargin   = 2
	bodyFrameChrome = 16 + 1 + 1 + focusRailCells + components.ScrollbarWidth // nav + │ + gap + rail + bar
)

// contentWidth returns the inner text column (right of sidebar, focus rail,
// and scrollbar gutter). Capped so the full shell stays on-screen.
func (m *App) contentWidth() int {
	w := m.width - 2*minBodyMargin - bodyFrameChrome
	if w > maxContentWidth {
		return maxContentWidth
	}
	if w < 12 {
		return 12
	}
	return w
}

// frameWidth is nav + rule + focus rail + content + scrollbar (before centering).
func (m *App) frameWidth() int {
	return bodyFrameChrome + m.contentWidth()
}

// composeBodyFrame lays the nav+content shell into the body region. The shell is
// horizontally CENTERED (so on wide terminals content does not pile up on the
// left with an empty right side) and top-aligned with a small top padding, so
// the nav column keeps a stable position. Overlong content is already clipped
// to the viewport in renderBodyCached.
func (m *App) composeBodyFrame(shell string, bodyH int) string {
	lines := strings.Split(shell, "\n")
	if len(lines) > bodyH {
		lines = lines[:bodyH]
	}

	shellW := m.frameWidth()
	if shellW > m.width-2*minBodyMargin {
		shellW = m.width - 2*minBodyMargin
		if shellW < 1 {
			shellW = 1
		}
	}
	pad := (m.width - shellW) / 2
	if pad < minBodyMargin {
		pad = minBodyMargin
	}

	topPad := m.bodyTopPad()

	// Record geometry for mouse hit-testing.
	m.shellLeft = pad
	m.shellWidth = shellW
	// Scrollbar is the rightmost column of the shell.
	m.scrollBarX = pad + shellW - 1
	m.scrollBarTop = m.bodyTop + topPad
	m.scrollBarH = len(lines)
	if m.scrollBarH > bodyH-topPad {
		m.scrollBarH = bodyH - topPad
	}

	// Top-align the block with responsive top padding for breathing room.
	out := make([]string, bodyH)
	for i := 0; i < topPad && i < len(out); i++ {
		out[i] = strings.Repeat(" ", pad)
	}
	for i, line := range lines {
		idx := topPad + i
		if idx >= len(out) {
			break
		}
		// Left-align within the fixed shell, then clip to the terminal so a
		// long line can never wrap and break the nav/content columns.
		out[idx] = components.FitLine(strings.Repeat(" ", pad)+line, m.width)
	}
	return strings.Join(out, "\n")
}

func stripANSI(s string) string {
	var out strings.Builder
	in := false
	for _, r := range s {
		if r == '\x1b' {
			in = true
			continue
		}
		if in {
			if r == 'm' {
				in = false
			}
			continue
		}
		out.WriteRune(r)
	}
	return out.String()
}

// centerBlock kept for tests.
func (m *App) centerBlock(block string, areaH int) string {
	return m.composeBodyFrame(block, areaH)
}
