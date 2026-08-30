package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// UseRenderer sets the active lipgloss renderer (SSH session or local terminal).
func UseRenderer(r *lipgloss.Renderer) {
	if r == nil {
		return
	}
	lipgloss.SetDefaultRenderer(r)
}

// ForceTrueColor upgrades styling to 24-bit color. SSH clients often report a
// low color profile even when the terminal supports TrueColor; without this,
// hex palette colors collapse to pale ANSI greys.
func ForceTrueColor() {
	lipgloss.SetColorProfile(termenv.TrueColor)
	if r := lipgloss.DefaultRenderer(); r != nil {
		r.SetColorProfile(termenv.TrueColor)
	}
}
