package data

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

// GetSkillCategories returns the unique skill categories in sorted order.
func GetSkillCategories() []string {
	categories := make([]string, 0, 4)
	seen := make(map[string]bool)
	for _, s := range GetSkills() {
		if !seen[s.Category] {
			seen[s.Category] = true
			categories = append(categories, s.Category)
		}
	}
	return categories
}
