// Package data contains bundled portfolio data for the TUI.
// No external API calls — all data is embedded in the binary.
package data

// Profile represents the developer profile.
type Profile struct {
	Name        string
	Title       string
	Description string
	Location    string
	Email       string
	GitHub      string
	LinkedIn    string
	Website     string
}

// Project represents a portfolio project.
type Project struct {
	Name        string
	Description string
	Stack       []string
	GitHub      string
	Live        string
	Featured    bool
}

// Skill represents a technical skill.
type Skill struct {
	Name       string
	Category   string
	Level      int // 1-5
}

// Experience represents work experience.
type Experience struct {
	Company    string
	Role       string
	Period     string
	Location   string
	Details    []string
}

// Certificate represents a certification.
type Certificate struct {
	Name   string
	Issuer string
	Date   string
	URL    string
}

// Social represents a social media link.
type Social struct {
	Name string
	URL  string
	Icon string
}

// GetProfile returns the developer profile data.
func GetProfile() Profile {
	return Profile{
		Name:        "Habibi Ahmad Aziz",
		Title:       "Full-Stack Web Developer",
		Description: "Passionate developer crafting digital experiences with modern web technologies.",
		Location:    "Indonesia",
		Email:       "habibi@habibiahmada.dev",
		GitHub:      "github.com/habibiahmada",
		LinkedIn:    "linkedin.com/in/habibiahmada",
		Website:     "habibiahmada.dev",
	}
}

// GetProjects returns all portfolio projects.
func GetProjects() []Project {
	return []Project{
		{
			Name:        "Renshuu",
			Description: "Japanese learning platform with spaced repetition, kanji tracking, and community features.",
			Stack:       []string{"Laravel", "React", "Inertia.js", "PostgreSQL"},
			GitHub:      "https://github.com/habibiahmada/renshuu",
			Live:        "",
			Featured:    true,
		},
		{
			Name:        "SmartFarm AI",
			Description: "AI-powered agriculture platform for crop monitoring, pest detection, and yield prediction.",
			Stack:       []string{"Next.js", "Python", "TensorFlow", "FastAPI"},
			GitHub:      "https://github.com/habibiahmada/smartfarm-ai",
			Live:        "",
			Featured:    true,
		},
		{
			Name:        "CultureConnect",
			Description: "Cultural exchange platform connecting people worldwide through language and traditions.",
			Stack:       []string{"Next.js", "TypeScript", "Supabase", "Tailwind CSS"},
			GitHub:      "https://github.com/habibiahmada/cultureconnect",
			Live:        "",
			Featured:    true,
		},
		{
			Name:        "Spacelab",
			Description: "Collaborative workspace management tool with real-time updates and analytics dashboard.",
			Stack:       []string{"React", "Node.js", "WebSocket", "MongoDB"},
			GitHub:      "https://github.com/habibiahmada/spacelab",
			Live:        "",
			Featured:    false,
		},
	}
}

// GetSkills returns all technical skills.
func GetSkills() []Skill {
	return []Skill{
		{Name: "Laravel", Category: "Backend", Level: 5},
		{Name: "Next.js", Category: "Frontend", Level: 5},
		{Name: "React", Category: "Frontend", Level: 5},
		{Name: "Node.js", Category: "Backend", Level: 4},
		{Name: "TypeScript", Category: "Language", Level: 5},
		{Name: "Go", Category: "Language", Level: 3},
		{Name: "PostgreSQL", Category: "Database", Level: 4},
		{Name: "Supabase", Category: "Database", Level: 4},
		{Name: "Docker", Category: "DevOps", Level: 3},
		{Name: "GCP", Category: "Cloud", Level: 3},
		{Name: "Tailwind CSS", Category: "Frontend", Level: 5},
		{Name: "Python", Category: "Language", Level: 3},
	}
}

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

// GetCertificates returns certifications.
func GetCertificates() []Certificate {
	return []Certificate{
		{
			Name:   "Laravel Certified Developer",
			Issuer: "Laravel",
			Date:   "2024",
			URL:    "",
		},
	}
}

// GetSocials returns social media links.
func GetSocials() []Social {
	return []Social{
		{Name: "GitHub", URL: "https://github.com/habibiahmada", Icon: "GH"},
		{Name: "LinkedIn", URL: "https://linkedin.com/in/habibiahmada", Icon: "LI"},
		{Name: "Website", URL: "https://habibiahmada.dev", Icon: "WEB"},
	}
}
