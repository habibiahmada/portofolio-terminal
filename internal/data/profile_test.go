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
	if p.Location == "" {
		t.Error("expected profile Location to be non-empty")
	}
	if p.Email == "" {
		t.Error("expected profile Email to be non-empty")
	}
}
