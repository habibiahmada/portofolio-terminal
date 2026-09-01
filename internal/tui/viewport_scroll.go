package tui

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
)

// ensureSelectionInViewport adjusts contentOffset so the given line range stays
// visible inside the content viewport.
func (m *App) ensureSelectionInViewport(startLine, endLine, totalLines int) {
	if startLine < 0 || endLine < startLine {
		return
	}

	maxH := m.contentHeight() - 1
	if maxH < 3 {
		maxH = 3
	}

	if startLine < m.contentOffset {
		m.contentOffset = startLine
	} else if endLine >= m.contentOffset+maxH {
		m.contentOffset = endLine - maxH + 1
	}

	if totalLines > maxH {
		scrollMax := totalLines - maxH
		if m.contentOffset > scrollMax {
			m.contentOffset = scrollMax
		}
		if m.contentOffset < 0 {
			m.contentOffset = 0
		}
	}

	m.invalidateLayoutCache()
}

func (m *App) homeContentTotalLines() int {
	w := m.contentWidth()
	content := components.ReflowBlock(m.renderHomeContent(), w)
	return lineCount(content)
}

func lineCount(s string) int {
	if s == "" {
		return 0
	}
	return len(strings.Split(s, "\n"))
}
