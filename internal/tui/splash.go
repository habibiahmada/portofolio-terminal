package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Splash: signature + brand + progress, then hand off to App.
var splashDelays = []time.Duration{
	500 * time.Millisecond,
	500 * time.Millisecond,
}

const (
	skipSplashWidth  = 40
	skipSplashHeight = 20
)

type splashTickMsg struct{}

func nextSplashTick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return splashTickMsg{} })
}

// Splash is the startup model.
type Splash struct {
	width  int
	height int
	frame  int
}

// NewSplash creates the splash model.
func NewSplash() *Splash {
	return &Splash{}
}

func (s *Splash) Init() tea.Cmd {
	return nextSplashTick(splashDelays[0])
}

func (s *Splash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		if s.skip() {
			return s.toApp(), nil
		}
		return s, nil

	case tea.KeyMsg:
		return s.toApp(), nil

	case splashTickMsg:
		s.frame++
		if s.frame >= len(splashDelays) {
			return s.toApp(), nil
		}
		return s, nextSplashTick(splashDelays[s.frame])
	}

	return s, nil
}

func (s *Splash) toApp() *App {
	app := New()
	app.width = s.width
	app.height = s.height
	return app
}

func (s *Splash) skip() bool {
	if s.width == 0 || s.height == 0 {
		return false
	}
	return s.width < skipSplashWidth || s.height < skipSplashHeight
}

func (s *Splash) View() string {
	if s.width == 0 || s.height == 0 {
		return styles.SplashTextStyle.Render("habibiahmada.")
	}

	barW := clampInt(16, s.width/3, 32)
	progress := 50
	if s.frame >= 1 {
		progress = 100
	}

	art := components.SignatureBlink(s.width/2, s.frame%2 == 0)
	if art == "" {
		art = styles.PromptStyle.Render(">_")
	}

	brand := styles.HeaderWordmark.Render("habibiahmada") + styles.HeaderDot.Render(".")
	sub := styles.MutedStyle.Render("terminal portfolio")

	block := lipgloss.JoinVertical(
		lipgloss.Center,
		art,
		brand,
		sub,
		components.ProgressBar(progress, barW),
	)

	if s.frame >= 1 {
		block = lipgloss.JoinVertical(lipgloss.Center, block, styles.MutedStyle.Render("any key to continue"))
	}

	bodyH := s.height
	if bodyH < 1 {
		bodyH = 1
	}
	return lipgloss.Place(s.width, bodyH, lipgloss.Center, lipgloss.Center, block)
}

func clampInt(min, v, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
