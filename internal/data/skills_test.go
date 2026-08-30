package data

import "testing"

func TestGetSkills(t *testing.T) {
	skills := GetSkills()
	if len(skills) != 16 {
		t.Errorf("expected exactly 16 skills, got %d", len(skills))
	}

	seen := make(map[string]bool)
	for i, s := range skills {
		if s.Name == "" {
			t.Errorf("skill[%d]: Name must not be empty", i)
		}
		if seen[s.Name] {
			t.Errorf("duplicate skill %q at index %d", s.Name, i)
		}
		seen[s.Name] = true
	}
}

func TestGetSkillCategories(t *testing.T) {
	cats := GetSkillCategories()
	if len(cats) == 0 {
		t.Fatal("expected non-empty skill categories")
	}
	for i, c := range cats {
		if c.Name == "" {
			t.Errorf("category[%d]: Name must not be empty", i)
		}
		if len(c.Skills) == 0 {
			t.Errorf("category %q has no skills", c.Name)
		}
	}
}
