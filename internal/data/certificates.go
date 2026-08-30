package data

// GetCertificates returns 52 certifications with 3 featured/pinned entries per
// docs/pages.md. The three pinned entries are the highlighted ones.
func GetCertificates() []Certificate {
	return []Certificate{
		{Name: "Best Capstone Project — Coding Camp 2025", Issuer: "DBS Foundation", Date: "2025", Pinned: true},
		{Name: "Coding Camp 2025 — Certificate of Completion", Issuer: "Coding Camp", Date: "2025", Pinned: true},
		{Name: "Country Award Winner — Intel AI for Youth", Issuer: "Intel", Date: "2025", Pinned: true},

		{Name: "Dicoding: Belajar Dasar Pemrograman Web", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Membangun Front-End Web untuk Pemula", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Fundamental Front-End Web Development", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Membuat Aplikasi Web dengan React", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Fundamental Aplikasi Web dengan React", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Menjadi Front-End Web Developer Expert", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Dasar JavaScript", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Membuat Front-End Web untuk Pemula", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Membuat Aplikasi Back-End untuk Pemula", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Fundamental Back-End", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Menjadi Back-End Developer Expert", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Cloud Practitioner Essentials", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Architecting on AWS", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Dasar Structured Query Language (SQL)", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Dasar Git dengan GitHub", Issuer: "Dicoding", Date: "2024"},
		{Name: "Dicoding: Belajar Dasar Pemrograman JavaScript", Issuer: "Dicoding", Date: "2024"},
		{Name: "AWS: AWS Academy Cloud Foundations", Issuer: "Amazon Web Services", Date: "2024"},
		{Name: "AWS: AWS Cloud Practitioner Essentials", Issuer: "Amazon Web Services", Date: "2024"},
		{Name: "AWS: AWS Partner: Sales Accreditation", Issuer: "Amazon Web Services", Date: "2024"},
		{Name: "AWS: AWS Partner: Economics Accreditation", Issuer: "Amazon Web Services", Date: "2024"},
		{Name: "AWS: AWS Technical Essentials", Issuer: "Amazon Web Services", Date: "2024"},
		{Name: "Kominfo: Junior Web Developer", Issuer: "Kementerian Kominfo", Date: "2024"},
		{Name: "Progate: HTML & CSS Course", Issuer: "Progate", Date: "2023"},
		{Name: "Progate: JavaScript Course", Issuer: "Progate", Date: "2023"},
		{Name: "Progate: React Course", Issuer: "Progate", Date: "2023"},
		{Name: "Progate: Ruby Course", Issuer: "Progate", Date: "2023"},
		{Name: "Progate: Git Course", Issuer: "Progate", Date: "2023"},
		{Name: "Coursera: Crash Course on Python", Issuer: "Google", Date: "2023"},
		{Name: "Coursera: Foundations of Cybersecurity", Issuer: "Google", Date: "2023"},
		{Name: "Coursera: System Administration and IT Infra Services", Issuer: "Google", Date: "2023"},
		{Name: "Coursera: HTML, CSS, and Javascript for Web Developers", Issuer: "Johns Hopkins", Date: "2023"},
		{Name: "Coursera: Meta Front-End Developer", Issuer: "Meta", Date: "2024"},
		{Name: "Coursera: Google UX Design", Issuer: "Google", Date: "2023"},
		{Name: "HackerRank: JavaScript (Basic)", Issuer: "HackerRank", Date: "2024"},
		{Name: "HackerRank: Problem Solving (Basic)", Issuer: "HackerRank", Date: "2024"},
		{Name: "HackerRank: Python (Basic)", Issuer: "HackerRank", Date: "2024"},
		{Name: "HackerRank: SQL (Basic)", Issuer: "HackerRank", Date: "2024"},
		{Name: "HackerRank: React (Basic)", Issuer: "HackerRank", Date: "2024"},
		{Name: "freeCodeCamp: Responsive Web Design", Issuer: "freeCodeCamp", Date: "2023"},
		{Name: "freeCodeCamp: JavaScript Algorithms and Data Structures", Issuer: "freeCodeCamp", Date: "2023"},
		{Name: "freeCodeCamp: Front End Development Libraries", Issuer: "freeCodeCamp", Date: "2023"},

		{Name: "SoloLearn: Web Development with PHP", Issuer: "SoloLearn", Date: "2023"},
		{Name: "SoloLearn: Introduction to Java", Issuer: "SoloLearn", Date: "2023"},
		{Name: "Buildwith Angga: Front End Web Development", Issuer: "Buildwith Angga", Date: "2023"},
		{Name: "Skilvul: Front End Web Development", Issuer: "Skilvul", Date: "2023"},
		{Name: "Udemy: The Complete 2024 Web Development Bootcamp", Issuer: "Udemy", Date: "2024"},
		{Name: "IDCamp: Cloud Developer", Issuer: "IDCamp", Date: "2024"},
		{Name: "Oracle: Oracle Cloud Infrastructure Foundations", Issuer: "Oracle", Date: "2024"},
		{Name: "BNSP: Junior Web Developer", Issuer: "BNSP", Date: "2024"},
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
