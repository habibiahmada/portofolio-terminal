// Package tui implements the terminal user interface using Bubble Tea.
// The App model is the top-level router that composes screens.
//
// Navigation and layout logic lives here; each screen's content renderer
// lives in its own file (home.go, about.go, projects.go, ...).
package tui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/data"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// contentOffsetX tracks the horizontal offset of the centered body for mouse
// click coordinate translation.
var contentOffsetX int

// contentOffsetY tracks the vertical offset of the centered body for mouse
// click coordinate translation.
var contentOffsetY int

// Screen identifiers.
type Screen int

const (
	ScreenHome Screen = iota
	ScreenAbout
	ScreenProjects
	ScreenProjectDetail
	ScreenSkills
	ScreenExperience
	ScreenCertificates
	ScreenServices
	ScreenContact
)

// ScreenNames maps screen identifiers to display names.
var ScreenNames = map[Screen]string{
	ScreenHome:          "Home",
	ScreenAbout:         "About",
	ScreenProjects:      "Projects",
	ScreenProjectDetail: "Project Detail",
	ScreenSkills:        "Skills",
	ScreenExperience:    "Experience",
	ScreenCertificates:  "Certificates",
	ScreenServices:      "Services",
	ScreenContact:       "Contact",
}

// menuItems is the ordered navigation list (Home first, then site sections).
var menuItems = []Screen{
	ScreenHome,
	ScreenAbout,
	ScreenProjects,
	ScreenSkills,
	ScreenExperience,
	ScreenCertificates,
	ScreenServices,
	ScreenContact,
}

// Layout dimensions.
const (
	mastheadLines = 1
	footerHeight  = 2
)

// footerTickInterval drives the footer animation. Kept slow to avoid rebuilding
// the full layout too often (was 120ms — caused high CPU/RAM on large screens).
const footerTickInterval = 600 * time.Millisecond

type App struct {
	width  int
	height int

	// Navigation state.
	currentScreen Screen
	selectedMenu  int
	prevScreen    Screen

	// Child screen state.
	selectedProject int
	projectDetail   data.Project

	// Data.
	profile      data.Profile
	projects     []data.Project
	skills       []data.Skill
	work         []data.ExperienceWork
	education    []data.ExperienceEducation
	certificates []data.Certificate
	socials      []data.Social
	companies    []data.Company
	services     []data.Service
	process      []data.ProcessStep
	press        []data.Press

	// View state.
	contentOffset int
	showHelp      bool

	// CV modal viewer (V from Home).
	cvModal bool

	// Footer animation frame.
	footerFrame int

	// Layout cache — excludes footerFrame so animation is cheap.
	cachedBodyKey string
	cachedBody    string

	// Flags.
	quitting bool
}

// New creates a new App model with bundled data.
func New() *App {
	return &App{
		currentScreen: ScreenHome,
		selectedMenu:  0,
		profile:       data.GetProfile(),
		projects:      data.GetProjects(),
		skills:        data.GetSkills(),
		work:          data.GetWorkExperience(),
		education:     data.GetEducation(),
		certificates:  data.GetCertificates(),
		socials:       data.GetSocials(),
		companies:     data.GetCompanies(),
		services:      data.GetServices(),
		process:       data.GetProcessSteps(),
		press:         data.GetPress(),
	}
}

// footerTickMsg advances the footer animation by one frame.
type footerTickMsg struct{}

func nextFooterTick() tea.Cmd {
	return tea.Tick(footerTickInterval, func(time.Time) tea.Msg { return footerTickMsg{} })
}

// Init implements tea.Model.
func (m *App) Init() tea.Cmd {
	return nextFooterTick()
}

// Update implements tea.Model.
func (m *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)

	case tea.MouseMsg:
		return m.handleMouse(msg)

	case footerTickMsg:
		m.footerFrame++
		return m, nextFooterTick()
	}

	return m, nil
}

// handleMouse processes mouse click and scroll events.
func (m *App) handleMouse(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Only handle left-click press and wheel scroll.
	if msg.Action != tea.MouseActionPress && msg.Button != tea.MouseButtonWheelUp && msg.Button != tea.MouseButtonWheelDown {
		return m, nil
	}

	// Help overlay intercepts clicks.
	if m.showHelp {
		m.showHelp = false
		return m, nil
	}

	// CV modal intercepts clicks.
	if m.cvModal {
		m.cvModal = false
		return m, nil
	}

	// Translate click coordinates relative to the centered body.
	_ = msg.X - contentOffsetX
	y := msg.Y - contentOffsetY

	// Wheel scroll.
	if msg.Button == tea.MouseButtonWheelUp {
		m.scrollOrSelectUp()
		return m, nil
	}
	if msg.Button == tea.MouseButtonWheelDown {
		m.scrollOrSelectDown()
		return m, nil
	}

	// Click content to open detail on list screens.
	if msg.Action == tea.MouseActionPress {
		if y >= 0 && m.currentScreen == ScreenProjects {
			return m.selectItem()
		}
	}

	return m, nil
}

// handleKey processes keyboard input.
func (m *App) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Help overlay intercepts keys first.
	if m.showHelp {
		if isBack(msg) || isQuit(msg) || isHelp(msg) {
			m.showHelp = false
		}
		return m, nil
	}

	// Global quit.
	if isQuit(msg) {
		m.quitting = true
		return m, tea.Quit
	}

	// Toggle help.
	if isHelp(msg) {
		m.showHelp = true
		return m, nil
	}

	// CV overlay intercepts keys first.
	if m.cvModal {
		if isBack(msg) || isSelect(msg) || isQuit(msg) || isHelp(msg) ||
			msg.String() == "P" || msg.String() == "p" ||
			msg.String() == "C" || msg.String() == "c" ||
			msg.String() == "V" || msg.String() == "v" {
			m.cvModal = false
		}
		return m, nil
	}

	// Home shortcuts: jump directly to key screens.
	if m.currentScreen == ScreenHome {
		switch msg.String() {
		case "P", "p":
			return m, m.enterProjects()
		case "C", "c":
			return m, m.enterContact()
		case "V", "v":
			m.cvModal = true
			return m, nil
		}
	}

	// Project detail prev/next navigation (`h`/`l`). `←`/Esc remains Back.
	if m.currentScreen == ScreenProjectDetail {
		if msg.String() == "h" {
			m.prevProject()
			return m, nil
		}
		if msg.String() == "l" || msg.Type == tea.KeyRight {
			m.nextProject()
			return m, nil
		}
	}

	// Screen switch (↑↓) — instant, no Enter required.
	if isNavigateUp(msg) {
		return m.switchScreen(-1)
	}
	if isNavigateDown(msg) {
		return m.switchScreen(1)
	}

	// Scroll / list selection (j/k).
	if isScrollUp(msg) {
		m.scrollOrSelectUp()
		return m, nil
	}
	if isScrollDown(msg) {
		m.scrollOrSelectDown()
		return m, nil
	}

	// Select / drill-down.
	if isSelect(msg) {
		return m.selectItem()
	}

	// Back.
	if isBack(msg) {
		return m.goBack()
	}

	return m, nil
}

// switchScreen moves to the next/previous nav item immediately.
func (m *App) switchScreen(delta int) (tea.Model, tea.Cmd) {
	idx := m.navSelectedIndex()
	idx += delta
	if idx < 0 {
		idx = len(menuItems) - 1
	}
	if idx >= len(menuItems) {
		idx = 0
	}
	m.selectedMenu = idx
	return m.enterMenuScreen(idx)
}

// syncMenuFromScreen aligns selectedMenu with the current screen.
func (m *App) syncMenuFromScreen() {
	m.selectedMenu = m.navSelectedIndex()
}

// enterMenuScreen activates the screen at menuItems[idx].
func (m *App) enterMenuScreen(idx int) (tea.Model, tea.Cmd) {
	m.prevScreen = ScreenHome
	m.resetScroll()

	screen := menuItems[idx]
	m.selectedMenu = idx
	m.currentScreen = screen

	// Always land on the projects list — never auto-open detail when switching screens.
	if screen == ScreenProjects {
		m.currentScreen = ScreenProjects
		m.selectedProject = 0
	}

	return m, nil
}

// scrollOrSelectUp handles j/k up: list selection or content scroll.
func (m *App) scrollOrSelectUp() {
	switch m.currentScreen {
	case ScreenProjects:
		m.selectedProject--
		if m.selectedProject < 0 {
			m.selectedProject = len(m.projects) - 1
		}
	default:
		m.scrollUp()
	}
}

// scrollOrSelectDown handles j/k down: list selection or content scroll.
func (m *App) scrollOrSelectDown() {
	switch m.currentScreen {
	case ScreenProjects:
		m.selectedProject++
		if m.selectedProject >= len(m.projects) {
			m.selectedProject = 0
		}
	default:
		m.scrollDown()
	}
}

// projectIndex returns the index of the currently open project in the list.
func (m *App) projectIndex() int {
	for i, p := range m.projects {
		if p.Slug == m.projectDetail.Slug {
			return i
		}
	}
	return m.selectedProject
}

// prevProject opens the previous project in the list.
func (m *App) prevProject() {
	if len(m.projects) > 0 {
		idx := m.projectIndex()
		idx--
		if idx < 0 {
			idx = len(m.projects) - 1
		}
		m.projectDetail = m.projects[idx]
		m.selectedProject = idx
		m.resetScroll()
	}
}

// nextProject opens the next project in the list.
func (m *App) nextProject() {
	if len(m.projects) > 0 {
		idx := m.projectIndex()
		idx++
		if idx >= len(m.projects) {
			idx = 0
		}
		m.projectDetail = m.projects[idx]
		m.selectedProject = idx
		m.resetScroll()
	}
}

// enterProjects jumps to the Projects screen.
func (m *App) enterProjects() tea.Cmd {
	m.syncMenuFromScreen()
	for i, s := range menuItems {
		if s == ScreenProjects {
			m.selectedMenu = i
			break
		}
	}
	_, cmd := m.enterMenuScreen(m.selectedMenu)
	return cmd
}

// enterContact jumps to the Contact screen.
func (m *App) enterContact() tea.Cmd {
	for i, s := range menuItems {
		if s == ScreenContact {
			m.selectedMenu = i
			break
		}
	}
	_, cmd := m.enterMenuScreen(m.selectedMenu)
	return cmd
}

// selectItem opens a drill-down view (e.g. project detail).
func (m *App) selectItem() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case ScreenHome:
		return m, nil
	case ScreenProjects:
		if len(m.projects) == 0 {
			return m, nil
		}
		m.prevScreen = m.currentScreen
		m.projectDetail = m.projects[m.selectedProject]
		m.resetScroll()
		m.currentScreen = ScreenProjectDetail
	}
	return m, nil
}

// goBack returns to the previous screen.
func (m *App) goBack() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case ScreenProjectDetail:
		m.currentScreen = ScreenProjects
	case ScreenHome:
		m.quitting = true
		return m, tea.Quit
	default:
		m.currentScreen = ScreenHome
	}
	m.resetScroll()
	return m, nil
}

// contentHeight returns visible lines for page content beside the sidebar.
func (m *App) contentHeight() int {
	bodyH := m.height - footerHeight
	h := bodyH - mastheadLines - 2
	if h < 6 {
		h = 6
	}
	return h
}

// resetScroll resets the content scroll offset to the top.
func (m *App) resetScroll() {
	m.contentOffset = 0
}

// scrollUp scrolls the content one line up.
func (m *App) scrollUp() {
	if m.contentOffset > 0 {
		m.contentOffset--
	}
}

// scrollDown scrolls the content one line down.
func (m *App) scrollDown() {
	maxH := m.contentHeight()
	if maxH <= 0 {
		m.contentOffset++
		return
	}
	// Soft cap: avoid unbounded offset growth on short content.
	if m.contentOffset < 10000 {
		m.contentOffset++
	}
}

// View implements tea.Model.
func (m *App) View() string {
	if m.quitting {
		return ""
	}

	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	layout := m.renderLayout()

	if m.showHelp {
		return m.renderHelpOverlay(layout)
	}

	if m.cvModal {
		return m.renderCVOverlay(layout)
	}

	return layout
}

// renderLayout pins footer to the bottom; body fills remaining height with ambient gutters.
func (m *App) renderLayout() string {
	footer := m.renderFooter()
	footerLines := strings.Count(footer, "\n") + 1
	bodyH := m.height - footerLines
	if bodyH < 1 {
		bodyH = 1
	}

	shell := m.renderBodyCached()
	body := m.composeBodyFrame(shell, bodyH)
	return body + "\n" + footer
}

// renderBodyCached builds sidebar + content column (cached, no footer frame).
func (m *App) renderBodyCached() string {
	key := m.layoutCacheKey()
	if key == m.cachedBodyKey && m.cachedBody != "" {
		return m.cachedBody
	}

	w := m.contentWidth()
	masthead := components.Masthead("habibiahmada", m.profile.Title, w)
	content := m.renderContentRaw()
	contentCol := lipgloss.JoinVertical(lipgloss.Left, masthead, content)

	rail := m.renderNavRail()
	shellLines := maxLineCount(rail, contentCol)
	block := components.JoinShell(rail, contentCol, shellLines)

	m.cachedBody = block
	m.cachedBodyKey = key
	return m.cachedBody
}

// renderNavRail renders the vertical side navigation.
func (m *App) renderNavRail() string {
	items := make([]components.SidebarItem, 0, len(menuItems))
	for _, screen := range menuItems {
		items = append(items, components.SidebarItem{
			Key:  ScreenNames[screen],
			Name: navLabel(screen),
		})
	}
	return components.NavRail(items, m.navSelectedIndex(), components.NavRailWidth())
}

// navLabel returns short nav text (skip "Project Detail").
func navLabel(s Screen) string {
	if s == ScreenProjectDetail {
		return "Projects"
	}
	return ScreenNames[s]
}

// navSelectedIndex returns the highlighted nav index for the current screen.
func (m *App) navSelectedIndex() int {
	if m.currentScreen == ScreenProjectDetail {
		for i, s := range menuItems {
			if s == ScreenProjects {
				return i
			}
		}
	}
	for i, s := range menuItems {
		if s == m.currentScreen {
			return i
		}
	}
	return m.selectedMenu
}

// renderFooter renders brand (left) and keyboard hints (right) on one row.
func (m *App) renderFooter() string {
	return components.FooterBar(m.footerFrame, m.width, footerHints())
}

// footerHints returns the context-sensitive key hints.
func footerHints() []components.FooterHint {
	return []components.FooterHint{
		{Key: "↑↓", Label: "Screens"},
		{Key: "j/k", Label: "Scroll"},
		{Key: "Click", Label: "Navigate"},
		{Key: "Enter", Label: "Open"},
		{Key: "←", Label: "Back"},
		{Key: "?", Label: "Help"},
		{Key: "q", Label: "Quit"},
	}
}

// renderContentRaw routes to the current screen and applies viewport clipping.
func (m *App) renderContentRaw() string {
	var content string
	switch m.currentScreen {
	case ScreenAbout:
		content = m.renderAboutContent()
	case ScreenProjects:
		content = m.renderProjectsContent()
	case ScreenProjectDetail:
		content = m.renderProjectDetailContent()
	case ScreenSkills:
		content = m.renderSkillsContent()
	case ScreenExperience:
		content = m.renderExperienceContent()
	case ScreenCertificates:
		content = m.renderCertificatesContent()
	case ScreenServices:
		content = m.renderServicesContent()
	case ScreenContact:
		content = m.renderContactContent()
	default:
		content = m.renderHomeContent()
	}
	return m.applyViewport(content)
}

// applyViewport clips long content to the visible area (top-aligned, stable layout).
func (m *App) applyViewport(content string) string {
	contentWidth := m.contentWidth()
	if contentWidth < 1 {
		contentWidth = m.width
	}

	maxH := m.contentHeight()
	clipped, _ := components.ClipContent(content, m.contentOffset, maxH)
	_ = contentWidth
	return clipped
}

// clipScroll is a test helper wrapping ClipContent.
func (m *App) clipScroll(content string) string {
	clipped, offset := components.ClipContent(content, m.contentOffset, m.contentHeight())
	m.contentOffset = offset
	return clipped
}

// renderHelpOverlay draws the keymap help as a modal on top of the layout.
func (m *App) renderHelpOverlay(layout string) string {
	lines := []string{
		"Navigation",
		"──────────",
		"↑ / ↓       Switch screen",
		"j / k       Scroll / select in list",
		"Enter       Open detail",
		"← / Esc       Back",
		"",
		"Mouse",
		"──────",
		"Click         Open / select item",
		"Click nav     Switch screen",
		"Scroll        Scroll content",
		"",
		"Screens",
		"───────",
		"   P / p      Projects (from Home)",
		"   C / c      Contact (from Home)",
		"   V / v      View CV (from Home)",
		"",
		"Detail",
		"──────",
		"h / l         Previous / next project (detail)",
		"→             Next project (detail)",
		"",
		"Other",
		"─────",
		"?             Show / hide this help",
		"q / Ctrl+C    Quit",
	}
	return components.Modal("Keymap Help", lines, m.width, m.height)
}

// renderCVOverlay draws the CV summary modal on top of the layout.
func (m *App) renderCVOverlay(layout string) string {
	lines := []string{
		styles.LabelStyle.Render("// CV"),
		"",
		styles.SectionTitleStyle.Render("Habibi Ahmad Aziz"),
		styles.MutedStyle.Render("Fullstack Developer · " + m.profile.Location),
		"",
		styles.SuccessStyle.Render("• " + m.profile.Availability),
		"",
		styles.NormalStyle.Render("Email: " + styles.LinkStyle.Render(m.profile.Email)),
		styles.NormalStyle.Render("Web:   " + styles.LinkStyle.Render(m.profile.Website)),
		"",
		styles.SubtitleStyle.Render("Recent role"),
		styles.NormalStyle.Render("Web Developer — " + m.profile.Employer),
		"",
		styles.MutedStyle.Render("(V opens this viewer · ← / Esc closes)"),
	}
	return components.Modal("Resume", lines, m.width, m.height)
}

func maxLineCount(a, b string) int {
	na := strings.Count(a, "\n") + 1
	nb := strings.Count(b, "\n") + 1
	if na > nb {
		return na
	}
	return nb
}
