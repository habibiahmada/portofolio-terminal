package tui

import "testing"

func TestEnsureProjectSelectionVisibleScrollsDown(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects
	m.focus = FocusContent
	m.selectedProject = 0
	m.contentOffset = 0

	// Simulate a long list: 30 projects, short viewport.
	for len(m.projects) < 30 {
		p := m.projects[len(m.projects)-1]
		p.Slug += "-copy"
		p.Name += " Copy"
		m.projects = append(m.projects, p)
	}

	m.ensureProjectSelectionVisible()
	startOffset := m.contentOffset

	m.selectedProject = 20
	m.ensureProjectSelectionVisible()

	if m.contentOffset <= startOffset {
		t.Fatalf("expected contentOffset to increase when selecting lower project, got %d -> %d", startOffset, m.contentOffset)
	}
}

func TestEnsureProjectSelectionVisibleScrollsUp(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects
	m.focus = FocusContent
	m.selectedProject = 25
	m.contentOffset = 100

	for len(m.projects) < 30 {
		p := m.projects[len(m.projects)-1]
		p.Slug += "-dup"
		p.Name += " Dup"
		m.projects = append(m.projects, p)
	}

	m.ensureProjectSelectionVisible()
	m.selectedProject = 0
	m.ensureProjectSelectionVisible()

	starts, _ := m.projectListMetrics()
	if m.contentOffset != starts[0] {
		t.Fatalf("expected contentOffset %d when selecting first project, got %d", starts[0], m.contentOffset)
	}
}
