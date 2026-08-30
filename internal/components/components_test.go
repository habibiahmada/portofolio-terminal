package components

import (
	"strings"
	"testing"
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

	// selectedIndex renders ">", activeKey renders "▸".
	out := Sidebar(items, 1, "Skills", 10)
	if !strings.Contains(out, "> Projects") {
		t.Errorf("expected selected item to be marked with '>', got %q", out)
	}
	if !strings.Contains(out, "▸ Skills") {
		t.Errorf("expected active item to be marked with '▸', got %q", out)
	}
	if !strings.Contains(out, "About") {
		t.Errorf("expected unselected item to still render, got %q", out)
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
