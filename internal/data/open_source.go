package data

// GetOpenSourceProjects returns live repositories when the portfolio API
// answered, otherwise the bundled copy.
func GetOpenSourceProjects() []OpenSourceProject {
	if live, ok := liveOpenSourceCopy(); ok {
		return live
	}
	return bundledOpenSourceProjects()
}

// bundledOpenSourceProjects is the offline fallback shipped with the binary.
// Keep it in sync with lib/data/open-source.ts in the portfolio repo.
func bundledOpenSourceProjects() []OpenSourceProject {
	return []OpenSourceProject{
		{
			Slug:        "autumnnote-plus",
			Name:        "autumnnote-plus",
			Kind:        "npm package",
			Description: "A WYSIWYG rich-text editor in plain JavaScript. No jQuery, no runtime dependencies, no execCommand, so Chrome, Firefox, and Safari produce the same markup.",
			Facts:       []string{"0 runtime deps", "React + Vue wrappers", "9 UI languages"},
			Language:    "JavaScript",
			License:     "MIT",
			Install:     "npm install autumnnote-plus",
			RepoURL:     "https://github.com/habibiahmada/autumnnote-plus",
			SiteURL:     "https://autumnnote.habibiahmada.dev",
			SiteLabel:   "Docs",
			ForkOf:      "Autumn Note 3.1.0 (https://github.com/cmm-cmm/Autumn-Note)",
		},
		{
			Slug:        "portofolio-terminal",
			Name:        "habibiahmada",
			Kind:        "CLI portfolio",
			Description: "This portfolio as a terminal app. A Go TUI built with Bubble Tea, shipped as one binary per platform through npx and served over SSH.",
			Facts:       []string{"Go + Bubble Tea", "5 platform builds", "Live portfolio sync"},
			Language:    "Go",
			License:     "MIT",
			Install:     "npx habibiahmada",
			RepoURL:     "https://github.com/habibiahmada/portofolio-terminal",
			SiteURL:     "https://www.npmjs.com/package/habibiahmada",
			SiteLabel:   "npm",
		},
		{
			Slug:        "emoji-picker",
			Name:        "emoji-picker",
			Kind:        "Linux desktop app",
			Description: "An emoji popup for low-RAM Linux X11 desktops. Press Win+. and the emoji lands in the window you were typing in. No Electron, no Flatpak.",
			Facts:       []string{"~225 KB binary", "~30 MB idle", "1,914 emoji"},
			Language:    "C + GTK3",
			License:     "MIT",
			RepoURL:     "https://github.com/habibiahmada/emoji-picker",
		},
	}
}
