package data

import "testing"

func TestGetWorkExperience(t *testing.T) {
	work := GetWorkExperience()
	if len(work) != 4 {
		t.Fatalf("expected 4 work entries, got %d", len(work))
	}

	for i, w := range work {
		if w.Period == "" {
			t.Errorf("work[%d]: Period must not be empty", i)
		}
		if w.Role == "" {
			t.Errorf("work[%d]: Role must not be empty", i)
		}
		if w.Company == "" {
			t.Errorf("work[%d]: Company must not be empty", i)
		}
		if len(w.Details) == 0 {
			t.Errorf("work[%d]: Details must not be empty", i)
		}
	}
}

func TestGetEducation(t *testing.T) {
	edu := GetEducation()
	if len(edu) != 2 {
		t.Fatalf("expected 2 education entries, got %d", len(edu))
	}

	for i, e := range edu {
		if e.Title == "" {
			t.Errorf("education[%d]: Title must not be empty", i)
		}
		if e.School == "" {
			t.Errorf("education[%d]: School must not be empty", i)
		}
	}
}
