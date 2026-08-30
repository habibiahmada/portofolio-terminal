package data

// GetSocials returns social media links.
func GetSocials() []Social {
	return []Social{
		{Name: "GitHub", URL: "https://github.com/habibiahmada", Icon: "GH"},
		{Name: "LinkedIn", URL: "https://linkedin.com/in/habibiahmada", Icon: "LI"},
		{Name: "Website", URL: "https://habibiahmada.dev", Icon: "WEB"},
	}
}
