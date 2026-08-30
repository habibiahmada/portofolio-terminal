package data

// GetExperiences returns work experiences.
func GetExperiences() []Experience {
	return []Experience{
		{
			Company:  "Freelance Developer",
			Role:     "Full-Stack Web Developer",
			Period:   "2022 — Present",
			Location: "Remote",
			Details: []string{
				"Built web applications for various clients using Laravel, Next.js, and React",
				"Developed RESTful APIs and integrated third-party services",
				"Implemented responsive UI/UX designs and optimized performance",
			},
		},
	}
}
