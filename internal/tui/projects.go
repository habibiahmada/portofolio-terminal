package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderProjectsContent() string {
	cw := m.contentWidth()

	// ── Header ────────────────────────────────────────────────────────────
	title := styles.SectionTitleStyle.Render("▸ Projects Archive")
	hint := styles.MutedStyle.Render(
		fmt.Sprintf("%d projects · ↑/↓ select · Enter open · ← back", len(m.projects)),
	)

	lines := []string{title, hint, ""}

	// ── Project rows ──────────────────────────────────────────────────────
	for i, p := range m.projects {
		sel := i == m.selectedProject
		focused := m.focus == FocusContent
		lines = append(lines, renderProjectListRow(p, i, sel, focused, cw))
		// Blank separator between items for breathing room.
		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}

// renderProjectListRow renders a single project entry.
// Selected + focused  → full highlight bar with border accent.
// Selected unfocused  → soft accent.
// Unselected          → dimmed minimal row.
func renderProjectListRow(p data.Project, idx int, selected, contentFocused bool, cw int) string {
	// ── Number badge ─────────────────────────────────────────────────────
	num := fmt.Sprintf("%02d", idx+1)
	var numStr string
	if selected && contentFocused {
		numStr = styles.PrimaryText.Render(num)
	} else {
		numStr = styles.MutedStyle.Render(num)
	}

	// ── Project name ──────────────────────────────────────────────────────
	var nameStr string
	if selected && contentFocused {
		nameStr = styles.ListSelectedStyle.Render(p.Name)
	} else if selected {
		nameStr = styles.SelectedStyle.Render(p.Name)
	} else {
		nameStr = styles.NormalStyle.Render(p.Name)
	}

	// ── Featured badge ────────────────────────────────────────────────────
	featuredBadge := ""
	if p.Featured {
		featuredBadge = "  " + styles.BadgeStyle.Render("★ Featured")
	}

	// ── Year ──────────────────────────────────────────────────────────────
	yearStr := styles.MutedStyle.Render(p.Year)

	// ── Tags (up to 3, inline) ────────────────────────────────────────────
	tagLimit := 3
	if len(p.Tags) < tagLimit {
		tagLimit = len(p.Tags)
	}
	tagParts := make([]string, tagLimit)
	for i, t := range p.Tags[:tagLimit] {
		tagParts[i] = styles.TagStyle.Render(t)
	}
	if len(p.Tags) > 3 {
		tagParts = append(tagParts, styles.MutedStyle.Render(fmt.Sprintf("+%d", len(p.Tags)-3)))
	}
	tagsStr := strings.Join(tagParts, styles.MutedStyle.Render(" · "))

	// ── Assemble top line: num + name + featured badge ────────────────────
	topLine := numStr + "  " + nameStr + featuredBadge

	// ── Assemble meta line: indent + year · tags ──────────────────────────
	metaLine := "    " + yearStr + styles.MutedStyle.Render("  ·  ") + tagsStr

	// ── Layout with selection indicator ───────────────────────────────────
	if selected && contentFocused {
		// Full-width highlight bar for active selection.
		// The ">" is required for test assertions (TestProjectsListFocusExclusive).
		indicator := styles.PrimaryText.Render("> ")
		topRendered := indicator + topLine
		metaRendered := "  " + metaLine

		// Clamp to content width.
		topRendered = components.FitLine(topRendered, cw)
		metaRendered = components.FitLine(metaRendered, cw)

		// Wrap with a thick left border accent.
		body := lipgloss.JoinVertical(lipgloss.Left, topRendered, metaRendered)
		return lipgloss.NewStyle().
			BorderLeft(true).
			BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(styles.ColorPrimary).
			PaddingLeft(1).
			Render(body)
	}

	if selected {
		// Soft selection — muted left accent.
		indicator := styles.MutedStyle.Render("│ ")
		topRendered := indicator + topLine
		metaRendered := "  " + metaLine
		body := lipgloss.JoinVertical(lipgloss.Left, topRendered, metaRendered)
		return lipgloss.NewStyle().
			PaddingLeft(1).
			Render(body)
	}

	// Unselected row — compact single line with meta inline.
	rowLine := "   " + topLine
	metaInline := "    " + metaLine
	return lipgloss.JoinVertical(lipgloss.Left, rowLine, metaInline)
}

// renderProjectCard is kept for backward-compat (used by mouse hit-testing).
func renderProjectCard(p data.Project, width int, selected, contentFocused bool) string {
	return renderProjectListRow(p, 0, selected, contentFocused, width)
}

func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail
	cw := m.contentWidth()

	// ── Back nav ──────────────────────────────────────────────────────────
	backHint := styles.MutedStyle.Render("← Back  (← / Esc)")

	// ── Title block ───────────────────────────────────────────────────────
	titleLine := styles.HeroTitleStyle.Render(p.Name)

	yearTags := p.Year
	if len(p.Tags) > 0 {
		yearTags += "  ·  " + strings.Join(p.Tags, " · ")
	}
	metaLine := styles.SubtitleStyle.Render(yearTags)

	// ── Featured marker ───────────────────────────────────────────────────
	var featLine string
	if p.Featured {
		featLine = styles.BadgeStyle.Render("★ Featured Project")
	}

	lines := []string{backHint, "", titleLine, metaLine}
	if featLine != "" {
		lines = append(lines, featLine)
	}
	lines = append(lines, "")

	// ── Hero / description ────────────────────────────────────────────────
	cs := data.GetCaseStudy(p.Slug)
	if cs != nil && cs.Hero != "" {
		wrapped := strings.Join(components.WrapText(cs.Hero, cw), "\n")
		lines = append(lines, styles.NormalStyle.Render(wrapped), "")
	} else if p.Description != "" {
		wrapped := strings.Join(components.WrapText(p.Description, cw), "\n")
		lines = append(lines, styles.NormalStyle.Render(wrapped), "")
	}

	// ── Case study sections ───────────────────────────────────────────────
	if cs != nil {
		for _, sec := range cs.Sections {
			sectionLabel := styles.PrimaryText.Render("▸ " + sec.Label)
			body := strings.Join(components.WrapText(sec.Body, cw-2), "\n  ")
			sectionBody := "  " + styles.NormalStyle.Render(body)
			lines = append(lines, sectionLabel, sectionBody, "")
		}
	}

	// ── Tech stack pills ──────────────────────────────────────────────────
	if len(p.Stack) > 0 {
		stackLabel := styles.MutedStyle.Render("Stack  ")
		stackPills := make([]string, len(p.Stack))
		for i, s := range p.Stack {
			stackPills[i] = styles.TagStyle.Render("[" + s + "]")
		}
		lines = append(lines, stackLabel+strings.Join(stackPills, " "), "")
	}

	// ── Links ─────────────────────────────────────────────────────────────
	if p.Live != "" {
		lines = append(lines, styles.MutedStyle.Render("Live  ")+styles.LinkStyle.Render(p.Live))
	}
	if p.GitHub != "" {
		lines = append(lines, styles.MutedStyle.Render("Repo  ")+styles.LinkStyle.Render(p.GitHub))
	}

	// ── Prev / Next navigation ────────────────────────────────────────────
	prev, next := m.projectNeighbors()
	lines = append(lines, "",
		styles.MutedStyle.Render("  ← "+prev+"   ·   "+next+" →   ·   [Esc] Back to list"),
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
