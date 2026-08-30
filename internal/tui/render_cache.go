package tui

import "fmt"

// layoutCacheKey builds a stable key for the main layout cache. footerFrame is
// intentionally excluded so footer animation does not rebuild the whole UI.
func (m *App) layoutCacheKey() string {
	return fmt.Sprintf(
		"%d|%dx%d|off%d|menu%d|proj%d|help%t|cv%t|f%d",
		m.currentScreen,
		m.width,
		m.height,
		m.contentOffset,
		m.selectedMenu,
		m.selectedProject,
		m.showHelp,
		m.cvModal,
		m.focus,
	)
}

// invalidateLayoutCache clears the cached main layout.
func (m *App) invalidateLayoutCache() {
	m.cachedBodyKey = ""
	m.cachedBody = ""
}
