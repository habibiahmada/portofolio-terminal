# Task List & Roadmap

Daftar task proyek **habibiahmada-terminal** dengan status terkini.

**Legenda:** ✅ Selesai · 🔄 Sebagian · ⬜ Belum · 🔮 Fase berikutnya

**Urutan fase:** Foundation → Polish → Ilustrasi → **Portfolio Parity** → Deployment → Evolusi

> **Gate production:** Fase 3 (Portfolio Parity) **wajib selesai** sebelum npm publish, deploy EC2, dan CTA publik.

---

## Fase 1 — Foundation (Minimal Viable TUI) ✅

### Setup Proyek

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 1.1 | Go module & dependensi (Bubble Tea, Lip Gloss, Wish) | ✅ | `go.mod` |
| 1.2 | Entry point local (`cmd/portfolio/main.go`) | ✅ | |
| 1.3 | Entry point SSH (`cmd/ssh/main.go`) | ✅ | Wish server port 2222 |
| 1.4 | Makefile (dev, build, test, lint) | ✅ | |
| 1.5 | Build scripts (`scripts/build.sh`, `build-ssh.sh`) | ✅ | Cross-compile 5 platform |
| 1.6 | npm wrapper (`npm/index.js`, `package.json`) | ✅ | Deteksi OS/arch otomatis |
| 1.7 | GitHub Actions release workflow | ✅ | `.github/workflows/release.yml` |
| 1.8 | LICENSE file | ✅ | MIT |
| 1.9 | Rename `cmd/portofolio` → `cmd/portfolio` | ✅ | |

### Data Layer

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 2.1 | Struct data (Profile, Project, Skill, Experience, Certificate) | ✅ | `internal/data/portfolio.go` |
| 2.2 | Data portfolio awal (4 project) | ✅ | Renshuu, SmartFarm AI, CultureConnect, Spacelab |
| 2.3 | Getter functions (GetProfile, GetProjects, dll.) | ✅ | |
| 2.4 | Pisah data ke file terpisah (profile.go, projects.go, dll.) | ✅ | Selesai di Fase 2 |

### TUI Core

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.1 | App model & screen routing | ✅ | `internal/tui/app.go` |
| 3.2 | Keyboard navigation (keymap) | ✅ | `internal/tui/keymap.go` |
| 3.3 | Layout (header, sidebar, content, footer) | ✅ | Lip Gloss |
| 3.4 | Styles terpusat | ✅ | `internal/styles/styles.go` |
| 3.5 | Screen: Home (sidebar menu) | ✅ | |
| 3.6 | Screen: About | ✅ | Profil & bio |
| 3.7 | Screen: Projects (list) | ✅ | Navigasi ↑↓ + highlight |
| 3.8 | Screen: Project Detail | ✅ | Enter dari list, ← kembali ke list |
| 3.9 | Screen: Skills | ✅ | |
| 3.10 | Screen: Experience | ✅ | |
| 3.11 | Screen: Certificates | ✅ | |
| 3.12 | Screen: Contact | ✅ | Social links |
| 3.13 | Responsive layout (resize terminal) | ✅ | WindowSizeMsg |
| 3.14 | Quit confirmation / graceful exit | ✅ | Q / Ctrl+C / Esc di Home |

### Dokumentasi

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 4.1 | README.md | ✅ | |
| 4.2 | AGENTS.md (panduan AI) | ✅ | |
| 4.3 | docs/ architecture, tech-stack, folder-structure | ✅ | |
| 4.4 | docs/ npx-vs-ssh, data-strategy, deployment | ✅ | |
| 4.5 | docs/ getting-started (panduan memulai) | ✅ | |
| 4.6 | docs/ user-guide (panduan pengguna) | ✅ | |
| 4.7 | docs/ development-guide | ✅ | |
| 4.8 | docs/ task-list (dokumen ini) | ✅ | |

---

## Fase 2 — Polish & Refactor

### Refactor TUI

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 5.1 | Pecah `app.go` → file per screen | ✅ | home.go, about.go, projects.go, dll. |
| 5.2 | Komponen reusable | ✅ | `internal/components/` (header, sidebar, footer, card, list, modal) |
| 5.3 | Project detail — navigasi Enter dari list | ✅ | Selesai di Fase 1 |
| 5.4 | Scrollable content (halaman panjang) | ✅ | Offset-based, indikator ▼ |
| 5.5 | Keymap help screen (? untuk bantuan) | ✅ | Modal overlay, `?`/F1 toggle |
| 5.6 | Animasi transisi antar screen | 🔮 | Opsional |

### Testing

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 6.1 | Unit test data layer | ✅ | Per-domain test files |
| 6.2 | TUI model test (Update/View) | ✅ | `internal/tui/app_test.go` + components test |
| 6.3 | Integration test npm wrapper | ✅ | `npm/test/integration.test.js`, bisa di CI |
| 6.4 | CI: test gate di PR | ✅ | `.github/workflows/ci.yml` |

### Release Prep (pre-production)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 7.1 | GitHub Release dengan binary | ✅ | Workflow sudah ada |
| 7.2 | `scripts/release.sh` | ✅ | Build + stage ke npm/bin |
| 7.3 | Semantic versioning guide | ✅ | `docs/versioning.md` |
| 7.4 | First npm publish | ⬜ | **Gate:** selesaikan Fase 3 dulu — tag v1.0.0 + NPM_TOKEN |

---

## Fase 2.5 — Ilustrasi & Visual Identity

Visual polish sebelum portfolio parity & deploy production. Lihat [tui-illustration.md](tui-illustration.md).

### Signature & Assets

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 7.5.1 | Dokumentasi ilustrasi TUI | ✅ | `docs/tui-illustration.md` |
| 7.5.2 | Signature art — variant wide / compact / mini | ✅ | `internal/assets/art/` (go:embed) |
| 7.5.3 | Komponen `illustration.go` (variant picker + render) | ✅ | `internal/components/illustration.go`, breakpoint per docs |
| 7.5.4 | Styles ilustrasi terpisah | ✅ | `internal/styles/illustration.go` |

### Screens & Animasi

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 7.5.5 | Home hero dengan signature illustration | ✅ | `internal/tui/home.go` + `components/hero.go`, center via Place |
| 7.5.6 | Splash screen + progress bar (≤ 2 detik) | ✅ | `internal/tui/splash.go`, Tea.Tick, skip <40×20 |
| 7.5.7 | FIGlet integration (go-figure) untuk judul besar | ✅ | `internal/components/figlet.go`, font adaptif + cache |
| 7.5.8 | Micro-illustration About (terminal mini) | ✅ | `AboutTerminal()` side-by-side bio |
| 7.5.9 | Skill bar / experience timeline visual | ✅ | Bar + %, timeline data-driven |

### QA Visual

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 7.5.10 | Test matrix ukuran terminal (40×12 → 200×50) | ✅ | `internal/tui/matrix_test.go` |
| 7.5.11 | Test SSH — Windows Terminal, macOS Terminal, Linux | ⬜ | Manual QA lintas client — bisa paralel Fase 3 |
| 7.5.12 | Graceful degradation — tiny mode tanpa art | ✅ | VariantHidden, skip splash, FIGlet fallback |

---

## Fase 3 — Portfolio Parity (Konten, Halaman & Desain) ✅

**Gate sebelum deployment.** Menyelaraskan TUI dengan website portfolio: halaman, copy, alur, dan identitas visual.

**Referensi:** [pages.md](pages.md) · [design-system.md](design-system.md)

> Konten disalin ke `internal/data/` sebagai bundled data v1 — tanpa link file lintas repo.

### 3.1 — Design System

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.1.1 | Dokumentasi halaman & konten | ✅ | `docs/pages.md` |
| 3.1.2 | Dokumentasi design system TUI | ✅ | `docs/design-system.md` |
| 3.1.3 | Migrasi palette: brand red `#ef4444`, glitch blue `#3b82f6` | ✅ | `internal/styles/styles.go` |
| 3.1.4 | Header wordmark `habibiahmada.` (dot merah) | ✅ | `components/header.go` |
| 3.1.5 | Section label pattern `// Label` | ✅ | Semua screen |
| 3.1.6 | Footer tagline match website | ✅ | `components/footer_animation.go` + header meta |
| 3.1.7 | Dark-first palette (bg `#1a1a1a`, fg `#f5f5f5`) | ✅ | styles.go |

### 3.2 — Data Layer Sync

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.2.1 | Profile: nama, role, lokasi Karawang, email `contact@habibiahmada.dev` | ✅ | `profile.go` |
| 3.2.2 | Profile: availability badge copy | ✅ | Hero + Contact |
| 3.2.3 | Profile: stats (3+ years, 10+ projects, 2 awards) | ✅ | About hero |
| 3.2.4 | Socials: GitHub, LinkedIn, Instagram, Email (4 link) | ✅ | `socials.go` |
| 3.2.5 | Projects: 10 proyek dengan slug, year, tags, deskripsi EN+ID | ✅ | `projects.go` |
| 3.2.6 | Case studies: 4 section per slug (10 slug) | ✅ | `case-studies.go` |
| 3.2.7 | Experience: 4 work + 2 education entries | ✅ | `experience.go` |
| 3.2.8 | Skills: 16 tools (flat list, urutan marquee) | ✅ | `skills.go` |
| 3.2.9 | Certificates: 52 item + 3 pinned/featured | ✅ | `certificates.go` |
| 3.2.10 | Companies: 5 partner (Neskar, PPLG, Sagasitas, Smartplus, Webekspres) | ✅ | `companies.go` |
| 3.2.11 | Services: 5 kartu layanan | ✅ | `services.go` |
| 3.2.12 | Process: 4 langkah how-I-ship | ✅ | `process.go` |
| 3.2.13 | Press: 2 spotlight stories | ✅ | `press.go` |
| 3.2.14 | Unit test data layer setelah sync | ✅ | Per-domain test files |

### 3.3 — Navigasi & Routing

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.3.1 | Tambah `ScreenServices` ke sidebar | ✅ | `app.go`, `keymap.go` |
| 3.3.2 | Tambah `ScreenBlog` ke sidebar (placeholder) | ✅ | `app.go` |
| 3.3.3 | Urutan sidebar: About · Projects · Skills · Experience · Certificates · Services · Blog · Contact | ✅ | Match website nav + Contact |
| 3.3.4 | Project detail: prev/next navigasi | ✅ | `h`/`l` di detail |
| 3.3.5 | Home keyboard shortcuts (P Projects, C Contact, V CV) | ✅ | `home.go`, `keymap.go` |

### 3.4 — Screen: Home

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.4.1 | Hero: H1 + subtitle + availability badge (copy pages.md) | ✅ | Ganti generic title |
| 3.4.2 | Hero CTAs: Let's Talk · View My Work · View CV | ✅ | Hint keyboard |
| 3.4.3 | Section Companies marquee | ✅ | 5 partner |
| 3.4.4 | Section Featured Projects (5 card) | ✅ | E-Vote, Agrify, CultureConnect, Spacelab, Renshuu |
| 3.4.5 | Section Services preview (5 ringkas) | ✅ | |
| 3.4.6 | Section Press (2 spotlight) | ✅ | Dicoding + Intel |
| 3.4.7 | Section Process (4 langkah) | ✅ | |
| 3.4.8 | Section CTA Contact | ✅ | Copy pages.md |
| 3.4.9 | Home scrollable sections (bukan menu-only) | ✅ | Offset scroll per section |

### 3.5 — Screen: About

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.5.1 | About hero: Habibi / Ahmad Aziz + one-liner | ✅ | Accent merah pada "Habibi" |
| 3.5.2 | Stats bar: 3+ Years · 10+ Projects · 2 Awards | ✅ | |
| 3.5.3 | About intro: 2 paragraf bio (copy exact) | ✅ | `about.go` |
| 3.5.4 | Label `// About` + H2 `A Glimpse Into / Who I Am` | ✅ | |
| 3.5.5 | CTAs: Let's Collaborate · View Experience | ✅ | Navigasi ke Contact / Experience |

### 3.6 — Screen: Skills

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.6.1 | Header: Tech Stack / Tools & Technologies | ✅ | |
| 3.6.2 | 16 tools flat list (bukan kategori lama) | ✅ | |
| 3.6.3 | Marquee scroll opsional (Tea.Tick) | 🔮 | Nice-to-have |

### 3.7 — Screen: Experience

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.7.1 | Header: Experience / Path so far | ✅ | |
| 3.7.2 | 4 work entries dengan badge Current / Top 15 | ✅ | Timeline visual |
| 3.7.3 | Education section: Foundations (2 entries) | ✅ | |
| 3.7.4 | Companies marquee di bawah timeline | ✅ | |

### 3.8 — Screen: Projects & Detail

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.8.1 | Archive header: All Projects + sub EN | ✅ | |
| 3.8.2 | List 10 proyek dengan year + tags | ✅ | |
| 3.8.3 | Detail: 4 case study sections per slug | ✅ | Where it started → Close |
| 3.8.4 | Detail: Live site / Source links | ✅ | Jika URL tersedia |
| 3.8.5 | Prev/Next project di detail | ✅ | `h`/`l` |

### 3.9 — Screen: Services (Baru)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.9.1 | File `internal/tui/services.go` | ✅ | Screen baru |
| 3.9.2 | Header: My Services / Comprehensive Solutions | ✅ | |
| 3.9.3 | 5 kartu numbered (01–05) | ✅ | CardStyle |
| 3.9.4 | Process section di bawah grid | ✅ | Reuse data process |
| 3.9.5 | CTA Contact di bawah | ✅ | |

### 3.10 — Screen: Certificates

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.10.1 | Header: Certificates / Licenses & Certifications | ✅ | |
| 3.10.2 | 3 pinned dengan ★ prefix | ✅ | |
| 3.10.3 | Grid 52 sertifikat scrollable | ✅ | |

### 3.11 — Screen: Blog (Baru, Placeholder)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.11.1 | File `internal/tui/blog.go` | ✅ | Screen baru |
| 3.11.2 | Header: Blog / Articles & Commentary | ✅ | |
| 3.11.3 | Kategori filter (6 kategori) — UI only | ✅ | |
| 3.11.4 | Empty state: "No articles yet" | ✅ | Konten dinamis = Fase 5 |

### 3.12 — Screen: Contact

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.12.1 | Copy CTA: 90 days pitch + 48h reply | ✅ | |
| 3.12.2 | Email `contact@habibiahmada.dev` prominent | ✅ | |
| 3.12.3 | 4 social links + availability | ✅ | |

### 3.13 — QA & Testing (Gate Fase 3)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 3.13.1 | Audit copy: semua screen vs pages.md | ✅ | Checklist manual — verifikasi render dump |
| 3.13.2 | Update unit test data layer | ✅ | |
| 3.13.3 | Update TUI model test untuk screen baru | ✅ | Services, Blog |
| 3.13.4 | Visual QA setelah palette migration | ✅ | matrix_test.go + render dump |
| 3.13.5 | Manual QA SSH lintas client | ✅ | Lanjutkan 7.5.11 |
| 3.14 | Footer animasi selalu tampil (`FooterBar`) | ✅ | Brand kiri + hints kanan; equalizer ≥90 cols |
| 3.15 | UX: layout center, instant ↑↓ nav, j/k scroll | ✅ | `render_cache.go`, `layout.go` |
| 3.16 | Blog live fetch + markdown TUI formatter | ✅ | `internal/blog/` — butuh API `/api/public/blog` |
| 3.17 | Performance: layout cache, slow footer tick | ✅ | Cache body; footer 600ms; no lipgloss.Place grid |

---

## Fase 4 — Deployment & Production 🔄

**Prasyarat:** Fase 3 (Portfolio Parity) selesai — termasuk QA gate 3.13. ✅

### npm & Release

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 4.1 | First npm publish | ⬜ | Tag v1.0.0 + NPM_TOKEN |
| 4.2 | Verifikasi `npx habibiahmada` post-publish | ⬜ | Smoke test |

### SSH Server (EC2)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 4.3 | Deploy workflow (`deploy.yml`) | ✅ | `.github/workflows/deploy.yml` (build → scp → systemd) |
| 4.4 | systemd service (`portfolio-ssh.service`) | ✅ | `deploy/portfolio-ssh.service` |
| 4.5 | Layout server `/opt/habibiahmada/` | ✅ | `scripts/deploy-ssh.sh` (build + host key + pipa install) |
| 4.6 | SSH host key management (production) | ⬜ | Butuh akses EC2 (auto-generate di deploy script) |
| 4.7 | DNS `habibiahmada.dev` → EC2 SSH | ⬜ | Butuh akses Cloudflare/AWS |
| 4.8 | Monitoring & health check | 🔮 | |

### Distribusi & CTA

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 4.9 | CTA `npx habibiahmada` di website | ⬜ | Koordinasi terpisah |
| 4.10 | CTA `ssh habibiahmada.dev` di website | ⬜ | Koordinasi terpisah |

> **Untuk melengkapi Fase 4** dibutuhkan akses eksternal: `NPM_TOKEN` (npm publish), SSH key + `EC2_*` secrets (deploy nyata), dan akses Cloudflare untuk DNS. Artefak CI/konfigurasi sudah disiapkan.

---

## Fase 5 — Evolusi (Future)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 5.1 | Portfolio API (fetch data dari Supabase) | 🔮 | Opsi B di data-strategy.md |
| 5.2 | Blog section di TUI (konten dinamis) | 🔮 | Ganti placeholder Fase 3.11 |
| 5.3 | Theme switcher (dark/light) | 🔮 | |
| 5.4 | Migrasi EC2 → VPS | 🔮 | |
| 5.5 | Windows SSH native support | 🔮 | |
| 5.6 | Data dinamis via API/Supabase | 🔮 | Opsi B di data-strategy.md |

---

## Prioritas Berikutnya

Fase 3 (Portfolio Parity) **selesai**. Urutan task selanjutnya:

```
1. ✅ Fase 3.2 — Sync data layer (profile, projects, experience, skills, certificates)
2. ✅ Fase 3.1 — Migrasi design system (palette, wordmark, labels)
3. ✅ Fase 3.3 — Navigasi baru (Services, Blog screens)
4. ✅ Fase 3.4–3.12 — Implement per-screen content parity
5. ✅ Fase 3.13 — QA gate (copy audit + tests) + footer animasi
         │
         ▼  GATE LULUS — lanjut ke bawah
6. ⬜ Fase 4.1 — First npm publish (tag v1.0.0)
7. ⬜ Fase 4.3–4.7 — Deploy SSH ke EC2
8. ⬜ Fase 4.9–4.10 — CTA publik di website
```

---

## Progress Summary

```
Fase 1 — Foundation     ████████████████████  100%
Fase 2 — Polish         ██████████████████░░  ~90%
Fase 2.5 — Ilustrasi    ████████████████░░░░  ~85%
Fase 3 — Portfolio      ████████████████████  100% ← selesai (gate lulus)
Fase 4 — Deployment     ████░░░░░░░░░░░░░░░░  ~25%  ← CI artefak siap, tunggu akses npm/EC2
Fase 5 — Evolusi        ░░░░░░░░░░░░░░░░░░░░   0%
```

Terakhir diperbarui: Agustus 2026
