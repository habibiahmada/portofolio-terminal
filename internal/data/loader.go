package data

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	defaultAPIBase        = "https://www.habibiahmada.dev"
	fetchTimeout          = 8 * time.Second
	terminalProjectsPath  = "/api/public/terminal/projects"
	caseStudiesPath       = "/api/public/case-studies"
	certificatesPath      = "/api/public/certificates?page_size=100"
)

type apiEnvelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type apiProject struct {
	Slug          string   `json:"slug"`
	Name          string   `json:"name"`
	Year          string   `json:"year"`
	Description   string   `json:"description"`
	DescriptionID string   `json:"description_id"`
	Tags          []string `json:"tags"`
	Live          string   `json:"live"`
	Featured      bool     `json:"featured"`
}

type apiCaseStudy struct {
	Slug     string                `json:"slug"`
	Hero     string                `json:"hero"`
	Sections []apiCaseStudySection `json:"sections"`
}

type apiCaseStudySection struct {
	Label string `json:"label"`
	Body  string `json:"body"`
}

type apiCertificate struct {
	Title     string   `json:"title"`
	Org       string   `json:"org"`
	Pages     []string `json:"pages"`
	Thumb     string   `json:"thumb"`
	IsPinned  bool     `json:"is_pinned"`
}

var yearInTitle = regexp.MustCompile(`20\d{2}`)

// APIBase returns the portfolio site origin (override with PORTFOLIO_API_URL).
func APIBase() string {
	if v := strings.TrimSpace(os.Getenv("PORTFOLIO_API_URL")); v != "" {
		return strings.TrimRight(v, "/")
	}
	return defaultAPIBase
}

// FetchPortfolio loads projects, case studies, and certificates in parallel.
func FetchPortfolio() ([]Project, []CaseStudy, []Certificate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), fetchTimeout)
	defer cancel()

	var (
		wg          sync.WaitGroup
		mu          sync.Mutex
		projects    []Project
		caseStudies []CaseStudy
		certificates []Certificate
		fetchErr    error
	)

	setErr := func(err error) {
		mu.Lock()
		if fetchErr == nil {
			fetchErr = err
		}
		mu.Unlock()
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		var rows []apiProject
		if err := getJSON(ctx, terminalProjectsPath, &rows); err != nil {
			setErr(fmt.Errorf("projects: %w", err))
			return
		}
		mapped := mapAPIProjects(rows)
		if len(mapped) == 0 {
			setErr(fmt.Errorf("projects: empty response"))
			return
		}
		mu.Lock()
		projects = mapped
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		var rows []apiCaseStudy
		if err := getJSON(ctx, caseStudiesPath, &rows); err != nil {
			setErr(fmt.Errorf("case studies: %w", err))
			return
		}
		mu.Lock()
		caseStudies = mapAPICaseStudies(rows)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		var rows []apiCertificate
		if err := getJSON(ctx, certificatesPath, &rows); err != nil {
			setErr(fmt.Errorf("certificates: %w", err))
			return
		}
		mu.Lock()
		certificates = mapAPICertificates(rows)
		mu.Unlock()
	}()

	wg.Wait()

	if fetchErr != nil {
		return nil, nil, nil, fetchErr
	}
	return projects, caseStudies, certificates, nil
}

func getJSON(ctx context.Context, path string, dest any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, APIBase()+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "habibiahmada-terminal/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d", resp.StatusCode)
	}

	var envelope apiEnvelope
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return err
	}
	if !envelope.Success || len(envelope.Data) == 0 {
		return fmt.Errorf("unsuccessful response")
	}
	return json.Unmarshal(envelope.Data, dest)
}

func mapAPIProjects(rows []apiProject) []Project {
	out := make([]Project, 0, len(rows))
	for _, row := range rows {
		if row.Slug == "" || row.Name == "" {
			continue
		}
		tags := append([]string(nil), row.Tags...)
		out = append(out, Project{
			Name:          row.Name,
			Slug:          row.Slug,
			Year:          row.Year,
			Description:   row.Description,
			DescriptionID: row.DescriptionID,
			Tags:          tags,
			Stack:         tags,
			Live:          row.Live,
			Featured:      row.Featured,
		})
	}
	return out
}

func mapAPICaseStudies(rows []apiCaseStudy) []CaseStudy {
	out := make([]CaseStudy, 0, len(rows))
	for _, row := range rows {
		if row.Slug == "" {
			continue
		}
		sections := make([]CaseStudySection, 0, len(row.Sections))
		for _, sec := range row.Sections {
			if sec.Label == "" && sec.Body == "" {
				continue
			}
			sections = append(sections, CaseStudySection{
				Label: sec.Label,
				Body:  sec.Body,
			})
		}
		out = append(out, CaseStudy{
			Slug:     row.Slug,
			Hero:     row.Hero,
			Sections: sections,
		})
	}
	return out
}

func mapAPICertificates(rows []apiCertificate) []Certificate {
	out := make([]Certificate, 0, len(rows))
	for _, row := range rows {
		if row.Title == "" {
			continue
		}
		url := row.Thumb
		if url == "" && len(row.Pages) > 0 {
			url = row.Pages[0]
		}
		out = append(out, Certificate{
			Name:   row.Title,
			Issuer: row.Org,
			Date:   extractCertYear(row.Title),
			URL:    url,
			Pinned: row.IsPinned,
		})
	}
	return out
}

func extractCertYear(title string) string {
	if match := yearInTitle.FindString(title); match != "" {
		return match
	}
	return ""
}
