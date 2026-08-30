package data

// GetSocials returns social media links, synced with docs/pages.md.
func GetSocials() []Social {
	return []Social{
		{Name: "GitHub", URL: "https://github.com/habibiahmada", Icon: "GH"},
		{Name: "LinkedIn", URL: "https://www.linkedin.com/in/habibi-ahmad-aziz", Icon: "LI"},
		{Name: "Instagram", URL: "https://instagram.com/habibiahmad.a", Icon: "IG"},
		{Name: "Email", URL: "mailto:contact@habibiahmada.dev", Icon: "@"},
	}
}
