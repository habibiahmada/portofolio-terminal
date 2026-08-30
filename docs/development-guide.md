# Panduan Development

Workflow sehari-hari untuk mengembangkan **habibiahmada-terminal**.

## Alur Kerja Harian

```
Edit kode → make dev → tes manual di terminal → make lint → commit
```

## Menjalankan TUI Lokal

```bash
# Hot reload manual: stop (Q) lalu jalankan ulang
go run ./cmd/portfolio

# Atau via Makefile
make dev
```

## Menjalankan SSH Server Lokal

Untuk menguji experience SSH tanpa deploy ke EC2:

```bash
# Terminal 1 — jalankan server
make dev-ssh
# Server listen di 0.0.0.0:2222

# Terminal 2 — connect sebagai client
ssh -p 2222 localhost
```

> Host key disimpan di `.ssh/term_info_ed25519` (auto-generated saat pertama kali).

## Mengubah Data Portfolio

Data v1 **dibundel di Go** — tidak ada API call. Semua data ada di:

```
internal/data/portfolio.go
```

### Menambah / mengubah project

Edit struct dan slice di `portfolio.go`:

```go
var projects = []Project{
    {
        Name:        "Nama Project",
        Description: "Deskripsi singkat",
        Stack:       []string{"Go", "React"},
        GitHub:      "https://github.com/...",
        Live:        "https://...",
        Featured:    true,
    },
}
```

Setelah edit, restart TUI (`make dev`) untuk melihat perubahan.

> Perubahan data memerlukan rebuild binary untuk pengguna npx (release baru).

## Mengubah TUI / Screen

Saat ini semua screen ada di satu file:

```
internal/tui/app.go    # App model, navigation, semua screen view
internal/tui/keymap.go # Keyboard handler
internal/styles/styles.go # Lip Gloss styles
```

### Screen yang ada

| Screen | Konstanta | Status |
|--------|-----------|--------|
| Home | `ScreenHome` | Sidebar + menu |
| About | `ScreenAbout` | Konten profil |
| Projects | `ScreenProjects` | Daftar project |
| Project Detail | `ScreenProjectDetail` | Detail project (Enter dari list) |
| Skills | `ScreenSkills` | Daftar skill |
| Experience | `ScreenExperience` | Riwayat kerja |
| Certificates | `ScreenCertificates` | Sertifikasi |
| Contact | `ScreenContact` | Social links |

### Menambah screen baru

1. Tambah konstanta di `Screen` enum (`app.go`)
2. Tambah ke `ScreenNames` dan `menuItems`
3. Tambah case di `renderContent()`
4. Update navigasi jika perlu

Target fase berikutnya: pecah ke file terpisah (`home.go`, `about.go`, dll.) — lihat [task-list.md](task-list.md).

## Styling

Styles terpusat di `internal/styles/styles.go`:

```go
styles.TitleStyle      // Judul halaman
styles.SubtitleStyle   // Subjudul
styles.SelectedStyle   // Item menu terpilih
styles.NormalStyle     // Teks biasa
styles.BorderStyle     // Border panel
styles.MutedStyle      // Teks secondary
styles.LinkStyle       // URL / link
```

Gunakan styles ini — jangan hard-code warna di screen.

## Quality Checks

```bash
# Format kode
make fmt
# atau: gofmt -w -s .

# Static analysis
make vet
# atau: go vet ./...

# Semua test
make test
# atau: go test ./...

# Semua sekaligus
make lint
```

## Build & Release

### Build lokal (semua platform)

```bash
make build
# Output: dist/
#   habibiahmada-linux-amd64
#   habibiahmada-linux-arm64
#   habibiahmada-darwin-amd64
#   habibiahmada-darwin-arm64
#   habibiahmada-windows-amd64.exe
```

### Build SSH server

```bash
make build-ssh
# Output: habibiahmada-ssh
```

### Release ke npm

Release otomatis via GitHub Actions saat push tag:

```bash
git tag v1.0.1
git push origin v1.0.1
```

Workflow `.github/workflows/release.yml` akan:
1. Run test & vet
2. Cross-compile semua platform
3. Copy binary ke `npm/bin/`
4. Publish ke npm registry
5. Buat GitHub Release dengan artefak

Detail deployment: [deployment.md](deployment.md).

## Konvensi Kode

### Arsitektur

- **Satu TUI core** — `internal/tui/` dipakai oleh `cmd/portfolio` dan `cmd/ssh`
- **Data terpisah dari UI** — TUI baca dari `internal/data/`, bukan hard-code
- **Entry point tipis** — `cmd/` hanya bootstrap, tidak ada logic UI

### Go

- Ikuti standar di skill `.cursor/skills/go-code-quality/SKILL.md`
- Package `internal/` — tidak diimport dari luar module
- Error handling eksplisit di entry point (`cmd/`)

### TUI

- Ikuti pola Bubble Tea: `Init`, `Update`, `View`
- Keyboard handling di `keymap.go`
- Layout: Header → Sidebar + Content → Footer

## Debugging

### TUI tidak render

Pastikan terminal cukup besar (min ~80×24). Bubble Tea butuh `WindowSizeMsg` sebelum render penuh.

### SSH lokal gagal connect

```bash
# Cek server jalan
ss -tlnp | grep 2222

# Connect dengan verbose
ssh -v -p 2222 localhost
```

### Build gagal

```bash
# Pastikan modul lengkap
go mod tidy
go mod download

# Build spesifik
go build ./cmd/portfolio
go build ./cmd/ssh
```

## Graphify (Codebase Navigation)

Sebelum explore codebase dengan grep, gunakan knowledge graph:

```bash
graphify query "what connects cmd/portfolio to internal/tui?" --graph graphify-out/graph.json
graphify explain "internal/tui/app.go" --graph graphify-out/graph.json
```

Update graph setelah perubahan signifikan:

```bash
graphify update .
```

Detail: `.cursor/skills/graphify-workflow/SKILL.md`.

## Referensi Skill Proyek

| Skill | Path | Kapan dipakai |
|-------|------|---------------|
| Go code quality | `.cursor/skills/go-code-quality/` | Menulis/review Go |
| Bubble Tea TUI | `.cursor/skills/bubble-tea-tui/` | UI, layout, navigasi |
| Graphify workflow | `.cursor/skills/graphify-workflow/` | Explore codebase |
