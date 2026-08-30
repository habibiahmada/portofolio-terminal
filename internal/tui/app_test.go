package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
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
	m.selectedMenu = 1

	m.Update(keyMsg(tea.KeyDown))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected Projects after ↓ from About, got %v", m.currentScreen)
	}
}

func TestSelectFromHomeNoOp(t *testing.T) {
	m := newTestApp()

	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenHome {
		t.Errorf("expected Enter on Home to stay on Home, got %v", m.currentScreen)
	}
}

func TestEnterProjectsShowsList(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjectDetail
	m.projectDetail = m.projects[0]

	_, _ = m.enterMenuScreen(2)
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected Projects list, got %v", m.currentScreen)
	}
}

func TestProjectDetailNavigation(t *testing.T) {
	m := newTestApp()

	m.currentScreen = ScreenProjects
	m.focus = FocusContent
	m.Update(keyMsg(tea.KeyEnter))
	if m.currentScreen != ScreenProjectDetail {
		t.Fatalf("expected ProjectDetail, got %v", m.currentScreen)
	}

	// Esc goes back to the project list (← now navigates prev/next project).
	m.Update(keyMsg(tea.KeyEsc))
	if m.currentScreen != ScreenProjects {
		t.Errorf("expected to go back to Projects, got %v", m.currentScreen)
	}
}

func TestNavigateProjectsList(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects

	// Arrow/right must first drop into the screen's content before the list
	// selection moves.
	m.Update(keyMsg(tea.KeyRight))
	if m.focus != FocusContent {
		t.Fatalf("expected focus content after →, got %v", m.focus)
	}

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

	// Top-level screens only return focus to the nav; switching screens is
	// done with ↑/↓ while the nav is focused.
	m.focus = FocusContent
	m.Update(keyMsg(tea.KeyLeft))
	if m.focus != FocusNav {
		t.Errorf("expected focus back to nav, got %v", m.focus)
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

func TestSelectModeToggle(t *testing.T) {
	m := newTestApp()

	model, cmd := m.Update(keyRunes('s'))
	app := model.(*App)
	if !app.selectMode {
		t.Error("expected selectMode to be enabled after 's'")
	}
	if cmd == nil {
		t.Error("expected a command (disable mouse) after enabling select mode")
	}

	model, cmd = app.Update(keyRunes('s'))
	if model.(*App).selectMode {
		t.Error("expected selectMode to be disabled after second 's'")
	}
	if cmd == nil {
		t.Error("expected a command (enable mouse) after disabling select mode")
	}
}

func TestMouseIgnoredInSelectMode(t *testing.T) {
	m := newTestApp()
	m.selectMode = true
	m.showHelp = true

	m.Update(tea.MouseMsg{
		X:      10,
		Y:      10,
		Action: tea.MouseActionPress,
		Button: tea.MouseButtonLeft,
	})
	if !m.showHelp {
		t.Error("expected mouse clicks to be ignored in select mode")
	}
}

func mousePress(x, y int) tea.MouseMsg {
	return tea.MouseMsg{
		X:      x,
		Y:      y,
		Action: tea.MouseActionPress,
		Button: tea.MouseButtonLeft,
	}
}

func TestNavClickSwitchesScreen(t *testing.T) {
	m := newTestApp()
	m.bodyTop = mastheadLines
	m.shellLeft = 2

	// menuItems order: Home(idx 0), About(idx 1), Projects(idx 2), ...
	projectsIdx := 2
	m.Update(mousePress(m.shellLeft+5, m.bodyTop+m.bodyTopPad()+projectsIdx))

	if m.currentScreen != ScreenProjects {
		t.Errorf("expected click on nav row %d to open Projects, got %v", projectsIdx, m.currentScreen)
	}
}

func TestScrollbarClickJumpsExactColumn(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenAbout
	m.scrollBarX = 20
	m.scrollBarTop = 3
	m.scrollBarH = 10
	m.scrollMax = 40

	// Top of the track maps to the top of the content.
	m.Update(mousePress(m.scrollBarX, m.scrollBarTop))
	if m.contentOffset != 0 {
		t.Errorf("expected offset 0 at track top, got %d", m.contentOffset)
	}

	// Bottom of the track maps to the bottom of the content.
	m.Update(mousePress(m.scrollBarX, m.scrollBarTop+m.scrollBarH-1))
	if m.contentOffset != 40 {
		t.Errorf("expected offset 40 at track bottom, got %d", m.contentOffset)
	}

	// One column to the right of the track must NOT trigger the scrollbar.
	m.contentOffset = 5
	m.Update(mousePress(m.scrollBarX+1, m.scrollBarTop+m.scrollBarH-1))
	if m.contentOffset != 5 {
		t.Errorf("expected offset unchanged when clicking beside the track, got %d", m.contentOffset)
	}
}

func TestProjectRowClickOpensDetail(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects
	m.bodyTop = mastheadLines
	m.shellLeft = 2
	m.contentOffset = 0

	textLeft := m.shellLeft + bodyFrameChrome - components.ScrollbarWidth
	i := 1
	m.Update(mousePress(textLeft+10, m.bodyTop+m.bodyTopPad()+4+i))

	if m.currentScreen != ScreenProjectDetail {
		t.Fatalf("expected ProjectDetail after clicking project row %d, got %v", i, m.currentScreen)
	}
	if m.projectDetail.Slug != m.projects[i].Slug {
		t.Errorf("expected project %q selected, got %q", m.projects[i].Slug, m.projectDetail.Slug)
	}
}

func TestScrollContent(t *testing.T) {
	m := newTestApp()

	// Only content taller than the viewport scrolls; short About content fits
	// so j/k are no-ops. Simulate a long screen via scrollMax.
	m.currentScreen = ScreenAbout
	m.focus = FocusContent
	m.scrollMax = 5
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

func TestScrollClampedAtBottom(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenAbout
	m.focus = FocusContent
	m.scrollMax = 5
	m.contentOffset = 5

	m.Update(keyRunes('j'))
	if m.contentOffset != 5 {
		t.Errorf("expected offset to stay 5 at bottom, got %d", m.contentOffset)
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
	want := m.height - mastheadLines - footerHeight - m.bodyTopPad()
	if m.contentHeight() != want {
		t.Errorf("expected content height %d, got %d", want, m.contentHeight())
	}
}

func TestResponsiveBodyTopPad(t *testing.T) {
	m := newTestApp()

	tests := []struct {
		height int
		want   int
	}{
		{18, 0},
		{22, 0},
		{24, 1},
		{30, 1},
		{35, 2},
		{40, 2},
		{45, 3},
		{52, 3},
		{60, 4},
		{70, 5},
	}

	for _, tt := range tests {
		m.height = tt.height
		if got := m.bodyTopPad(); got != tt.want {
			t.Errorf("height %d: expected top pad %d, got %d", tt.height, tt.want, got)
		}
	}
}

func TestMenuItemsOrder(t *testing.T) {
	want := []Screen{
		ScreenHome,
		ScreenAbout,
		ScreenProjects,
		ScreenSkills,
		ScreenExperience,
		ScreenCertificates,
		ScreenServices,
		ScreenContact,
	}
	for i, s := range want {
		if menuItems[i] != s {
			t.Errorf("menuItems[%d] = %v, want %v", i, menuItems[i], s)
		}
	}
}

func TestSelectEntersServices(t *testing.T) {
	cases := []struct {
		index int
		want  Screen
	}{
		{6, ScreenServices},
	}
	for _, c := range cases {
		m := newTestApp()
		_, _ = m.enterMenuScreen(c.index)
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

func TestViewLinesFitTerminal(t *testing.T) {
	sizes := []struct{ w, h int }{
		{80, 24},
		{100, 30},
		{120, 40},
	}
	screens := []Screen{ScreenHome, ScreenCertificates, ScreenAbout}
	for _, size := range sizes {
		for _, s := range screens {
			m := New()
			m.width, m.height = size.w, size.h
			m.currentScreen = s
			view := m.View()
			for i, ln := range strings.Split(stripANSI(view), "\n") {
				if w := lipgloss.Width(ln); w > size.w {
					t.Errorf("%v %dx%d line %d width %d > %d: %q", s, size.w, size.h, i, w, size.w, ln)
				}
			}
		}
	}
}

func TestProjectsListFocusExclusive(t *testing.T) {
	m := newTestApp()
	m.currentScreen = ScreenProjects
	m.selectedProject = 0

	// Nav focused: sidebar owns the strong cue, list stays plain.
	m.focus = FocusNav
	for _, ln := range strings.Split(stripANSI(m.View()), "\n") {
		if strings.Contains(ln, m.projects[0].Name) && strings.HasPrefix(strings.TrimSpace(ln), ">") {
			t.Error("expected no list cursor while nav is focused")
		}
	}

	// Content focused: list owns the strong cue, nav selection stays dim.
	m.focus = FocusContent
	found := false
	for _, ln := range strings.Split(stripANSI(m.View()), "\n") {
		if strings.Contains(ln, m.projects[0].Name) && strings.Contains(ln, ">") {
			found = true
		}
		if strings.Contains(ln, "▌ Projects") || strings.Contains(ln, "▍ Projects") {
			t.Error("expected no animated nav cursor while content is focused")
		}
	}
	if !found {
		t.Error("expected list cursor when content is focused")
	}
}

func TestNavigateChipNotOnName(t *testing.T) {
	m := newTestApp()
	plain := stripANSI(m.View())
	for _, ln := range strings.Split(plain, "\n") {
		if strings.Contains(ln, "NAVIGATE") && strings.Contains(ln, "Habibi") {
			t.Errorf("NAVIGATE overlapped the name: %q", ln)
		}
	}
	if !strings.Contains(plain, "NAVIGATE") {
		t.Error("expected NAVIGATE chip on its own row")
	}
	if !strings.Contains(plain, "Habibi Ahmad Aziz") {
		t.Error("expected full name on one line")
	}
}

func TestMascotMouseInteractivity(t *testing.T) {
	m := newTestApp()
	m.width = 140
	m.height = 30
	_ = m.View() // Render to compute geometry

	if m.mascotWidth == 0 || m.mascotHeight == 0 {
		t.Fatal("expected mascot to be positioned in right margin")
	}

	// 1 click: triggers Blink state
	clickMsg := mousePress(m.mascotLeft+1, m.mascotTop+1)
	m.Update(clickMsg)
	if m.mascotState != components.MascotStateBlink {
		t.Errorf("expected MascotStateBlink after 1 click, got %d", m.mascotState)
	}

	// 2nd click: remains Blink / refresh timer
	m.Update(clickMsg)
	if m.mascotState != components.MascotStateBlink {
		t.Errorf("expected MascotStateBlink after 2 clicks, got %d", m.mascotState)
	}

	// 3rd click: triggers Angry state!
	m.Update(clickMsg)
	if m.mascotState != components.MascotStateAngry {
		t.Errorf("expected MascotStateAngry after 3 clicks, got %d", m.mascotState)
	}

	view := stripANSI(m.View())
	if !strings.Contains(view, "╬") && !strings.Contains(view, "⑊") {
		t.Errorf("expected angry anger mark in rendered view when angry, got %q", view)
	}
}
