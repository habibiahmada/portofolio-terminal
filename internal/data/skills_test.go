package data

import "testing"

func TestGetSkills(t *testing.T) {
	skills := GetSkills()
	if len(skills) == 0 {
		t.Fatal("expected at least one skill")
	}

	for i, s := range skills {
		if s.Name == "" {
			t.Errorf("skill[%d]: Name must not be empty", i)
		}
		if s.Category == "" {
			t.Errorf("skill %q: Category must not be empty", s.Name)
		}
		if s.Level < 1 || s.Level > 5 {
			t.Errorf("skill %q: Level %d out of range (1-5)", s.Name, s.Level)
		}
	}
}

func TestGetSkillCategories(t *testing.T) {
	categories := GetSkillCategories()
	if len(categories) == 0 {
		t.Fatal("expected at least one skill category")
	}

	seen := make(map[string]bool)
	for i, c := range categories {
		if seen[c] {
			t.Errorf("duplicate category %q at index %d", c, i)
		}
		seen[c] = true
	}
}
