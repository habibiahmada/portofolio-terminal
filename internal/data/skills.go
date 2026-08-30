package data

// skillOrder is the flat marquee order of tools from docs/pages.md. Category
// and level bars were removed to match the website's flat marquee.
var skillOrder = []string{
	"React", "Next.js", "Node.js", "TypeScript", "PostgreSQL",
	"Tailwind CSS", "PHP", "Laravel", "WordPress", "Elementor",
	"Astra", "Git", "GitHub", "Bootstrap", "Vercel", "JavaScript",
}

// GetSkills returns the flat list of technical tools in marquee order.
func GetSkills() []Skill {
	skills := make([]Skill, 0, len(skillOrder))
	for _, name := range skillOrder {
		skills = append(skills, Skill{Name: name})
	}
	return skills
}

// GetSkillCategories returns structured skill groups with icons and contextual summaries.
func GetSkillCategories() []SkillCategory {
	return []SkillCategory{
		{
			Name:        "Frontend & UI Development",
			Icon:        "⚛",
			Description: "Responsive interfaces, component architecture, and modern styling systems.",
			Skills:      []string{"React", "Next.js", "TypeScript", "JavaScript", "Tailwind CSS", "Bootstrap"},
		},
		{
			Name:        "Backend & API Engineering",
			Icon:        "⬡",
			Description: "Scalable server logic, RESTful APIs, authentication, and secure data handling.",
			Skills:      []string{"Node.js", "Express", "PHP", "Laravel"},
		},
		{
			Name:        "Databases & Data Modeling",
			Icon:        "◉",
			Description: "Relational schema design, query optimization, migrations, and ORM workflows.",
			Skills:      []string{"PostgreSQL", "MySQL", "Prisma"},
		},
		{
			Name:        "Cloud, DevOps & Tooling",
			Icon:        "▲",
			Description: "Automated deployments, serverless hosting, version control, and CI/CD pipelines.",
			Skills:      []string{"AWS", "Vercel", "Docker", "Git", "GitHub"},
		},
		{
			Name:        "CMS & Web Platforms",
			Icon:        "◆",
			Description: "Tailored client CMS solutions, custom WordPress themes, and production web delivery.",
			Skills:      []string{"WordPress", "Elementor", "Astra"},
		},
	}
}
