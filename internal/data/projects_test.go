package data

import "testing"

func TestGetProjects(t *testing.T) {
	projects := GetProjects()
	if len(projects) != 10 {
		t.Fatalf("expected 10 projects, got %d", len(projects))
	}

	seenSlug := make(map[string]bool)
	for i, p := range projects {
		if p.Name == "" {
			t.Errorf("project[%d]: Name must not be empty", i)
		}
		if p.Slug == "" {
			t.Errorf("project[%d] %q: Slug must not be empty", i, p.Name)
		}
		if seenSlug[p.Slug] {
			t.Errorf("duplicate slug %q", p.Slug)
		}
		seenSlug[p.Slug] = true
		if p.Description == "" {
			t.Errorf("project[%d] %q: Description must not be empty", i, p.Name)
		}
		if p.DescriptionID == "" {
			t.Errorf("project[%d] %q: DescriptionID must not be empty", i, p.Name)
		}
		if p.Year == "" {
			t.Errorf("project[%d] %q: Year must not be empty", i, p.Name)
		}
		if len(p.Tags) == 0 {
			t.Errorf("project[%d] %q: Tags must not be empty", i, p.Name)
		}
	}
}

func TestGetFeaturedProjects(t *testing.T) {
	featured := GetFeaturedProjects()
	if len(featured) != 5 {
		t.Fatalf("expected 5 featured projects, got %d", len(featured))
	}

	for _, p := range featured {
		if !p.Featured {
			t.Errorf("project %q returned by GetFeaturedProjects is not flagged Featured", p.Name)
		}
	}
}

func TestEveryProjectHasCaseStudy(t *testing.T) {
	for _, p := range GetProjects() {
		cs := GetCaseStudy(p.Slug)
		if cs == nil {
			t.Errorf("project %q has no case study", p.Slug)
			continue
		}
		if len(cs.Sections) != 4 {
			t.Errorf("case study %q expected 4 sections, got %d", p.Slug, len(cs.Sections))
		}
	}
}
