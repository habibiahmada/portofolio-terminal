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
			Foreground(ColorMuted)

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

	ContentStyle = lipgloss.NewStyle()

	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Padding(1, 2).
			MarginBottom(1)

	TagStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	// HomeCardTitleStyle and HomeCardMetaStyle fill the Featured Projects card
	// tiles on the Home screen.
	HomeCardTitleStyle = lipgloss.NewStyle().
				Foreground(ColorText).
				Bold(true)

	HomeCardMetaStyle = lipgloss.NewStyle().
				Foreground(ColorSecondary)
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
			Bold(true)

	BadgeAccentStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	BadgeNeutralStyle = lipgloss.NewStyle().
				Foreground(ColorMuted)

	PrimaryCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorPrimary).
				Padding(1, 2).
				MarginTop(1).
				MarginBottom(1)

	RuleStyle = lipgloss.NewStyle().
			Foreground(ColorBorder)

	// NavActiveStyle is the currently selected nav item. When the nav zone has
	// focus it is a solid highlighted bar (filled background) so it is obvious
	// which item is active; otherwise it stays bold red without the fill.
	NavActiveStyle = lipgloss.NewStyle().
			Foreground(ColorBackground).
			Background(ColorPrimary).
			Bold(true)

	// NavSelectedInactive is the selected item while the content zone has
	// focus — dimmed so the content list owns the strong selection cue.
	NavSelectedInactive = lipgloss.NewStyle().
				Foreground(ColorMuted)

	NavItemStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	// NavZoneHighlight draws a colored left rail on the nav column when it has
	// focus, so nav vs. content zones are visually distinct.
	NavZoneHighlight = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	// ContentZoneHighlight draws the matching left rail on the content column
	// when the content zone has focus.
	ContentZoneHighlight = lipgloss.NewStyle().
				Foreground(ColorSecondary).
				Bold(true)
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

	// FooterArtMuted is a dimmed outline style for the decorative mascot /
	// corner illustration so it stays a subtle hiasan, not a distraction.
	FooterArtMuted = lipgloss.NewStyle().
			Foreground(ColorSecondary)
)

// Mascot illustration — layered thin/thick lines with saturation steps.
var (
	MascotBold = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Bold(true)

	MascotThin = lipgloss.NewStyle().
			Foreground(ColorBorder)

	MascotDim = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#1e3a5f"))

	MascotMid = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	MascotBright = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#60a5fa")).
			Bold(true)

	MascotEye = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	MascotCursor = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)
)

// Fixed header bar and body frame styles (decorative "hiasan" around content).
var (
	HeaderBarStyle = lipgloss.NewStyle().
			Foreground(ColorText)

	FrameBorderStyle = lipgloss.NewStyle().
				Foreground(ColorBorder)

	FrameCornerStyle = lipgloss.NewStyle().
				Foreground(ColorSecondary)

	FrameAccentStyle = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	// HeaderScanStyle is the bright dash that travels along the header rule.
	HeaderScanStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)
)

// Scrollbar styles.
var (
	ScrollThumb = lipgloss.NewStyle().
			Foreground(ColorSecondary).
			Bold(true).
			Render("█")

	ScrollTrack = lipgloss.NewStyle().
			Foreground(ColorBorder).
			Render("░")
)
