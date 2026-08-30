package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

// Helpers to construct keyboard messages for tests.

func keyRunes(r ...rune) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: r}
}

func keyMsg(t tea.KeyType) tea.KeyMsg {
	return tea.KeyMsg{Type: t}
}

func teaMsgWindow(w, h int) tea.WindowSizeMsg {
	return tea.WindowSizeMsg{Width: w, Height: h}
}

func newTestApp() *App {
	m := New()
	m.width, m.height = 100, 30
	return m
}

func TestNew(t *testing.T) {
	m := New()

	if m.currentScreen != ScreenHome {
		t.Errorf("expected currentScreen Home, got %v", m.currentScreen)
	}
	if len(m.projects) == 0 {
		t.Error("expected projects to be loaded")
	}
	if m.profile.Name == "" {
		t.Error("expected profile to be loaded")
	}
	if m.profile.Email == "" {
		t.Error("expected profile Email to be set")
	}
	if m.quitting {
		t.Error("expected app to start not quitting")
	}
}

func TestNavigateMenu(t *testing.T) {
	m := newTestApp()

	m.Update(keyMsg(tea.KeyDown))
	if m.selectedMenu != 1 {
		t.Errorf("expected selectedMenu 1 after down, got %d", m.selectedMenu)
	}

	// Wraps to the last item.
	m.selectedMenu = 0
	m.Update(keyMsg(tea.KeyUp))
	if m.selectedMenu != len(menuItems)-1 {
		t.Errorf("expected up at top to wrap to %d, got %d", len(menuItems)-1, m.selectedMenu)
	}

	// Wraps to the first item.
	m.selectedMenu = len(menuItems) - 1
	m.Update(keyMsg(tea.KeyDown))
	if m.selectedMenu != 0 {
		t.Errorf("expected down at bottom to wrap to 0, got %d", m.selectedMenu)
	}
}

func TestSelectEntersAbout(t *testing.T) {
	m := newTestApp()

	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenAbout {
		t.Errorf("expected to enter About, got %v", m.currentScreen)
	}
}

func TestSelectEntersProjects(t *testing.T) {
	m := newTestApp()
	m.selectedMenu = 1 // Projects

	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected to enter Projects, got %v", m.currentScreen)
	}
	if m.selectedProject != 0 {
		t.Errorf("expected selectedProject reset to 0, got %d", m.selectedProject)
	}
}

func TestProjectDetailNavigation(t *testing.T) {
	m := newTestApp()

	// Go to Projects then select the first project.
	m.selectedMenu = 1
	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenProjects {
		t.Fatalf("expected Projects, got %v", m.currentScreen)
	}

	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenProjectDetail {
		t.Fatalf("expected ProjectDetail, got %v", m.currentScreen)
	}

	// Back to the list.
	m.Update(keyMsg(tea.KeyLeft))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected to go back to Projects, got %v", m.currentScreen)
	}
}

func TestNavigateProjectsList(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects

	m.Update(keyMsg(tea.KeyDown))
	if m.selectedProject != 1 {
		t.Errorf("expected selectedProject 1 after down, got %d", m.selectedProject)
	}

	m.Update(keyMsg(tea.KeyUp))
	if m.selectedProject != 0 {
		t.Errorf("expected selectedProject 0 after up, got %d", m.selectedProject)
	}
}

func TestGoBackFromScreen(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenSkills

	m.Update(keyMsg(tea.KeyLeft))
	if m.currentScreen != ScreenHome {
		t.Errorf("expected back to Home, got %v", m.currentScreen)
	}
}

func TestQuit(t *testing.T) {
	m := newTestApp()

	model, _ := m.Update(keyRunes('q'))
	if !model.(*App).quitting {
		t.Error("expected quitting after 'q'")
	}

	m = newTestApp()
	_, cmd := m.Update(keyMsg(tea.KeyCtrlC))
	if cmd == nil {
		t.Error("expected a quit command after Ctrl+C")
	}
}

func TestHelpToggle(t *testing.T) {
	m := newTestApp()

	m.Update(keyRunes('?'))
	if !m.showHelp {
		t.Error("expected help to be shown after '?'")
	}

	// '?' or Esc closes it.
	m.Update(keyRunes('?'))
	if m.showHelp {
		t.Error("expected help to be closed after second '?'")
	}

	m.Update(keyRunes('?'))
	m.Update(keyMsg(tea.KeyEsc))
	if m.showHelp {
		t.Error("expected help to close on Esc")
	}
}

func TestScrollContent(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenAbout

	// Push an artificially long content.
	longContent := strings.Repeat("line\n", 60)
	m.profile.Description = longContent

	m.Update(keyMsg(tea.KeyDown))
	if m.contentOffset != 1 {
		t.Errorf("expected offset 1 after down, got %d", m.contentOffset)
	}

	m.Update(keyMsg(tea.KeyUp))
	if m.contentOffset != 0 {
		t.Errorf("expected offset 0 after up, got %d", m.contentOffset)
	}

	// Up at top stays at 0.
	m.Update(keyMsg(tea.KeyUp))
	if m.contentOffset != 0 {
		t.Errorf("expected offset to stay 0 at top, got %d", m.contentOffset)
	}
}

func TestViewRendersScreens(t *testing.T) {
	screens := []Screen{
		ScreenHome,
		ScreenAbout,
		ScreenProjects,
		ScreenProjectDetail,
		ScreenSkills,
		ScreenExperience,
		ScreenCertificates,
		ScreenContact,
	}

	for _, s := range screens {
		m := newTestApp()
		m.currentScreen = s
		m.projectDetail = m.projects[0]

		view := m.View()
		if view == "" {
			t.Errorf("view for screen %v must not be empty", s)
		}
	}
}

func TestViewInitializing(t *testing.T) {
	m := New()
	if got := m.View(); !strings.Contains(got, "Initializing") {
		t.Errorf("expected Initializing view, got %q", got)
	}
}

func TestContentHeight(t *testing.T) {
	m := newTestApp()
	want := m.height - headerHeight - footerHeight
	if m.contentHeight() != want {
		t.Errorf("expected content height %d, got %d", want, m.contentHeight())
	}
}
