package components

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// ProgressBar renders a horizontal progress bar at the given width (in cells).
// percent must be in [0, 100].
func ProgressBar(percent, width int) string {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	if width < 1 {
		width = 1
	}

	filled := (width * percent) / 100
	empty := width - filled

	// Round the boundary: use a half block when there is remainder.
	var half string
	remainder := (width * percent) % 100
	if remainder > 0 && filled < width {
		half = styles.ProgressBarStyle.Render("▓")
		empty--
	}

	bar := styles.ProgressBarStyle.Render(strings.Repeat("█", filled))
	track := styles.ProgressTrackStyle.Render(strings.Repeat("░", max(0, empty)))

	return "[" + bar + half + track + "]"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Percentage returns a percentage label, e.g. "50%".
func Percentage(percent int) string {
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	return styles.SplashTextStyle.Render(fmt.Sprintf("%d%%", percent))
}
