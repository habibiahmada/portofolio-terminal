package data

import "testing"

func TestGetExperiences(t *testing.T) {
	experiences := GetExperiences()
	if len(experiences) == 0 {
		t.Fatal("expected at least one experience")
	}

	for i, e := range experiences {
		if e.Company == "" {
			t.Errorf("experience[%d]: Company must not be empty", i)
		}
		if e.Role == "" {
			t.Errorf("experience[%d]: Role must not be empty", i)
		}
		if e.Period == "" {
			t.Errorf("experience[%d]: Period must not be empty", i)
		}
		if len(e.Details) == 0 {
			t.Errorf("experience[%d]: Details must not be empty", i)
		}
	}
}
