# Arsitektur

Dokumen ini merangkum arsitektur platform portfolio Habibi Ahmad Aziz.

## Visi

Satu ekosistem portfolio dengan tiga interface:

```
                         HABIBI AHMAD AZIZ
                         Developer Platform
                                │
             ┌──────────────────┼──────────────────┐
             │                  │                  │
             ▼                  ▼                  ▼
          WEBSITE              CLI                SSH
             │                  │                  │
     habibiahmada.dev    npx habibiahmada    ssh habibiahmada.dev
             │                  │                  │
             ▼                  │                  ▼
          Cloudflare            │                AWS EC2
             │                  │                  │
             ▼                  │                  │
           Vercel               │                Go TUI
             │                  │                  │
          Next.js               │                  │
             │                  ▼                  │
             │             Go Binary               │
             │                  │                  │
             └──────────────────┼──────────────────┘
                                │
                                ▼
                           Supabase (website)
                                │
                                ▼
                         PostgreSQL
```

**Prinsip:** Website, CLI, dan SSH bukan tiga portfolio berbeda. Mereka adalah tiga presentation layer dengan identitas yang sama.

## Layer Pengguna

| Tipe user | Akses | Pengalaman |
|-----------|-------|------------|
| Pengguna biasa | `https://habibiahmada.dev` | Website portfolio modern |
| Developer | `npx habibiahmada` | TUI lokal |
| Developer/Linux enthusiast | `ssh habibiahmada.dev` | TUI melalui SSH |

## Website Stack

```
Browser → Cloudflare (DNS, TLS, proxy) → Vercel → Next.js → Supabase → PostgreSQL
```

Cloudflare berfungsi sebagai edge/network layer, bukan tempat utama aplikasi.

Website menangani: Home, About, Projects, Experience, Skills, Certificates, Blog, Contact.

Repository target: `portfolio-web/` (terpisah dari terminal).

## Terminal Stack

```
                    ┌───────────────┐
                    │ Portfolio Data│ (bundled di Go, v1)
                    └───────┬───────┘
                            │
                            ▼
                    ┌───────────────┐
                    │   TUI Core    │
                    │  Bubble Tea   │
                    └───────┬───────┘
                            │
             ┌──────────────┴──────────────┐
             │                             │
             ▼                             ▼
       Local Runtime                SSH Runtime
       (cmd/portfolio)             (cmd/ssh + Wish)
             │                             │
             ▼                             ▼
          npx                           EC2
```

## Data Flow — Koreksi Arsitektur

CLI tidak perlu melalui Supabase di v1:

```
Website ───────► Supabase
                    ▲
                    │
Telegram Agent ─────┘

CLI ───────────► Local portfolio data (bundled)
SSH ───────────► Local/server portfolio data (bundled)
```

Opsi sinkronisasi penuh via Portfolio API ke Supabase ada, tetapi untuk v1 dipilih **data statis dibundel ke binary/package**.

## EC2 Workload

```
AWS EC2
├── Telegram Agent
│   ├── Finance
│   ├── News
│   └── Blog Automation
└── SSH Portfolio
    └── Go TUI (Wish)
```

Scheduler EC2 menyalakan instance 3× sehari untuk Telegram Agent. SSH portfolio ikut tersedia saat EC2 ON.

Migrasi ke VPS tidak mengubah konsep aplikasi Go TUI.

## Blog Automation Flow

```
Telegram Agent → Generate Article → Supabase/CMS → Next.js → habibiahmada.dev/blog
```

## GitHub & CI/CD

```
                         GitHub
                           │
          ┌────────────────┼────────────────┐
          │                │                │
          ▼                ▼                ▼
    portfolio-web    terminal-portfolio   agent
          │                │                │
          ▼                ▼                ▼
        Vercel            npm             AWS EC2
```

| Repo | Pipeline |
|------|----------|
| portfolio-web | Git push → GitHub → Vercel → Production |
| terminal-portfolio | Git push → GitHub Actions → Go build → Release → npm publish |
| agent | Git push → GitHub Actions → Deploy → EC2 |

## Dependency Isolation

```
                    ┌──────────────┐
                    │   Supabase   │
                    └──────▲───────┘
                           │ optional data
             ┌─────────────┴─────────────┐
             │                           │
        ┌────▼─────┐                ┌────▼─────┐
        │ Next.js  │                │ Go TUI   │
        └────▲─────┘                └────▲─────┘
             │                           │
          Vercel                   npx / SSH
             │                           │
             ▼                           ▼
         Browser                       User
```

## Satu TUI, Dua Transport

```
              TUI CORE
             /        \
            /          \
         npx            SSH
```

Jangan membuat `npx portfolio → TUI A` dan `ssh portfolio → TUI B`.
