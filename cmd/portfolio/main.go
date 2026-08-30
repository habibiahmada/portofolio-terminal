// Package main is the entry point for the local/npx experience.
// It starts the Bubble Tea TUI directly on the user's terminal.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
	"github.com/habibiahmada/habibiahmada-terminal/internal/tui"
)

func main() {
	styles.ForceTrueColor()
	// Splash transitions into the full TUI core.
	model := tui.NewSplash()

	// Mouse capture is enabled so the wheel can scroll the content and the
	// scrollbar can be dragged. Press `s` to toggle "select mode", which
	// releases the mouse so the terminal's native text selection works; press
	// `s` again to re-enable mouse scrolling. Everything is also reachable
	// with the keyboard.
	p := tea.NewProgram(model,
		tea.WithAltScreen(),       // clean full-screen painting
		tea.WithMouseCellMotion(), // wheel scroll / scrollbar drag (toggled by `s`)
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
