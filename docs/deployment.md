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
├── habibiahmada-linux-x64
├── habibiahmada-linux-arm64
├── habibiahmada-darwin-x64
├── habibiahmada-darwin-arm64
└── habibiahmada-win-x64.exe
```

`scripts/release.sh` memasukkan binary dari `dist/` ke `npm/bin/`.

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
│   ├── habibiahmada-ssh
│   ├── portfolio-ssh.service   # salinan unit (sumber: deploy/portfolio-ssh.service)
│   └── .ssh/
│       └── term_info_ed25519   # Wish SSH host key (di-generate otomatis saat deploy)
│
└── systemd
    └── portfolio-ssh.service
```

Systemd menjalankan SSH application (Wish + Go TUI).

### Artefak Deployment

| File | Isi |
|------|-----|
| `.github/workflows/deploy.yml` | CI: build `cmd/ssh` → scp ke EC2 → install unit → restart (+ cek port 2222) |
| `deploy/portfolio-ssh.service` | systemd unit (WorkingDirectory `/opt/habibiahmada`) |
| `scripts/deploy-ssh.sh` | Deploy manual dari mesin lokal (tanpa GitHub Actions) |

### CI Secrets yang dibutuhkan (`deploy.yml`)

| Secret | Deskripsi |
|--------|-----------|
| `EC2_HOST` | IP / hostname EC2 |
| `EC2_USER` | SSH username (mis. `ubuntu`) |
| `EC2_SSH_KEY` | Isi private key SSH (pastikan public key ada di `~/.ssh/authorized_keys` server) |
| `EC2_PORT` | (opsional) port SSH, default `22` |

Deploy dipicu otomatis saat push ke `main` yang menyentuh kode TUI/core, atau manual via **Actions → Deploy SSH Server → Run workflow**.

### Host Key Wish

Server Wish butuh key host di `.ssh/term_info_ed25519` (relatif ke CWD = `/opt/habibiahmada`).
`scripts/deploy-ssh.sh` meng-generate-nya otomatis bila belum ada:

```bash
ssh-keygen -t ed25519 -f .ssh/term_info_ed25519 -N "" -C "wish-portfolio@habibiahmada.dev"
```

> Sama halnya dengan SSH server Linux, key Wish **jangan pernah di-commit ke git** (sudah masuk `.gitignore` melalui `/.ssh/`). Regenerate sesekali sebagai rotasi key.

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
