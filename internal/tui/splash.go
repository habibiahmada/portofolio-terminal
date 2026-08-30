package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Splash frames follow docs/tui-illustration.md:
//
//	0ms    Initializing portfolio...
//	400ms  progress 50%
//	800ms  progress 100%
//	1000ms signature + cursor blink
//	1800ms transition -> Home
var splashDelays = []time.Duration{
	400 * time.Millisecond,
	400 * time.Millisecond,
	200 * time.Millisecond,
	800 * time.Millisecond,
}

// splashProgress is the progress bar percentage per frame.
var splashProgress = []int{0, 50, 100, 100, 100}

// Splash is skipped on terminals below these sizes to protect small SSH UX.
const (
	skipSplashWidth  = 40
	skipSplashHeight = 20
)

// splashTickMsg advances the splash sequence one frame.
type splashTickMsg struct{}

func nextSplashTick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg { return splashTickMsg{} })
}

// Splash is the startup model. It plays a short sequence then hands the
// session over to the main App model.
type Splash struct {
	width  int
	height int
	frame  int
}

// NewSplash creates the splash model that transitions into the full App.
func NewSplash() *Splash {
	return &Splash{}
}

// Init implements tea.Model.
func (s *Splash) Init() tea.Cmd {
	return nextSplashTick(splashDelays[0])
}

// Update implements tea.Model.
func (s *Splash) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		if s.skip() {
			app := New()
			app.width = s.width
			app.height = s.height
			return app, nil
		}
		return s, nil

	case splashTickMsg:
		s.frame++
		if s.frame >= len(splashDelays) {
			// Sequence complete — transition to the main app.
			app := New()
			app.width = s.width
			app.height = s.height
			return app, nil
		}
		return s, nextSplashTick(splashDelays[s.frame])
	}

	return s, nil
}

// skip reports whether the splash should be skipped on this terminal size.
func (s *Splash) skip() bool {
	if s.width == 0 || s.height == 0 {
		return false
	}
	return s.width < skipSplashWidth || s.height < skipSplashHeight
}

// View implements tea.Model.
func (s *Splash) View() string {
	if s.width == 0 || s.height == 0 {
		return styles.SplashTextStyle.Render("Initializing portfolio...")
	}

	progress := splashProgress[s.frame]

	status := "Initializing portfolio..."
	if s.frame >= 1 {
		status = "Loading portfolio data..."
	}

	var art string
	if s.frame >= 3 {
		art = components.SignatureBlink(s.width, s.frame%2 == 1)
		if art != "" {
			art += "\n\n"
		} else {
			art = styles.HeroTitleStyle.Render("Habibi Ahmad Aziz") + "\n\n"
		}
	}

	barWidth := clampInt(20, s.width-8, 40)
	block := lipgloss.JoinVertical(
		lipgloss.Left,
		art,
		styles.SplashTextStyle.Render(status),
		components.ProgressBar(progress, barWidth)+" "+components.Percentage(progress),
	)

	// Center block but keep it near the vertical middle of the screen.
	return lipgloss.Place(s.width, s.height, lipgloss.Center, lipgloss.Center, block)
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
