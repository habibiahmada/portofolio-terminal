package data

import "testing"

func TestGetProjects(t *testing.T) {
	projects := GetProjects()
	if len(projects) == 0 {
		t.Fatal("expected at least one project")
	}

	for i, p := range projects {
		if p.Name == "" {
			t.Errorf("project[%d]: Name must not be empty", i)
		}
		if p.Description == "" {
			t.Errorf("project[%d] %q: Description must not be empty", i, p.Name)
		}
		if len(p.Stack) == 0 {
			t.Errorf("project[%d] %q: Stack must not be empty", i, p.Name)
		}
	}
}

func TestGetFeaturedProjects(t *testing.T) {
	featured := GetFeaturedProjects()
	if len(featured) == 0 {
		t.Fatal("expected at least one featured project")
	}

	for _, p := range featured {
		if !p.Featured {
			t.Errorf("project %q returned by GetFeaturedProjects is not flagged Featured", p.Name)
		}
	}
}
