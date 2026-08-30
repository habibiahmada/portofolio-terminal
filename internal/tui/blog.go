package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/blog"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// blogCategories are the filter labels shown on the Blog screen.
var blogCategories = []string{
	"programming", "education", "web", "career", "opinion", "news-commentary",
}

type blogPostsMsg struct {
	posts []blog.Post
	err   error
}

func (m *App) fetchBlogCmd() tea.Cmd {
	return func() tea.Msg {
		posts, err := blog.NewClient().FetchPublished()
		return blogPostsMsg{posts: posts, err: err}
	}
}

// renderBlogContent renders the blog list or article detail.
func (m *App) renderBlogContent() string {
	contentWidth := m.width - sidebarWidth - 2

	if m.blogDetail {
		return styles.ContentStyle.Render(m.renderBlogDetail(contentWidth))
	}
	return styles.ContentStyle.Render(m.renderBlogList(contentWidth))
}

func (m *App) renderBlogList(width int) string {
	label := styles.LabelStyle.Render("// Blog")
	title := styles.SectionTitleStyle.Render("Articles & Commentary")
	sub := styles.MutedStyle.Render(
		"Technical writing on web development, programming, and the craft of shipping production software.",
	)

	filters := make([]string, 0, len(blogCategories))
	for _, c := range blogCategories {
		filters = append(filters, styles.TagStyle.Render(blog.CategoryLabel(c)))
	}

	var body string
	switch {
	case m.blogLoading:
		body = styles.MutedStyle.Render("Loading articles…")
	case m.blogError != "":
		body = lipgloss.JoinVertical(
			lipgloss.Left,
			styles.NormalStyle.Render("Could not load blog posts."),
			styles.MutedStyle.Render(m.blogError),
			styles.MutedStyle.Render("Retry: leave this screen and navigate back."),
		)
	case len(m.blogPosts) == 0:
		body = lipgloss.JoinVertical(
			lipgloss.Left,
			styles.NormalStyle.Render("No articles yet."),
			styles.MutedStyle.Render("Published posts will appear here automatically."),
		)
	default:
		lines := make([]string, 0, len(m.blogPosts))
		for i, p := range m.blogPosts {
			prefix := "  "
			if i == m.blogSelected {
				prefix = "> "
			}
			meta := blog.CategoryLabel(p.Category)
			if p.PublishedAt != "" {
				meta += " · " + blog.FormatDate(p.PublishedAt)
			}
			if p.ReadingTimeMinutes > 0 {
				meta += fmt.Sprintf(" · %d min", p.ReadingTimeMinutes)
			}
			head := prefix + styles.NormalStyle.Render(p.Title)
			if i == m.blogSelected {
				head = prefix + styles.ListSelectedStyle.Render(p.Title)
			}
			desc := ""
			if p.Description != "" && width >= 60 {
				desc = "\n    " + styles.MutedStyle.Render(truncate(p.Description, min(width-6, 72)))
			}
			lines = append(lines, head+desc, "    "+styles.MutedStyle.Render(meta), "")
		}
		body = strings.Join(lines, "\n")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		title,
		sub,
		"",
		styles.SectionTitleStyle.Render("Categories"),
		joinInline(filters, " "),
		"",
		styles.SectionTitleStyle.Render("Feed"),
		body,
		"",
		styles.MutedStyle.Render("Enter open · j/k select"),
	)
}

func (m *App) renderBlogDetail(width int) string {
	p := m.blogPosts[m.blogSelected]
	label := styles.LabelStyle.Render("// Blog")
	meta := blog.CategoryLabel(p.Category)
	if p.PublishedAt != "" {
		meta += " · " + blog.FormatDate(p.PublishedAt)
	}
	if p.ReadingTimeMinutes > 0 {
		meta += fmt.Sprintf(" · %d min read", p.ReadingTimeMinutes)
	}

	tags := make([]string, 0, len(p.Tags))
	for _, t := range p.Tags {
		tags = append(tags, styles.TagStyle.Render(t))
	}

	body := blog.FormatMarkdown(p.BodyMD, width)
	if body == "" && p.Description != "" {
		body = styles.NormalStyle.Render(p.Description)
	}

	parts := []string{
		label,
		styles.MutedStyle.Render("← back to list"),
		"",
		styles.SectionTitleStyle.Render(p.Title),
		styles.MutedStyle.Render(meta),
	}
	if len(tags) > 0 {
		parts = append(parts, joinInline(tags, " "))
	}
	if p.Description != "" {
		parts = append(parts, "", styles.NormalStyle.Render(p.Description))
	}
	parts = append(parts, "", body)

	return strings.Join(parts, "\n")
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
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
