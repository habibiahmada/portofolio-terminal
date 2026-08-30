package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/habibiahmada/habibiahmada-terminal/internal/assets"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

// Variant describes which illustration fidelity fits the current terminal width.
type Variant int

const (
	VariantWide        Variant = iota // >= 100 cols — full signature art
	VariantComfortable                // 80–99 — compact signature
	VariantNarrow                     // 60–79 — inline mini signature
	VariantHidden                     // < 60 — no art, content first
)

// IllustrationVariant picks an illustration variant for a given width.
func IllustrationVariant(width int) Variant {
	switch {
	case width >= 100:
		return VariantWide
	case width >= 80:
		return VariantComfortable
	case width >= 60:
		return VariantNarrow
	default:
		return VariantHidden
	}
}

// signatureArt returns the raw art file matching the variant for the width.
func signatureArt(width int) string {
	switch IllustrationVariant(width) {
	case VariantWide:
		return assets.Art("signature_wide.txt")
	case VariantComfortable:
		return assets.Art("signature_compact.txt")
	case VariantNarrow:
		return assets.Art("signature_mini.txt")
	default:
		return ""
	}
}

// renderArt styles each art line according to its content.
func renderArt(art string) string {
	lines := strings.Split(art, "\n")
	styled := make([]string, 0, len(lines))
	for _, line := range lines {
		styled = append(styled, styleArtLine(line))
	}
	return lipgloss.JoinVertical(lipgloss.Left, styled...)
}

// Signature renders the "Habibi Terminal" signature art, styled per line.
// It returns an empty string when the terminal is too narrow for art.
func Signature(width int) string {
	return renderArt(signatureArt(width))
}

// SignatureBlink renders the signature with a blinking cursor. When showCursor
// is false the ">_" prompt is rendered without its underscore.
func SignatureBlink(width int, showCursor bool) string {
	art := signatureArt(width)
	if art == "" {
		return ""
	}
	if !showCursor {
		art = strings.ReplaceAll(art, ">_", "> ")
	}
	return renderArt(art)
}

// AboutTerminal renders the small terminal illustration for the About screen.
func AboutTerminal() string {
	return renderArt(assets.Art("about_terminal.txt"))
}

// heroFont is a 5-row block-letter font used for the Home hero banner.
// █ marks letter body; empty cells are structural voids that receive contour
// shadow lines (─ │) from heroCastShadows, never solid accent fills.
var heroFont = map[rune][5]string{
	'H': {"██   ██", "██   ██", "███████", "██   ██", "██   ██"},
	'A': {"███████", "██   ██", "███████", "██   ██", "██   ██"},
	'B': {"██████ ", "██   ██ ", "██████ ", "██   ██ ", "██████ "},
	'I': {"███████", "  ███  ", "  ███  ", "  ███  ", "███████"},
}

const (
	heroRows      = 5
	heroShadowDX  = 1
	heroShadowDY  = 1
	heroCellFill  = '█'
	heroCellEmpty = ' '
)

type heroLayer int

const (
	heroEmpty heroLayer = iota
	heroShadowLineH
	heroShadowLineV
	heroShadowCross
	heroFill
)

type heroCell struct {
	layer  heroLayer
	style  lipgloss.Style
	letter int
}

type heroLetterMask struct {
	originX int
	originY int
	width   int
	height  int
	filled  [][]bool
}

// CenterText centers a (possibly ANSI-styled) string within width using a
// simple left-padding of spaces. Multi-line strings should be centered line by
// line. When the text already fits width it is returned unchanged.
func CenterText(s string, width int) string {
	if width <= 0 {
		return s
	}
	w := lipgloss.Width(s)
	if w >= width {
		return s
	}
	return strings.Repeat(" ", (width-w)/2) + s
}

// HeroBanner renders word vertically-centered block letters (FIGlet style),
// centered within width. Letter bodies are thick white blocks with internal
// red/blue accents and contour-following shadow lines offset to the
// bottom-right. It returns an empty string when the terminal is too narrow for
// the banner.
func HeroBanner(word string, width int) string {
	glyphs := []rune(strings.ToUpper(word))
	letters := make([]rune, 0, len(glyphs))
	for _, g := range glyphs {
		if _, ok := heroFont[g]; ok {
			letters = append(letters, g)
		}
	}
	if len(letters) == 0 {
		return ""
	}

	gridW := 0
	for li, g := range letters {
		gridW += len([]rune(heroFont[g][0]))
		if li > 0 {
			gridW++
		}
	}
	gridW += heroShadowDX
	gridH := heroRows + heroShadowDY

	grid := make([][]heroCell, gridH)
	for y := range grid {
		grid[y] = make([]heroCell, gridW)
	}

	palette := []lipgloss.Style{styles.ArtHeroShadowPrimary, styles.ArtHeroShadowSecondary}
	masks := make([]heroLetterMask, 0, len(letters))

	x := 0
	for li, g := range letters {
		if li > 0 {
			x++
		}
		lines := heroFont[g]
		width := len([]rune(lines[0]))
		mask := heroLetterMask{
			originX: x,
			originY: 0,
			width:   width,
			height:  heroRows,
			filled:  make([][]bool, heroRows),
		}
		for y := range mask.filled {
			mask.filled[y] = make([]bool, width)
		}

		for y, line := range lines {
			for dx, ch := range []rune(line) {
				gx := x + dx
				if ch == heroCellFill {
					grid[y][gx] = heroCell{heroFill, styles.ArtHeroFillStyle, li}
					mask.filled[y][dx] = true
				}
			}
		}
		masks = append(masks, mask)
		x += width
	}

	heroCastShadows(grid, masks, palette)
	heroCastInnerContours(grid, masks, palette)

	if gridW > width {
		return ""
	}

	out := make([]string, gridH)
	for y := 0; y < gridH; y++ {
		var b strings.Builder
		for x := 0; x < gridW; x++ {
			b.WriteString(heroRenderCell(grid, x, y))
		}
		out[y] = CenterText(b.String(), width)
	}
	return strings.Join(out, "\n")
}

func heroRenderCell(grid [][]heroCell, x, y int) string {
	c := grid[y][x]
	switch c.layer {
	case heroEmpty:
		return string(heroCellEmpty)
	case heroFill:
		return c.style.Render(string(heroCellFill))
	case heroShadowLineH, heroShadowLineV, heroShadowCross:
		return c.style.Render(string(heroShadowRune(grid, x, y)))
	default:
		return string(heroCellEmpty)
	}
}

// heroCellShadow reports whether a grid cell participates in horizontal and/or
// vertical shadow strokes.
func heroCellShadow(grid [][]heroCell, x, y int) (horizontal, vertical bool) {
	if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[0]) {
		return false, false
	}
	switch grid[y][x].layer {
	case heroShadowLineH:
		return true, false
	case heroShadowLineV:
		return false, true
	case heroShadowCross:
		return true, true
	default:
		return false, false
	}
}

// heroShadowRune renders shadow strokes. Neighbouring perpendicular lines keep
// their own cell glyph (─ or │) so corners never curl upward into letter fills.
// Only same-cell crosses pick a corner, preferring the bottom shadow row.
func heroShadowRune(grid [][]heroCell, x, y int) rune {
	c := grid[y][x]
	switch c.layer {
	case heroShadowLineH:
		return '─'
	case heroShadowLineV:
		return '│'
	case heroShadowCross:
		return heroShadowCornerRune(grid, x, y)
	default:
		return ' '
	}
}

func heroShadowCornerRune(grid [][]heroCell, x, y int) rune {
	// Bottom shadow row: horizontal stroke continues outward, not up into the glyph.
	if y >= heroRows {
		return '─'
	}
	// Right-edge column inside the glyph band: keep the vertical drop.
	_, upV := heroCellShadow(grid, x, y-1)
	_, downV := heroCellShadow(grid, x, y+1)
	if upV || downV {
		return '│'
	}
	return '─'
}

func heroCastShadows(grid [][]heroCell, masks []heroLetterMask, palette []lipgloss.Style) {
	occupied := func(x, y int) bool {
		if y < 0 || y >= heroRows || x < 0 || x >= len(grid[0]) {
			return false
		}
		switch grid[y][x].layer {
		case heroFill:
			return true
		default:
			return false
		}
	}

	markShadow := func(x, y, letter int, horizontal bool) {
		if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
			return
		}
		if occupied(x, y) {
			return
		}
		shadow := palette[letter%len(palette)]
		c := &grid[y][x]
		switch c.layer {
		case heroEmpty:
			if horizontal {
				c.layer = heroShadowLineH
			} else {
				c.layer = heroShadowLineV
			}
			c.style = shadow
			c.letter = letter
		case heroShadowLineH:
			if !horizontal {
				// Prefer horizontal on the bottom shadow row, vertical above it.
				if y >= heroRows {
					c.layer = heroShadowLineH
				} else {
					c.layer = heroShadowLineV
				}
				c.style = shadow
				c.letter = letter
			}
		case heroShadowLineV:
			if horizontal {
				if y >= heroRows {
					c.layer = heroShadowLineH
				} else {
					c.layer = heroShadowLineV
				}
				c.style = shadow
				c.letter = letter
			}
		}
	}

	for li, mask := range masks {
		for y := 0; y < mask.height; y++ {
			for lx := 0; lx < mask.width; lx++ {
				if !mask.filled[y][lx] {
					continue
				}
				gx := mask.originX + lx
				gy := mask.originY + y

				if !heroMaskFilled(mask, lx, y+1) && !heroMaskStructuralVoid(mask, lx, y+1) {
					markShadow(gx+heroShadowDX, gy+heroShadowDY, li, true)
				}
				if !heroMaskFilled(mask, lx+1, y) && !heroMaskStructuralVoid(mask, lx+1, y) {
					markShadow(gx+heroShadowDX, gy, li, false)
				}
			}
		}
	}
}

// heroCastInnerContours draws horizontal contour lines inside structural voids
// (B loops, A aperture). Vertical strokes are omitted to avoid plus-shaped joints.
func heroCastInnerContours(grid [][]heroCell, masks []heroLetterMask, palette []lipgloss.Style) {
	for li, mask := range masks {
		shadow := palette[li%len(palette)]
		for y := 0; y < mask.height; y++ {
			for lx := 0; lx < mask.width; lx++ {
				if !heroMaskStructuralVoid(mask, lx, y) {
					continue
				}
				if !heroMaskFilled(mask, lx, y-1) && !heroMaskFilled(mask, lx, y+1) {
					continue
				}
				gx := mask.originX + lx
				gy := mask.originY + y
				heroMarkVoidContour(grid, gx, gy, li, shadow)
			}
		}
	}
}

func heroMarkVoidContour(grid [][]heroCell, x, y, letter int, shadow lipgloss.Style) {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
		return
	}
	c := &grid[y][x]
	if c.layer == heroFill {
		return
	}
	switch c.layer {
	case heroEmpty:
		c.layer = heroShadowLineH
		c.style = shadow
		c.letter = letter
	case heroShadowLineV:
		// Inner void lip is horizontal; do not merge into a vertical corner.
		c.layer = heroShadowLineH
		c.style = shadow
		c.letter = letter
	}
}

func heroMaskFilled(mask heroLetterMask, x, y int) bool {
	if y < 0 || y >= mask.height || x < 0 || x >= mask.width {
		return false
	}
	return mask.filled[y][x]
}

// heroMaskStructuralVoid reports letter-local empty cells sandwiched inside the
// glyph, such as the waist indent of B, so they do not receive shadow strokes.
func heroMaskStructuralVoid(mask heroLetterMask, x, y int) bool {
	if heroMaskFilled(mask, x, y) {
		return false
	}
	up := heroMaskFilled(mask, x, y-1)
	down := heroMaskFilled(mask, x, y+1)
	left := heroMaskFilled(mask, x-1, y)
	right := heroMaskFilled(mask, x+1, y)
	return (up && down) || (left && right)
}

// styleArtLine applies the palette to a single art line based on its content.
func styleArtLine(line string) string {
	switch {
	case strings.Contains(line, ">_"):
		return styles.PromptStyle.Render(line)
	case strings.Contains(line, "HABIBI"):
		return styles.ArtBrandPrimary.Render(line)
	case strings.Contains(line, "TERMINAL"):
		return styles.ArtBrandSecondary.Render(line)
	case strings.ContainsAny(line, "╭╮╰╯│"):
		return styles.ArtBorderStyle.Render(line)
	default:
		return line
	}
}
