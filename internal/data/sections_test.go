package data

import "testing"

func TestGetCompanies(t *testing.T) {
	companies := GetCompanies()
	if len(companies) != 5 {
		t.Fatalf("expected 5 companies, got %d", len(companies))
	}
	for i, c := range companies {
		if c.Name == "" {
			t.Errorf("company[%d]: Name must not be empty", i)
		}
	}
}

func TestGetServices(t *testing.T) {
	services := GetServices()
	if len(services) != 5 {
		t.Fatalf("expected 5 services, got %d", len(services))
	}
	for i, s := range services {
		if s.Number == "" {
			t.Errorf("service[%d]: Number must not be empty", i)
		}
		if s.Title == "" {
			t.Errorf("service[%d]: Title must not be empty", i)
		}
		if s.Description == "" {
			t.Errorf("service[%d]: Description must not be empty", i)
		}
	}
}

func TestGetProcessSteps(t *testing.T) {
	steps := GetProcessSteps()
	if len(steps) != 4 {
		t.Fatalf("expected 4 process steps, got %d", len(steps))
	}
	for i, s := range steps {
		if s.Title == "" || s.Description == "" {
			t.Errorf("process step %d must have Title and Description", i)
		}
	}
}

func TestGetPress(t *testing.T) {
	press := GetPress()
	if len(press) != 2 {
		t.Fatalf("expected 2 press stories, got %d", len(press))
	}
	for i, p := range press {
		if p.Title == "" || p.Body == "" || p.CTALabel == "" {
			t.Errorf("press %d must have Title, Body, and CTALabel", i)
		}
	}
}

func TestGetCaseStudiesCoverAllProjects(t *testing.T) {
	for _, p := range bundledProjects() {
		if bundledCaseStudy(p.Slug) == nil {
			t.Errorf("missing case study for slug %q", p.Slug)
		}
	}
}
