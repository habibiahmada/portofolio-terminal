package tui

import (
	"strings"

	"github.com/habibiahmada/habibiahmada-terminal/internal/components"
	"github.com/habibiahmada/habibiahmada-terminal/internal/styles"
)

func (m *App) renderPortfolioSyncBanner(width int) string {
	switch m.portfolioSync {
	case PortfolioSyncPending:
		return ""
	case PortfolioSyncCached:
		msg := "⚠ Offline — connect to the internet to load the latest project archive. Showing a local copy. Press R to retry."
		return strings.Join(components.WrapText(styles.WarningStyle.Render(msg), width), "\n")
	default:
		return ""
	}
}
