// Package styles provides centralized Lip Gloss style definitions for the TUI.
// All visual styling lives here — no raw ANSI in screen files.
package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette.
var (
	ColorPrimary    = lipgloss.Color("#FF6B6B")
	ColorSecondary  = lipgloss.Color("#4ECDC4")
	ColorAccent     = lipgloss.Color("#FFE66D")
	ColorMuted      = lipgloss.Color("#636e72")
	ColorSuccess    = lipgloss.Color("#00B894")
	ColorLink       = lipgloss.Color("#0984e3")
	ColorBorder     = lipgloss.Color("#dfe6e9")
	ColorText       = lipgloss.Color("#2d3436")
	ColorBackground = lipgloss.Color("#ffffff")
)

// Core styles.
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			MarginBottom(1)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true).
			PaddingLeft(1)

	NormalStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			PaddingLeft(1)

	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(0, 1)

	MutedStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Italic(true)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)

	LinkStyle = lipgloss.NewStyle().
			Foreground(ColorLink).
			Underline(true)
)

// Layout styles.
var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginBottom(1)

	FooterStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			PaddingTop(1)

	SidebarItemStyle = lipgloss.NewStyle().
				PaddingLeft(1).
				Width(20)

	SidebarItemSelectedStyle = lipgloss.NewStyle().
					PaddingLeft(1).
					Width(20).
					Foreground(ColorAccent).
					Bold(true)

	ContentStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(1, 2).
			MarginBottom(1)

	TagStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSecondary).
			Padding(0, 1).
			MarginRight(1)
)
