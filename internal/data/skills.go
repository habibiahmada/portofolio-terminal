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
