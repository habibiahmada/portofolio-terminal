// Package main is the entry point for the local/npx experience.
// It starts the Bubble Tea TUI directly on the user's terminal.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/habibiahmada/habibiahmada-terminal/internal/tui"
)

func main() {
	// Start the Bubble Tea program with a splash that transitions to the App.
	model := tui.NewSplash()

	// Start the Bubble Tea program.
	p := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
