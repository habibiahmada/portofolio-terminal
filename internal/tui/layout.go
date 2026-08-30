package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
)

const maxContentWidth = 76

// contentWidth returns page content width (right of sidebar + divider).
func (m *App) contentWidth() int {
	// rail + rule + gap + margin
	shellPad := components.NavRailWidth() + 3 + 4
	w := m.width - shellPad
	if w > maxContentWidth {
		return maxContentWidth
	}
	if w < 32 {
		return 32
	}
	return w
}

// composeBodyFrame center-aligns the shell vertically and horizontally so the
// content sits in the middle of the viewport with a stable nav position.
func (m *App) composeBodyFrame(shell string, bodyH int) string {
	lines := strings.Split(shell, "\n")
	if len(lines) > bodyH {
		lines = lines[:bodyH]
	}
	if len(lines) == 0 {
		lines = []string{""}
	}

	// Vertical centering padding above the content block.
	padTop := 0
	if len(lines) < bodyH {
		padTop = (bodyH - len(lines)) / 2
	}

	maxW := 0
	for _, line := range lines {
		if w := lipgloss.Width(line); w > maxW {
			maxW = w
		}
	}

	padX := (m.width - maxW) / 2
	if padX < 0 {
		padX = 0
	}
	contentOffsetX = padX
	contentOffsetY = padTop

	out := make([]string, 0, bodyH)
	for i := 0; i < padTop; i++ {
		out = append(out, "")
	}
	for _, line := range lines {
		out = append(out, strings.Repeat(" ", padX)+line)
	}
	for len(out) < bodyH {
		out = append(out, "")
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
