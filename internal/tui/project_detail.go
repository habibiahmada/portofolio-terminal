package tui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/sanitize"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

var caseStudyPhaseLabels = []string{
	"Opening",
	"Reality",
	"Build",
	"Close",
}

func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail
	cw := m.contentWidth()
	inner := cw - 4
	if inner < 24 {
		inner = 24
	}

	var lines []string
	lines = append(lines, styles.MutedStyle.Render("← Back to list  (← / Esc)"))
	lines = append(lines, "")

	if banner := m.renderPortfolioSyncBanner(cw); banner != "" {
		lines = append(lines, banner, "")
	}

	lines = append(lines, styles.HeroTitleStyle.Render(p.Name))
	lines = append(lines, renderProjectDetailMeta(p, cw)...)
	if p.Featured {
		lines = append(lines, styles.BadgeStyle.Render("★ Featured Project"))
	}
	lines = append(lines, "")

	cs := data.GetCaseStudy(p.Slug)
	if cs != nil {
		sanitized := sanitize.CaseStudy(*cs)
		cs = &sanitized
	}

	lead := ""
	if cs != nil && cs.Hero != "" {
		lead = cs.Hero
	} else if p.Description != "" {
		lead = p.Description
	}
	if lead != "" {
		lines = append(lines, renderDetailLead(lead, inner), "")
	}

	if cs != nil {
		for i, sec := range cs.Sections {
			lines = append(lines, renderCaseStudySection(i, sec, inner)...)
			lines = append(lines, "")
		}
	}

	if len(p.Stack) > 0 {
		lines = append(lines, renderDetailStack(p.Stack, inner)...)
		lines = append(lines, "")
	}

	if p.Live != "" {
		lines = append(lines, styles.MutedStyle.Render("Live  ")+styles.LinkStyle.Render(p.Live))
	}

	prev, next := m.projectNeighbors()
	lines = append(lines, "",
		styles.RuleStyle.Render(strings.Repeat("─", min(cw, 48))),
		styles.MutedStyle.Render(
			fmt.Sprintf("  ← %s   ·   %s →   ·   h / l prev/next   ·   Esc back",
				truncateLabel(prev, 22), truncateLabel(next, 22)),
		),
	)

	return strings.Join(lines, "\n")
}

func renderProjectDetailMeta(p data.Project, cw int) []string {
	parts := []string{styles.MutedStyle.Render(p.Year)}
	for _, tag := range p.Tags {
		parts = append(parts, styles.TagStyle.Render(tag))
	}
	meta := strings.Join(parts, styles.MutedStyle.Render("  ·  "))
	var out []string
	for _, wl := range components.WrapText(meta, cw) {
		out = append(out, wl)
	}
	return out
}

func renderDetailLead(text string, inner int) string {
	var plain []string
	for _, wl := range components.WrapText(text, inner-2) {
		plain = append(plain, "  "+wl)
	}
	body := styles.NormalStyle.Render(strings.Join(plain, "\n"))
	return lipgloss.NewStyle().
		BorderLeft(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(styles.ColorSecondary).
		PaddingLeft(1).
		Render(body)
}

func renderCaseStudySection(idx int, sec data.CaseStudySection, inner int) []string {
	phase := "Section"
	if idx >= 0 && idx < len(caseStudyPhaseLabels) {
		phase = caseStudyPhaseLabels[idx]
	}
	num := fmt.Sprintf("%02d", idx+1)

	header := []string{
		styles.MutedStyle.Render(fmt.Sprintf("%s · %s", num, strings.ToUpper(phase))),
		styles.PrimaryText.Render(sec.Label),
	}
	panel := renderDetailPanel(formatSectionBody(sec.Body, inner), inner)

	return append(header, panel)
}

func renderDetailPanel(content string, inner int) string {
	if strings.TrimSpace(content) == "" {
		return ""
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.ColorBorder).
		Padding(0, 1).
		Render(content)
}

func formatSectionBody(body string, inner int) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}

	var blocks []string
	for _, chunk := range strings.Split(body, "\n\n") {
		chunk = strings.TrimSpace(chunk)
		if chunk == "" {
			continue
		}
		if bullets := parseBulletBlock(chunk); len(bullets) > 0 {
			blocks = append(blocks, renderBulletList(bullets, inner))
			continue
		}
		if title, bodyText, ok := parseTitledBlock(chunk); ok {
			blocks = append(blocks, renderTitledBlock(title, bodyText, inner))
			continue
		}
		blocks = append(blocks, strings.Join(indentWrapped(chunk, inner-2, ""), "\n"))
	}

	return strings.Join(blocks, "\n\n")
}

func parseBulletBlock(chunk string) []string {
	lines := strings.Split(chunk, "\n")
	var bullets []string
	for _, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" {
			continue
		}
		switch {
		case strings.HasPrefix(ln, "• "):
			bullets = append(bullets, strings.TrimPrefix(ln, "• "))
		case strings.HasPrefix(ln, "- "):
			bullets = append(bullets, strings.TrimPrefix(ln, "- "))
		default:
			if len(bullets) > 0 {
				bullets[len(bullets)-1] += " " + ln
			} else {
				return nil
			}
		}
	}
	return bullets
}

func parseTitledBlock(chunk string) (title, body string, ok bool) {
	lines := strings.Split(chunk, "\n")
	var clean []string
	for _, ln := range lines {
		if t := strings.TrimSpace(ln); t != "" {
			clean = append(clean, t)
		}
	}
	if len(clean) < 2 {
		return "", "", false
	}
	// Architecture blocks from the API: short title line + body paragraph(s).
	if len(clean) == 2 && len(clean[0]) <= 48 && len(clean[1]) > len(clean[0]) {
		return clean[0], clean[1], true
	}
	return "", "", false
}

func renderBulletList(items []string, inner int) string {
	lines := make([]string, 0, len(items))
	bulletIndent := "  • "
	bodyWidth := inner - len(bulletIndent)
	if bodyWidth < 12 {
		bodyWidth = 12
	}
	for _, item := range items {
		wrapped := components.WrapText(item, bodyWidth)
		for i, wl := range wrapped {
			prefix := bulletIndent
			if i > 0 {
				prefix = strings.Repeat(" ", len(bulletIndent))
			}
			lines = append(lines, prefix+styles.NormalStyle.Render(wl))
		}
	}
	return strings.Join(lines, "\n")
}

func renderTitledBlock(title, body string, inner int) string {
	titleLine := styles.PrimaryText.Render(title)
	bodyLines := indentWrapped(body, inner-2, "  ")
	return strings.Join(append([]string{titleLine}, bodyLines...), "\n")
}

func renderDetailStack(stack []string, inner int) []string {
	label := styles.MutedStyle.Render("Stack")
	var lines []string
	line := label + "  "
	for i, s := range stack {
		pill := styles.TagStyle.Render("[" + s + "]")
		sep := " "
		if i == 0 {
			sep = "  "
		}
		candidate := line + sep + pill
		if lipgloss.Width(candidate) > inner && i > 0 {
			lines = append(lines, line)
			line = "       " + pill
		} else {
			line = candidate
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

func indentWrapped(text string, width int, prefix string) []string {
	if width < 8 {
		width = 8
	}
	var out []string
	for _, wl := range components.WrapText(text, width) {
		out = append(out, prefix+styles.NormalStyle.Render(wl))
	}
	return out
}

func truncateLabel(s string, max int) string {
	if s == "" {
		return "—"
	}
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	return string(runes[:max-1]) + "…"
}
