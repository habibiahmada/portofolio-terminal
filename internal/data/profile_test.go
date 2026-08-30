package data

import "testing"

func TestGetProfile(t *testing.T) {
	p := GetProfile()

	if p.Name == "" {
		t.Error("expected profile Name to be non-empty")
	}
	if p.Title == "" {
		t.Error("expected profile Title to be non-empty")
	}
	if p.Location != "Karawang, Indonesia" {
		t.Errorf("expected location Karawang, Indonesia, got %q", p.Location)
	}
	if p.Email != "contact@habibiahmada.dev" {
		t.Errorf("expected email contact@habibiahmada.dev, got %q", p.Email)
	}
	if p.Availability == "" {
		t.Error("expected Availability to be non-empty")
	}
	if len(p.Stats) != 3 {
		t.Errorf("expected 3 stats, got %d", len(p.Stats))
	}
	for _, s := range p.Stats {
		if s.Value == "" || s.Label == "" {
			t.Errorf("stat must have both Value and Label, got %+v", s)
		}
	}
}
