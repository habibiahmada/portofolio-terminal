package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// projectListMetrics mirrors renderProjectsContent layout to map each project
// index to its starting line after reflow (used for scroll-into-view).
func (m *App) projectListMetrics() (starts []int, totalLines int) {
	if len(m.projects) == 0 {
		return nil, 0
	}

	cw := m.contentWidth()
	header := []string{
		styles.SectionTitleStyle.Render("▸ Projects Archive"),
	}
	if m.portfolioSync == PortfolioSyncPending {
		header = append(header, styles.MutedStyle.Render("Syncing latest projects…"))
	} else {
		header = append(header, styles.MutedStyle.Render(projectsArchiveHint(len(m.projects))))
	}
	if banner := m.renderPortfolioSyncBanner(cw); banner != "" {
		header = append(header, "", banner)
	}
	header = append(header, "")

	cursor := len(strings.Split(components.ReflowBlock(strings.Join(header, "\n"), cw), "\n"))
	starts = make([]int, len(m.projects))

	for i, p := range m.projects {
		starts[i] = cursor
		sel := i == m.selectedProject
		row := renderProjectListRow(p, i, sel, m.focus == FocusContent, cw)
		rowLines := len(strings.Split(components.ReflowBlock(row, cw), "\n"))
		cursor += rowLines + 1 // blank separator between items
	}

	return starts, cursor
}

func (m *App) projectRowLineCount(idx int) int {
	if idx < 0 || idx >= len(m.projects) {
		return 0
	}
	cw := m.contentWidth()
	row := renderProjectListRow(
		m.projects[idx],
		idx,
		idx == m.selectedProject,
		m.focus == FocusContent,
		cw,
	)
	return len(strings.Split(components.ReflowBlock(row, cw), "\n"))
}

// ensureProjectSelectionVisible adjusts contentOffset so the selected project
// row stays inside the viewport when navigating with j/k or arrow keys.
func (m *App) ensureProjectSelectionVisible() {
	if m.currentScreen != ScreenProjects || len(m.projects) == 0 {
		return
	}

	starts, total := m.projectListMetrics()
	idx := m.selectedProject
	if idx < 0 || idx >= len(starts) {
		return
	}

	startLine := starts[idx]
	endLine := startLine + m.projectRowLineCount(idx) - 1
	m.ensureSelectionInViewport(startLine, endLine, total)
}

func projectsArchiveHint(count int) string {
	return fmt.Sprintf(
		"%d projects in the archive. Use j/k or arrow keys to browse, Enter to open a case study, and Esc to go back.",
		count,
	)
}
