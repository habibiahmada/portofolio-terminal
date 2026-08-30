// Package data contains bundled portfolio data for the TUI.
// No external API calls — all data is embedded in the binary.
//
// Shared struct definitions live here; concrete data and getters live in
// per-domain files: profile.go, projects.go, skills.go, experience.go,
// certificates.go, socials.go.
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
	Name     string
	Category string
	Level    int // 1-5
}

// Experience represents work experience.
type Experience struct {
	Company  string
	Role     string
	Period   string
	Location string
	Details  []string
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
