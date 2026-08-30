package data

import "testing"

func TestGetCertificates(t *testing.T) {
	certificates := GetCertificates()
	if len(certificates) != 52 {
		t.Fatalf("expected 52 certificates, got %d", len(certificates))
	}

	for i, c := range certificates {
		if c.Name == "" {
			t.Errorf("certificate[%d]: Name must not be empty", i)
		}
		if c.Issuer == "" {
			t.Errorf("certificate[%d]: Issuer must not be empty", i)
		}
		if c.Date == "" {
			t.Errorf("certificate[%d]: Date must not be empty", i)
		}
	}
}

func TestGetPinnedCertificates(t *testing.T) {
	pinned := GetPinnedCertificates()
	if len(pinned) != 3 {
		t.Fatalf("expected 3 pinned certificates, got %d", len(pinned))
	}

	for _, c := range pinned {
		if !c.Pinned {
			t.Errorf("certificate %q should be flagged Pinned", c.Name)
		}
	}
}
