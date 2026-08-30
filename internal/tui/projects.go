package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderProjectsContent() string {
	lines := []string{
		styles.SectionTitleStyle.Render("▸ Projects Archive"),
		styles.MutedStyle.Render(fmt.Sprintf("%d shipped projects · ↑/↓ select · Enter for case study", len(m.projects))),
		"",
	}
	for i, p := range m.projects {
		lines = append(lines, renderProjectListRow(p, i == m.selectedProject, m.focus == FocusContent))
	}
	return strings.Join(lines, "\n")
}

func renderProjectListRow(p data.Project, selected, contentFocused bool) string {
	prefix := "  "
	nameStyle := styles.NormalStyle
	if selected && contentFocused {
		prefix = "> "
		nameStyle = styles.ListSelectedStyle
	} else if selected {
		prefix = "· "
		nameStyle = styles.SelectedStyle
	}
	meta := p.Year
	if len(p.Tags) > 0 {
		meta += " · " + strings.Join(p.Tags[:min(3, len(p.Tags))], ", ")
	}
	return prefix + nameStyle.Render(p.Name) + "  " + styles.MutedStyle.Render(meta)
}

func renderProjectCard(p data.Project, width int, selected, contentFocused bool) string {
	return renderProjectListRow(p, selected, contentFocused)
}

func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail
	cw := m.contentWidth()

	lines := []string{
		styles.MutedStyle.Render("← Back to projects (← / Esc)"),
		"",
		styles.HeroTitleStyle.Render(p.Name),
		styles.SubtitleStyle.Render(p.Year + " · " + strings.Join(p.Tags, " · ")),
		"",
	}

	cs := data.GetCaseStudy(p.Slug)
	if cs != nil && cs.Hero != "" {
		lines = append(lines, styles.NormalStyle.Render(strings.Join(components.WrapText(cs.Hero, cw), "\n")), "")
	} else if p.Description != "" {
		lines = append(lines, styles.NormalStyle.Render(strings.Join(components.WrapText(p.Description, cw), "\n")), "")
	}

	if cs != nil {
		for _, sec := range cs.Sections {
			body := strings.Join(components.WrapText(sec.Body, cw-2), "\n  ")
			lines = append(lines,
				styles.PrimaryText.Render("▸ "+sec.Label),
				"  "+styles.NormalStyle.Render(body),
				"",
			)
		}
	}

	if len(p.Stack) > 0 {
		lines = append(lines, styles.MutedStyle.Render("Tech Stack: "+strings.Join(p.Stack, ", ")))
	}
	if p.Live != "" {
		lines = append(lines, styles.LinkStyle.Render("Live Demo:  "+p.Live))
	}
	if p.GitHub != "" {
		lines = append(lines, styles.LinkStyle.Render("Repository: "+p.GitHub))
	}

	prev, next := m.projectNeighbors()
	lines = append(lines, "",
		styles.MutedStyle.Render(fmt.Sprintf("Shortcuts: [h] %s  ·  [l] %s  ·  [←] List", prev, next)),
	)
	return strings.Join(lines, "\n")
}

func (m *App) projectNeighbors() (prev, next string) {
	if len(m.projects) == 0 {
		return "", ""
	}
	idx := m.projectIndex()
	pIdx := idx - 1
	if pIdx < 0 {
		pIdx = len(m.projects) - 1
	}
	nIdx := idx + 1
	if nIdx >= len(m.projects) {
		nIdx = 0
	}
	return m.projects[pIdx].Name, m.projects[nIdx].Name
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
