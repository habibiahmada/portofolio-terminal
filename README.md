# habibiahmada-terminal

Interactive CLI portfolio untuk **Habibi Ahmad Aziz** — Full-Stack Web Developer.

Proyek ini adalah satu Go TUI yang dapat diakses melalui dua jalur distribusi:

```bash
npx habibiahmada          # local experience
ssh habibiahmada.dev      # remote experience
```

Website portfolio (`https://habibiahmada.dev`) dan terminal portfolio berbagi identitas yang sama, tetapi merupakan presentation layer terpisah.

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
| CLI (local) | `npx habibiahmada` | Laptop user → Go binary → Bubble Tea | Node.js/npm (distribusi), internet (unduh pertama kali) |
| SSH (remote) | `ssh habibiahmada.dev` | EC2 → Wish → Go TUI | EC2 harus hidup, SSH client |

## Tech Stack (Terminal)

| Layer | Teknologi | Fungsi |
|-------|-----------|---------|
| Runtime | Go | Aplikasi TUI |
| TUI engine | Bubble Tea | Terminal UI |
| Styling | Lip Gloss | Terminal styling |
| SSH | Wish | SSH application server |
| Distribusi | npm | `npx habibiahmada` (wrapper, bukan runtime TUI) |
| Infrastructure | AWS EC2 | SSH portfolio + Telegram Agent |

**Prinsip:** npx adalah distribution layer, bukan teknologi utama TUI. TUI dibuat dengan Go.

## Struktur Proyek (Target)

```
habibiahmada-terminal/
├── cmd/
│   ├── portfolio/main.go    # entry point local / npx
│   └── ssh/main.go          # entry point SSH server
├── internal/
│   ├── tui/                 # screens & navigation
│   ├── components/          # header, sidebar, footer, dll.
│   ├── styles/              # Lip Gloss styles
│   └── data/                # portfolio data (bundled)
├── npm/                     # npm package wrapper + binaries
├── scripts/                 # build & release
├── .github/workflows/       # CI/CD
├── go.mod
└── Makefile
```

Versi awal disarankan lebih sederhana — lihat [docs/folder-structure.md](docs/folder-structure.md).

## Menu TUI

- About
- Projects
- Skills
- Experience
- Certificates
- Contact

Navigasi: `↑↓` Navigate · `Enter` Select · `←` Back · `Q` Quit

## Data Portfolio (v1)

Untuk versi pertama, data TUI **dibundel di Go binary** (Opsi A):

- Sederhana, cepat, bisa offline setelah binary tersedia
- SSH tidak perlu akses Supabase
- Perubahan portfolio memerlukan release binary baru

Website tetap mengambil data dari Supabase. CLI/SSH tidak wajib melalui Supabase.

## Proyek Contoh (Portfolio Data)

- Renshuu — Japanese Learning Platform
- SmartFarm AI — AI-powered agriculture platform
- CultureConnect — Cultural exchange platform
- Spacelab

## CI/CD

| Target | Alur |
|--------|------|
| CLI | Git push → GitHub Actions → Go build → Release → npm publish |
| SSH | Git push → GitHub Actions → Deploy → EC2 |

## Dokumentasi

| File | Isi |
|------|-----|
| [AGENTS.md](AGENTS.md) | Panduan untuk AI coding assistant |
| [docs/index.md](docs/index.md) | Indeks dokumentasi |
| [docs/architecture.md](docs/architecture.md) | Arsitektur lengkap |
| [docs/tech-stack.md](docs/tech-stack.md) | Pembagian tech stack |
| [docs/folder-structure.md](docs/folder-structure.md) | Struktur folder & evolusi |
| [docs/npx-vs-ssh.md](docs/npx-vs-ssh.md) | Perbandingan local vs remote |
| [docs/data-strategy.md](docs/data-strategy.md) | Strategi data TUI vs website |
| [docs/deployment.md](docs/deployment.md) | Deployment EC2, npm, CI/CD |

## Mulai Development

Proyek masih dalam tahap persiapan dokumentasi. Implementasi dimulai dari struktur minimal:

```
habibiahmada-terminal/
├── cmd/portfolio/main.go
├── cmd/ssh/main.go
├── internal/tui/
├── internal/data/portfolio.go
├── npm/
├── go.mod
└── Makefile
```

## Lisensi

TBD
