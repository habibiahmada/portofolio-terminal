package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderProjectsContent renders the projects archive: all 10 projects with
// year + tags, navigable with ↑↓.
func (m *App) renderProjectsContent() string {
	title := styles.SectionTitleStyle.Render("All Projects")
	sub := styles.MutedStyle.Render(
		"Web projects by Habibi Ahmad Aziz. Production and capstone work across school systems, AI products, payments, and fullstack apps. Each project links to a detailed case study.",
	)

	cards := make([]string, 0, len(m.projects))
	for i, p := range m.projects {
		selected := i == m.selectedProject
		card := renderProjectCard(p, m.width, selected)
		cards = append(cards, card)
	}

	hint := styles.MutedStyle.Render("↑↓ browse · Enter view case study · ← back")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		sub,
		"",
		strings.Join(cards, "\n"),
		"",
		hint,
	)

	return styles.ContentStyle.Render(content)
}

// renderProjectCard renders a single project row card. When selected the title
// is highlighted and a "▸" marker is added.
func renderProjectCard(p data.Project, width int, selected bool) string {
	nameStyle := styles.SectionTitleStyle
	marker := ""
	if selected {
		nameStyle = styles.ListSelectedStyle
		marker = "▸ "
	}

	meta := styles.MutedStyle.Render(p.Year)
	tags := components.TagList(p.Tags)

	head := lipgloss.JoinHorizontal(lipgloss.Left, marker+nameStyle.Render(p.Name), "   "+meta)
	body := styles.NormalStyle.Render(p.Description)

	inner := lipgloss.JoinVertical(
		lipgloss.Left,
		head,
		"",
		body,
		"",
		tags,
	)

	cardStyle := styles.CardStyle
	if selected {
		cardStyle = styles.PrimaryCardStyle
	}
	return cardStyle.Render(inner)
}

// renderProjectDetailContent renders the case study for the open project.
func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail

	back := styles.MutedStyle.Render("← All projects")

	title := styles.SectionTitleStyle.Render(p.Name)
	year := styles.MutedStyle.Render(p.Year)
	desc := styles.NormalStyle.Render(p.Description)

	tags := components.TagList(p.Tags)

	// Links.
	links := make([]string, 0, 2)
	if p.Live != "" {
		links = append(links, styles.LinkStyle.Render("Live site: "+p.Live))
	}
	if p.GitHub != "" {
		links = append(links, styles.LinkStyle.Render("Source: "+p.GitHub))
	}

	// Case study sections.
	var sections []string
	cs := data.GetCaseStudy(p.Slug)
	if cs != nil && len(cs.Sections) > 0 {
		blocks := make([]string, 0, len(cs.Sections)*2)
		for _, sec := range cs.Sections {
			blocks = append(blocks,
				styles.LabelStyle.Render("── "+sec.Label+" ──"),
				styles.NormalStyle.Render(sec.Body),
				"",
			)
		}
		sections = append(sections, strings.Join(blocks, "\n"))
	}

	// Prev / Next navigation hints.
	prev, next := m.projectNeighbors()
	nav := fmt.Sprintf("%s   %s",
		styles.MutedStyle.Render("← "+prev),
		styles.MutedStyle.Render(" "+next+" →"),
	)

	footer := styles.MutedStyle.Render("h / l prev·next · ← back to projects")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		back,
		"",
		title,
		year,
		"",
		desc,
		"",
		tags,
		"",
		strings.Join(links, "\n"),
	)

	if len(sections) > 0 {
		content = lipgloss.JoinVertical(lipgloss.Left, content, "", strings.Join(sections, "\n"))
	}

	content = lipgloss.JoinVertical(lipgloss.Left, content, "", nav, footer)

	return styles.ContentStyle.Render(content)
}

// projectNeighbors returns the display names of the previous and next projects.
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
