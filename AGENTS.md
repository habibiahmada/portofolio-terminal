# AGENTS.md — habibiahmada-terminal

Panduan untuk AI coding assistant yang bekerja pada proyek terminal portfolio ini.

## Konteks Proyek

Membangun **satu Go TUI** (Bubble Tea + Lip Gloss) yang dijalankan melalui:

1. **`npx habibiahmada`** — TUI berjalan di laptop user (local experience)
2. **`ssh habibiahmada.dev`** — TUI berjalan di AWS EC2 via Wish (remote experience)

Jangan membuat dua TUI terpisah. Satu TUI core, dua transport/runtime.

## Aturan Arsitektur

### Wajib

- **Satu TUI core** dipakai oleh `cmd/portfolio` dan `cmd/ssh`
- **Data v1 dibundel** di Go (`internal/data/`), bukan fetch API/Supabase
- **npm hanya distribution wrapper** — bukan runtime TUI
- **Pisahkan data dari UI** — TUI membaca data, bukan hard-code di setiap screen
- **Styles terpusat** di `internal/styles/` (Lip Gloss)
- **Komponen reusable** di `internal/components/` (header, sidebar, footer, card, list, modal)

### Jangan

- Membuat TUI A untuk npx dan TUI B untuk SSH
- Over-engineering struktur folder sejak hari pertama
- Menambahkan ketergantungan Supabase/API ke CLI/SSH di v1
- Menggunakan Node.js/TypeScript untuk logic TUI (hanya npm wrapper)

## Struktur Entry Point

```
cmd/portfolio/main.go  →  npx habibiahmada  →  internal/tui
cmd/ssh/main.go        →  ssh habibiahmada.dev  →  internal/tui (sama)
```

## Screens TUI

| Screen | File (target) |
|--------|---------------|
| Root / Home | `internal/tui/app.go`, `home.go` |
| About | `internal/tui/about.go` |
| Projects | `internal/tui/projects.go`, `project_detail.go` |
| Skills | `internal/tui/skills.go` |
| Experience | `internal/tui/experience.go` |
| Certificates | `internal/tui/certificates.go` |
| Contact | `internal/tui/contact.go` |
| Keymap | `internal/tui/keymap.go` |

## Navigasi & Input

```
↑ ↓     Navigate
Enter   Select
←       Back
q       Quit
```

## Data Model (Contoh)

```go
type Project struct {
    Name        string
    Description string
    Stack       []string
    GitHub      string
    Live        string
}
```

Portfolio data awal mencakup: Renshuu, SmartFarm AI, CultureConnect, Spacelab.

## Failure Isolation

| Jika down | Website | CLI (npx) | SSH |
|-----------|---------|-----------|-----|
| Supabase | mungkin terdampak | tetap jalan | tetap jalan |
| Vercel | down | tetap jalan | tetap jalan |
| EC2 | tetap jalan | tetap jalan | down |

## npm Package Flow

```
npx habibiahmada
    → npm package
    → deteksi OS + architecture
    → jalankan Go binary
    → Bubble Tea → TUI
```

Binaries target: linux-x64, linux-arm64, darwin-x64, darwin-arm64, win-x64.

## SSH Deployment (EC2)

```
/opt/habibiahmada/habibiahmada-ssh
systemd: portfolio-ssh.service
```

EC2 juga menjalankan Telegram Agent (Finance, News, Blog Automation). SSH portfolio menumpang pada EC2 yang sama.

## Graphify

Sebelum explore codebase dengan grep/file read, gunakan graphify:

```bash
graphify query "..." --graph graphify-out/graph.json
graphify update .
```

Rule Cursor: `.cursor/rules/graphify.mdc` (alwaysApply).

## Skills Proyek

| Skill | Path | Kapan dipakai |
|-------|------|---------------|
| Graphify workflow | `.cursor/skills/graphify-workflow/` | Sebelum explore codebase |
| Go code quality | `.cursor/skills/go-code-quality/` | Saat menulis/mereview Go |
| Bubble Tea TUI | `.cursor/skills/bubble-tea-tui/` | Saat membangun UI terminal |

## Evolusi Struktur

**Fase 1 (minimal):**

```
cmd/portfolio, cmd/ssh
internal/tui/ (app, home, projects, styles)
internal/data/portfolio.go
npm/
```

**Fase 2 (setelah TUI matang):** pecah ke `components/`, `styles/`, `data/`, `scripts/`, `.github/`.

## Referensi

Semua spesifikasi detail ada di `docs/`. Jangan menambah fitur/tech stack di luar yang terdokumentasi di sana.
