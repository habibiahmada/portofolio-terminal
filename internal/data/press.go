package data

// GetPress returns the two spotlight stories shown on Home.
func GetPress() []Press {
	return []Press{
		{
			Title:    "From zero tech background to shipping under real deadlines.",
			Body:     "Coding Camp did not hand me confidence. It forced me to build, present, and intern at the same time. That pressure is still how I work: clear scope, thin vertical slice, ship.",
			CTALabel: "Read the Dicoding story",
		},
		{
			Title:    "Indonesia country award for Agrify at a global AI festival.",
			Body:     "Our team built Agrify for AI Changemakers. I owned the product surface: turn model output into advice a farmer can act on, not another pretty dashboard.",
			CTALabel: "See the Intel winners list",
		},
	}
}
