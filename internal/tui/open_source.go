package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// renderOpenSourceContent lists public repositories: what each one is, a few
// facts, the command to try it, and where the code lives.
func (m *App) renderOpenSourceContent() string {
	cw := m.contentWidth()
	lines := []string{
		styles.SectionTitleStyle.Render("▸ Built in public, free to fork"),
	}
	lines = append(lines, components.WrapText(styles.MutedStyle.Render(
		"Tools I maintain on GitHub under the MIT license. Read the code, try it, or open an issue when it breaks."), cw)...)
	lines = append(lines, "")

	for i, p := range m.openSource {
		num := styles.BadgeAccentStyle.Render(fmt.Sprintf("[%02d]", i+1))
		kind := styles.BadgeStyle.Render("[" + p.Kind + "]")
		header := fmt.Sprintf("  %s %s %s", num, styles.PrimaryText.Render(p.Name), kind)
		lines = append(lines, components.WrapText(header, cw)...)

		// Body lines sit under the name, six columns in.
		const indent = "      "
		body := func(text string) {
			for _, l := range components.WrapText(text, cw-len(indent)) {
				lines = append(lines, indent+l)
			}
		}

		text := styles.NormalStyle.PaddingLeft(0)
		body(text.Render(p.Description))
		if p.ForkOf != "" {
			body(styles.MutedStyle.Render("Security-hardened fork of " + p.ForkOf))
		}
		if len(p.Facts) > 0 {
			body(styles.MutedStyle.Render(strings.Join(p.Facts, " · ")))
		}
		meta := p.Language
		if p.License != "" {
			meta += " · " + p.License
		}
		body(styles.MutedStyle.Render(meta))
		if p.Install != "" {
			body(styles.PromptStyle.Render("$ ") + text.Render(p.Install))
		}
		body(styles.MutedStyle.Render("Source ") + styles.LinkStyle.Render(p.RepoURL))
		if p.SiteURL != "" {
			label := p.SiteLabel
			if label == "" {
				label = "Site"
			}
			body(styles.MutedStyle.Render(label+" ") + styles.LinkStyle.Render(p.SiteURL))
		}
		lines = append(lines, "")
	}

	more := styles.MutedStyle.Render("  More on GitHub: ") + styles.LinkStyle.Render("https://github.com/habibiahmada")
	lines = append(lines, components.WrapText(more, cw)...)

	return strings.Join(lines, "\n")
}
