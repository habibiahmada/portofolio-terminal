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
	"github.com/habibiahmada/habibiahmada-terminal/internal/sanitize"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Focus identifies the active interaction zone (nav rail vs screen content).
type Focus int

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

// Layout dimensions. The header (masthead) and footer are fixed bars; the body
// between them holds a fixed nav column and a scrollable content viewport.
const (
	mastheadLines = 2 // brand row + separator
	footerHeight  = 2
)

// bodyTopPad returns a responsive top padding between the header and content
// shell based on the terminal height.
// Small/short terminals (IDE split pane / <=22 rows) get 0-1 padding so content
// isn't clipped; standard and tall terminals scale smoothly.
func (m *App) bodyTopPad() int {
	switch {
	case m.height <= 22:
		return 0
	case m.height <= 30:
		return 1
	case m.height <= 40:
		return 2
	case m.height <= 52:
		return 3
	case m.height <= 65:
		return 4
	default:
		return 5
	}
}

// Focus identifiers — the top-level interaction zones. When focus is on the
// navigation rail, arrow keys switch screens; pressing → enters the screen's
// content where arrows scroll/select. ←/Esc returns to the nav.
const (
	FocusNav Focus = iota
	FocusContent
)

// footerTickInterval drives the App-wide animation loop (splash, nav selector,
// header accent, footer equalizer). One tick redraws the fixed bars and the
// animated accents; the cached content body stays cheap.
const footerTickInterval = 350 * time.Millisecond

type App struct {
	width  int
	height int

	// Navigation state.
	currentScreen Screen
	selectedMenu  int
	prevScreen    Screen

	// Focus zone: nav rail or screen content.
	focus Focus

	// Child screen state.
	selectedProject  int
	projectDetail    data.Project
	selectedFeatured int // highlighted featured card on Home

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
	scrollMax     int // max contentOffset for the current screen (for scrollbar)
	showHelp      bool

	// selectMode releases the mouse so the terminal's native text selection
	// works. When off, the mouse is captured for wheel scroll / scrollbar drag.
	selectMode bool

	// CV modal viewer (V from Home).
	cvModal bool

	// Footer animation frame.
	footerFrame int

	// Last-painted layout geometry for mouse hit-testing (scrollbar drag,
	// nav click). Set during renderLayout/renderBodyCached each frame.
	bodyTop      int // y of the first body row (below the header)
	shellLeft    int // x where the nav+content shell starts
	shellWidth   int // total shell width (nav + divider + content)
	contentTop   int // y of the first content row (including top padding)
	scrollBarX   int // x column of the scrollbar (right of content)
	scrollBarTop int // y of the first scrollbar row
	scrollBarH   int // visible height of the scrollbar

	// Layout cache — excludes footerFrame so animation is cheap.
	cachedBodyKey string
	cachedBody    string

	// Mascot interactive state.
	mascotLeft   int
	mascotTop    int
	mascotWidth  int
	mascotHeight int
	mascotClicks int
	mascotState  int // 0: normal, 1: blink/wink, 2: angry
	mascotTimer  int // ticks remaining in interactive state

	// Flags.
	quitting bool
}

// New creates a new App model with bundled data.
func New() *App {
	return &App{
		currentScreen: ScreenHome,
		selectedMenu:  0,
		focus:         FocusNav,
		profile:       sanitize.Profile(data.GetProfile()),
		projects:      sanitize.Projects(data.GetProjects()),
		skills:        sanitize.Skills(data.GetSkills()),
		work:          sanitize.WorkExperience(data.GetWorkExperience()),
		education:     sanitize.Education(data.GetEducation()),
		certificates:  sanitize.Certificates(data.GetCertificates()),
		socials:       sanitize.Socials(data.GetSocials()),
		companies:     sanitize.Companies(data.GetCompanies()),
		services:      sanitize.Services(data.GetServices()),
		process:       sanitize.ProcessSteps(data.GetProcessSteps()),
		press:         sanitize.PressItems(data.GetPress()),
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
		if m.mascotTimer > 0 {
			m.mascotTimer--
			if m.mascotTimer == 0 {
				m.mascotClicks = 0
				m.mascotState = components.MascotStateNormal
			}
		}
		return m, nextFooterTick()
	}

	return m, nil
}

// handleMouse processes wheel scroll, scrollbar drag, and nav clicks. Mouse
// capture is enabled so these work; press `s` to toggle select mode and use
// the terminal's native text selection instead.
func (m *App) handleMouse(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Select mode releases the mouse (tea.DisableMouse), so no events should
	// arrive; guard defensively in case any are still in flight.
	if m.selectMode {
		return m, nil
	}

	// Help overlay / CV modal intercept clicks.
	if msg.Action == tea.MouseActionPress {
		m.showHelp = false
		m.cvModal = false
	}

	// Wheel scroll moves content only; it must not change focus or list selection.
	if msg.Button == tea.MouseButtonWheelUp {
		m.mouseScrollUp()
		return m, nil
	}
	if msg.Button == tea.MouseButtonWheelDown {
		m.mouseScrollDown()
		return m, nil
	}

	// Scrollbar drag / click: jump scroll position by fraction. The track is a
	// single column at the right edge of the shell (scrollBarX), so only that
	// exact column counts — clicks in the right margin must not trigger it.
	if msg.X == m.scrollBarX && msg.Y >= m.scrollBarTop &&
		msg.Y < m.scrollBarTop+m.scrollBarH {
		frac := 0.0
		if m.scrollBarH > 1 {
			frac = float64(msg.Y-m.scrollBarTop) / float64(m.scrollBarH-1)
		}
		m.setScrollToLineFraction(frac)
		return m, nil
	}

	// Click on mascot in the right margin triggers interactive reactions:
	// 1-2 clicks: eyes blink / happy wink!
	// 3+ clicks: becomes angry with ╬ vein mark!
	if msg.Action == tea.MouseActionPress && m.mascotWidth > 0 &&
		msg.X >= m.mascotLeft && msg.X < m.mascotLeft+m.mascotWidth &&
		msg.Y >= m.mascotTop && msg.Y < m.mascotTop+m.mascotHeight {
		m.mascotClicks++
		if m.mascotClicks >= 3 {
			m.mascotState = components.MascotStateAngry
			m.mascotTimer = 10 // ~3.5s
		} else {
			m.mascotState = components.MascotStateBlink
			m.mascotTimer = 5 // ~1.7s
		}
		return m, nil
	}

	// Click on the nav rail focuses navigation; clicking a menu row switches
	// to that screen. Rows map 1:1 to menuItems (NavRail renders one line per
	// item starting at the top of the shell), so hit-testing is exact.
	if msg.Action == tea.MouseActionPress && msg.X >= m.shellLeft &&
		msg.X <= m.shellLeft+m.navRailWidth()-1 && msg.Y >= m.bodyTop {
		m.focus = FocusNav
		if row := msg.Y - m.bodyTop - m.bodyTopPad(); row >= 0 && row < len(menuItems) {
			m.selectedMenu = row
			_, cmd := m.enterMenuScreen(row)
			return m, cmd
		}
		return m, nil
	}

	// Click a project row on the Projects list to open its detail view. Rows
	// map exactly to the rendered list: focus-rail chip (+1) and the
	// title/meta/blank headers (+3) sit above the list, then the current
	// scroll offset is subtracted.
	if msg.Action == tea.MouseActionPress && m.currentScreen == ScreenProjects &&
		msg.Y >= m.bodyTop+m.bodyTopPad() && len(m.projects) > 0 {
		textLeft := m.shellLeft + m.bodyFrameChrome() - components.ScrollbarWidth
		textRight := textLeft + m.contentWidth() - 1
		if textRight > m.width-1 {
			textRight = m.width - 1
		}
		if msg.X >= textLeft && msg.X <= textRight {
			if i := msg.Y - m.bodyTop - m.bodyTopPad() - 4 + m.contentOffset; i >= 0 && i < len(m.projects) {
				m.focus = FocusContent
				m.selectedProject = i
				return m.selectItem()
			}
		}
		// Fall through: a click on any other content row just focuses content.
	}

	// Click anywhere in content focuses it.
	if msg.Action == tea.MouseActionPress && msg.Y >= m.contentTop {
		m.focus = FocusContent
		return m, nil
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

	// Toggle select mode: releases mouse capture so the terminal's native text
	// selection (click + drag) works. Press again to re-enable mouse scrolling.
	if isSelectModeToggle(msg) {
		m.selectMode = !m.selectMode
		if m.selectMode {
			return m, tea.DisableMouse
		}
		return m, tea.EnableMouseCellMotion
	}

	// Mascot interaction key (M) — easter egg for keyboard users.
	if (msg.String() == "m" || msg.String() == "M") && !msg.Alt {
		m.mascotClicks++
		if m.mascotClicks >= 3 {
			m.mascotState = components.MascotStateAngry
			m.mascotTimer = 10
		} else {
			m.mascotState = components.MascotStateBlink
			m.mascotTimer = 5
		}
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

	// Project detail navigation.
	// ← / h   → previous project
	// → / l   → next project
	// Esc     → back to list (NOT ← which now navigates projects)
	if m.currentScreen == ScreenProjectDetail {
		if msg.Type == tea.KeyLeft || msg.String() == "h" {
			m.prevProject()
			return m, nil
		}
		if msg.Type == tea.KeyRight || msg.String() == "l" {
			m.nextProject()
			return m, nil
		}
		if msg.Type == tea.KeyEsc {
			return m.goBack()
		}
	}

	// Focus-based navigation.
	//
	// Nav focus (default): ↑↓ switch screens, → drops into the screen content.
	if m.focus == FocusNav {
		if msg.Type == tea.KeyRight || isSelect(msg) {
			m.focus = FocusContent
			return m, nil
		}
		if isNavigateUp(msg) {
			return m.switchScreen(-1)
		}
		if isNavigateDown(msg) {
			return m.switchScreen(1)
		}
		return m, nil
	}

	// Content focus: ↑↓ / j/k scroll or select within the screen; ←/Esc return.
	if isBack(msg) {
		m.focus = FocusNav
		return m, nil
	}
	if isPageUp(msg) {
		m.scrollPageUp()
		return m, nil
	}
	if isPageDown(msg) {
		m.scrollPageDown()
		return m, nil
	}
	if isScrollHome(msg) {
		m.scrollTop()
		return m, nil
	}
	if isScrollEnd(msg) {
		m.scrollBottom()
		return m, nil
	}
	if isNavigateUp(msg) || isScrollUp(msg) {
		m.scrollOrSelectUp()
		return m, nil
	}
	if isNavigateDown(msg) || isScrollDown(msg) {
		m.scrollOrSelectDown()
		return m, nil
	}

	// Select / drill-down.
	if isSelect(msg) {
		return m.selectItem()
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
	m.focus = FocusNav

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

// mouseScrollUp scrolls the viewport one line without changing focus.
func (m *App) mouseScrollUp() {
	m.scrollUp()
}

// mouseScrollDown scrolls the viewport one line without changing focus.
func (m *App) mouseScrollDown() {
	m.scrollDown()
}

// scrollOrSelectUp handles j/k up: list selection or content scroll.
func (m *App) scrollOrSelectUp() {
	switch m.currentScreen {
	case ScreenHome:
		featured := m.featuredProjects()
		if len(featured) > 0 {
			m.selectedFeatured--
			if m.selectedFeatured < 0 {
				m.selectedFeatured = len(featured) - 1
			}
		}
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
	case ScreenHome:
		featured := m.featuredProjects()
		if len(featured) > 0 {
			m.selectedFeatured++
			if m.selectedFeatured >= len(featured) {
				m.selectedFeatured = 0
			}
		}
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
		// Enter on a featured card → open its project detail.
		featured := m.featuredProjects()
		if len(featured) == 0 {
			return m, nil
		}
		if m.selectedFeatured < 0 || m.selectedFeatured >= len(featured) {
			m.selectedFeatured = 0
		}
		// Find the project index in the full list.
		fp := featured[m.selectedFeatured]
		for i, p := range m.projects {
			if p.Slug == fp.Slug {
				m.selectedProject = i
				break
			}
		}
		m.prevScreen = ScreenHome
		m.projectDetail = fp
		m.resetScroll()
		m.currentScreen = ScreenProjectDetail
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
		// Return to wherever we came from (Home or Projects list).
		m.currentScreen = m.prevScreen
		if m.currentScreen == ScreenProjectDetail {
			m.currentScreen = ScreenProjects // safety fallback
		}
	case ScreenHome:
		m.quitting = true
		return m, tea.Quit
	default:
		m.currentScreen = ScreenHome
	}
	m.resetScroll()
	m.focus = FocusNav
	return m, nil
}

// contentHeight returns visible lines for page content beside the sidebar.
func (m *App) contentHeight() int {
	bodyH := m.height - mastheadLines - footerHeight
	h := bodyH - m.bodyTopPad()
	if h < 4 {
		h = 4
	}
	return h
}

// resetScroll resets the content scroll offset to the top.
func (m *App) resetScroll() {
	m.contentOffset = 0
	m.scrollMax = 0
}

// scrollUp scrolls the content one line up.
func (m *App) scrollUp() {
	if m.contentOffset > 0 {
		m.contentOffset--
	}
}

// scrollDown scrolls the content one line down.
func (m *App) scrollDown() {
	if m.scrollMax > 0 && m.contentOffset < m.scrollMax {
		m.contentOffset++
	}
}

// scrollTop jumps to the start of the screen content.
func (m *App) scrollTop() {
	m.contentOffset = 0
}

// scrollBottom jumps to the end of the screen content.
func (m *App) scrollBottom() {
	m.contentOffset = m.scrollMax
}

// setScrollToLineFraction maps a 0..1 fraction (cursor position within the
// scrollbar) to a content offset.
func (m *App) setScrollToLineFraction(f float64) {
	if f < 0 {
		f = 0
	}
	if f > 1 {
		f = 1
	}
	m.contentOffset = int(f * float64(m.scrollMax))
}

// scrollPageDown scrolls down by roughly one visible page.
func (m *App) scrollPageDown() {
	page := m.contentHeight() - 2
	if page < 1 {
		page = 1
	}
	m.contentOffset += page
	if m.contentOffset > m.scrollMax {
		m.contentOffset = m.scrollMax
	}
}

// scrollPageUp scrolls up by roughly one visible page.
func (m *App) scrollPageUp() {
	page := m.contentHeight() - 2
	if page < 1 {
		page = 1
	}
	m.contentOffset -= page
	if m.contentOffset < 0 {
		m.contentOffset = 0
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

// renderLayout pins a fixed header on top and footer on the bottom; the body
// between them is a decorated, top-aligned area that scrolls.
func (m *App) renderLayout() string {
	header := m.renderHeaderBar()
	headerLines := strings.Count(header, "\n") + 1
	footer := m.renderFooter()
	footerLines := strings.Count(footer, "\n") + 1

	bodyH := m.height - headerLines - footerLines
	if bodyH < 5 {
		bodyH = 5
	}

	shell := m.renderBodyCached()
	m.bodyTop = headerLines
	m.contentTop = headerLines + m.bodyTopPad()
	body := m.composeBodyFrame(shell, bodyH)
	body = m.overlayMascot(body, bodyH)
	return header + "\n" + body + "\n" + footer
}

// overlayMascot draws the decorative mascot in the RIGHT MARGIN of the body,
// never over content. If the shell already fills the terminal, the mascot is
// skipped so certificate boxes and the scrollbar stay intact.
func (m *App) overlayMascot(body string, bodyH int) string {
	figure := components.MascotFigure(m.footerFrame, m.mascotState)
	if figure == "" {
		m.mascotWidth = 0
		m.mascotHeight = 0
		return body
	}
	fLines := strings.Split(figure, "\n")
	mascotW := 0
	for _, fl := range fLines {
		if w := lipgloss.Width(fl); w > mascotW {
			mascotW = w
		}
	}
	if mascotW == 0 || len(fLines) == 0 || len(fLines) > bodyH {
		m.mascotWidth = 0
		m.mascotHeight = 0
		return body
	}

	rightStart := m.shellLeft + m.shellWidth + 1
	if rightStart+mascotW > m.width {
		m.mascotWidth = 0
		m.mascotHeight = 0
		return body
	}

	bodyLines := strings.Split(body, "\n")
	if len(bodyLines) < bodyH {
		for len(bodyLines) < bodyH {
			bodyLines = append(bodyLines, strings.Repeat(" ", m.width))
		}
	}

	startRow := bodyH - len(fLines)
	m.mascotLeft = rightStart
	m.mascotTop = m.bodyTop + startRow
	m.mascotWidth = mascotW
	m.mascotHeight = len(fLines)

	for i, fl := range fLines {
		row := components.FitLine(bodyLines[startRow+i], rightStart)
		bodyLines[startRow+i] = components.FitLine(row+fl, m.width)
	}
	return strings.Join(bodyLines, "\n")
}

// renderHeaderBar renders the fixed top bar (brand + role) with an animated
// accent, spanning the full terminal width.
func (m *App) renderHeaderBar() string {
	return components.HeaderBar("habibiahmada", m.profile.Title, m.profile.Location, m.width, m.footerFrame)
}

// renderBodyCached builds the nav column + content viewport (cached). The
// content is clipped to a FIXED viewport height so the nav column stays put
// regardless of how long a screen's content is.
func (m *App) renderBodyCached() string {
	key := m.layoutCacheKey()
	if key == m.cachedBodyKey && m.cachedBody != "" {
		return m.cachedBody
	}

	h := m.contentHeight()
	content := m.renderContentRaw()
	content = m.renderFocusRail(content, h)
	rail := m.renderNavRail()
	block := components.JoinShell(rail, content, h)

	m.cachedBody = block
	m.cachedBodyKey = key
	return m.cachedBody
}

// renderFocusRail adds a one-cell focus indicator column on the LEFT edge of
// the content area. The zone chip ("► NAVIGATE" / "► CONTENT") sits on its own
// first row so it never collides with the page heading.
func (m *App) renderFocusRail(content string, h int) string {
	lines := strings.Split(content, "\n")

	var railStyle lipgloss.Style
	var chipStyle lipgloss.Style
	chip := "CONTENT"
	railGlyph := "▎"
	if m.focus == FocusNav {
		railStyle = styles.RuleStyle
		chipStyle = styles.MutedStyle
		chip = "NAVIGATE"
		railGlyph = "·"
	} else {
		railStyle = styles.ContentZoneHighlight
		chipStyle = styles.ContentZoneHighlight
	}

	prefix := railStyle.Render(railGlyph) + " "
	rowWidth := 0
	if len(lines) > 0 {
		rowWidth = lipgloss.Width(lines[0])
	}
	target := focusRailCells + rowWidth

	out := make([]string, h)
	out[0] = components.FitLine(prefix+chipStyle.Render("► "+chip), target)
	src := lines
	if len(src) > h-1 {
		src = src[:h-1]
	}
	for i := 0; i < h-1; i++ {
		ln := ""
		if i < len(src) {
			ln = src[i]
		}
		out[i+1] = components.FitLine(prefix+ln, target)
	}
	return strings.Join(out, "\n")
}

// renderNavRail renders the vertical side navigation with an animated
// selection indicator.
func (m *App) renderNavRail() string {
	items := make([]components.SidebarItem, 0, len(menuItems))
	for _, screen := range menuItems {
		items = append(items, components.SidebarItem{
			Key:  ScreenNames[screen],
			Name: navLabel(screen, m.width),
		})
	}
	return components.NavRail(items, m.navSelectedIndex(), m.navRailWidth(), m.footerFrame, m.focus == FocusNav)
}

// navLabel returns nav text, shortened on compact terminals.
func navLabel(s Screen, termWidth int) string {
	if s == ScreenProjectDetail {
		return navLabel(ScreenProjects, termWidth)
	}
	switch {
	case termWidth < 40:
		switch s {
		case ScreenHome:
			return "Ho"
		case ScreenAbout:
			return "Ab"
		case ScreenProjects:
			return "Pr"
		case ScreenSkills:
			return "Sk"
		case ScreenExperience:
			return "Ex"
		case ScreenCertificates:
			return "Ce"
		case ScreenServices:
			return "Sv"
		case ScreenContact:
			return "Ct"
		}
	case termWidth < 60:
		switch s {
		case ScreenProjects:
			return "Proj"
		case ScreenSkills:
			return "Skill"
		case ScreenExperience:
			return "Work"
		case ScreenCertificates:
			return "Cert"
		case ScreenServices:
			return "Svc"
		case ScreenContact:
			return "Talk"
		}
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
	return components.FooterBar(m.footerFrame, m.width, m.footerHints())
}

// footerHints returns the context-sensitive key hints.
func (m *App) footerHints() []components.FooterHint {
	if m.focus == FocusNav {
		return []components.FooterHint{
			{Key: "↑↓", Label: "Screens"},
			{Key: "→", Label: "Focus screen"},
			{Key: "s", Label: "Select text"},
			{Key: "?", Label: "Help"},
			{Key: "q", Label: "Quit"},
		}
	}
	return []components.FooterHint{
		{Key: "↑↓/jk", Label: "Scroll"},
		{Key: "PgUp/PgDn", Label: "Page"},
		{Key: "Enter", Label: "Open"},
		{Key: "←/Esc", Label: "Nav"},
		{Key: "s", Label: "Select text"},
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
	return m.applyViewport(components.ReflowBlock(content, m.contentWidth()))
}

// applyViewport clips long content to the visible area (top-aligned, stable
// layout) and appends a scrollbar column when the screen is taller than the
// viewport. The nav column stays fixed; only this content viewport scrolls.
func (m *App) applyViewport(content string) string {
	maxH := m.contentHeight() - 1 // first viewport row is the zone chip
	if maxH < 3 {
		maxH = 3
	}
	cw := m.contentWidth()

	res := components.ClipContentFull(content, m.contentOffset, maxH)
	m.contentOffset = res.Offset
	m.scrollMax = res.Max

	// Pad every line to a FIXED content width so the shell (nav + rule +
	// content) has a stable width regardless of screen. Always reserve the
	// scrollbar gutter so the right edge does not jump between screens.
	body := padToWidth(res.Text, maxH, cw)
	return components.AddScrollbar(body, res.Offset, res.Max, cw)
}

// padToWidth ensures text is exactly maxH lines, each padded to cw visible
// cells (right-padded) so the content column is a fixed width.
func padToWidth(s string, maxH, cw int) string {
	lines := strings.Split(s, "\n")
	if len(lines) > maxH {
		lines = lines[:maxH]
	}
	for i, ln := range lines {
		lines[i] = components.FitLine(ln, cw)
	}
	for len(lines) < maxH {
		lines = append(lines, strings.Repeat(" ", cw))
	}
	return strings.Join(lines, "\n")
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
		"On nav:",
		"↑ / ↓       Switch screen",
		"Click item  Switch screen (mouse)",
		"→ / Enter   Focus into screen content",
		"On screen:",
		"↑ / ↓ / j/k  Scroll / select in list",
		"Mouse wheel  Scroll content (does not change focus)",
		"PgUp/PgDn   Scroll a page",
		"Home / End  Jump top / bottom",
		"Enter       Open detail",
		"← / Esc     Back to nav",
		"",
		"Text selection",
		"──────────────",
		"s            Toggle select mode (release mouse)",
		"             then click & drag to select text.",
		"             Press s again to re-enable mouse",
		"             scroll (wheel + scrollbar drag).",
		"Shift+drag   Also select text while mouse is on",
		"             (terminals bypass the app's mouse).",
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
		styles.NormalStyle.Render("Web Developer at " + m.profile.Employer),
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
