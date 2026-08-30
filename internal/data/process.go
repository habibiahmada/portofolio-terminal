package data

// GetProcessSteps returns the four-step "How I ship" sequence.
func GetProcessSteps() []ProcessStep {
	return []ProcessStep{
		{
			Number:      "1",
			Title:       "Scope the real job",
			Description: "I write down the user, the constraint, and the non-goals before opening an editor. Fancy stacks wait until the problem is boringly clear.",
		},
		{
			Number:      "2",
			Title:       "Ship a thin vertical slice",
			Description: "One path that works end to end (auth, data, UI) beats a polished shell. Preview deploys keep feedback honest.",
		},
		{
			Number:      "3",
			Title:       "Harden what users can break",
			Description: "Validation at trust boundaries, fail-closed writes, and empty/error states you can explain to a non-engineer.",
		},
		{
			Number:      "4",
			Title:       "Measure, then decorate",
			Description: "Motion and polish come after the page is readable and fast enough on a mid-range phone. If it is not measurable, it is not a claim.",
		},
	}
}
