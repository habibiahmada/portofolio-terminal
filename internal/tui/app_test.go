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

func TestInstantScreenSwitchFromHome(t *testing.T) {
	m := newTestApp()

	m.Update(keyMsg(tea.KeyDown))
	if m.currentScreen != ScreenAbout {
		t.Errorf("expected About after ↓ from Home, got %v", m.currentScreen)
	}
}

func TestInstantScreenSwitchCycle(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenAbout
	m.selectedMenu = 0

	m.Update(keyMsg(tea.KeyDown))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected Projects after ↓ from About, got %v", m.currentScreen)
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

	m.currentScreen = ScreenProjects
	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenProjectDetail {
		t.Fatalf("expected ProjectDetail, got %v", m.currentScreen)
	}

	m.Update(keyMsg(tea.KeyLeft))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected to go back to Projects, got %v", m.currentScreen)
	}
}

func TestNavigateProjectsList(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects

	m.Update(keyRunes('j'))
	if m.selectedProject != 1 {
		t.Errorf("expected selectedProject 1 after j, got %d", m.selectedProject)
	}

	m.Update(keyRunes('k'))
	if m.selectedProject != 0 {
		t.Errorf("expected selectedProject 0 after k, got %d", m.selectedProject)
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
	m.Update(keyRunes('j'))
	if m.contentOffset != 1 {
		t.Errorf("expected offset 1 after j, got %d", m.contentOffset)
	}

	m.Update(keyRunes('k'))
	if m.contentOffset != 0 {
		t.Errorf("expected offset 0 after k, got %d", m.contentOffset)
	}

	m.Update(keyRunes('k'))
	if m.contentOffset != 0 {
		t.Errorf("expected offset to stay 0 at top, got %d", m.contentOffset)
	}
}

func TestClipScrollResetsOnShortContent(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenAbout
	m.contentOffset = 5
	out := m.clipScroll("short")
	if m.contentOffset != 0 {
		t.Errorf("expected offset reset to 0 for short content, got %d", m.contentOffset)
	}
	if out == "" {
		t.Error("expected clipped content to be non-empty")
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
		ScreenServices,
		ScreenBlog,
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

func TestMenuItemsOrder(t *testing.T) {
	want := []Screen{
		ScreenAbout,
		ScreenProjects,
		ScreenSkills,
		ScreenExperience,
		ScreenCertificates,
		ScreenServices,
		ScreenBlog,
		ScreenContact,
	}
	for i, s := range want {
		if menuItems[i] != s {
			t.Errorf("menuItems[%d] = %v, want %v", i, menuItems[i], s)
		}
	}
}

func TestSelectEntersServicesAndBlog(t *testing.T) {
	cases := []struct {
		index int
		want  Screen
	}{
		{5, ScreenServices},
		{6, ScreenBlog},
	}
	for _, c := range cases {
		m := newTestApp()
		m.selectedMenu = c.index
		m.Update(keyMsg(tea.KeyEnter))
		if m.currentScreen != c.want {
			t.Errorf("expected menu index %d to enter %v, got %v", c.index, c.want, m.currentScreen)
		}
	}
}

func TestHomeShortcuts(t *testing.T) {
	m := newTestApp()
	m.Update(keyRunes('P'))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected P to jump to Projects, got %v", m.currentScreen)
	}

	m = newTestApp()
	m.Update(keyRunes('C'))
	if m.currentScreen != ScreenContact {
		t.Errorf("expected C to jump to Contact, got %v", m.currentScreen)
	}

	m = newTestApp()
	m.Update(keyRunes('V'))
	if !m.cvModal {
		t.Error("expected V to open the CV modal")
	}
	m.Update(keyRunes('v'))
	if m.cvModal {
		t.Error("expected v to close the CV modal")
	}
}

func TestProjectPrevNext(t *testing.T) {
	m := newTestApp()
	m.projectDetail = m.projects[0]
	m.selectedProject = 0
	m.currentScreen = ScreenProjectDetail

	m.nextProject()
	if m.projectDetail.Slug != m.projects[1].Slug {
		t.Errorf("expected next project %q, got %q", m.projects[1].Slug, m.projectDetail.Slug)
	}

	m.prevProject()
	if m.projectDetail.Slug != m.projects[0].Slug {
		t.Errorf("expected prev project %q, got %q", m.projects[0].Slug, m.projectDetail.Slug)
	}
}

func TestFooterTickAdvances(t *testing.T) {
	m := newTestApp()
	start := m.footerFrame
	model, cmd := m.Update(footerTickMsg{})
	next := model.(*App)
	if next.footerFrame != start+1 {
		t.Errorf("expected footerFrame to advance by 1, got %d", next.footerFrame)
	}
	if cmd == nil {
		t.Error("expected a follow-up tick command after footerTickMsg")
	}
}

func TestFooterShowsBrandAndHints(t *testing.T) {
	m := newTestApp()
	view := m.View()
	if !strings.Contains(view, "habibiahmada") {
		t.Error("expected footer brand in view")
	}
	if !strings.Contains(view, "Screens") {
		t.Error("expected footer screen hint in view")
	}
}
