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
	footerHeight = 3
	sidebarWidth = 20
)

// footerTickInterval drives the footer animation frame rate.
const footerTickInterval = 120 * time.Millisecond

// App is the top-level Bubble Tea model.
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
			m.enterProjects()
			return m, nil
		case "C", "c":
			m.enterContact()
			return m, nil
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

	// Navigation.
	if isNavigateUp(msg) {
		m.navigateUp()
		return m, nil
	}

	if isNavigateDown(msg) {
		m.navigateDown()
		return m, nil
	}

	// Select.
	if isSelect(msg) {
		return m.selectItem()
	}

	// Back.
	if isBack(msg) {
		return m.goBack()
	}

	return m, nil
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
func (m *App) enterProjects() {
	m.prevScreen = m.currentScreen
	m.resetScroll()
	m.selectedProject = 0
	m.currentScreen = ScreenProjects
}

// enterContact jumps to the Contact screen.
func (m *App) enterContact() {
	m.prevScreen = m.currentScreen
	m.resetScroll()
	m.currentScreen = ScreenContact
}

// navigateUp moves the selection or content up.
func (m *App) navigateUp() {
	switch m.currentScreen {
	case ScreenHome:
		m.selectedMenu--
		if m.selectedMenu < 0 {
			m.selectedMenu = len(menuItems) - 1
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

// navigateDown moves the selection or content down.
func (m *App) navigateDown() {
	switch m.currentScreen {
	case ScreenHome:
		m.selectedMenu++
		if m.selectedMenu >= len(menuItems) {
			m.selectedMenu = 0
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

// selectItem enters the selected screen.
func (m *App) selectItem() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case ScreenHome:
		m.prevScreen = m.currentScreen
		m.resetScroll()
		m.currentScreen = menuItems[m.selectedMenu]
		if m.currentScreen == ScreenProjects {
			m.selectedProject = 0
		}
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
	m.contentOffset++
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

// renderLayout composes the full layout with header, sidebar, content, footer.
func (m *App) renderLayout() string {
	header := components.Header("habibiahmada", m.profile.Title, m.width)
	sidebar := m.renderSidebar()
	content := m.renderContent()
	footer := m.renderFooter()

	contentWidth := m.width - sidebarWidth - 1

	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidebar,
		lipgloss.NewStyle().Width(contentWidth).Render(content),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		mainContent,
		footer,
	)
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

	// On Home the sidebar shows the highlighted menu selection; on other
	// screens it highlights the currently active screen.
	activeKey := ""
	selectedIndex := -1
	if m.currentScreen == ScreenHome {
		selectedIndex = m.selectedMenu
	} else {
		activeKey = ScreenNames[m.currentScreen]
		if m.currentScreen == ScreenProjectDetail {
			activeKey = ScreenNames[ScreenProjects]
		}
	}

	return components.Sidebar(items, selectedIndex, activeKey, m.height-4)
}

// renderFooter renders the animated illustration plus the navigation hints.
func (m *App) renderFooter() string {
	art := components.FooterArtline(m.footerFrame, m.width, m.profile.Website)
	return components.FooterWithHint(art, footerHints(), m.width-2)
}

// footerHints returns the context-sensitive key hints.
func footerHints() []components.FooterHint {
	return []components.FooterHint{
		{Key: "↑↓", Label: "Navigate"},
		{Key: "Enter", Label: "Select"},
		{Key: "←", Label: "Back"},
		{Key: "?", Label: "Help"},
		{Key: "q", Label: "Quit"},
	}
}

// renderContent routes to the current screen's content renderer.
func (m *App) renderContent() string {
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
	return m.clipScroll(content)
}

// clipScroll shows a viewport window of long content based on contentOffset.
func (m *App) clipScroll(content string) string {
	maxH := m.contentHeight()
	if maxH <= 0 {
		return content
	}

	lines := strings.Split(content, "\n")
	if len(lines) <= maxH {
		m.contentOffset = 0
		return content
	}

	// Clamp the offset to a valid window.
	maxOffset := len(lines) - maxH
	if m.contentOffset > maxOffset {
		m.contentOffset = maxOffset
	}

	window := lines[m.contentOffset : m.contentOffset+maxH]

	// Append a scroll indicator when there is hidden content below.
	indicator := ""
	if m.contentOffset < maxOffset {
		indicator = "\n" + styles.MutedStyle.Render("▼ scroll with ↑↓")
	}

	return strings.Join(window, "\n") + indicator
}

// renderHelpOverlay draws the keymap help as a modal on top of the layout.
func (m *App) renderHelpOverlay(layout string) string {
	lines := []string{
		"Navigation",
		"──────────",
		"↑ / k         Move up",
		"↓ / j         Move down / scroll",
		"Enter         Select",
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
