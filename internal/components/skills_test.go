package components

import (
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestSkillGridWraps(t *testing.T) {
	names := []string{"React", "Next.js", "Node.js", "TypeScript"}
	out := SkillGrid(names, 30)
	if !strings.Contains(out, "React") {
		t.Errorf("expected skills in grid, got %q", out)
	}
	if strings.Count(out, "\n") < 1 {
		t.Error("expected wrapped rows at narrow width")
	}
}

func TestReflowBlockWrapsLongLines(t *testing.T) {
	long := "This is a long sentence that should wrap across multiple terminal columns when the width is narrow."
	out := ReflowBlock(long, 20)
	lines := strings.Split(out, "\n")
	if len(lines) < 2 {
		t.Errorf("expected wrapped lines, got %q", out)
	}
	for i, ln := range lines {
		if w := lipgloss.Width(ln); w > 20 {
			t.Errorf("line %d width %d > 20: %q", i, w, ln)
		}
	}
}

func TestNavRailWidthFor(t *testing.T) {
	cases := []struct {
		w, want int
	}{
		{30, 8},
		{39, 8},
		{40, 10},
		{50, 10},
		{59, 10},
		{60, 13},
		{70, 13},
		{79, 13},
		{80, 16},
		{100, 16},
	}
	for _, c := range cases {
		if got := NavRailWidthFor(c.w); got != c.want {
			t.Errorf("NavRailWidthFor(%d) = %d, want %d", c.w, got, c.want)
		}
	}
}

func TestFitLinePadsAndClips(t *testing.T) {
	if got := FitLine("ab", 4); got != "ab  " {
		t.Errorf("pad: got %q", got)
	}
	if w := lipgloss.Width(FitLine("abcdef", 3)); w != 3 {
		t.Errorf("clip width: got %d", w)
	}
}

func TestTruncate(t *testing.T) {
	if got := Truncate("hello", 10); got != "hello" {
		t.Errorf("short: got %q", got)
	}
	got := Truncate("hello world", 8)
	if lipgloss.Width(got) > 8 {
		t.Errorf("truncated width %d > 8 (%q)", lipgloss.Width(got), got)
	}
	if !strings.Contains(got, "…") {
		t.Errorf("expected ellipsis, got %q", got)
	}
}

func TestTagColumnUniformWidth(t *testing.T) {
	out := TagColumn([]string{"short", "a much longer certificate name"}, 40)
	lines := strings.Split(out, "\n")
	if len(lines) < 2 {
		t.Fatalf("expected multiple lines, got %q", out)
	}
	w0 := lipgloss.Width(lines[0])
	for i, ln := range lines {
		if lipgloss.Width(ln) != w0 {
			t.Errorf("line %d width %d != %d (%q)", i, lipgloss.Width(ln), w0, ln)
		}
	}
}

func TestJoinShellHasDivider(t *testing.T) {
	rail := NavRail([]SidebarItem{{Name: "Home"}}, 0, navRailWidth, 0, true)
	content := "content line"
	out := JoinShell(rail, content, 3)
	if !strings.Contains(out, "│") {
		t.Errorf("expected vertical divider, got %q", out)
	}
}
