// Package sanitize strips terminal escape sequences and control characters from
// plain text before it is rendered in the TUI.
package sanitize

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
)

// Plain removes ANSI/OSC escape sequences and C0 control characters (except
// newline, tab, and carriage return) from s. Use on any user-controllable or
// externally sourced string before lipgloss styling.
func Plain(s string) string {
	s = ansi.Strip(s)
	return strings.Map(func(r rune) rune {
		switch r {
		case '\n', '\t', '\r':
			return r
		}
		if r < 0x20 || r == 0x7f {
			return -1
		}
		return r
	}, s)
}

// Strings applies Plain to each element.
func Strings(ss []string) []string {
	if len(ss) == 0 {
		return ss
	}
	out := make([]string, len(ss))
	for i, s := range ss {
		out[i] = Plain(s)
	}
	return out
}
