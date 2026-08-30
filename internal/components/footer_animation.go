package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// wave is a single-line equalizer wave whose visible window cycles with frame.
var wave = []string{
	"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁",
}

// FooterArtline renders a compact, always-present animated illustration for the
// footer: a single-line terminal badge with a blinking prompt and a cycling
// equalizer wave. On narrower terminals it degrades progressively so content
// stays the priority.
func FooterArtline(frame int, width int, website string) string {
	if width < 40 {
		return styles.FooterArtStyle.Render(styles.PromptStyle.Render(">_"))
	}

	// On very narrow widths keep only the brand mark.
	if width < 60 {
		mark := styles.PromptStyle.Render(">_") + " " +
			styles.FooterWordmark.Render("habibiahmada") +
			styles.HeaderDot.Render(".")
		return styles.FooterArtStyle.Render(mark + " " + PromptCursor(frame))
	}

	left := styles.PromptStyle.Render(">_") + " " +
		styles.FooterWordmark.Render("habibiahmada") +
		styles.HeaderDot.Render(".") + " " +
		styles.MutedStyle.Render("· "+website)

	eq := ""
	if width >= 70 {
		eq = "  " + equalizerWave(frame)
	} else {
		eq = "  " + styles.PromptStyle.Render("~")
	}

	return styles.FooterArtStyle.Render(left + eq + "  " + PromptCursor(frame))
}

// PromptCursor toggles a blinking cursor underscore.
func PromptCursor(frame int) string {
	if frame%2 == 0 {
		return styles.PromptStyle.Render("_")
	}
	return styles.PromptStyle.Render(" ")
}

// equalizerWave renders a single line of wave blocks whose position shifts
// with the frame, giving a gentle "alive" animation in the footer.
func equalizerWave(frame int) string {
	n := 7
	start := frame % len(wave)
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		c := styles.FooterBarStyle.Render(wave[(start+i)%len(wave)])
		parts = append(parts, c)
	}
	return strings.Join(parts, "")
}

// FooterWithHint renders the footer: the animated illustration on top and the
// keyboard hints below.
func FooterWithHint(artline string, hints []FooterHint, width int) string {
	hintParts := make([]string, 0, len(hints))
	for _, h := range hints {
		hintParts = append(hintParts, h.Key+" "+h.Label)
	}
	hintsLine := styles.FooterStyle.Render(strings.Join(hintParts, "  ·  "))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		artline,
		hintsLine,
	)
}
