package data

// GetServices returns the five service offerings per docs/pages.md.
func GetServices() []Service {
	return []Service{
		{
			Number:      "01",
			Category:    "Design",
			Title:       "Web Design & Mobile-First",
			Description: "Translating ideas into pixel-perfect responsive interfaces. Wireframes to production-ready layouts that feel intuitive on every device.",
		},
		{
			Number:      "02",
			Category:    "Engineering",
			Title:       "Frontend Development",
			Description: "High-quality UIs with React & Next.js. Clean components, reusable logic, and fluid state management.",
		},
		{
			Number:      "03",
			Category:    "Performance",
			Title:       "Web Performance",
			Description: "Core Web Vitals optimization for instant load. SEO-ready architecture that ranks and converts.",
		},
		{
			Number:      "04",
			Category:    "Backend",
			Title:       "APIs & Databases",
			Description: "Robust REST APIs, relational databases, and secure auth systems built to scale with your product.",
		},
		{
			Number:      "05",
			Category:    "DevOps",
			Title:       "CI/CD & Deployment",
			Description: "Automated pipelines, container-ready apps, serverless hosting, and zero-downtime production deploys.",
		},
	}
}
