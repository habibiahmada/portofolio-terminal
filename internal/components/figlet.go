package components

import (
	"strings"
	"sync"

	"github.com/charmbracelet/lipgloss"

	"github.com/common-nighthawk/go-figure"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

var (
	figletMu    sync.Mutex
	figletCache = map[string]string{}
)

// figletFontFor picks a FIGlet font based on available width.
func figletFontFor(width int) string {
	switch {
	case width >= 100:
		return "slant"
	case width >= 80:
		return "standard"
	default:
		return "small"
	}
}

// Figlet renders text in a large terminal font that fits the given width,
// falling back to smaller fonts and finally to plain styled text. Renders are
// cached so the View does not regenerate them on every redraw.
func Figlet(text string, width int) string {
	font := figletFontFor(width)

	key := text + "|" + font
	figletMu.Lock()
	if v, ok := figletCache[key]; ok {
		figletMu.Unlock()
		return v
	}
	figletMu.Unlock()

	out := renderFigletFont(text, font)
	if figletWidth(out) > width {
		out = renderFigletFont(text, "small")
	}
	if figletWidth(out) > width {
		return styles.HeroTitleStyle.Render(text)
	}

	figletMu.Lock()
	figletCache[key] = out
	figletMu.Unlock()
	return out
}

func renderFigletFont(text, font string) string {
	f := figure.NewFigure(text, font, true)
	return styles.HeroTitleStyle.Render(f.String())
}

func figletWidth(rendered string) int {
	w := 0
	for _, line := range strings.Split(rendered, "\n") {
		if lw := lipgloss.Width(line); lw > w {
			w = lw
		}
	}
	return w
}
