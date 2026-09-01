package data

import "sync"

// Source identifies where portfolio project data came from.
type Source int

const (
	// SourceBundled is embedded fallback data shipped with the binary.
	SourceBundled Source = iota
	// SourceLive was fetched from the portfolio public API.
	SourceLive
)

var (
	storeMu            sync.RWMutex
	liveProjects       []Project
	liveCaseStudies    map[string]CaseStudy
	liveCertificates   []Certificate
	liveSource         Source = SourceBundled
)

// DataSource returns the active portfolio data source.
func DataSource() Source {
	storeMu.RLock()
	defer storeMu.RUnlock()
	return liveSource
}

// SetLiveData replaces in-memory project, case study, and certificate data from the API.
func SetLiveData(projects []Project, caseStudies []CaseStudy, certificates []Certificate) {
	storeMu.Lock()
	defer storeMu.Unlock()

	liveProjects = make([]Project, len(projects))
	copy(liveProjects, projects)

	liveCaseStudies = make(map[string]CaseStudy, len(caseStudies))
	for _, cs := range caseStudies {
		liveCaseStudies[cs.Slug] = cs
	}

	if len(certificates) > 0 {
		liveCertificates = make([]Certificate, len(certificates))
		copy(liveCertificates, certificates)
	}
	liveSource = SourceLive
}

// SetLivePortfolio is an alias for SetLiveData without certificates.
func SetLivePortfolio(projects []Project, caseStudies []CaseStudy) {
	SetLiveData(projects, caseStudies, nil)
}

func liveProjectsCopy() ([]Project, bool) {
	storeMu.RLock()
	defer storeMu.RUnlock()
	if liveSource != SourceLive || len(liveProjects) == 0 {
		return nil, false
	}
	out := make([]Project, len(liveProjects))
	copy(out, liveProjects)
	return out, true
}

func liveCaseStudy(slug string) (*CaseStudy, bool) {
	storeMu.RLock()
	defer storeMu.RUnlock()
	if liveSource != SourceLive {
		return nil, false
	}
	cs, ok := liveCaseStudies[slug]
	if !ok {
		return nil, false
	}
	copy := cs
	return &copy, true
}

func liveCertificatesCopy() ([]Certificate, bool) {
	storeMu.RLock()
	defer storeMu.RUnlock()
	if liveSource != SourceLive || len(liveCertificates) == 0 {
		return nil, false
	}
	out := make([]Certificate, len(liveCertificates))
	copy(out, liveCertificates)
	return out, true
}
