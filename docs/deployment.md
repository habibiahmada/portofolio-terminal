# Deployment

Panduan deployment CLI (npm) dan SSH (EC2).

## CLI — npm Distribution

### npm Package Structure

```
npm/
├── package.json
├── index.js
└── bin/
    ├── habibiahmada-linux-x64
    ├── habibiahmada-linux-arm64
    ├── habibiahmada-darwin-x64
    ├── habibiahmada-darwin-arm64
    └── habibiahmada-win-x64.exe
```

### Flow

```
npx habibiahmada
    → npm package
    → deteksi OS + architecture
    → jalankan Go binary
    → Bubble Tea → TUI
```

### Build Output

`scripts/build.sh` menghasilkan:

```
dist/
├── linux-amd64
├── linux-arm64
├── darwin-amd64
├── darwin-arm64
└── windows-amd64.exe
```

`scripts/release.sh` memasukkan binary ke npm package.

## CI/CD — GitHub Actions

### Release CLI (release.yml)

Trigger: git tag (contoh `v1.0.0`)

```
Git tag
   │
   ▼
GitHub Actions
   ├── Go test
   ├── Go build
   ├── Build Linux
   ├── Build macOS
   ├── Build Windows
   │
   ▼
npm publish
```

### Deploy SSH (deploy.yml)

```
Git push → GitHub Actions → Deploy → EC2
```

## SSH — AWS EC2

### Build

```bash
go build ./cmd/ssh
# output: habibiahmada-ssh
```

### Layout di Server

```
AWS EC2
│
├── /opt/habibiahmada/
│   └── habibiahmada-ssh
│
└── systemd
    └── portfolio-ssh.service
```

Systemd menjalankan SSH application (Wish + Go TUI).

### Wish Stack

```
User → ssh habibiahmada.dev → EC2 → SSH Server → Wish → Bubble Tea → Go TUI
```

TUI identik dengan versi npx.

## EC2 — Shared Workload

EC2 menjalankan dua workload:

| Workload | Fungsi |
|----------|--------|
| Telegram Agent | Finance, News, Blog Automation |
| SSH Portfolio | Go TUI via Wish |

Scheduler: EC2 dinyalakan 3× sehari untuk Telegram Agent. SSH portfolio tersedia saat EC2 ON.

## Migrasi EC2 → VPS

```
AWS EC2 → migration → VPS
                         ├── Telegram Agent
                         └── SSH Portfolio
```

Konsep aplikasi Go TUI tidak berubah.

## Cloudflare & Website (Referensi)

Website deployment terpisah:

```
Browser → Cloudflare → Vercel → Next.js → Supabase
```

Domain: `habibiahmada.dev`

DNS di Cloudflare diarahkan ke Vercel.

## Recruiter-Facing Commands

| Channel | Command |
|---------|---------|
| Website | `https://habibiahmada.dev` |
| Terminal (local) | `npx habibiahmada` |
| Terminal (remote) | `ssh habibiahmada.dev` |

Portfolio project ini sendiri dapat menjadi entry portfolio: *"Interactive CLI Portfolio built with Go/Bubble Tea"*.
