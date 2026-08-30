package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// FitLine pads or clips s to exactly width visible cells, preserving ANSI.
func FitLine(s string, width int) string {
	if width <= 0 {
		return ""
	}
	w := lipgloss.Width(s)
	if w == width {
		return s
	}
	if w < width {
		return s + strings.Repeat(" ", width-w)
	}
	clipped := lipgloss.NewStyle().MaxWidth(width).Render(s)
	first, _, _ := strings.Cut(clipped, "\n")
	w = lipgloss.Width(first)
	if w < width {
		first += strings.Repeat(" ", width-w)
	}
	return first
}

// Truncate cuts plain text to at most n visible cells, adding an ellipsis
// when the original is longer.
func Truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if lipgloss.Width(s) <= n {
		return s
	}
	if n == 1 {
		return "…"
	}
	runes := []rune(s)
	for i := len(runes); i >= 0; i-- {
		cand := string(runes[:i]) + "…"
		if lipgloss.Width(cand) <= n {
			return cand
		}
	}
	return "…"
}

// FitBlock pads or clips every line of block to exactly width cells.
func FitBlock(block string, width int) string {
	lines := strings.Split(block, "\n")
	for i, ln := range lines {
		lines[i] = FitLine(ln, width)
	}
	return strings.Join(lines, "\n")
}

// CenterInViewport vertically centers short content in the viewport using
// lightweight newline padding. Avoids lipgloss.Place which allocates a full
// width×height buffer and is expensive when called every frame.
func CenterInViewport(content string, width, height int) string {
	_ = width // horizontal centering handled by ContentStyle padding
	if height <= 0 || content == "" {
		return content
	}

	lines := strings.Split(content, "\n")
	if len(lines) >= height {
		return content
	}

	topPad := (height - len(lines)) / 2
	if topPad <= 0 {
		return content
	}
	return strings.Repeat("\n", topPad) + content
}

// WrapText wraps plain text to the given width using word boundaries.
func WrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}

	var lines []string
	line := words[0]
	for _, w := range words[1:] {
		candidate := line + " " + w
		if len(candidate) > width {
			lines = append(lines, line)
			line = w
			continue
		}
		line = candidate
	}
	lines = append(lines, line)
	return lines
}

// WrapTextLine wraps a single line of text to fit width (word boundaries).
func WrapTextLine(text string, width int) string {
	lines := WrapText(text, width)
	return strings.Join(lines, "\n")
}

// ClipResult describes a clipped viewport window.
type ClipResult struct {
	Text   string
	Offset int
	Max    int
}

// clipContent returns a viewport window without mutating external state.
func ClipContent(content string, offset, maxH int) (clipped string, clampedOffset int) {
	r := ClipContentFull(content, offset, maxH)
	return r.Text, r.Offset
}

// ClipContentFull returns the clipped window plus its applied offset and the
// maximum possible offset, so callers can render a scrollbar.
func ClipContentFull(content string, offset, maxH int) ClipResult {
	r := ClipResult{}
	if maxH <= 0 {
		r.Text = content
		return r
	}

	lines := strings.Split(content, "\n")
	if len(lines) <= maxH {
		r.Text = content
		return r
	}

	r.Max = len(lines) - maxH
	if offset > r.Max {
		offset = r.Max
	}
	if offset < 0 {
		offset = 0
	}
	r.Offset = offset
	r.Text = strings.Join(lines[offset:offset+maxH], "\n")
	return r
}

// ScrollbarWidth is the reserved gutter (gap + thumb) appended to content.
const ScrollbarWidth = 2

// AddScrollbar appends a one-cell scrollbar in a fixed 2-cell gutter so every
// line stays the same width. When max is 0 the gutter is blank (stable layout).
func AddScrollbar(text string, offset, max int, _ int) string {
	lines := strings.Split(text, "\n")
	var bar []string
	if max <= 0 {
		bar = make([]string, len(lines))
		for i := range bar {
			bar[i] = " "
		}
	} else {
		bar = buildScrollbar(offset, max, len(lines))
	}
	out := make([]string, len(lines))
	for i, ln := range lines {
		if i >= len(bar) {
			break
		}
		out[i] = ln + " " + bar[i]
	}
	return strings.Join(out, "\n")
}

// buildScrollbar builds the per-line cells for the scrollbar lane.
func buildScrollbar(offset, max, h int) []string {
	bar := make([]string, h)
	for i := range bar {
		bar[i] = styles.ScrollTrack
	}
	if max <= 0 || h <= 0 {
		return bar
	}

	thumb := 1
	if h >= 4 {
		thumb = h / 4
		if thumb < 1 {
			thumb = 1
		}
	}

	pos := 0
	if max > 0 {
		pos = int(float64(offset) / float64(max) * float64(h-thumb))
	}
	for i := pos; i < pos+thumb && i < h; i++ {
		bar[i] = styles.ScrollThumb
	}
	return bar
}
