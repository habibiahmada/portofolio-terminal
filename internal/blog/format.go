package blog

import (
	"regexp"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

var (
	reHeader   = regexp.MustCompile(`^(#{1,3})\s+(.+)$`)
	reBullet   = regexp.MustCompile(`^[-*+]\s+(.+)$`)
	reOrdered  = regexp.MustCompile(`^\d+\.\s+(.+)$`)
	reImage    = regexp.MustCompile(`^!\[.*?\]\(.*?\)\s*$`)
	reHTMLTag  = regexp.MustCompile(`<[^>]+>`)
	reBold     = regexp.MustCompile(`\*\*(.+?)\*\*`)
	reCode     = regexp.MustCompile("`([^`]+)`")
	reLink     = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	reHR       = regexp.MustCompile(`^---+$|^\*\*\*+$|^___+$`)
)

// maxMarkdownRunes limits blog body parsing to keep TUI responsive.
const maxMarkdownRunes = 32 * 1024

// FormatMarkdown converts markdown body text into styled TUI lines.
func FormatMarkdown(md string, width int) string {
	if width < 20 {
		width = 20
	}
	runes := []rune(md)
	if len(runes) > maxMarkdownRunes {
		md = string(runes[:maxMarkdownRunes]) + "\n\n…"
	}
	bodyWidth := width - 4
	if bodyWidth < 16 {
		bodyWidth = 16
	}

	var out []string
	inCodeBlock := false

	for _, raw := range strings.Split(md, "\n") {
		line := strings.TrimRight(raw, " \t")

		if strings.HasPrefix(line, "```") {
			inCodeBlock = !inCodeBlock
			if inCodeBlock {
				out = append(out, styles.MutedStyle.Render("── code ──"))
			} else {
				out = append(out, styles.MutedStyle.Render("── /code ──"))
			}
			continue
		}

		if inCodeBlock {
			out = append(out, styles.MutedStyle.Render("  "+line))
			continue
		}

		if line == "" {
			out = append(out, "")
			continue
		}

		if reImage.MatchString(line) {
			continue
		}
		if reHR.MatchString(line) {
			out = append(out, styles.RuleStyle.Render(strings.Repeat("─", min(bodyWidth, 40))))
			continue
		}

		if m := reHeader.FindStringSubmatch(line); m != nil {
			text := inlineStyles(m[2])
			switch len(m[1]) {
			case 1:
				out = append(out, styles.SectionTitleStyle.Render(text))
			case 2:
				out = append(out, styles.SubtitleStyle.Render(text))
			default:
				out = append(out, styles.NormalStyle.Render(text))
			}
			continue
		}

		prefix := "  "
		content := line
		if m := reBullet.FindStringSubmatch(line); m != nil {
			prefix = "  • "
			content = m[1]
		} else if m := reOrdered.FindStringSubmatch(line); m != nil {
			prefix = "  · "
			content = m[1]
		}

		wrapped := components.WrapText(stripBlockMarkdown(content), bodyWidth-len(prefix))
		for i, w := range wrapped {
			p := prefix
			if i > 0 {
				p = strings.Repeat(" ", len(prefix))
			}
			out = append(out, styles.NormalStyle.Render(p+inlineStyles(w)))
		}
	}

	return strings.Join(out, "\n")
}

func inlineStyles(s string) string {
	s = reHTMLTag.ReplaceAllString(s, "")
	s = reLink.ReplaceAllString(s, "$1")
	s = reCode.ReplaceAllStringFunc(s, func(m string) string {
		inner := reCode.FindStringSubmatch(m)
		if len(inner) < 2 {
			return m
		}
		return styles.TagStyle.Render(inner[1])
	})
	s = reBold.ReplaceAllStringFunc(s, func(m string) string {
		inner := reBold.FindStringSubmatch(m)
		if len(inner) < 2 {
			return m
		}
		return styles.SelectedStyle.Render(inner[1])
	})
	return s
}

func stripBlockMarkdown(s string) string {
	s = reLink.ReplaceAllString(s, "$1")
	s = strings.Trim(s, "*_`")
	return s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
