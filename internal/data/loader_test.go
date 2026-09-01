package data

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchPortfolio(t *testing.T) {
	projectsPayload := []apiProject{
		{
			Slug:        "e-vote",
			Name:        "E-Vote",
			Year:        "2025",
			Description: "Digital elections",
			Tags:        []string{"Laravel"},
			Featured:    true,
		},
	}
	caseStudiesPayload := []apiCaseStudy{
		{
			Slug: "e-vote",
			Hero: "Students needed digital elections.",
			Sections: []apiCaseStudySection{
				{Label: "Where it started", Body: "Paper ballots were slow."},
			},
		},
	}
	certsPayload := []apiCertificate{
		{
			Title:    "AWS Cloud Practitioner",
			Org:      "AWS",
			IsPinned: true,
			Thumb:    "https://example.com/cert.webp",
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc(terminalProjectsPath, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"data":    projectsPayload,
			"error":   nil,
		})
	})
	mux.HandleFunc(caseStudiesPath, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"data":    caseStudiesPayload,
			"error":   nil,
		})
	})
	mux.HandleFunc("/api/public/certificates", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{
			"success": true,
			"data":    certsPayload,
			"error":   nil,
		})
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	t.Setenv("PORTFOLIO_API_URL", srv.URL)

	projects, caseStudies, certificates, err := FetchPortfolio()
	if err != nil {
		t.Fatalf("FetchPortfolio: %v", err)
	}
	if len(projects) != 1 || projects[0].Slug != "e-vote" {
		t.Fatalf("unexpected projects: %+v", projects)
	}
	if len(caseStudies) != 1 || caseStudies[0].Slug != "e-vote" {
		t.Fatalf("unexpected case studies: %+v", caseStudies)
	}
	if len(certificates) != 1 || certificates[0].Name != "AWS Cloud Practitioner" {
		t.Fatalf("unexpected certificates: %+v", certificates)
	}
}

func TestGetCaseStudyLiveStore(t *testing.T) {
	t.Cleanup(func() {
		storeMu.Lock()
		liveProjects = nil
		liveCaseStudies = nil
		liveCertificates = nil
		liveSource = SourceBundled
		storeMu.Unlock()
	})

	SetLiveData(
		[]Project{{Slug: "live-slug", Name: "Live", Year: "2026"}},
		[]CaseStudy{{Slug: "live-slug", Hero: "Live hero"}},
		[]Certificate{{Name: "Live Cert", Issuer: "Org", Date: "2025"}},
	)

	cs := GetCaseStudy("live-slug")
	if cs == nil || cs.Hero != "Live hero" {
		t.Fatalf("expected live case study, got %+v", cs)
	}
	certs := GetCertificates()
	if len(certs) != 1 || certs[0].Name != "Live Cert" {
		t.Fatalf("expected live certificates, got %+v", certs)
	}
	if DataSource() != SourceLive {
		t.Fatalf("expected SourceLive, got %v", DataSource())
	}
}
