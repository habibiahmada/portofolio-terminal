package data

import (
	"strings"
	"testing"
)

func TestGetSocials(t *testing.T) {
	socials := GetSocials()
	if len(socials) == 0 {
		t.Fatal("expected at least one social link")
	}

	for i, s := range socials {
		if s.Name == "" {
			t.Errorf("social[%d]: Name must not be empty", i)
		}
		if !strings.HasPrefix(s.URL, "https://") {
			t.Errorf("social %q: URL %q must start with https://", s.Name, s.URL)
		}
	}
}
