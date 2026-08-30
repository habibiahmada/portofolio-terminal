// Package tui implements the terminal user interface using Bubble Tea.
// The App model is the top-level router that composes screens.
package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

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
	ScreenHome:           "Home",
	ScreenAbout:          "About",
	ScreenProjects:       "Projects",
	ScreenProjectDetail:  "Project Detail",
	ScreenSkills:         "Skills",
	ScreenExperience:     "Experience",
	ScreenCertificates:   "Certificates",
	ScreenContact:        "Contact",
}

// Menu items displayed in the sidebar.
var menuItems = []Screen{
	ScreenAbout,
	ScreenProjects,
	ScreenSkills,
	ScreenExperience,
	ScreenCertificates,
	ScreenContact,
}

// App is the top-level Bubble Tea model.
type App struct {
	// Layout dimensions.
	width  int
	height int

	// Navigation state.
	currentScreen Screen
	selectedMenu  int
	prevScreen    Screen

	// Child screen state.
	selectedProject int
	projectDetail   data.Project

	// Profile data.
	profile    data.Profile
	projects   []data.Project
	skills     []data.Skill
	experiences []data.Experience
	certificates []data.Certificate
	socials    []data.Social

	// Flags.
	quitting bool
}

// New creates a new App model with bundled data.
func New() *App {
	profile := data.GetProfile()
	projects := data.GetProjects()
	skills := data.GetSkills()
	experiences := data.GetExperiences()
	certificates := data.GetCertificates()
	socials := data.GetSocials()

	return &App{
		currentScreen: ScreenHome,
		selectedMenu:  0,
		profile:       profile,
		projects:      projects,
		skills:        skills,
		experiences:   experiences,
		certificates:  certificates,
		socials:       socials,
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
	// Global quit.
	if isQuit(msg) {
		m.quitting = true
		return m, tea.Quit
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

// navigateUp moves the selection up.
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
	}
}

// navigateDown moves the selection down.
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
	}
}

// selectItem enters the selected screen.
func (m *App) selectItem() (tea.Model, tea.Cmd) {
	switch m.currentScreen {
	case ScreenHome:
		m.prevScreen = m.currentScreen
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
	return m, nil
}

// View implements tea.Model.
func (m *App) View() string {
	if m.quitting {
		return ""
	}

	if m.width == 0 || m.height == 0 {
		return "Initializing..."
	}

	return m.renderLayout()
}

// renderLayout composes the full layout with header, sidebar, content, and footer.
func (m *App) renderLayout() string {
	header := m.renderHeader()
	sidebar := m.renderSidebar()
	content := m.renderContent()
	footer := m.renderFooter()

	// Sidebar width.
	sidebarWidth := 20
	contentWidth := m.width - sidebarWidth - 1

	// Compose sidebar and content vertically aligned.
	mainContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sidebar,
		lipgloss.NewStyle().Width(contentWidth).Render(content),
	)

	// Full layout.
	fullContent := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		mainContent,
		footer,
	)

	return fullContent
}

// renderHeader renders the top header bar.
func (m *App) renderHeader() string {
	name := styles.TitleStyle.Render(m.profile.Name)
	title := styles.SubtitleStyle.Render(m.profile.Title)
	header := lipgloss.JoinVertical(lipgloss.Left, name, title)
	return styles.BorderStyle.Width(m.width - 2).Render(header)
}

// renderSidebar renders the navigation sidebar.
func (m *App) renderSidebar() string {
	items := make([]string, 0, len(menuItems))
	for i, screen := range menuItems {
		name := ScreenNames[screen]
		if i == m.selectedMenu && m.currentScreen == ScreenHome {
			items = append(items, styles.SidebarItemSelectedStyle.Render("> "+name))
		} else if screen == m.currentScreen ||
			(screen == ScreenProjects && m.currentScreen == ScreenProjectDetail) {
			items = append(items, styles.SidebarItemSelectedStyle.Render("▸ "+name))
		} else {
			items = append(items, styles.SidebarItemStyle.Render("  "+name))
		}
	}

	sidebar := lipgloss.JoinVertical(lipgloss.Left, items...)
	return styles.BorderStyle.
		Width(20).
		Height(m.height - 4).
		Render(sidebar)
}

// renderContent renders the current screen content.
func (m *App) renderContent() string {
	switch m.currentScreen {
	case ScreenHome:
		return m.renderHomeContent()
	case ScreenAbout:
		return m.renderAboutContent()
	case ScreenProjects:
		return m.renderProjectsContent()
	case ScreenProjectDetail:
		return m.renderProjectDetailContent()
	case ScreenSkills:
		return m.renderSkillsContent()
	case ScreenExperience:
		return m.renderExperienceContent()
	case ScreenCertificates:
		return m.renderCertificatesContent()
	case ScreenContact:
		return m.renderContactContent()
	default:
		return m.renderHomeContent()
	}
}

// renderFooter renders the bottom navigation hints.
func (m *App) renderFooter() string {
	hints := []string{
		"↑↓ Navigate",
		"Enter Select",
		"← Back",
		"q Quit",
	}

	footer := strings.Join(hints, "  •  ")
	return styles.FooterStyle.Render(footer)
}

// renderHomeContent renders the home/welcome screen.
func (m *App) renderHomeContent() string {
	welcome := styles.TitleStyle.Render("Welcome!")
	subtitle := styles.MutedStyle.Render("Use ↑↓ to navigate, Enter to select")
	bio := styles.NormalStyle.Render(m.profile.Description)

	// Social links.
	socialLines := make([]string, 0, len(m.socials))
	for _, s := range m.socials {
		socialLines = append(socialLines, styles.LinkStyle.Render(fmt.Sprintf("→ %s: %s", s.Name, s.URL)))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		welcome,
		"",
		bio,
		"",
		subtitle,
		"",
		strings.Join(socialLines, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderAboutContent renders the about screen.
func (m *App) renderAboutContent() string {
	title := styles.TitleStyle.Render("About Me")
	name := styles.NormalStyle.Render(fmt.Sprintf("Name: %s", m.profile.Name))
	titleLine := styles.NormalStyle.Render(fmt.Sprintf("Title: %s", m.profile.Title))
	location := styles.NormalStyle.Render(fmt.Sprintf("Location: %s", m.profile.Location))
	email := styles.LinkStyle.Render(fmt.Sprintf("Email: %s", m.profile.Email))

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		name,
		titleLine,
		location,
		email,
	)

	return styles.ContentStyle.Render(content)
}

// renderProjectsContent renders the projects list.
func (m *App) renderProjectsContent() string {
	title := styles.TitleStyle.Render("Projects")
	hint := styles.MutedStyle.Render("↑↓ to browse • Enter to view details")

	cards := make([]string, 0, len(m.projects))
	for i, p := range m.projects {
		stack := strings.Join(p.Stack, " • ")
		nameStyle := styles.NormalStyle
		if i == m.selectedProject {
			nameStyle = styles.SelectedStyle
		}
		card := styles.CardStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				nameStyle.Render(p.Name),
				styles.NormalStyle.Render(p.Description),
				styles.MutedStyle.Render(stack),
			),
		)
		cards = append(cards, card)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		hint,
		"",
		strings.Join(cards, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderProjectDetailContent renders a single project detail view.
func (m *App) renderProjectDetailContent() string {
	p := m.projectDetail
	title := styles.TitleStyle.Render(p.Name)
	desc := styles.NormalStyle.Render(p.Description)

	tags := make([]string, 0, len(p.Stack))
	for _, s := range p.Stack {
		tags = append(tags, styles.TagStyle.Render(s))
	}
	stack := strings.Join(tags, " ")

	links := make([]string, 0, 2)
	if p.GitHub != "" {
		links = append(links, styles.LinkStyle.Render("GitHub: "+p.GitHub))
	}
	if p.Live != "" {
		links = append(links, styles.LinkStyle.Render("Live: "+p.Live))
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		desc,
		"",
		stack,
		"",
		strings.Join(links, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderSkillsContent renders the skills screen.
func (m *App) renderSkillsContent() string {
	title := styles.TitleStyle.Render("Skills")

	// Group skills by category.
	categories := make(map[string][]data.Skill)
	for _, s := range m.skills {
		categories[s.Category] = append(categories[s.Category], s)
	}

	lines := make([]string, 0)
	for cat, skills := range categories {
		lines = append(lines, styles.SubtitleStyle.Render(cat))
		for _, s := range skills {
			bar := renderSkillBar(s.Level)
			lines = append(lines, fmt.Sprintf("  %s %s", styles.NormalStyle.Render(s.Name), bar))
		}
		lines = append(lines, "")
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(lines, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderSkillBar renders a visual skill level bar.
func renderSkillBar(level int) string {
	filled := strings.Repeat("█", level)
	empty := strings.Repeat("░", 5-level)
	return styles.SuccessStyle.Render(filled) + styles.MutedStyle.Render(empty)
}

// renderExperienceContent renders the experience screen.
func (m *App) renderExperienceContent() string {
	title := styles.TitleStyle.Render("Experience")

	cards := make([]string, 0, len(m.experiences))
	for _, e := range m.experiences {
		details := make([]string, 0, len(e.Details))
		for _, d := range e.Details {
			details = append(details, "  • "+styles.NormalStyle.Render(d))
		}

		card := styles.CardStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				styles.SelectedStyle.Render(e.Role),
				styles.SubtitleStyle.Render(e.Company),
				styles.MutedStyle.Render(fmt.Sprintf("%s • %s", e.Period, e.Location)),
				"",
				strings.Join(details, "\n"),
			),
		)
		cards = append(cards, card)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(cards, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderCertificatesContent renders the certificates screen.
func (m *App) renderCertificatesContent() string {
	title := styles.TitleStyle.Render("Certificates")

	items := make([]string, 0, len(m.certificates))
	for _, c := range m.certificates {
		item := styles.CardStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				styles.SelectedStyle.Render(c.Name),
				styles.MutedStyle.Render(fmt.Sprintf("%s • %s", c.Issuer, c.Date)),
			),
		)
		items = append(items, item)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(items, "\n"),
	)

	return styles.ContentStyle.Render(content)
}

// renderContactContent renders the contact screen.
func (m *App) renderContactContent() string {
	title := styles.TitleStyle.Render("Contact")

	contactLines := []string{
		styles.NormalStyle.Render(fmt.Sprintf("Email:    %s", styles.LinkStyle.Render(m.profile.Email))),
		styles.NormalStyle.Render(fmt.Sprintf("GitHub:   %s", styles.LinkStyle.Render(m.profile.GitHub))),
		styles.NormalStyle.Render(fmt.Sprintf("LinkedIn: %s", styles.LinkStyle.Render(m.profile.LinkedIn))),
		styles.NormalStyle.Render(fmt.Sprintf("Website:  %s", styles.LinkStyle.Render(m.profile.Website))),
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		strings.Join(contactLines, "\n"),
		"",
		styles.MutedStyle.Render("Feel free to reach out!"),
	)

	return styles.ContentStyle.Render(content)
}
