package components

import (
	"regexp"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestHeader(t *testing.T) {
	out := Header("habibiahmada", "Fullstack Developer", 80)
	if !strings.Contains(out, "habibiahmada") {
		t.Errorf("expected header to contain wordmark, got %q", out)
	}
	if !strings.Contains(out, "Fullstack Developer") {
		t.Errorf("expected header to contain title, got %q", out)
	}
}

func TestFooterArtlineRenders(t *testing.T) {
	out := FooterArtline(0, 90, "habibiahmada.dev")
	if out == "" {
		t.Fatal("expected footer brand to render")
	}
	if !strings.Contains(out, ">_") {
		t.Errorf("expected footer brand to contain the >_ prompt, got %q", out)
	}
	if !strings.Contains(out, "habibiahmada") {
		t.Errorf("expected footer brand to show wordmark, got %q", out)
	}
	if !strings.Contains(out, "▁") {
		t.Errorf("expected footer brand to include the equalizer wave, got %q", out)
	}
}

func TestFooterArtlineAnimates(t *testing.T) {
	a := FooterArtline(0, 90, "habibiahmada.dev")
	b := FooterArtline(3, 90, "habibiahmada.dev")
	if a == b {
		t.Errorf("expected footer artline to change with frame (got identical output %q)", a)
	}
}

func TestFooterArtlineTinyDegrades(t *testing.T) {
	out := FooterArtline(3, 30, "habibiahmada.dev")
	if !strings.Contains(out, "habibiahmada") {
		t.Errorf("expected tiny footer to keep the wordmark, got %q", out)
	}
}

func TestFooterBarBrandAndHints(t *testing.T) {
	hints := []FooterHint{{Key: "↑↓", Label: "Screens"}}
	out := FooterBar(0, 80, hints)
	if !strings.Contains(out, "habibiahmada") {
		t.Errorf("expected brand on the left, got %q", out)
	}
	if !strings.Contains(out, "Screens") {
		t.Errorf("expected hints on the right, got %q", out)
	}
}

func TestSidebarHighlights(t *testing.T) {
	items := []SidebarItem{
		{Key: "About", Name: "About"},
		{Key: "Projects", Name: "Projects"},
		{Key: "Skills", Name: "Skills"},
	}

	out := NavRail(items, 1, navRailWidth, 0, true)
	if !strings.Contains(out, "Projects") {
		t.Errorf("expected selected item present, got %q", out)
	}
	if !strings.Contains(out, "About") {
		t.Errorf("expected unselected item to still render, got %q", out)
	}
	if !strings.Contains(out, "◤ NAV") {
		t.Errorf("expected nav focus hint, got %q", out)
	}
}

func TestFooter(t *testing.T) {
	hints := []FooterHint{
		{Key: "↑↓", Label: "Navigate"},
		{Key: "Enter", Label: "Select"},
	}
	out := Footer(hints)
	if !strings.Contains(out, "Navigate") || !strings.Contains(out, "Select") {
		t.Errorf("expected footer to contain hints, got %q", out)
	}
}

func TestCard(t *testing.T) {
	out := Card("Title", "Body")
	if out == "" {
		t.Error("expected card to render content")
	}
	if !strings.Contains(out, "Title") || !strings.Contains(out, "Body") {
		t.Errorf("expected card to contain lines, got %q", out)
	}
}

func TestListHighlightsSelected(t *testing.T) {
	out := List([]string{"one", "two", "three"}, 1)
	if !strings.Contains(out, "▸ two") {
		t.Errorf("expected selected item to be marked, got %q", out)
	}
}

func TestTagList(t *testing.T) {
	out := TagList([]string{"Go", "React"})
	if !strings.Contains(out, "Go") || !strings.Contains(out, "React") {
		t.Errorf("expected tag list to contain tags, got %q", out)
	}
}

func TestSection(t *testing.T) {
	out := Section("Header", "body line")
	if !strings.Contains(out, "Header") || !strings.Contains(out, "body line") {
		t.Errorf("expected section to contain title and body, got %q", out)
	}
}

func TestModal(t *testing.T) {
	out := Modal("Title", []string{"line one", "line two"}, 80, 24)
	if out == "" {
		t.Error("expected modal to render")
	}
	if !strings.Contains(out, "line one") {
		t.Errorf("expected modal to contain lines, got %q", out)
	}
}

func TestIllustrationVariant(t *testing.T) {
	cases := []struct {
		width int
		want  Variant
	}{
		{200, VariantWide},
		{120, VariantWide},
		{100, VariantWide},
		{99, VariantComfortable},
		{80, VariantComfortable},
		{79, VariantNarrow},
		{60, VariantNarrow},
		{59, VariantHidden},
		{30, VariantHidden},
	}
	for _, c := range cases {
		if got := IllustrationVariant(c.width); got != c.want {
			t.Errorf("IllustrationVariant(%d) = %v, want %v", c.width, got, c.want)
		}
	}
}

func TestSignatureHiddenOnNarrow(t *testing.T) {
	if got := Signature(30); got != "" {
		t.Errorf("expected empty signature on 30 cols, got %q", got)
	}
}

func TestSignatureRendersWide(t *testing.T) {
	out := Signature(120)
	if out == "" {
		t.Fatal("expected wide signature to render")
	}
	// The signature must keep its branding text.
	if !strings.Contains(out, "HABIBI") {
		t.Errorf("expected signature to contain HABIBI, got %q", out)
	}
}

func TestProgressBar(t *testing.T) {
	out := ProgressBar(50, 20)
	if !strings.Contains(out, "█") || !strings.Contains(out, "░") {
		t.Errorf("expected progress bar with filled and track segments, got %q", out)
	}

	if full := ProgressBar(100, 20); strings.Contains(full, "░") {
		t.Errorf("expected full bar to have no track, got %q", full)
	}
}

func TestMascotFigureRendersLayeredMonitor(t *testing.T) {
	out := MascotFigure(1)
	if out == "" {
		t.Fatal("expected mascot to render")
	}
	if !strings.Contains(out, "╔") || !strings.Contains(out, "╭") {
		t.Errorf("expected thick and thin frame glyphs in mascot, got %q", out)
	}
	if !strings.Contains(out, "●") {
		t.Errorf("expected centered eyes in mascot, got %q", out)
	}
	if strings.Contains(out, ">_") {
		t.Errorf("expected face without prompt text, got %q", out)
	}
}

func TestMascotFigureAnimates(t *testing.T) {
	a := MascotFigure(0)
	b := MascotFigure(5)
	if a == b {
		t.Errorf("expected mascot to change with frame, got identical output")
	}
}

func TestMascotFigureInteractiveStates(t *testing.T) {
	// Blink state (clicked 1-2 times) has blink / happy squint / sparkle
	blink := MascotFigure(0, MascotStateBlink)
	if !strings.Contains(blink, "─") && !strings.Contains(blink, "^") && !strings.Contains(blink, "★") {
		t.Errorf("expected blinking/happy glyphs in blink state, got %q", blink)
	}

	// Angry state (clicked repeatedly) has anger mark ╬ and angry eyes, no emojis
	angry := MascotFigure(0, MascotStateAngry)
	if !strings.Contains(angry, "╬") && !strings.Contains(angry, "⑊") && !strings.Contains(angry, "⁑") {
		t.Errorf("expected non-emoji anger mark (╬/⑊/⁑) in angry state, got %q", angry)
	}
	if !strings.Contains(angry, "◣") && !strings.Contains(angry, "\\") {
		t.Errorf("expected angry eyes in angry state, got %q", angry)
	}
}

func TestMascotFigureDimensionsAndAlignment(t *testing.T) {
	states := []int{MascotStateNormal, MascotStateBlink, MascotStateAngry}
	for _, st := range states {
		for f := 0; f < 20; f++ {
			fig := MascotFigure(f, st)
			lines := strings.Split(fig, "\n")
			if len(lines) != 6 {
				t.Fatalf("expected 6 lines for state %d frame %d, got %d", st, f, len(lines))
			}
			for i, ln := range lines {
				if w := lipgloss.Width(ln); w != 14 {
					t.Errorf("state %d frame %d line %d: expected width 14, got %d: %q", st, f, i, w, ln)
				}
			}
		}
	}
}

func TestFigureFiglet(t *testing.T) {
	out := Figlet("HABIBI", 120)
	if out == "" {
		t.Fatal("expected figlet output")
	}
	// Large-font rendering spreads letters across multiple lines/blocks.
	if len(out) < len("HABIBI")*4 {
		t.Errorf("expected substantial figlet output, got %q", out)
	}
}

func TestHeroBannerUsesContourShadowsNotAccentBlocks(t *testing.T) {
	out := HeroBanner("HABIBI", 120)
	if out == "" {
		t.Fatal("expected hero banner to render")
	}
	plain := regexp.MustCompile("\x1b\\[[0-9;]*m").ReplaceAllString(out, "")
	if strings.Contains(plain, "▓") {
		t.Errorf("hero banner must not use solid accent blocks, got %q", plain)
	}
	if !strings.Contains(plain, "─") && !strings.Contains(plain, "│") {
		t.Errorf("hero banner should include contour shadow lines, got %q", plain)
	}
	if !strings.Contains(plain, "█") {
		t.Errorf("hero banner should include white letter fills, got %q", plain)
	}
	if strings.Contains(plain, "┼") {
		t.Errorf("hero banner must not use plus-shaped shadow joints, got %q", plain)
	}
	if strings.ContainsAny(plain, "┘┐└┌") {
		t.Errorf("hero banner must not curl shadow corners into letter fills, got %q", plain)
	}
}
