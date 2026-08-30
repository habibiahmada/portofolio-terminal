// Package tui implements the terminal user interface using Bubble Tea.
// The App model is the top-level router that composes screens.
//
// Navigation and layout logic lives here; each screen's content renderer
// lives in its own file (home.go, about.go, projects.go, ...).
package tui

import (
	"strings"

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
	ScreenContact:       "Contact",
}

// menuItems is the ordered list of screens reachable from the Home sidebar.
var menuItems = []Screen{
	ScreenAbout,
	ScreenProjects,
	ScreenSkills,
	ScreenExperience,
	ScreenCertificates,
	ScreenContact,
}

// Layout dimensions (best-effort; used for scrolling).
const (
	headerHeight = 4
	footerHeight = 2
	sidebarWidth = 20
)

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
	experiences  []data.Experience
	certificates []data.Certificate
	socials      []data.Social

	// View state.
	contentOffset int
	showHelp      bool

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
		experiences:   data.GetExperiences(),
		certificates:  data.GetCertificates(),
		socials:       data.GetSocials(),
	}
}

// Init implements tea.Model.
func (m *App) Init() tea.Cmd {
	return nil
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

	return layout
}

// renderLayout composes the full layout with header, sidebar, content, footer.
func (m *App) renderLayout() string {
	header := components.Header(m.profile.Name, m.profile.Title, m.width)
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

// renderFooter renders the bottom navigation hints.
func (m *App) renderFooter() string {
	hints := []components.FooterHint{
		{Key: "↑↓", Label: "Navigate"},
		{Key: "Enter", Label: "Select"},
		{Key: "←", Label: "Back"},
		{Key: "?", Label: "Help"},
		{Key: "q", Label: "Quit"},
	}
	return components.Footer(hints)
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
		"↑ / k    Move up",
		"↓ / j    Move down / scroll",
		"Enter    Select",
		"← / Esc  Back",
		"",
		"Other",
		"─────",
		"?        Show / hide this help",
		"q / Ctrl+C  Quit",
	}
	return components.Modal("Keymap Help", lines, m.width, m.height)
}
