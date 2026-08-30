package data

import "testing"

func TestGetCertificates(t *testing.T) {
	certificates := GetCertificates()
	if len(certificates) == 0 {
		t.Fatal("expected at least one certificate")
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
