// Package data contains bundled portfolio data for the TUI.
// No external API calls — all data is embedded in the binary.
//
// Shared struct definitions live here; concrete data and getters live in
// per-domain files: profile.go, projects.go, skills.go, experience.go,
// certificates.go, socials.go.
package data

// Profile represents the developer profile.
type Profile struct {
	Name         string
	BoostName    string // short display name used in the brand/hero
	Title        string
	Location     string
	Email        string
	Availability string
	Website      string
	Employer     string
	School       string
	Stats        []Stat
}

// Stat is a headline figure shown on the About hero (e.g. "3+ Years").
type Stat struct {
	Value string
	Label string
}

// Project represents a portfolio project.
type Project struct {
	Name          string
	Slug          string
	Year          string
	Description   string // EN description
	DescriptionID string // ID description (stored, shown in later phase)
	Tags          []string
	Stack         []string
	GitHub        string
	Live          string
	Featured      bool
}

// CaseStudy is one narrative section of a project's case study. The four
// sections follow docs/pages.md: opening, constraints, build, close.
type CaseStudy struct {
	Slug     string
	Title    string
	Year     string
	Tags     []string
	Hero     string
	Sections []CaseStudySection
	Live     string
	Source   string
}

// CaseStudySection is a hook-label/narrative pair inside a case study.
type CaseStudySection struct {
	Label string // e.g. "Where it started"
	Body  string
}

// Skill represents a technical skill.
type Skill struct {
	Name string
}

// SkillCategory groups related skills with contextual description.
type SkillCategory struct {
	Name        string
	Icon        string
	Description string
	Skills      []string
}

// ExperienceWork represents a work experience entry.
type ExperienceWork struct {
	Period   string
	Role     string
	Company  string
	Location string
	Badge    string
	Details  []string
}

// ExperienceEducation represents an education entry.
type ExperienceEducation struct {
	Period      string
	Title       string
	School      string
	Description string
}

// Certificate represents a certification.
type Certificate struct {
	Name   string
	Issuer string
	Date   string
	URL    string
	Pinned bool
}

// Social represents a social media link.
type Social struct {
	Name string
	URL  string
	Icon string
}

// Company represents a partner/collaboration shown in the marquee.
type Company struct {
	Name string
}

// Service represents a single labelled service offering.
type Service struct {
	Number      string // e.g. "01"
	Category    string // e.g. "Design"
	Title       string
	Description string
}

// ProcessStep is one step of the "How I ship" sequence.
type ProcessStep struct {
	Number      string
	Title       string
	Description string
}

// Press is a spotlight story (featured press / award).
type Press struct {
	Title    string
	Body     string
	CTALabel string
	URL      string
}
