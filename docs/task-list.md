# Task List & Roadmap

Daftar task proyek **habibiahmada-terminal** dengan status terkini.

**Legenda:** ✅ Selesai · 🔄 Sebagian · ⬜ Belum · 🔮 Fase berikutnya

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

### Release & Distribution

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 7.1 | First npm publish | ⬜ | Tag v1.0.0 + NPM_TOKEN |
| 7.2 | GitHub Release dengan binary | ✅ | Workflow sudah ada |
| 7.3 | `scripts/release.sh` | ✅ | Build + stage ke npm/bin |
| 7.4 | Semantic versioning guide | ✅ | `docs/versioning.md` |

---

## Fase 2.5 — Ilustrasi & Visual Identity

Visual polish sebelum deploy production. Lihat [tui-illustration.md](tui-illustration.md).

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
| 7.5.11 | Test SSH — Windows Terminal, macOS Terminal, Linux | ⬜ | Manual QA lintas client |
| 7.5.12 | Graceful degradation — tiny mode tanpa art | ✅ | VariantHidden, skip splash, FIGlet fallback |

---

## Fase 3 — Deployment & Production

### SSH Server (EC2)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 8.1 | Deploy workflow (`deploy.yml`) | ⬜ | Git push → EC2 |
| 8.2 | systemd service (`portfolio-ssh.service`) | ⬜ | |
| 8.3 | Layout server `/opt/habibiahmada/` | ⬜ | |
| 8.4 | SSH host key management (production) | ⬜ | |
| 8.5 | DNS `habibiahmada.dev` → EC2 SSH | ⬜ | |
| 8.6 | Monitoring & health check | 🔮 | |

### Integrasi Website

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 9.1 | CTA di website: `npx habibiahmada` | ⬜ | Repo terpisah (portfolio-web) |
| 9.2 | CTA di website: `ssh habibiahmada.dev` | ⬜ | |
| 9.3 | Sinkronisasi data website ↔ terminal | 🔮 | API/Supabase — bukan v1 |

---

## Fase 4 — Evolusi (Future)

| # | Task | Status | Catatan |
|---|------|--------|---------|
| 10.1 | Portfolio API (fetch data dari Supabase) | 🔮 | Opsi B di data-strategy.md |
| 10.2 | Blog section di TUI | 🔮 | |
| 10.3 | Theme switcher (dark/light) | 🔮 | |
| 10.4 | Migrasi EC2 → VPS | 🔮 | |
| 10.5 | Windows SSH native support | 🔮 | |

---

## Prioritas Berikutnya (Fase 2 → 2.5)

Urutan task yang disarankan:

```
1. ✅ Pecah app.go ke file per screen (selesai)
2. ✅ Buat internal/components/ (header, sidebar, footer, dll.) (selesai)
3. ✅ Unit tests untuk data layer + TUI model (selesai)
4. ⬜ Signature illustration + home hero (Fase 2.5)
5. ⬜ Splash screen startup
6. ⬜ Test matrix visual (terminal sizes + SSH)
7. ⬜ First npm publish (tag v1.0.0) — butuh NPM_TOKEN
8. ⬜ Deploy SSH ke EC2 (deploy.yml + systemd)
```

> Ilustrasi (Fase 2.5) selesai **sebelum** npm publish dan deploy EC2.

---

## Progress Summary

```
Fase 1 — Foundation     ████████████████████  100%
Fase 2 — Polish         ██████████████████░░  ~90%
Fase 2.5 — Ilustrasi    ████████████████░░░░  ~80%
Fase 3 — Deployment     ██░░░░░░░░░░░░░░░░░░  ~10%
Fase 4 — Evolusi        ░░░░░░░░░░░░░░░░░░░░   0%
```

Terakhir diperbarui: Agustus 2026
