# Struktur Folder

Struktur target dan evolusi proyek `habibiahmada-terminal/`.

## Struktur Lengkap (Target Akhir)

```
habibiahmada-terminal/
│
├── cmd/
│   ├── portfolio/
│   │   └── main.go              # Entry point local / npx
│   │
│   └── ssh/
│       └── main.go              # Entry point SSH server
│
├── internal/
│   │
│   ├── tui/
│   │   ├── app.go               # Root TUI / state
│   │   ├── splash.go            # Splash startup
│   │   ├── home.go              # Home screen
│   │   ├── about.go             # About
│   │   ├── projects.go          # Projects
│   │   ├── project_detail.go    # Detail project
│   │   ├── skills.go            # Skills
│   │   ├── experience.go        # Experience
│   │   ├── certificates.go      # Certificates
│   │   ├── contact.go           # Contact
│   │   └── keymap.go            # Keyboard shortcuts
│   │
│   ├── components/
│   │   ├── header.go
│   │   ├── sidebar.go
│   │   ├── footer.go
│   │   ├── card.go
│   │   ├── list.go
│   │   ├── modal.go
│   │   ├── illustration.go      # Signature + variant picker
│   │   ├── hero.go              # Home hero block
│   │   ├── figlet.go            # FIGlet besar (go-figure)
│   │   └── progress_bar.go      # Splash progress
│   │
│   ├── styles/
│   │   ├── styles.go            # Lip Gloss styles (umum)
│   │   └── illustration.go      # Lip Gloss styles (ilustrasi)
│   │
│   ├── assets/
│   │   ├── assets.go            # go:embed
│   │   └── art/                 # Signature, about art
│   │
│   └── data/
│       ├── portfolio.go         # Struct definitions
│       ├── profile.go
│       ├── projects.go
│       ├── skills.go
│       ├── experience.go
│       └── certificates.go
│
├── npm/
│   ├── package.json
│   ├── index.js
│   └── bin/
│       ├── habibiahmada-linux-x64
│       ├── habibiahmada-linux-arm64
│       ├── habibiahmada-darwin-x64
│       ├── habibiahmada-darwin-arm64
│       └── habibiahmada-win-x64.exe
│
├── assets/
│   └── ...
│
├── scripts/
│   ├── build.sh
│   └── release.sh
│
├── .github/
│   └── workflows/
│       ├── release.yml
│       └── deploy.yml
│
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── LICENSE
```

## Struktur Minimal (Fase 1 — Mulai Di Sini)

```
habibiahmada-terminal/
│
├── cmd/
│   ├── portfolio/main.go
│   └── ssh/main.go
│
├── internal/
│   ├── tui/
│   │   ├── app.go
│   │   ├── home.go
│   │   ├── projects.go
│   │   └── styles.go
│   │
│   └── data/
│       └── portfolio.go
│
├── npm/
│   ├── package.json
│   └── index.js
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

**Jangan over-engineering sejak awal.** Setelah TUI matang, pecah ke `components/`, `styles/`, `data/`, `scripts/`, `.github/`.

## Peran Setiap Direktori

### `cmd/`

Hanya entry point. Tidak ada logic UI di sini.

| File | Alur |
|------|------|
| `cmd/portfolio/main.go` | npx → Go binary → internal/tui → Bubble Tea |
| `cmd/ssh/main.go` | SSH → Wish → SSH session → internal/tui → Bubble Tea |

### `internal/tui/`

Jantung aplikasi. Bubble Tea model utama:

- Current Screen
- Selected Menu
- Window Size
- Keyboard Input
- Navigation

```
                  App
                   │
       ┌───────────┼────────────┐
       │           │            │
      Home      Projects      Skills
                   │
                   ▼
             Project Detail
```

### `internal/components/`

Layout reusable:

```
┌───────────────────────────────────────────────┐
│ HABIBI AHMAD AZIZ                             │ ← Header
├──────────────┬────────────────────────────────┤
│ > About      │                                │
│   Projects   │       Content                  │ ← Sidebar + Content
│   Skills     │                                │
├──────────────┴────────────────────────────────┤
│ ↑↓ Navigate  Enter Select  Q Quit             │ ← Footer
└───────────────────────────────────────────────┘
```

### `internal/data/`

Data dipisah dari UI:

```
Data → TUI
```

Bukan hard-code di setiap halaman.

### `npm/`

Wrapper distribusi:

```
npx habibiahmada → npm package → deteksi OS/arch → Go binary → TUI
```

### `scripts/`

`build.sh` menghasilkan:

```
dist/
├── habibiahmada-linux-x64
├── habibiahmada-linux-arm64
├── habibiahmada-darwin-x64
├── habibiahmada-darwin-arm64
└── habibiahmada-win-x64.exe
```

`release.sh` memasukkan binary ke npm package (langsung dari `dist/`).

### EC2 — Bukan di Repo

Infrastructure server tidak masuk repo sebagai folder rumit. EC2 menjalankan binary `cmd/ssh` yang dibuild terpisah.

## Dependency Flow

```
                ┌───────────────┐
                │ Portfolio Data│
                └───────┬───────┘
                        │
                        ▼
                ┌───────────────┐
                │   TUI Core    │
                └───────┬───────┘
                        │
             ┌──────────┴──────────┐
             │                     │
             ▼                     ▼
       Local Runtime          SSH Runtime
             │                     │
             ▼                     ▼
          npx                   EC2
```
