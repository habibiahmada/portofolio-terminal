package tui

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderCertificatesContent() string {
	pinned := make([]string, 0, 3)
	rest := make([]string, 0, len(m.certificates))
	for _, c := range m.certificates {
		if c.Pinned {
			pinned = append(pinned, "★ "+c.Name)
			continue
		}
		rest = append(rest, c.Name)
	}

	lines := []string{styles.SectionTitleStyle.Render("Certificates")}
	lines = append(lines, pinned...)
	if len(rest) > 0 {
		lines = append(lines, components.TagGrid(rest, m.contentWidth()))
	}
	return strings.Join(lines, "\n")
}
