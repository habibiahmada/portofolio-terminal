# Halaman & Konten TUI

Dokumentasi halaman **habibiahmada-terminal** yang diselaraskan dengan website portfolio. Semua teks intro, tagline, bio, dan data di bawah ini adalah **sumber kebenaran konten** — implementasi TUI harus mengikuti copy ini.

**Bahasa UI:** English (sama dengan halaman produksi website).  
**Data bilingual:** Deskripsi proyek tersedia EN + ID; TUI v1 menampilkan EN, ID disimpan di data layer untuk fase berikutnya.

---

## Peta Navigasi

### Website (referensi struktur)

```
Header/Footer: Home · About · Work · Blog · Services
CTA global: "Let's Talk" → contact@habibiahmada.dev
```

### Terminal (target)

```
Sidebar: Home · About · Projects · Skills · Experience · Certificates · Services · Contact
Home sections (scroll): Hero · Companies · Featured Projects · Services · Press · Process · CTA
```

| Website | Terminal screen | Catatan |
|---------|-----------------|---------|
| `/` Home | `ScreenHome` | Hero + section scroll (bukan menu-only) |
| `/about` | `ScreenAbout` | Bio, stats, tech stack preview |
| `/about` #about-timeline | `ScreenExperience` | Timeline + education |
| `/about` #about-techstack | `ScreenSkills` | 16 tools marquee → list/grid |
| `/projects` | `ScreenProjects` | 10 proyek |
| `/projects/[slug]` | `ScreenProjectDetail` | Case study 4 bagian |
| `/services` | `ScreenServices` | 5 kartu layanan |
| CTA / Contact | `ScreenContact` | Email + social + availability |
| — | `ScreenPress` | 2 spotlight (embedded di Home atau screen terpisah) |
| — | `ScreenProcess` | 4 langkah (embedded di Home/Services) |
| — | `ScreenCompanies` | Marquee partner (embedded di Home/About) |

**Tidak ada screen:** Admin, Login, Blog Preview, locale stub (`/en/about`, `/id/contact`).

---

## Profil & Identitas

| Field | Nilai |
|-------|-------|
| Nama lengkap | Habibi Ahmad Aziz |
| Nama pendek | Habibi Ahmad |
| Role | Fullstack Developer |
| Title (TUI header) | Fullstack Developer |
| Lokasi | Karawang, Indonesia |
| Email | contact@habibiahmada.dev |
| Availability | Open to freelance & full-time · Remote (WIB) |
| Website | habibiahmada.dev |
| GitHub | https://github.com/habibiahmada |
| LinkedIn | https://www.linkedin.com/in/habibi-ahmad-aziz |
| Instagram | https://instagram.com/habibiahmad.a |
| Employer | PT Webekspres Teknologi Indonesia |
| School | SMKN 1 Karawang (Software Engineering) |

### Stats (About hero)

| Stat | Label |
|------|-------|
| 3+ | Years Building |
| 10+ | Projects Shipped |
| 2 | Awards Won |

---

## Screen: Home

**Setara dengan:** `/` — semua section fold.

### Hero (`#hero`)

| Elemen | Copy |
|--------|------|
| Badge | `Open to freelance & full-time · Remote (WIB)` |
| H1 | `Building digital experiences that actually matter` |
| Accent words | *experiences*, *matter* |
| Subtitle | Frontend-leaning full-stack developer. I craft clear interfaces and the APIs behind them, with a bias for performance you can measure, not just claim. |
| CTA primary | Let's Talk |
| CTA secondary | View My Work |
| CTA tertiary | View CV |

**TUI mapping:**
- Badge → bar availability di bawah header atau di hero
- H1 → FIGlet/signature + subtitle Lip Gloss
- CTAs → hint keyboard: `Enter` Explore · `P` Projects · `C` Contact · `V` CV (modal/teks)

### Companies (`#companies`)

| Elemen | Copy |
|--------|------|
| Heading | Collaborations & Trusted By |

**Partner:** Neskar · PPLG · Sagasitas · Smartplus · Webekspres

### Featured Projects (`#projects`)

**Featured (urutan home):** E-Vote · Smartfarm AI (Agrify) · CultureConnect · Spacelab · Renshuu

| Elemen | Copy |
|--------|------|
| Section label | Featured Work |
| Link | All Projects → navigasi ke ScreenProjects |

### Services preview (`#services`)

5 kartu ringkas (lihat Screen Services). Di Home: tampilkan judul + 1 baris deskripsi per kartu.

### Press / Spotlights (`#press`)

**1. Dicoding Blog**
- Title: From zero tech background to shipping under real deadlines.
- Body: Coding Camp did not hand me confidence. It forced me to build, present, and intern at the same time. That pressure is still how I work: clear scope, thin vertical slice, ship.
- CTA label: Read the Dicoding story

**2. Intel AI Festival**
- Title: Indonesia country award for Agrify at a global AI festival.
- Body: Our team built Agrify for AI Changemakers. I owned the product surface: turn model output into advice a farmer can act on, not another pretty dashboard.
- CTA label: See the Intel winners list

### Process (`#process`)

| Elemen | Copy |
|--------|------|
| Label | How I ship |
| H2 | Process over performance theater |
| Sub | A short operating manual, the same sequence I use on school systems, internships, and competition deadlines. |

| # | Title | Description |
|---|-------|-------------|
| 1 | Scope the real job | I write down the user, the constraint, and the non-goals before opening an editor. Fancy stacks wait until the problem is boringly clear. |
| 2 | Ship a thin vertical slice | One path that works end to end (auth, data, UI) beats a polished shell. Preview deploys keep feedback honest. |
| 3 | Harden what users can break | Validation at trust boundaries, fail-closed writes, and empty/error states you can explain to a non-engineer. |
| 4 | Measure, then decorate | Motion and polish come after the page is readable and fast enough on a mid-range phone. If it is not measurable, it is not a claim. |

### CTA (`#cta`)

| Elemen | Copy |
|--------|------|
| Label | Contact |
| H2 | Need a web product that actually ships in the next 90 days? |
| Body | Open to freelance and full-time. Remote (WIB). Write to contact@habibiahmada.dev. I usually reply within 48 hours. |
| CTA | Let's talk · Browse projects |

### Footer (TUI)

| Elemen | Copy |
|--------|------|
| Brand | habibiahmada. |
| Tagline | Frontend-leaning full-stack developer. I ship web products end to end, from UI systems to APIs, with a bias for clarity and measurable performance. |
| Copyright | © {year} Habibi Ahmad Aziz. All rights reserved. |

---

## Screen: About

**Setara dengan:** `/about`

### About Hero

| Elemen | Copy |
|--------|------|
| H1 line 1 | Habibi (accent merah) |
| H1 line 2 | Ahmad Aziz |
| One-liner | Full-Stack Web Developer experienced in building responsive apps and CMS products. Skilled in crafting end-to-end features using Next.js, React, Laravel, Node.js, and WordPress to deliver production-ready solutions. |
| CTA 1 | Let's Collaborate |
| CTA 2 | View Experience |

### About Intro (`#about-intro`)

| Elemen | Copy |
|--------|------|
| Label | // About |
| H2 | A Glimpse Into / Who I Am |
| Para 1 | As a Software Engineering graduate from SMKN 1 Karawang, I currently work as a Web Developer at PT Webekspres Teknologi Indonesia. My expertise lies in developing tailored client websites, architecting CMS platforms, and deploying scalable full-stack features for production environments. |
| Para 2 | Driven by a deep passion for software architecture and modern web technologies, I thrive on solving complex technical challenges. My development philosophy centers on writing clean, scalable code and crafting intuitive digital experiences that deliver tangible impact. |
| CTA | View my journey |

**TUI:** Side-by-side dengan micro-illustration terminal (`AboutTerminal()`).

---

## Screen: Skills

**Setara dengan:** `/about` #about-techstack

| Elemen | Copy |
|--------|------|
| Label | Tech Stack |
| H2 | Tools & Technologies |
| Sub | The technologies I use daily to turn ideas into functional, high-performing digital reality. |

**16 tools (urutan marquee):** React · Next.js · Node.js · TypeScript · PostgreSQL · Tailwind CSS · PHP · Laravel · WordPress · Elementor · Astra · Git · GitHub · Bootstrap · Vercel · JavaScript

**TUI:** Grid/list dengan kategori opsional; bar level hanya jika ada data level di website (saat ini tidak — tampilkan flat list).

---

## Screen: Experience

**Setara dengan:** `/about` #about-timeline

| Elemen | Copy |
|--------|------|
| Label | Experience |
| H2 | Path so far |
| Sub | Roles and programs that taught me to scope, ship, and explain the trade-offs. |

### Work

| Period | Title | Company | Location | Badge | Description |
|--------|-------|---------|----------|-------|-------------|
| May 2026 – Now | Web Developer | PT Webekspres Technology Indonesia | Karawang · On site | Current | Client and internal web work on WordPress, CMS platforms, and modern stacks. I own features end to end, from brief to deploy. |
| Jun – Aug 2025 | Cloud Computing Trainer Intern | Yayasan Sagasitas Indonesia | Jakarta · On site | — | Taught Cloud Computing and Generative AI in schools. Built AWS PartyRock labs and kept teaching teams aligned with partner schools. |
| Jan – Apr 2025 | Student Member | Coding Camp powered by DBS Foundation | Bandung · Remote | Top 15 Capstone | Full-stack track under real deadlines. CultureConnect landed in the Top 15 Best Capstone Projects. |
| Jan – May 2025 | Web Developer Intern | CV. SmartPlus Indonesia | Karawang · Remote | — | Full-stack intern on internal company web projects. Shipped features end to end with modern stacks. |

### Education — Foundations

| Period | Title | School | Description |
|--------|-------|--------|-------------|
| 2023 – 2026 | Software Engineering | SMK Negeri 1 Karawang | Software development, programming, systems, and networking. Active in tech projects and competitions. |
| 2020 – 2023 | Arabic Language and Literature | MTSS Darunnadwah 01 | Arabic language and literature with a focus on communication and text analysis. |

---

## Screen: Projects

**Setara dengan:** `/projects`

| Elemen | Copy |
|--------|------|
| H1 | All Projects |
| Sub (EN) | Web projects by Habibi Ahmad Aziz. Production and capstone work across school systems, AI products, payments, and fullstack apps. Each project links to a detailed case study. |
| Sub (ID) | Proyek web oleh Habibi Ahmad Aziz. Karya produksi dan capstone: sistem sekolah, produk AI, pembayaran, dan aplikasi fullstack. Setiap proyek punya case study detail. |

### Daftar proyek (10)

| Slug | Title | Year | Tags |
|------|-------|------|------|
| aksara-pustaka | Aksara Pustaka | 2026 | Laravel, MySQL, PHP, Tailwind CSS |
| sipadu | SiPadu | 2026 | Laravel, Tailwind CSS, MySQL, PHP |
| parking-app | ParkingApp | 2026 | PHP |
| inventoryflow | Inventoryflow | 2026 | Laravel, Tailwind CSS, PHP, MySQL |
| bagiberkah | BagiBerkah | 2026 | Next.js, Express, Prisma, Mayar, Xendit |
| e-vote | E-Democracy (E-Vote) | 2025 | PHP, Laravel, Bootstrap, MySQL |
| agrify | Smartfarm AI (Agrify) | 2025 | React, Python, Machine Learning, JavaScript |
| culture-connect | CultureConnect. | 2025 | Python, React, Machine Learning, Express, Prisma |
| spacelab | Spacelab | 2025 | Laravel, PHP, JavaScript |
| renshuu | Renshuu web | 2025 | React, Laravel |

### Deskripsi singkat (EN)

| Project | Description |
|---------|-------------|
| Aksara Pustaka | I built a web library system to manage books, members, loans, returns, stock, and history in one place. |
| SiPadu | I built SiPadu so school facility reports can be filed and tracked in real time instead of disappearing into paper trails. |
| ParkingApp | I built a parking web app for check-in, check-out, reporting, and audit in one flow. |
| Inventoryflow | I built Inventoryflow to handle school/lab equipment loans, inventory, approvals, and returns, without spreadsheet chaos. |
| BagiBerkah | I built BagiBerkah, a digital THR experience with mini-games and smarter allocation recommendations. |
| E-Vote | I built E-Vote so students can run digital OSIS elections with confidence, from ballots to real-time results. |
| Agrify | On a team project, I helped ship Smartfarm AI, combining modern web UI with ML-assisted insights for Indonesian farmers. |
| CultureConnect | With a distributed team, I helped build CultureConnect, an AI platform for more personal, community-aware cultural travel. |
| Spacelab | I built Spacelab to keep school schedules, rooms, and teachers conflict-free without spreadsheet chaos. |
| Renshuu | During PKL at CV Smartplus, my team and I built Renshuu, a job-search web app assigned for SMKN 1 Karawang students. |

---

## Screen: Project Detail

**Setara dengan:** `/projects/[slug]`

### Struktur case study (semua slug)

| Section | Hook label |
|---------|------------|
| Opening | Where it started |
| Constraints | What boxed the work in |
| Build | How it came together |
| Close | What I can stand behind |

**TUI:** Back link `← All projects` · title · stack tags · Live site / Source (jika ada URL) · 4 section narrative · prev/next project.

Narrative lengkap per slug disimpan di `internal/data/case-studies.go` (target implementasi).

---

## Screen: Services

**Setara dengan:** `/services`

| Elemen | Copy |
|--------|------|
| Label | My Services |
| H2 | Comprehensive Solutions |
| Sub | From wireframe concepts to fully animated frontends and scalable servers. I build performant products that stand out. |

| # | Title | Description |
|---|-------|-------------|
| 01 / Design | Web Design & Mobile-First | Translating ideas into pixel-perfect responsive interfaces. Wireframes to production-ready layouts that feel intuitive on every device. |
| 02 / Engineering | Frontend Development | High-quality UIs with React & Next.js. Clean components, reusable logic, and fluid state management. |
| 03 / Performance | Web Performance | Core Web Vitals optimization for instant load. SEO-ready architecture that ranks and converts. |
| 04 / Backend | APIs & Databases | Robust REST APIs, relational databases, and secure auth systems built to scale with your product. |
| 05 / DevOps | CI/CD & Deployment | Automated pipelines, container-ready apps, serverless hosting, and zero-downtime production deploys. |

**Di bawah grid:** ulangi section Process (4 langkah) + CTA Contact.

---

## Screen: Certificates

**Setara dengan:** `/about` #certificates

| Elemen | Copy |
|--------|------|
| Label | Certificates |
| H2 | Licenses & Certifications |
| Sub | Professional certifications and awards that validate my expertise in software development, cloud computing, and technology innovation. |

### Pinned (Featured)

1. Best Capstone Project — Coding Camp 2025
2. Coding Camp 2025 — Certificate of Completion
3. Country Award Winner — Intel AI for Youth

**Total:** 52 sertifikat (lihat `internal/data/certificates.go` setelah sync).

---

## Screen: Contact

**Setara dengan:** CTA blocks (bukan halaman `/contact` stub)

| Elemen | Copy |
|--------|------|
| H2 | Need a web product that actually ships in the next 90 days? |
| Body | Open to freelance and full-time. Remote (WIB). Write to contact@habibiahmada.dev. I usually reply within 48 hours. |
| Primary CTA | Let's talk |

**Social Profiles:**

| Platform | URL |
|----------|-----|
| GitHub | https://github.com/habibiahmada |
| LinkedIn | https://www.linkedin.com/in/habibi-ahmad-aziz |
| Instagram | https://instagram.com/habibiahmad.a |
| Email | mailto:contact@habibiahmada.dev |

---

## Alur Pengguna (User Flow)

```
                         ┌──────────┐
                         │   HOME   │
                         │ (scroll) │
                         └────┬─────┘
              ┌────────────────┼────────────────┐
              ▼                ▼                ▼
         ┌────────┐      ┌──────────┐     ┌──────────┐
         │ ABOUT  │      │ PROJECTS │     │ SERVICES │
         └────┬───┘      └────┬─────┘     └────┬─────┘
              │               │                │
    ┌─────────┼─────────┐     ▼                │
    ▼         ▼         ▼  ┌────────────┐        │
 SKILLS   EXPERIENCE  CERT │  PROJECT   │        │
                           │   DETAIL   │        │
                           └─────┬──────┘        │
                                  │ prev/next     │
              ┌──────────────────┼───────────────┘
              ▼
         ┌──────────┐
         │ CONTACT  │
         └──────────┘
```

**Journey umum:**
- **Rekruter:** Home hero → Featured Projects → Project Detail → Contact
- **Kredibilitas:** About → Experience → Certificates → Press (Home)
- **Klien:** Services → Process → Contact

---

## Gap Implementasi Saat Ini

| Area | Status terminal | Target |
|------|-----------------|--------|
| Profile copy | Generic / email salah | Sync ke tabel Profil di atas |
| Projects | 4 proyek, deskripsi tidak match | 10 proyek + case studies |
| Experience | 1 entry freelance | 4 work + 2 education |
| Skills | 12 item, kategori berbeda | 16 tools flat list |
| Certificates | 1 dummy | 52 + 3 pinned |
| Services screen | Tidak ada | Screen baru |
| Home sections | Hero + social only | + Companies, Featured, Press, Process, CTA |
| Social links | 3 link | + Instagram, email benar |
| Design colors | #FF6B6B palette | Brand red #ef4444 (lihat design-system.md) |

---

Terakhir diperbarui: Agustus 2026
