# Dokumentasi habibiahmada-terminal

Indeks dokumentasi resmi proyek habibiahmada-terminal.

## Mulai Di Sini

| Dokumen | Untuk siapa | Isi |
|---------|-------------|-----|
| [user-guide.md](user-guide.md) | **Pengguna / recruiter** | Cara akses portfolio di terminal (`npx`, `ssh`) |
| [getting-started.md](getting-started.md) | **Developer baru** | Clone, setup, jalankan TUI pertama kali |
| [development-guide.md](development-guide.md) | **Kontributor** | Workflow harian, edit data/TUI, build, release |
| [task-list.md](task-list.md) | **Semua** | Roadmap, task list, progress proyek |

## Arsitektur & Desain

| Dokumen | Topik |
|---------|-------|
| [architecture.md](architecture.md) | Arsitektur platform, layer, CI/CD, isolation |
| [tech-stack.md](tech-stack.md) | Teknologi per layer |
| [folder-structure.md](folder-structure.md) | Struktur repo, fase 1 vs target |
| [npx-vs-ssh.md](npx-vs-ssh.md) | Perbandingan local vs remote experience |
| [data-strategy.md](data-strategy.md) | Bundled data v1, Supabase untuk website |

## Deployment & Ops

| Dokumen | Topik |
|---------|-------|
| [deployment.md](deployment.md) | npm publish, EC2, GitHub Actions |

## Dokumen Lain

| Dokumen | Isi |
|---------|-----|
| [README.md](../README.md) | Overview proyek |
| [AGENTS.md](../AGENTS.md) | Panduan AI coding assistant |

## Quick Reference

```bash
# Pengguna — akses portfolio
npx habibiahmada
ssh habibiahmada.dev

# Developer — jalankan lokal
go run ./cmd/portfolio
make dev

# Developer — build & test
make build
make lint
```
