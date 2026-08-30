package tui

import "github.com/charmbracelet/bubbletea"

// isQuit returns true if the key press should quit the application.
func isQuit(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyCtrlC ||
		(msg.String() == "q" && !msg.Alt) ||
		(msg.String() == "Q" && !msg.Alt)
}

// isNavigateUp returns true if the key switches the active screen up.
func isNavigateUp(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyUp
}

// isNavigateDown returns true if the key switches the active screen down.
func isNavigateDown(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyDown
}

// isScrollUp returns true if the key scrolls content or moves list selection up.
func isScrollUp(msg tea.KeyMsg) bool {
	return msg.String() == "k"
}

// isScrollDown returns true if the key scrolls content or moves list selection down.
func isScrollDown(msg tea.KeyMsg) bool {
	return msg.String() == "j"
}

// isSelect returns true if the key press selects an item.
func isSelect(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyEnter || msg.Type == tea.KeySpace
}

// isBack returns true if the key press goes back.
func isBack(msg tea.KeyMsg) bool {
	return msg.Type == tea.KeyLeft || msg.Type == tea.KeyEsc
}

// isHelp returns true if the key press toggles the help overlay.
func isHelp(msg tea.KeyMsg) bool {
	return msg.String() == "?" || msg.Type == tea.KeyF1
}
