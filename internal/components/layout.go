package components

import (
	"strings"
)

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

// clipContent returns a viewport window without mutating external state.
func ClipContent(content string, offset, maxH int) (clipped string, clampedOffset int) {
	if maxH <= 0 {
		return content, 0
	}

	lines := strings.Split(content, "\n")
	if len(lines) <= maxH {
		return content, 0
	}

	maxOffset := len(lines) - maxH
	if offset > maxOffset {
		offset = maxOffset
	}
	if offset < 0 {
		offset = 0
	}

	window := lines[offset : offset+maxH]
	result := strings.Join(window, "\n")
	if offset < maxOffset {
		result += "\n▼ scroll with j/k"
	}
	return result, offset
}
