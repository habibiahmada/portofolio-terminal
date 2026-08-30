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
			"Full-stack developer building clear UI and fast APIs. Press P for projects.",
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
// Height equalization is done on the raw body BEFORE the border is applied
// so both cards in each row always end up with identical outer dimensions.
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

	// Build raw bodies (no border yet) and selection flags.
	type cardData struct {
		body     string
		selected bool
	}
	cds := make([]cardData, 0, len(featured))
	for i, p := range featured {
		sel := m.focus == FocusContent && m.currentScreen == ScreenHome && i == m.selectedFeatured
		cds = append(cds, cardData{body: buildCardBody(p, inner, sel), selected: sel})
	}

	// Render rows of two cards with equalized heights.
	rows := []string{}
	for i := 0; i < len(cds); i += 2 {
		if i+1 < len(cds) {
			a, b := cds[i], cds[i+1]
			// Equalize body height BEFORE adding border.
			ha := lipgloss.Height(a.body)
			hb := lipgloss.Height(b.body)
			maxH := ha
			if hb > maxH {
				maxH = hb
			}
			aBody := lipgloss.NewStyle().Height(maxH).Width(inner).Render(a.body)
			bBody := lipgloss.NewStyle().Height(maxH).Width(inner).Render(b.body)
			// Now apply border to the height-equalized bodies.
			cardA := applyCardBorder(aBody, inner, a.selected)
			cardB := applyCardBorder(bBody, inner, b.selected)
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cardA, " ", cardB))
		} else {
			// Odd card at the end: just render as-is.
			cd := cds[i]
			rows = append(rows, applyCardBorder(cd.body, inner, cd.selected))
		}
	}

	return "\n" + heading + "\n" + sub + "\n\n" + strings.Join(rows, "\n")
}

// buildCardBody returns the unstyled text body (title + meta) for a featured card.
// inner is the content width (sans border and padding).
func buildCardBody(p data.Project, inner int, selected bool) string {
	// Title: always use 2-char prefix so all cards have the same indent.
	prefix := "  "
	if selected {
		prefix = styles.PrimaryText.Render("▸ ")
	}
	title := prefix + styles.HomeCardTitleStyle.Render(p.Name)

	// Meta: hard-truncate tags to fit the content width.
	tagsStr := p.Year + " · " + strings.Join(p.Tags, " · ")
	meta := "  " + styles.HomeCardMetaStyle.Render(components.Truncate(tagsStr, inner-2))

	return lipgloss.JoinVertical(lipgloss.Left, title, meta)
}

// applyCardBorder wraps a pre-rendered body with a rounded border.
// The border color changes to the primary accent when the card is selected.
func applyCardBorder(body string, inner int, selected bool) string {
	borderColor := styles.ColorBorder
	if selected {
		borderColor = styles.ColorPrimary
	}
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Padding(0, 1).
		Width(inner). // outer = inner + border(2) + padding(2)
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
