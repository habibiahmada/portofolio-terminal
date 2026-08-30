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
	"github.com/habibiahmada/habibiahmada-terminal/internal/blog"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

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
	ScreenBlog
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
	ScreenBlog:          "Blog",
	ScreenContact:       "Contact",
}

// menuItems is the ordered list of screens reachable from the Home sidebar.
// Ordered to match the website navigation plus Contact.
var menuItems = []Screen{
	ScreenAbout,
	ScreenProjects,
	ScreenSkills,
	ScreenExperience,
	ScreenCertificates,
	ScreenServices,
	ScreenBlog,
	ScreenContact,
}

// Layout dimensions (best-effort; used for scrolling).
const (
	headerHeight = 4
	footerHeight = 2
	sidebarWidth = 20
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

	// Blog (live fetch).
	blogPosts    []blog.Post
	blogLoading  bool
	blogLoaded   bool
	blogError    string
	blogSelected int
	blogDetail   bool

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

	case footerTickMsg:
		m.footerFrame++
		return m, nextFooterTick()

	case blogPostsMsg:
		m.blogLoading = false
		m.blogLoaded = true
		if msg.err != nil {
			m.blogError = msg.err.Error()
			return m, nil
		}
		m.blogPosts = msg.posts
		m.blogError = ""
		if m.blogSelected >= len(m.blogPosts) {
			m.blogSelected = 0
		}
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

	// Blog detail: back to list without leaving the screen.
	if m.currentScreen == ScreenBlog && m.blogDetail && isBack(msg) {
		m.blogDetail = false
		m.resetScroll()
		return m, nil
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

// switchScreen moves to the next/previous sidebar screen immediately.
func (m *App) switchScreen(delta int) (tea.Model, tea.Cmd) {
	if m.currentScreen == ScreenProjectDetail {
		m.currentScreen = ScreenProjects
	}

	if m.currentScreen == ScreenHome {
		if delta > 0 {
			return m.enterMenuScreen(m.selectedMenu)
		}
		return m.enterMenuScreen(len(menuItems) - 1)
	}

	m.syncMenuFromScreen()

	m.selectedMenu += delta
	if m.selectedMenu < 0 {
		m.selectedMenu = len(menuItems) - 1
	}
	if m.selectedMenu >= len(menuItems) {
		m.selectedMenu = 0
	}
	return m.enterMenuScreen(m.selectedMenu)
}

// syncMenuFromScreen aligns selectedMenu with the current screen.
func (m *App) syncMenuFromScreen() {
	if m.currentScreen == ScreenHome || m.currentScreen == ScreenProjectDetail {
		return
	}
	for i, s := range menuItems {
		if s == m.currentScreen {
			m.selectedMenu = i
			return
		}
	}
}

// enterMenuScreen activates the screen at menuItems[idx].
func (m *App) enterMenuScreen(idx int) (tea.Model, tea.Cmd) {
	m.prevScreen = ScreenHome
	m.resetScroll()
	m.blogDetail = false
	m.currentScreen = menuItems[idx]
	if m.currentScreen == ScreenProjects {
		m.selectedProject = 0
	}
	if m.currentScreen == ScreenBlog && !m.blogLoaded && !m.blogLoading {
		m.blogLoading = true
		return m, m.fetchBlogCmd()
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
	case ScreenBlog:
		if !m.blogDetail && len(m.blogPosts) > 0 {
			m.blogSelected--
			if m.blogSelected < 0 {
				m.blogSelected = len(m.blogPosts) - 1
			}
		} else {
			m.scrollUp()
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
	case ScreenBlog:
		if !m.blogDetail && len(m.blogPosts) > 0 {
			m.blogSelected++
			if m.blogSelected >= len(m.blogPosts) {
				m.blogSelected = 0
			}
		} else {
			m.scrollDown()
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

// selectItem opens a drill-down view (project detail, blog article).
func (m *App) selectItem() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case ScreenHome:
		return m.enterMenuScreen(m.selectedMenu)
	case ScreenProjects:
		if len(m.projects) == 0 {
			return m, nil
		}
		m.prevScreen = m.currentScreen
		m.projectDetail = m.projects[m.selectedProject]
		m.resetScroll()
		m.currentScreen = ScreenProjectDetail
	case ScreenBlog:
		if m.blogDetail || len(m.blogPosts) == 0 || m.blogLoading {
			return m, nil
		}
		m.blogDetail = true
		m.resetScroll()
	}
	return m, nil
}

// goBack returns to the previous screen.
func (m *App) goBack() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case ScreenProjectDetail:
		m.currentScreen = ScreenProjects
	case ScreenBlog:
		if m.blogDetail {
			m.blogDetail = false
			m.resetScroll()
			return m, nil
		}
		m.currentScreen = ScreenHome
	case ScreenHome:
		m.quitting = true
		return m, tea.Quit
	default:
		m.currentScreen = ScreenHome
	}
	m.resetScroll()
	return m, nil
}

// contentHeight returns the number of visible lines available for content.
func (m *App) contentHeight() int {
	return m.height - headerHeight - footerHeight
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

// renderLayout composes header, sidebar, content, and footer. The main body is
// cached; only the footer re-renders on animation ticks.
func (m *App) renderLayout() string {
	body := m.renderBodyCached()
	footer := m.renderFooter()
	return lipgloss.JoinVertical(lipgloss.Left, body, footer)
}

// renderBodyCached returns header + sidebar + content, using a cache key that
// ignores footerFrame so animation does not rebuild heavy screen content.
func (m *App) renderBodyCached() string {
	key := m.layoutCacheKey()
	if key == m.cachedBodyKey && m.cachedBody != "" {
		return m.cachedBody
	}

	header := components.Header("habibiahmada", m.profile.Title, m.width)
	sidebar := m.renderSidebar()
	content := m.renderContentRaw()
	contentWidth := m.width - sidebarWidth - 1
	if contentWidth < 1 {
		contentWidth = 1
	}

	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidebar,
		lipgloss.NewStyle().Width(contentWidth).Render(content),
	)

	m.cachedBody = lipgloss.JoinVertical(lipgloss.Left, header, mainContent)
	m.cachedBodyKey = key
	return m.cachedBody
}

// renderSidebar renders the navigation sidebar.
func (m *App) renderSidebar() string {
	items := make([]components.SidebarItem, 0, len(menuItems))
	for _, screen := range menuItems {
		items = append(items, components.SidebarItem{
			Key:  ScreenNames[screen],
			Name: ScreenNames[screen],
		})
	}

	selectedIndex := m.selectedMenu
	activeKey := ""
	if m.currentScreen == ScreenProjectDetail {
		activeKey = ScreenNames[ScreenProjects]
		for i, s := range menuItems {
			if s == ScreenProjects {
				selectedIndex = i
				break
			}
		}
	} else if m.currentScreen != ScreenHome {
		for i, s := range menuItems {
			if s == m.currentScreen {
				selectedIndex = i
				break
			}
		}
	}

	sidebarH := m.contentHeight()
	if sidebarH < 1 {
		sidebarH = m.height - headerHeight - footerHeight
	}
	return components.Sidebar(items, selectedIndex, activeKey, sidebarH)
}

// renderFooter renders brand (left) and keyboard hints (right) on one row.
func (m *App) renderFooter() string {
	return components.FooterBar(m.footerFrame, m.width-2, footerHints())
}

// footerHints returns the context-sensitive key hints.
func footerHints() []components.FooterHint {
	return []components.FooterHint{
		{Key: "↑↓", Label: "Screens"},
		{Key: "j/k", Label: "Scroll"},
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
	case ScreenBlog:
		content = m.renderBlogContent()
	case ScreenContact:
		content = m.renderContactContent()
	default:
		content = m.renderHomeContent()
	}
	return m.applyViewport(content)
}

// applyViewport clips long content, centers short content, and respects width.
func (m *App) applyViewport(content string) string {
	contentWidth := m.width - sidebarWidth - 2
	if contentWidth < 1 {
		contentWidth = m.width
	}

	maxH := m.contentHeight()
	clipped, offset := components.ClipContent(content, m.contentOffset, maxH)
	_ = offset // display-only clamp; contentOffset corrected on next scroll

	if m.currentScreen == ScreenHome {
		return clipped
	}

	lines := strings.Split(clipped, "\n")
	if len(lines) < maxH && maxH > 0 {
		return components.CenterInViewport(clipped, contentWidth, maxH)
	}
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
