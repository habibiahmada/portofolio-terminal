package tui

import "testing"

func TestEnsureFeaturedSelectionVisibleScrollsDown(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenHome
	m.focus = FocusContent
	m.selectedFeatured = 0
	m.contentOffset = 0
	m.height = 22 // short viewport forces scrolling

	for _, p := range m.projects {
		p.Featured = true
	}
	// Ensure at least 4 featured cards.
	for len(m.featuredProjects()) < 4 {
		p := m.projects[0]
		p.Slug += "-feat"
		p.Name += " Featured"
		p.Featured = true
		m.projects = append(m.projects, p)
	}

	m.ensureFeaturedSelectionVisible()
	startOffset := m.contentOffset

	m.selectedFeatured = len(m.featuredProjects()) - 1
	m.ensureFeaturedSelectionVisible()

	if m.contentOffset <= startOffset {
		t.Fatalf("expected contentOffset to increase when selecting lower featured card, got %d -> %d", startOffset, m.contentOffset)
	}
}

func TestEnsureFeaturedSelectionVisibleScrollsUp(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenHome
	m.focus = FocusContent
	m.height = 22

	for i := range m.projects {
		m.projects[i].Featured = true
	}
	for len(m.featuredProjects()) < 4 {
		p := m.projects[0]
		p.Slug += "-x"
		p.Featured = true
		m.projects = append(m.projects, p)
	}

	m.selectedFeatured = len(m.featuredProjects()) - 1
	m.contentOffset = 200
	m.ensureFeaturedSelectionVisible()

	m.selectedFeatured = 0
	m.ensureFeaturedSelectionVisible()

	starts, _ := m.homeFeaturedLineMetrics()
	if m.contentOffset != starts[0] {
		t.Fatalf("expected contentOffset %d for first featured card, got %d", starts[0], m.contentOffset)
	}
}
