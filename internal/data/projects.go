package data

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

// GetFeaturedProjects returns only the projects marked as featured.
func GetFeaturedProjects() []Project {
	all := GetProjects()
	featured := make([]Project, 0, len(all))
	for _, p := range all {
		if p.Featured {
			featured = append(featured, p)
		}
	}
	return featured
}
