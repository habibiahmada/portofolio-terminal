package data

// GetWorkExperience returns the four work entries per docs/pages.md.
func GetWorkExperience() []ExperienceWork {
	return []ExperienceWork{
		{
			Period:   "May - Nov 2026",
			Role:     "Web Developer",
			Company:  "PT Webekspres Teknologi Indonesia",
			Location: "Karawang · On site",
			Details: []string{
				"Client sites, news portals, and catalogs, plus my part in MIRU, HK GOLD VIP, and JepangKu.",
				"Owned features end to end, from brief to deploy, and built the team's shared CI/CD platform.",
			},
		},
		{
			Period:   "Jun - Aug 2025",
			Role:     "Cloud Computing Trainer Intern",
			Company:  "Yayasan Sagasitas Indonesia",
			Location: "Jakarta · On site",
			Badge:    "",
			Details: []string{
				"Taught Cloud Computing and Generative AI in schools.",
				"Built AWS PartyRock labs and kept teaching teams aligned with partner schools.",
			},
		},
		{
			Period:   "Jan - Apr 2025",
			Role:     "Student Member",
			Company:  "Coding Camp powered by DBS Foundation",
			Location: "Bandung · Remote",
			Badge:    "Top 15 Capstone",
			Details: []string{
				"Full-stack track under real deadlines.",
				"CultureConnect landed in the Top 15 Best Capstone Projects.",
			},
		},
		{
			Period:   "Jan - May 2025",
			Role:     "Web Developer Intern",
			Company:  "CV. SmartPlus Indonesia",
			Location: "Karawang · Remote",
			Badge:    "",
			Details: []string{
				"Full-stack intern on internal company web projects.",
				"Shipped features end to end with modern stacks.",
			},
		},
	}
}

// GetEducation returns the education "Foundations" entries.
func GetEducation() []ExperienceEducation {
	return []ExperienceEducation{
		{
			Period:      "2023 - 2026",
			Title:       "Software Engineering",
			School:      "SMK Negeri 1 Karawang",
			Description: "Software development, programming, systems, and networking. Active in tech projects and competitions.",
		},
		{
			Period:      "2020 - 2023",
			Title:       "Arabic Language and Literature",
			School:      "MTSS Darunnadwah 01",
			Description: "Arabic language and literature with a focus on communication and text analysis.",
		},
	}
}
