package tui

import (
	"strings"
	"testing"

	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
)

func TestFormatSectionBodyBullets(t *testing.T) {
	body := "Lead stays separate.\n\n• First constraint\n• Second constraint with detail"
	// First chunk is paragraph, second is bullets - full body split:
	body = "• Alpha\n• Beta"
	out := formatSectionBody(body, 60)
	if !strings.Contains(out, "Alpha") || !strings.Contains(out, "Beta") {
		t.Fatalf("expected bullets, got %q", out)
	}
}

func TestFormatSectionBodyTitledBlock(t *testing.T) {
	body := "Dashboard\nNext.js UI for operators to monitor sensor feeds in real time."
	out := formatSectionBody(body, 60)
	if !strings.Contains(out, "Dashboard") {
		t.Fatalf("expected titled block, got %q", out)
	}
}

func TestRenderProjectDetailContent(t *testing.T) {
	m := newTestApp()
	m.width, m.height = 100, 40
	m.currentScreen = ScreenProjectDetail
	m.projectDetail = m.projects[0]

	out := m.renderProjectDetailContent()
	if !strings.Contains(out, m.projectDetail.Name) {
		t.Fatalf("expected project name in detail, got %q", out)
	}
	// Section headers are numbered and driven by the section's own label.
	if !strings.Contains(out, "01 · ") {
		t.Fatalf("expected numbered section header, got %q", out)
	}
	if len(m.projectDetail.Slug) > 0 {
		if cs := data.GetCaseStudy(m.projectDetail.Slug); cs != nil && len(cs.Sections) > 0 {
			if !strings.Contains(out, strings.ToUpper(cs.Sections[0].Label)) {
				t.Fatalf("expected first section label in detail, got %q", out)
			}
		}
	}
}
