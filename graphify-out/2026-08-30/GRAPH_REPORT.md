# Graph Report - terminal  (2026-08-30)

## Corpus Check
- 23 files · ~8,048 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 239 nodes · 290 edges · 21 communities (17 shown, 4 thin omitted)
- Extraction: 98% EXTRACTED · 2% INFERRED · 0% AMBIGUOUS · INFERRED: 5 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `98813ef0`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- AGENTS.md — habibiahmada-terminal
- Deployment
- Bubble Tea TUI Design
- npx vs SSH
- Peran Setiap Direktori
- Go Code Quality — habibiahmada-terminal
- Arsitektur
- Strategi Data
- habibiahmada-terminal
- index.md
- Graphify Workflow
- Tech Stack
- App
- package.json
- .handleKey
- index.js
- teaHandler
- build.sh
- build-ssh.sh
- github.com/habibiahmada/habibiahmada-terminal

## God Nodes (most connected - your core abstractions)
1. `App` - 29 edges
2. `AGENTS.md — habibiahmada-terminal` - 14 edges
3. `Bubble Tea TUI Design` - 13 edges
4. `habibiahmada-terminal` - 11 edges
5. `Arsitektur` - 11 edges
6. `New()` - 10 edges
7. `Go Code Quality — habibiahmada-terminal` - 10 edges
8. `Strategi Data` - 9 edges
9. `Graphify Workflow` - 8 edges
10. `Deployment` - 8 edges

## Surprising Connections (you probably didn't know these)
- `main()` --calls--> `New()`  [EXTRACTED]
  cmd/portofolio/main.go → internal/tui/app.go
- `teaHandler()` --calls--> `New()`  [EXTRACTED]
  cmd/ssh/main.go → internal/tui/app.go
- `App` --references--> `Profile`  [EXTRACTED]
  internal/tui/app.go → internal/data/portfolio.go
- `App` --references--> `Project`  [EXTRACTED]
  internal/tui/app.go → internal/data/portfolio.go
- `App` --references--> `Skill`  [EXTRACTED]
  internal/tui/app.go → internal/data/portfolio.go

## Import Cycles
- None detected.

## Communities (21 total, 4 thin omitted)

### Community 0 - "AGENTS.md — habibiahmada-terminal"
Cohesion: 0.12
Nodes (16): AGENTS.md — habibiahmada-terminal, Aturan Arsitektur, Data Model (Contoh), Evolusi Struktur, Failure Isolation, Graphify, Jangan, Konteks Proyek (+8 more)

### Community 1 - "Deployment"
Cohesion: 0.12
Nodes (16): Build, Build Output, CI/CD — GitHub Actions, CLI — npm Distribution, Cloudflare & Website (Referensi), Deploy SSH (deploy.yml), Deployment, EC2 — Shared Workload (+8 more)

### Community 2 - "Bubble Tea TUI Design"
Cohesion: 0.14
Nodes (13): Anti-Patterns, Best Practices (Charm / Community), Bubble Tea Architecture, Bubble Tea TUI Design, Layout Pattern, Lip Gloss Layout, Navigation, Referensi Proyek (+5 more)

### Community 3 - "npx vs SSH"
Cohesion: 0.14
Nodes (14): Alur, Alur, Internet, Kelebihan, Kelemahan, Konsep Bersama, npx Bukan Runtime TUI, npx — Local Experience (+6 more)

### Community 4 - "Peran Setiap Direktori"
Cohesion: 0.17
Nodes (12): `cmd/`, Dependency Flow, EC2 — Bukan di Repo, `internal/components/`, `internal/data/`, `internal/tui/`, `npm/`, Peran Setiap Direktori (+4 more)

### Community 5 - "Go Code Quality — habibiahmada-terminal"
Cohesion: 0.18
Nodes (10): Anti-Patterns, Arsitektur Package, Build & Release, Checklist Review, Data Layer, Entry Points, Error Handling, Go Code Quality — habibiahmada-terminal (+2 more)

### Community 6 - "Arsitektur"
Cohesion: 0.18
Nodes (11): Arsitektur, Blog Automation Flow, Data Flow — Koreksi Arsitektur, Dependency Isolation, EC2 Workload, GitHub & CI/CD, Layer Pengguna, Satu TUI, Dua Transport (+3 more)

### Community 7 - "Strategi Data"
Cohesion: 0.18
Nodes (11): Alternatif Source of Truth (Dibahas, Bukan v1 Terminal), Data Flow per Interface, Evolusi Masa Depan, Failure Isolation, Konteks, Model Data Go (Contoh), Opsi A — Data Dibundel di Go (Dipilih untuk v1), Opsi B — TUI Fetch dari API (Tidak dipilih v1) (+3 more)

### Community 8 - "habibiahmada-terminal"
Cohesion: 0.18
Nodes (11): CI/CD, Data Portfolio (v1), Dokumentasi, Ekosistem Portfolio, habibiahmada-terminal, Lisensi, Menu TUI, Mulai Development (+3 more)

### Community 10 - "Graphify Workflow"
Cohesion: 0.22
Nodes (8): Catatan, Cursor Install/Uinstall, Graphify Workflow, Hook (opsional), Prioritas, Sebelum Explore Codebase, Setup (sudah dikonfigurasi), Update Graph

### Community 11 - "Tech Stack"
Cohesion: 0.22
Nodes (9): Automation (EC2), Bubble Tea — Screen Map, Catatan npx vs Node.js, Lip Gloss — Style Variables (Target), Source Control, Supabase — Schema Portfolio (Website), Tech Stack, Terminal Portfolio (Repositori Ini) (+1 more)

### Community 12 - "App"
Cohesion: 0.12
Nodes (17): main(), GetCertificates(), GetExperiences(), GetProfile(), GetProjects(), GetSkills(), GetSocials(), Certificate (+9 more)

### Community 13 - "package.json"
Cohesion: 0.08
Nodes (25): author, bin, habibiahmada, description, engines, node, files, keywords (+17 more)

### Community 14 - ".handleKey"
Cohesion: 0.20
Nodes (10): github.com/charmbracelet/bubbletea.Cmd, github.com/charmbracelet/bubbletea.KeyMsg, github.com/charmbracelet/bubbletea.Model, github.com/charmbracelet/bubbletea.Msg, tea.KeyMsg, isBack(), isNavigateDown(), isNavigateUp() (+2 more)

### Community 15 - "index.js"
Cohesion: 0.38
Nodes (6): detectPlatform(), { execFileSync }, fs, getBinaryPath(), main(), path

### Community 16 - "teaHandler"
Cohesion: 0.40
Nodes (3): teaHandler(), github.com/charmbracelet/bubbletea.ProgramOption, github.com/charmbracelet/ssh.Session

## Knowledge Gaps
- **137 isolated node(s):** `github.com/habibiahmada/habibiahmada-terminal`, `{ execFileSync }`, `path`, `fs`, `name` (+132 more)
  These have ≤1 connection - possible missing edges or undocumented components.
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `AGENTS.md — habibiahmada-terminal` connect `AGENTS.md — habibiahmada-terminal` to `index.md`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Why does `Deployment` connect `Deployment` to `index.md`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **Why does `npx vs SSH` connect `npx vs SSH` to `index.md`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **What connects `github.com/habibiahmada/habibiahmada-terminal`, `{ execFileSync }`, `path` to the rest of the system?**
  _137 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `AGENTS.md — habibiahmada-terminal` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Deployment` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._
- **Should `Bubble Tea TUI Design` be split into smaller, more focused modules?**
  _Cohesion score 0.14285714285714285 - nodes in this community are weakly interconnected._