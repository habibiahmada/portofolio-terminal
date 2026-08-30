// Package animation provides a tiny frame-based animation engine for the TUI.
//
// A TUI animation is just a sequence of frames redrawn on a timer. This package
// centralises that concept so screens can drive subtle looping animations
// (spinners, cursor blinks, moving selectors, splash sequences) without each
// screen wiring its own timer. The actual timing is handled by the App's single
// footer tick; an Animation merely advances its frame index on demand.
//
//	anim := animation.New([]string{"⠋", "⠙", "⠹"}) // 3 frames, loops by default
//	anim.Next() // "⠋"
//	anim.Next() // "⠙"
//	anim.Next() // "⠹"
//	anim.Next() // "⠋" (wraps)
package animation

// Animation is a frame-based animation controller.
type Animation struct {
	frames  []string
	Loop    bool // wrap around (true) or stop at the end (false)
	Playing bool // paused when false, Next() freezes the current frame
	index   int
}

// New returns an animation over the given frames. It loops by default.
func New(frames []string) *Animation {
	if len(frames) == 0 {
		frames = []string{" "}
	}
	return &Animation{frames: frames, Loop: true, Playing: true}
}

// Next advances to the next frame and returns it. When paused it just returns
// the current frame; when looping it wraps; otherwise it clamps to the last.
func (a *Animation) Next() string {
	if a.Playing {
		if a.Loop {
			a.index = (a.index + 1) % len(a.frames)
		} else if a.index < len(a.frames)-1 {
			a.index++
		}
	}
	return a.frames[a.index]
}

// Current returns the frame without advancing.
func (a *Animation) Current() string {
	return a.frames[a.index]
}

// Frame returns the frame at a fixed index (independent of the running index).
func (a *Animation) Frame(i int) string {
	n := len(a.frames)
	if n == 0 {
		return " "
	}
	return a.frames[((i%n)+n)%n]
}

// Start resumes playback.
func (a *Animation) Start() *Animation {
	a.Playing = true
	return a
}

// Stop pauses playback.
func (a *Animation) Stop() *Animation {
	a.Playing = false
	return a
}

// Reset returns to the first frame.
func (a *Animation) Reset() *Animation {
	a.index = 0
	return a
}

// Len returns the number of frames.
func (a *Animation) Len() int {
	return len(a.frames)
}

// Preset spinner frame sets.
var (
	// Spinner is a classic braille spinner.
	Spinner = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	// BlockPulse cycles a small block through four partial-fill states.
	BlockPulse = []string{"▏", "▎", "▍", "▌", "▍", "▎"}

	// Cursor is a two-state blink.
	Cursor = []string{"▊", " "}

	// Shimmer moves a highlight across a run of the same width (for accents).
	Shimmer = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"}
)
