// Package styles provides centralized Lip Gloss style definitions for the TUI.
// All visual styling lives here — no raw ANSI in screen files.
package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// Color palette — synced with docs/design-system.md (Red & Blue Glitch).
var (
	ColorPrimary    = lipgloss.Color("#ef4444") // brand red
	ColorSecondary  = lipgloss.Color("#3b82f6") // glitch blue
	ColorAccent     = lipgloss.Color("#f87171") // soft red (selection highlight)
	ColorMuted      = lipgloss.Color("#71717a") // zinc-500
	ColorSuccess    = lipgloss.Color("#00B894")
	ColorLink       = lipgloss.Color("#3b82f6")
	ColorBorder     = lipgloss.Color("#3f3f46") // zinc-800
	ColorText       = lipgloss.Color("#f5f5f5") // dark-first foreground
	ColorBackground = lipgloss.Color("#1a1a1a") // dark-first background
	ColorWarning    = lipgloss.Color("#fbbf24")
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

	HeaderWordmark = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorText)

	HeaderDot = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary)

	HeaderMeta = lipgloss.NewStyle().
			Foreground(ColorMuted)

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

	ListSelectedStyle = lipgloss.NewStyle().
				Foreground(ColorAccent).
				Bold(true).
				PaddingLeft(1)

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

// Section label ("// Label") and badge variants, per docs/design-system.md.
var (
	LabelStyle = lipgloss.NewStyle().
			Foreground(ColorMuted).
			Bold(true).
			MarginBottom(1)

	SectionTitleStyle = lipgloss.NewStyle().
				Foreground(ColorText).
				Bold(true).
				MarginBottom(1)

	PrimaryText = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	BadgeStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSuccess).
			Padding(0, 1)

	BadgeAccentStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorPrimary).
				Padding(0, 1).
				Bold(true)

	BadgeNeutralStyle = lipgloss.NewStyle().
				Foreground(ColorMuted).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorBorder).
				Padding(0, 1)

	PrimaryCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorPrimary).
				Padding(1, 2).
				MarginTop(1).
				MarginBottom(1)

	RuleStyle = lipgloss.NewStyle().
			Foreground(ColorBorder)
)

// Footer animation styles — the always-present animated illustration.
var (
	FooterArtStyle = lipgloss.NewStyle().
			Foreground(ColorText).
			Background(ColorBorder).
			Padding(0, 2).
			Inline(true)

	FooterBarStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	FooterWordmark = lipgloss.NewStyle().
			Foreground(ColorText).
			Bold(true)
)
