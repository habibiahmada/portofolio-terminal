# Tech Stack

Pembagian teknologi per layer proyek.

## Terminal Portfolio (Repositori Ini)

| Layer | Teknologi | Fungsi |
|-------|-----------|---------|
| Runtime | Go | Aplikasi TUI |
| TUI engine | Bubble Tea | Terminal UI framework |
| Styling | Lip Gloss | Terminal styling |
| SSH server | Wish | SSH application server |
| Distribusi | npm | `npx habibiahmada` sebagai wrapper binary Go |
| Build/Release | GitHub Actions | Cross-compile & npm publish |
| Infrastructure | AWS EC2 | Host SSH portfolio |

### Bubble Tea — Screen Map

```
             Bubble Tea
                  │
       ┌──────────┼──────────┐
       │          │          │
       ▼          ▼          ▼
     Home      Projects     Skills
       │          │          │
       └──────────┼──────────┘
                  │
               Contact
```

### Lip Gloss — Style Variables (Target)

Definisi di `internal/styles/styles.go`:

- `TitleStyle`
- `SubtitleStyle`
- `SelectedStyle`
- `NormalStyle`
- `BorderStyle`
- `MutedStyle`
- `SuccessStyle`
- `LinkStyle`

## Website Portfolio (Repo Terpisah)

| Layer | Teknologi | Fungsi |
|-------|-----------|---------|
| Domain | Cloudflare | DNS, TLS, edge/security |
| Web | Next.js + React + TypeScript | Web portfolio |
| Styling | Tailwind CSS | UI styling |
| Hosting | Vercel | Deployment web |
| Database | Supabase PostgreSQL | Portfolio/content data |

## Supabase — Schema Portfolio (Website)

Tabel yang direncanakan:

| Tabel | Field contoh |
|-------|--------------|
| projects | id, slug, name, description, thumbnail, github_url, live_url, featured, created_at |
| skills | — |
| experiences | — |
| certificates | — |
| articles | — |
| social_links | — |
| profile | — |

## Automation (EC2)

| Komponen | Fungsi |
|----------|--------|
| Telegram Agent | Finance, News, Blog Automation |

## Source Control

| Tool | Fungsi |
|------|--------|
| GitHub | Repository pusat |
| GitHub Actions | Build, release, deployment |

## Catatan npx vs Node.js

Stack awal percakapan menyebut Node.js/TypeScript/Commander/Inquirer/Chalk untuk pendekatan npm-native. **Keputusan final:** TUI dibuat dengan Go; npm hanya sebagai distribution channel untuk Go binary.

Pendekatan alternatif yang dibahas tapi tidak dipilih sebagai core:

- Node.js CLI dengan Commander.js, @inquirer/prompts, Chalk, Figlet, Ora
- SSH portfolio sebagai satu-satunya interface

Keputusan: **npx (local) + SSH (remote)**, keduanya menjalankan Go TUI yang sama.

## Ilustrasi TUI

Visual identity dan ASCII art: [tui-illustration.md](tui-illustration.md).

Stack visual:

| Layer | Teknologi |
|-------|-----------|
| Typography besar | go-figure (FIGlet) |
| Signature art | Custom ASCII (wide / compact / mini) |
| Animasi startup | Bubble Tea tick + Bubbles spinner |
| Layout ilustrasi | Lip Gloss `Place`, `MaxWidth` |
