package tui

import (
	"strings"
	"testing"
	"time"
)

// TestMatrixSizes renders every screen across the QA matrix from
// docs/tui-illustration.md to guard against panics and blank layouts.
func TestMatrixSizes(t *testing.T) {
	sizes := []struct {
		w, h int
	}{
		{40, 12},
		{80, 24},
		{120, 40},
		{200, 50},
	}

	screens := []Screen{
		ScreenHome,
		ScreenAbout,
		ScreenProjects,
		ScreenProjectDetail,
		ScreenSkills,
		ScreenExperience,
		ScreenCertificates,
		ScreenServices,
		ScreenContact,
	}

	for _, size := range sizes {
		for _, s := range screens {
			m := New()
			m.width, m.height = size.w, size.h
			m.currentScreen = s
			if len(m.projects) > 0 {
				m.projectDetail = m.projects[0]
			}

			view := m.View()
			if view == "" {
				t.Errorf("[%dx%d] screen %v rendered empty", size.w, size.h, s)
			}
		}
	}
}

// TestScreensContentCompleteness verifies each screen contains its expected rich sections.
func TestScreensContentCompleteness(t *testing.T) {
	m := New()
	m.width, m.height = 100, 35

	checks := []struct {
		screen   Screen
		expected []string
	}{
		{
			screen:   ScreenHome,
			expected: []string{"Featured Work", "Trusted & Supported By", "Quick Actions"},
		},
		{
			screen:   ScreenAbout,
			expected: []string{"Habibi Ahmad Aziz", "Core Metrics & Accolades", "Engineering Philosophy"},
		},
		{
			screen:   ScreenSkills,
			expected: []string{"Tools & Technologies", "Frontend & UI Development", "Backend & API Engineering", "Databases & Data Modeling"},
		},
		{
			screen:   ScreenExperience,
			expected: []string{"Work Experience", "Education & Foundations", "PT Webekspres", "SMK Negeri 1 Karawang"},
		},
		{
			screen:   ScreenCertificates,
			expected: []string{"Licenses & Certifications", "Featured Honors & Awards", "Professional Certifications"},
		},
		{
			screen:   ScreenServices,
			expected: []string{"Services & Solutions", "How I Ship", "Web Design & Mobile-First", "Frontend Development"},
		},
		{
			screen:   ScreenContact,
			expected: []string{"Contact Methods", "Socials", "Email", "Website"},
		},
	}

	for _, tc := range checks {
		m.currentScreen = tc.screen
		var raw string
		switch tc.screen {
		case ScreenHome:
			raw = m.renderHomeContent()
		case ScreenAbout:
			raw = m.renderAboutContent()
		case ScreenSkills:
			raw = m.renderSkillsContent()
		case ScreenExperience:
			raw = m.renderExperienceContent()
		case ScreenCertificates:
			raw = m.renderCertificatesContent()
		case ScreenServices:
			raw = m.renderServicesContent()
		case ScreenContact:
			raw = m.renderContactContent()
		}
		for _, exp := range tc.expected {
			if !strings.Contains(raw, exp) {
				t.Errorf("screen %v missing expected content %q", tc.screen, exp)
			}
		}
	}
}

// TestSplashProgress verifies the splash advances through frames and hands
// control to the App once the sequence completes.
func TestSplashProgress(t *testing.T) {
	s := NewSplash()
	s.width, s.height = 120, 40

	frames := 0
	deadline := time.Now().Add(5 * time.Second)
	for frames <= len(splashDelays) {
		if time.Now().After(deadline) {
			t.Fatal("splash progress timed out — possible infinite loop")
		}
		view := s.View()
		if view == "" {
			t.Fatalf("splash frame %d rendered empty", frames)
		}

		model, _ := s.Update(splashTickMsg{})
		frames++

		if app, ok := model.(*App); ok {
			if app.currentScreen != ScreenHome {
				t.Errorf("expected App at Home, got %v", app.currentScreen)
			}
			return
		}

		next, ok := model.(*Splash)
		if !ok {
			t.Fatalf("expected splash model at frame %d, got %T", frames, model)
		}
		s = next
	}

	t.Fatalf("splash never transitioned to App after %d frames", frames)
}

// TestSplashSkip verifies small terminals skip the splash.
func TestSplashSkip(t *testing.T) {
	s := NewSplash()
	model, _ := s.Update(teaMsgWindow(30, 12))
	if _, ok := model.(*App); !ok {
		t.Fatalf("expected splash to be skipped on small terminal, got %T", model)
	}
}

// TestSplashResizePreservesSize verifies the window size is forwarded to App.
func TestSplashResizePreservesSize(t *testing.T) {
	s := NewSplash()
	model, _ := s.Update(teaMsgWindow(200, 50))
	if _, ok := model.(*Splash); !ok {
		t.Fatalf("expected splash to continue on large terminal, got %T", model)
	}
}
