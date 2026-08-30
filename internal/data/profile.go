package data

// GetProfile returns the developer profile data, synced with docs/pages.md.
func GetProfile() Profile {
	return Profile{
		Name:         "Habibi Ahmad Aziz",
		BoostName:    "Habibi Ahmad",
		Title:        "Fullstack Developer",
		Location:     "Karawang, Indonesia",
		Email:        "contact@habibiahmada.dev",
		Availability: "Open to freelance & full-time · Remote (WIB)",
		Website:      "habibiahmada.dev",
		Employer:     "PT Webekspres Teknologi Indonesia",
		School:       "SMKN 1 Karawang (Software Engineering)",
		Stats: []Stat{
			{Value: "3+", Label: "Years Building"},
			{Value: "10+", Label: "Projects Shipped"},
			{Value: "2", Label: "Awards Won"},
		},
	}
}
