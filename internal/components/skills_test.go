package components

import (
	"strings"
	"testing"
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

func TestJoinShellHasDivider(t *testing.T) {
	rail := NavRail([]SidebarItem{{Name: "Home"}}, 0, navRailWidth)
	content := "content line"
	out := JoinShell(rail, content, 3)
	if !strings.Contains(out, "│") {
		t.Errorf("expected vertical divider, got %q", out)
	}
}
