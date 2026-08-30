package styles

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

func TestStylesHaveColorInDaemonEnvironment(t *testing.T) {
	out := TitleStyle.Render("Test Title")
	if !strings.Contains(out, "\x1b[") {
		t.Errorf("expected TitleStyle to contain ANSI escape code, got %q", out)
	}

	outNormal := NormalStyle.Render("Test Normal")
	if !strings.Contains(outNormal, "\x1b[") {
		t.Errorf("expected NormalStyle to contain ANSI escape code, got %q", outNormal)
	}

	outCard := CardStyle.Render("Card content")
	if !strings.Contains(outCard, "\x1b[") {
		t.Errorf("expected CardStyle to contain ANSI escape code, got %q", outCard)
	}
}

func TestStylesWithSessionRenderer(t *testing.T) {
	sessRenderer := lipgloss.NewRenderer(nil)
	sessRenderer.SetColorProfile(termenv.TrueColor)
	UseRenderer(sessRenderer)

	out := TitleStyle.Render("Test Title")
	if !strings.Contains(out, "\x1b[") {
		t.Errorf("expected TitleStyle to contain ANSI escape code after UseRenderer, got %q", out)
	}
}
