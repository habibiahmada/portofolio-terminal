package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// homeFeaturedLineMetrics maps each featured card index to its start/end line
// in the reflowed home content (for scroll-into-view).
func (m *App) homeFeaturedLineMetrics() (starts []int, ends []int) {
	w := m.contentWidth()
	featured := m.featuredProjects()
	if len(featured) == 0 {
		return nil, nil
	}

	heroLines := lineCount(components.ReflowBlock(m.renderHomeHero(w), w))

	twoCol := w >= 60
	inner := w - 4
	if twoCol {
		cardOuter := (w - 1) / 2
		inner = cardOuter - 4
	}
	if inner < 8 {
		inner = 8
	}

	heading := styles.SectionTitleStyle.Render("▸ Featured Work")
	sub := styles.MutedStyle.Render(
		"Select a card with arrow keys, press Enter to read the full case study, or press P to browse all projects.",
	)
	subLines := components.WrapText(sub, w)

	header := "\n" + heading + "\n" + strings.Join(subLines, "\n")
	if banner := m.renderPortfolioSyncBanner(w); banner != "" {
		header += "\n\n" + banner
	}
	header += "\n\n"

	cursor := heroLines + lineCount(components.ReflowBlock(header, w))

	starts = make([]int, len(featured))
	ends = make([]int, len(featured))
	blank := strings.Repeat(" ", inner)

	type rawCard struct {
		lines    []string
		selected bool
	}
	raws := make([]rawCard, 0, len(featured))
	for i, p := range featured {
		sel := m.focus == FocusContent && m.currentScreen == ScreenHome && i == m.selectedFeatured
		raws = append(raws, rawCard{buildCardLines(p, inner, sel), sel})
	}

	if !twoCol {
		for i, cd := range raws {
			starts[i] = cursor
			card := applyCardBorder(strings.Join(cd.lines, "\n"), inner, cd.selected)
			lines := lineCount(card)
			ends[i] = cursor + lines - 1
			cursor += lines
		}
		return starts, ends
	}

	for i := 0; i < len(raws); i += 2 {
		rowStart := cursor
		if i+1 < len(raws) {
			a, b := raws[i], raws[i+1]
			for len(a.lines) < len(b.lines) {
				a.lines = append(a.lines, blank)
			}
			for len(b.lines) < len(a.lines) {
				b.lines = append(b.lines, blank)
			}
			cardA := applyCardBorder(strings.Join(a.lines, "\n"), inner, a.selected)
			cardB := applyCardBorder(strings.Join(b.lines, "\n"), inner, b.selected)
			joined := lipgloss.JoinHorizontal(lipgloss.Top, cardA, " ", cardB)
			rowLines := lineCount(joined)
			starts[i] = rowStart
			ends[i] = rowStart + rowLines - 1
			starts[i+1] = rowStart
			ends[i+1] = rowStart + rowLines - 1
			cursor += rowLines
		} else {
			cd := raws[i]
			card := applyCardBorder(strings.Join(cd.lines, "\n"), inner, cd.selected)
			rowLines := lineCount(card)
			starts[i] = rowStart
			ends[i] = rowStart + rowLines - 1
			cursor += rowLines
		}
	}

	return starts, ends
}

// ensureFeaturedSelectionVisible keeps the highlighted featured card in view.
func (m *App) ensureFeaturedSelectionVisible() {
	if m.currentScreen != ScreenHome {
		return
	}

	featured := m.featuredProjects()
	if len(featured) == 0 {
		return
	}

	idx := m.selectedFeatured
	if idx < 0 || idx >= len(featured) {
		return
	}

	starts, ends := m.homeFeaturedLineMetrics()
	if len(starts) == 0 || idx >= len(starts) {
		return
	}

	m.ensureSelectionInViewport(starts[idx], ends[idx], m.homeContentTotalLines())
}
