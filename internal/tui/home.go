package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderHomeContent — Home = concise welcome hero, clean featured project cards,
// trusted-by partners strip, and quick keyboard hints.
func (m *App) renderHomeContent() string {
	w := m.contentWidth()
	parts := []string{
		m.renderHomeHero(w),
		m.renderHomeFeatured(w),
		m.renderHomeTrusted(w),
		m.renderHomeShortcuts(w),
	}
	return strings.Join(parts, "\n")
}

// renderHomeHero — availability status badge, wordmark banner, and brief bio.
func (m *App) renderHomeHero(width int) string {
	badge := styles.SuccessStyle.Render(
		components.WrapTextLine(
			"● "+m.profile.Availability, width,
		),
	)

	word := strings.ToUpper(strings.Fields(m.profile.BoostName)[0])
	banner := components.HeroBanner(word, width)
	caption := components.CenterText(styles.PromptStyle.Render(">_ "+m.profile.Website), width)

	name := styles.HeroTitleStyle.Render(m.profile.Name)
	meta := styles.MutedStyle.Render(m.profile.Title + " · " + m.profile.Location)
	bio := styles.NormalStyle.Render(
		components.WrapTextLine(
			"Welcome! You're browsing my portfolio straight from your shell. Every project, skill, and piece of experience I have is accessible right here without a browser.",
			width-1,
		),
	)

	lines := []string{badge}
	if banner != "" {
		lines = append(lines, "", banner, caption)
	}
	lines = append(lines, "", name, meta, bio)
	return strings.Join(lines, "\n")
}

// featuredProjects returns the featured project slice (up to 4).
func (m *App) featuredProjects() []data.Project {
	out := make([]data.Project, 0, 4)
	for _, p := range m.projects {
		if p.Featured {
			out = append(out, p)
			if len(out) == 4 {
				break
			}
		}
	}
	return out
}

// renderHomeFeatured — compact 2-col × 2-row card grid.
// Each line is FitLine'd to exactly `inner` chars before the border is applied,
// ensuring no lipgloss re-wrapping. Heights are equalized at the line-slice level.
func (m *App) renderHomeFeatured(width int) string {
	heading := styles.SectionTitleStyle.Render("▸ Featured Work")
	sub := styles.MutedStyle.Render("↑↓ navigate  ·  Enter to open  ·  P for all projects")

	featured := m.featuredProjects()
	if len(featured) == 0 {
		return ""
	}

	// cardOuter: each card's total outer width (border + padding included).
	// Two cards + 1-col gutter must fit in `width`.
	cardOuter := (width - 1) / 2
	inner := cardOuter - 4
	if inner < 16 {
		inner = 16
	}

	type rawCard struct {
		lines    []string
		selected bool
	}
	raws := make([]rawCard, 0, len(featured))
	for i, p := range featured {
		sel := m.focus == FocusContent && m.currentScreen == ScreenHome && i == m.selectedFeatured
		raws = append(raws, rawCard{buildCardLines(p, inner, sel), sel})
	}

	// Build rows: equalize line-counts per pair, then apply border.
	rows := []string{}
	blank := strings.Repeat(" ", inner)
	for i := 0; i < len(raws); i += 2 {
		if i+1 < len(raws) {
			a, b := raws[i], raws[i+1]
			// Equalize heights by appending blank lines to the shorter card.
			for len(a.lines) < len(b.lines) {
				a.lines = append(a.lines, blank)
			}
			for len(b.lines) < len(a.lines) {
				b.lines = append(b.lines, blank)
			}
			cardA := applyCardBorder(strings.Join(a.lines, "\n"), inner, a.selected)
			cardB := applyCardBorder(strings.Join(b.lines, "\n"), inner, b.selected)
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cardA, " ", cardB))
		} else {
			cd := raws[i]
			rows = append(rows, applyCardBorder(strings.Join(cd.lines, "\n"), inner, cd.selected))
		}
	}

	return "\n" + heading + "\n" + sub + "\n\n" + strings.Join(rows, "\n")
}

// buildCardLines builds the card content as a slice of lines, each FitLine'd to
// exactly `inner` visible cells. No JoinVertical → no trailing-space inflation.
func buildCardLines(p data.Project, inner int, selected bool) []string {
	// Line 1 — title with selection indicator.
	prefix := "  " // 2 visible cols, same as "▸ "
	if selected {
		prefix = styles.PrimaryText.Render("▸ ")
	}
	// Render title first, then FitLine the combined line.
	rawTitle := prefix + styles.HomeCardTitleStyle.Render(
		components.Truncate(p.Name, inner-2),
	)
	line1 := components.FitLine(rawTitle, inner)

	// Line 2 — year + tags, truncated then fit to width.
	tagsStr := p.Year + " · " + strings.Join(p.Tags, " · ")
	rawMeta := "  " + styles.HomeCardMetaStyle.Render(
		components.Truncate(tagsStr, inner-2),
	)
	line2 := components.FitLine(rawMeta, inner)

	return []string{line1, line2}
}

// applyCardBorder wraps a pre-built body string with a rounded border.
// Width(inner) is applied exactly once here; body lines are already FitLine'd.
func applyCardBorder(body string, inner int, selected bool) string {
	borderColor := styles.ColorBorder
	if selected {
		borderColor = styles.ColorPrimary
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(inner). // outer = inner + padding(2) + border(2) = cardOuter
		Render(body)
}

// renderHomeTrusted — centered "Trusted & Supported By" heading followed by company strip.
func (m *App) renderHomeTrusted(width int) string {
	names := make([]string, 0, len(m.companies))
	for _, c := range m.companies {
		names = append(names, c.Name)
	}
	line := strings.Join(names, " · ")

	heading := components.CenterText(
		styles.SectionTitleStyle.Render("Trusted & Supported By"), width,
	)

	var centered []string
	for _, ln := range strings.Split(components.WrapTextLine(line, width), "\n") {
		centered = append(centered, components.CenterText(styles.MutedStyle.Render(ln), width))
	}

	return "\n" + heading + "\n" + strings.Join(centered, "\n")
}

// renderHomeShortcuts renders a concise bottom keyboard hint strip.
func (m *App) renderHomeShortcuts(width int) string {
	hints := styles.MutedStyle.Render("Quick Actions: [P] All Projects  ·  [C] Contact  ·  [V] View CV  ·  [?] Help")
	return "\n" + components.CenterText(hints, width)
}
