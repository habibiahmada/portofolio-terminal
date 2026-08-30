# habibiahmada-terminal

Interactive CLI portfolio untuk **Habibi Ahmad Aziz** — Full-Stack Web Developer.

Proyek ini adalah satu Go TUI yang dapat diakses melalui dua jalur distribusi:

```bash
npx habibiahmada          # local experience
ssh habibiahmada.dev      # remote experience
```

Website portfolio (`https://habibiahmada.dev`) dan terminal portfolio berbagi identitas yang sama, tetapi merupakan presentation layer terpisah.

## Akses Cepat

| Siapa | Perintah | Keterangan |
|-------|----------|------------|
| Pengguna | `npx habibiahmada` | Jalankan TUI di laptop Anda |
| Pengguna | `ssh habibiahmada.dev` | TUI via SSH (EC2 harus ON) |
| Developer | `go run ./cmd/portfolio` | Development mode |
| Developer | `make dev` | Sama, via Makefile |

Navigasi TUI: `↑↓` Navigate · `Enter` Select · `←` Back · `Q` Quit

## Ekosistem Portfolio

```
                     habibiahmada.dev
                            │
              ┌─────────────┴─────────────┐
              │                           │
           Website                     Terminal
              │                           │
   habibiahmada.dev             npx habibiahmada
              │                           │
              └───────────┬───────────────┘
                          │
                     Same Portfolio
```

| Interface | Perintah | Runtime | Ketergantungan |
|-----------|----------|---------|----------------|
| Website | `https://habibiahmada.dev` | Browser → Cloudflare → Vercel → Next.js | Supabase |
| CLI (local) | `npx habibiahmada` | Laptop user → Go binary → Bubble Tea | Node.js/npm (distribusi) |
| SSH (remote) | `ssh habibiahmada.dev` | EC2 → Wish → Go TUI | EC2 harus hidup |

## Tech Stack

| Layer | Teknologi | Fungsi |
|-------|-----------|---------|
| Runtime | Go | Aplikasi TUI |
| TUI engine | Bubble Tea | Terminal UI |
| Styling | Lip Gloss | Terminal styling |
| SSH | Wish | SSH application server |
| Distribusi | npm | `npx habibiahmada` (wrapper, bukan runtime TUI) |
| Infrastructure | AWS EC2 | SSH portfolio + Telegram Agent |

## Struktur Proyek

```
habibiahmada-terminal/
├── cmd/
│   ├── portfolio/main.go    # entry point local / npx
│   └── ssh/main.go           # entry point SSH server
├── internal/
│   ├── tui/                  # TUI core (screens & navigation)
│   ├── styles/               # Lip Gloss styles
│   └── data/                 # portfolio data (bundled)
├── npm/                      # npm package wrapper
├── scripts/                  # build scripts
├── docs/                     # dokumentasi
├── go.mod
└── Makefile
```

## Menu TUI

- About
- Projects (Renshuu, SmartFarm AI, CultureConnect, Spacelab)
- Skills
- Experience
- Certificates
- Contact

## Dokumentasi

### Panduan

| Dokumen | Isi |
|---------|-----|
| [docs/user-guide.md](docs/user-guide.md) | Cara akses portfolio di terminal |
| [docs/getting-started.md](docs/getting-started.md) | Setup proyek untuk developer baru |
| [docs/development-guide.md](docs/development-guide.md) | Workflow development sehari-hari |
| [docs/task-list.md](docs/task-list.md) | Task list & roadmap proyek |

### Arsitektur

| Dokumen | Isi |
|---------|-----|
| [docs/index.md](docs/index.md) | Indeks semua dokumentasi |
| [docs/architecture.md](docs/architecture.md) | Arsitektur lengkap |
| [docs/tech-stack.md](docs/tech-stack.md) | Pembagian tech stack |
| [docs/folder-structure.md](docs/folder-structure.md) | Struktur folder & evolusi |
| [docs/npx-vs-ssh.md](docs/npx-vs-ssh.md) | Perbandingan local vs remote |
| [docs/data-strategy.md](docs/data-strategy.md) | Strategi data TUI vs website |
| [docs/deployment.md](docs/deployment.md) | Deployment EC2, npm, CI/CD |
| [AGENTS.md](AGENTS.md) | Panduan untuk AI coding assistant |

## Development

```bash
# Clone & setup
git clone https://github.com/habibiahmada/habibiahmada-terminal.git
cd habibiahmada-terminal
go mod download

# Jalankan TUI
make dev

# Build semua platform
make build

# Quality checks
make lint
```

## CI/CD

| Target | Alur |
|--------|------|
| CLI | Git tag → GitHub Actions → Go build → npm publish |
| SSH | Git push → GitHub Actions → Deploy → EC2 (planned) |

## Lisensi

MIT — lihat [LICENSE](../LICENSE).
