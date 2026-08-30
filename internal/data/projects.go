package data

// GetProjects returns all portfolio projects synced with docs/pages.md. Order
// matches the website archive (newest first) so prev/next navigation flows
// naturally through the list.
func GetProjects() []Project {
	return []Project{
		{
			Name:          "Aksara Pustaka",
			Slug:          "aksara-pustaka",
			Year:          "2026",
			Description:   "I built a web library system to manage books, members, loans, returns, stock, and history in one place.",
			DescriptionID: "Saya membangun sistem perpustakaan web untuk mengelola buku, anggota, peminjaman, pengembalian, stok, dan riwayat dalam satu tempat.",
			Tags:          []string{"Laravel", "MySQL", "PHP", "Tailwind CSS"},
			Stack:         []string{"Laravel", "MySQL", "PHP", "Tailwind CSS"},
			Featured:      false,
		},
		{
			Name:          "SiPadu",
			Slug:          "sipadu",
			Year:          "2026",
			Description:   "I built SiPadu so school facility reports can be filed and tracked in real time instead of disappearing into paper trails.",
			DescriptionID: "Saya membangun SiPadu agar laporan fasilitas sekolah bisa diajukan dan dilacak real-time, bukan hilang di jejak kertas.",
			Tags:          []string{"Laravel", "Tailwind CSS", "MySQL", "PHP"},
			Stack:         []string{"Laravel", "Tailwind CSS", "MySQL", "PHP"},
			Featured:      false,
		},
		{
			Name:          "ParkingApp",
			Slug:          "parking-app",
			Year:          "2026",
			Description:   "I built a parking web app for check-in, check-out, reporting, and audit in one flow.",
			DescriptionID: "Saya membangun aplikasi parkir web untuk check-in, check-out, pelaporan, dan audit dalam satu alur.",
			Tags:          []string{"PHP"},
			Stack:         []string{"PHP"},
			Featured:      false,
		},
		{
			Name:          "Inventoryflow",
			Slug:          "inventoryflow",
			Year:          "2026",
			Description:   "I built Inventoryflow to handle school/lab equipment loans, inventory, approvals, and returns, without spreadsheet chaos.",
			DescriptionID: "Saya membangun Inventoryflow untuk peminjaman alat sekolah/lab, inventaris, persetujuan, dan pengembalian, tanpa kekacauan spreadsheet.",
			Tags:          []string{"Laravel", "Tailwind CSS", "PHP", "MySQL"},
			Stack:         []string{"Laravel", "Tailwind CSS", "PHP", "MySQL"},
			Featured:      false,
		},
		{
			Name:          "BagiBerkah",
			Slug:          "bagiberkah",
			Year:          "2026",
			Description:   "I built BagiBerkah, a digital THR experience with mini-games and smarter allocation recommendations.",
			DescriptionID: "Saya membangun BagiBerkah, pengalaman THR digital dengan mini-game dan rekomendasi alokasi yang lebih cerdas.",
			Tags:          []string{"Next.js", "Express", "Prisma", "Mayar", "Xendit"},
			Stack:         []string{"Next.js", "Express", "Prisma", "Mayar", "Xendit"},
			Featured:      false,
		},
		{
			Name:          "E-Democracy (E-Vote)",
			Slug:          "e-vote",
			Year:          "2025",
			Description:   "I built E-Vote so students can run digital OSIS elections with confidence, from ballots to real-time results.",
			DescriptionID: "Saya membangun E-Vote agar siswa bisa menjalankan pemilihan OSIS digital dengan percaya diri, dari surat suara hingga hasil real-time.",
			Tags:          []string{"PHP", "Laravel", "Bootstrap", "MySQL"},
			Stack:         []string{"PHP", "Laravel", "Bootstrap", "MySQL"},
			Featured:      true,
		},
		{
			Name:          "Smartfarm AI (Agrify)",
			Slug:          "agrify",
			Year:          "2025",
			Description:   "On a team project, I helped ship Smartfarm AI, combining modern web UI with ML-assisted insights for Indonesian farmers.",
			DescriptionID: "Dalam proyek tim, saya membantu menghadirkan Smartfarm AI, menggabungkan UI web modern dengan insight berbantuan ML untuk petani Indonesia.",
			Tags:          []string{"React", "Python", "Machine Learning", "JavaScript"},
			Stack:         []string{"React", "Python", "Machine Learning", "JavaScript"},
			Featured:      true,
		},
		{
			Name:          "CultureConnect.",
			Slug:          "culture-connect",
			Year:          "2025",
			Description:   "With a distributed team, I helped build CultureConnect, an AI platform for more personal, community-aware cultural travel.",
			DescriptionID: "Bersama tim terdistribusi, saya membantu membangun CultureConnect, platform AI untuk wisata budaya yang lebih personal dan berdampak bagi komunitas.",
			Tags:          []string{"Python", "React", "Machine Learning", "Express", "Prisma"},
			Stack:         []string{"Python", "React", "Machine Learning", "Express", "Prisma"},
			Featured:      true,
		},
		{
			Name:          "Spacelab",
			Slug:          "spacelab",
			Year:          "2025",
			Description:   "I built Spacelab to keep school schedules, rooms, and teachers conflict-free without spreadsheet chaos.",
			DescriptionID: "Saya membangun Spacelab agar jadwal, ruangan, dan guru sekolah bebas konflik tanpa kekacauan spreadsheet.",
			Tags:          []string{"Laravel", "PHP", "JavaScript"},
			Stack:         []string{"Laravel", "PHP", "JavaScript"},
			Featured:      true,
		},
		{
			Name:          "Renshuu web",
			Slug:          "renshuu",
			Year:          "2025",
			Description:   "During PKL at CV Smartplus, my team and I built Renshuu, a job-search web app assigned for SMKN 1 Karawang students.",
			DescriptionID: "Saat PKL di CV Smartplus, saya dan tim mengerjakan Renshuu, aplikasi pencarian kerja yang ditugaskan untuk siswa SMKN 1 Karawang.",
			Tags:          []string{"React", "Laravel"},
			Stack:         []string{"React", "Laravel"},
			Featured:      true,
		},
	}
}

// GetFeaturedProjects returns only the projects marked as featured, in the
// order they appear on the website home section.
func GetFeaturedProjects() []Project {
	ordered := []string{"e-vote", "agrify", "culture-connect", "spacelab", "renshuu"}
	bySlug := map[string]Project{}
	for _, p := range GetProjects() {
		bySlug[p.Slug] = p
	}
	featured := make([]Project, 0, len(ordered))
	for _, slug := range ordered {
		p, ok := bySlug[slug]
		if ok {
			featured = append(featured, p)
		}
	}
	return featured
}
