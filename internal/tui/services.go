package tui

import (
	"fmt"
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderServicesContent() string {
	lines := []string{styles.SectionTitleStyle.Render("Services")}
	for _, sv := range m.services {
		lines = append(lines, fmt.Sprintf("%s %s — %s",
			styles.MutedStyle.Render(sv.Number),
			styles.NormalStyle.Render(sv.Title),
			styles.MutedStyle.Render(truncatePlain(sv.Description, 48)),
		))
	}
	return strings.Join(lines, "\n")
}

func truncatePlain(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}
