package tui

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// blogCategories are the filter labels shown on the Blog placeholder.
var blogCategories = []string{
	"programming", "education", "web", "career", "opinion", "news-commentary",
}

// renderBlogContent renders the blog placeholder screen. Content is dynamic in
// a later phase; the empty state explains that articles are coming soon.
func (m *App) renderBlogContent() string {
	label := styles.LabelStyle.Render("// Blog")
	title := styles.SectionTitleStyle.Render("Articles & Commentary")
	sub := styles.MutedStyle.Render(
		"Technical writing on web development, programming, and the craft of shipping production software.",
	)

	filters := make([]string, 0, len(blogCategories))
	for _, c := range blogCategories {
		filters = append(filters, styles.TagStyle.Render(c))
	}

	empty := lipgloss.JoinVertical(
		lipgloss.Left,
		styles.NormalStyle.Render("No articles yet."),
		styles.MutedStyle.Render("I am drafting posts, mostly on web development and shipping production software. First articles land soon."),
	)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		title,
		sub,
		"",
		styles.SectionTitleStyle.Render("Categories"),
		joinInline(filters, " "),
		"",
		styles.SectionTitleStyle.Render("Feed"),
		empty,
	)

	return styles.ContentStyle.Render(content)
}

func joinInline(parts []string, sep string) string {
	out := ""
	for _, p := range parts {
		if out != "" {
			out += sep
		}
		out += p
	}
	return out
}
