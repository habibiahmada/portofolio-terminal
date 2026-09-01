package data

// GetCertificates returns certifications — live API data when available, else bundled.
func GetCertificates() []Certificate {
	if live, ok := liveCertificatesCopy(); ok {
		return live
	}
	return bundledCertificates()
}

// bundledCertificates is the offline fallback shipped with the binary.
func bundledCertificates() []Certificate {
	return []Certificate{
		{Name: "Best Capstone Project — Coding Camp 2025", Issuer: "Coding Camp", Date: "2025", Pinned: true},
		{Name: "Coding Camp 2025 — Certificate of Completion", Issuer: "Coding Camp", Date: "2025", Pinned: true},
		{Name: "Country Award Winner — Intel AI for Youth", Issuer: "Intel", Date: "2025", Pinned: true},
		{Name: "Alibaba Cloud Certification", Issuer: "Alibaba Cloud", Date: "2024"},
		{Name: "AWS Fargate — Overview", Issuer: "AWS", Date: "2024"},
		{Name: "AWS Cloud Practitioner Essentials", Issuer: "AWS", Date: "2024"},
		{Name: "Cloud Essential Knowledge Badge", Issuer: "AWS", Date: "2024"},
		{Name: "Introduction to Containers", Issuer: "AWS", Date: "2024"},
		{Name: "Belajar Dasar Bootstrap", Issuer: "Codepolitan", Date: "2024"},
		{Name: "Belajar Dasar CSS", Issuer: "Codepolitan", Date: "2024"},
		{Name: "Belajar Dasar HTML", Issuer: "Codepolitan", Date: "2024"},
		{Name: "Git Dasar", Issuer: "Codepolitan", Date: "2024"},
		{Name: "JavaScript Dasar", Issuer: "Codepolitan", Date: "2024"},
		{Name: "Mengenal Dasar Pemrograman Komputer", Issuer: "Codepolitan", Date: "2024"},
		{Name: "SIC 7 — Student Internship Certification", Issuer: "Dibimbing.id", Date: "2024"},
		{Name: "Digital Skill Fair — Frontend Development", Issuer: "Dibimbing.id", Date: "2024"},
		{Name: "Dibimbing.id — Internship Program", Issuer: "Dibimbing.id", Date: "2024"},
		{Name: "Belajar Back-End Pemula dengan JavaScript", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Back-End Pemula dengan JavaScript (2)", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Dasar AI", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Dasar Git dengan GitHub", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Dasar JavaScript", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Dasar Manajemen Proyek", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Fundamental Front-End Web Development", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar KIRo — Dicoding", Issuer: "Dicoding", Date: "2024"},
		{Name: "Belajar Visualisasi Data", Issuer: "Dicoding", Date: "2024"},
		{Name: "Cloud Practitioner Essential", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dasar Pemrograman JavaScript", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dasar Pemrograman JavaScript (2)", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dasar Pemrograman Web", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dasar Pemrograman Web (2)", Issuer: "Dicoding", Date: "2024"},
		{Name: "DevCoach — Web Component", Issuer: "Dicoding", Date: "2024"},
		{Name: "DevCoach — REST API Caching", Issuer: "Dicoding", Date: "2024"},
		{Name: "Financial Literacy 101", Issuer: "Dicoding", Date: "2024"},
		{Name: "From Generalist to Back-End Specialist", Issuer: "Dicoding", Date: "2024"},
		{Name: "Membuat Front-End Web Pemula", Issuer: "Dicoding", Date: "2024"},
		{Name: "Membuat Front-End Web Pemula (2)", Issuer: "Dicoding", Date: "2024"},
		{Name: "Memulai Dasar Pemrograman untuk Jadi Pengembang Software", Issuer: "Dicoding", Date: "2024"},
		{Name: "Memulai Pemrograman dengan Dart", Issuer: "Dicoding", Date: "2024"},
		{Name: "Baparekraf Developer Day 2024", Issuer: "Dicoding", Date: "2024"},
		{Name: "Pengenalan ke Logika Pemrograman (Programming Logic 101)", Issuer: "Dicoding", Date: "2024"},
		{Name: "Fundamental Junior Web Developer", Issuer: "Digitalent", Date: "2024"},
		{Name: "Konsep Pemrograman", Issuer: "Digitalent", Date: "2024"},
		{Name: "Intermediate Junior Web Developer", Issuer: "Digitalent", Date: "2024"},
		{Name: "Google Analytics", Issuer: "Google", Date: "2024"},
		{Name: "Maju Bareng AI — Hacktiv8", Issuer: "Hacktiv8", Date: "2024"},
		{Name: "IBM SkillsBuild — Completion Certificate", Issuer: "IBM", Date: "2024"},
		{Name: "Developing Sites for the Web — IBM", Issuer: "IBM", Date: "2024"},
		{Name: "Fullstack Web Development — Kemnaker", Issuer: "Kemnaker", Date: "2024", Pinned: true},
		{Name: "Internet Introduction — MySkills", Issuer: "MySkills", Date: "2024"},
		{Name: "Pertamina — Technology Program", Issuer: "Pertamina", Date: "2024"},
		{Name: "Web3 HackQuest — Participation", Issuer: "Web3 HackQuest", Date: "2024"},
	}
}

// GetPinnedCertificates returns the featured certificates.
func GetPinnedCertificates() []Certificate {
	var pinned []Certificate
	for _, c := range GetCertificates() {
		if c.Pinned {
			pinned = append(pinned, c)
		}
	}
	return pinned
}
