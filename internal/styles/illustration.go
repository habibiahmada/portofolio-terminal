package styles

import "github.com/charmbracelet/lipgloss"

// Illustration styles — visual identity for signature art, hero, and splash.
var (
	// PromptStyle highlights the ">_" prompt cursor.
	PromptStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	// ArtBorderStyle colors the signature box border.
	ArtBorderStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	// ArtBrandStyle colors the "HABIBI TERMINAL" brand text.
	ArtBrandPrimary = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	// ArtBrandSecondary colors the second brand line.
	ArtBrandSecondary = lipgloss.NewStyle().
				Foreground(ColorSecondary).
				Bold(true)

	// ArtHeroFillStyle renders the thick hero banner letter bodies.
	ArtHeroFillStyle = lipgloss.NewStyle().
				Foreground(ColorText).
				Bold(true)

	// ArtHeroShadowPrimary and ArtHeroShadowSecondary color the orthogonal
	// contour shadow lines (─ │ ┼) on the hero banner.
	ArtHeroShadowPrimary = lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true)

	ArtHeroShadowSecondary = lipgloss.NewStyle().
				Foreground(ColorSecondary).
				Bold(true)

	// HeroTitleStyle renders the large hero name.
	HeroTitleStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginBottom(1)

	// HeroCTAStyle renders the "Explore Portfolio" prompt.
	HeroCTAStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true).
			PaddingLeft(1)

	// ProgressBarStyle renders filled progress blocks.
	ProgressBarStyle = lipgloss.NewStyle().
				Foreground(ColorSuccess).
				Bold(true)

	// ProgressTrackStyle renders empty progress blocks.
	ProgressTrackStyle = lipgloss.NewStyle().
				Foreground(ColorMuted)

	// SplashTextStyle renders status text during the splash sequence.
	SplashTextStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary)
)
