package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderProjectsContent() string {
	lines := []string{
		styles.SectionTitleStyle.Render("Projects"),
		styles.MutedStyle.Render(fmt.Sprintf("%d shipped · j/k select · Enter detail", len(m.projects))),
		"",
	}
	for i, p := range m.projects {
		lines = append(lines, renderProjectListRow(p, i == m.selectedProject))
	}
	return strings.Join(lines, "\n")
}

func renderProjectListRow(p data.Project, selected bool) string {
	prefix := "  "
	nameStyle := styles.NormalStyle
	if selected {
		prefix = "> "
		nameStyle = styles.ListSelectedStyle
	}
	meta := p.Year
	if len(p.Tags) > 0 {
		meta += " · " + strings.Join(p.Tags[:min(3, len(p.Tags))], ", ")
	}
	return prefix + nameStyle.Render(p.Name) + "  " + styles.MutedStyle.Render(meta)
}

func renderProjectCard(p data.Project, width int, selected bool) string {
	return renderProjectListRow(p, selected)
}

func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail
	lines := []string{
		styles.MutedStyle.Render("← back to list"),
		styles.SectionTitleStyle.Render(p.Name),
		styles.MutedStyle.Render(p.Year + " · " + strings.Join(p.Tags, ", ")),
	}

	cs := data.GetCaseStudy(p.Slug)
	if cs != nil && cs.Hero != "" {
		lines = append(lines, "", styles.NormalStyle.Render(componentsWrap(cs.Hero, m.contentWidth())))
	} else if p.Description != "" {
		lines = append(lines, "", styles.NormalStyle.Render(componentsWrap(p.Description, m.contentWidth())))
	}

	if cs != nil {
		for _, sec := range cs.Sections {
			body := componentsWrap(sec.Body, m.contentWidth())
			lines = append(lines, "",
				styles.PrimaryText.Render(sec.Label),
				styles.NormalStyle.Render(body),
			)
		}
	}

	if len(p.Stack) > 0 {
		lines = append(lines, "", styles.MutedStyle.Render("Stack: "+strings.Join(p.Stack, ", ")))
	}
	if p.Live != "" {
		lines = append(lines, styles.LinkStyle.Render("Live: "+p.Live))
	}
	if p.GitHub != "" {
		lines = append(lines, styles.LinkStyle.Render("Source: "+p.GitHub))
	}

	prev, next := m.projectNeighbors()
	lines = append(lines, "",
		styles.MutedStyle.Render(fmt.Sprintf("h/l %s · %s · ← list", prev, next)),
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
