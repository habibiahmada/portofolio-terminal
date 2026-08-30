---
name: go-code-quality
description: >-
  Go coding standards for habibiahmada-terminal. Use when writing or reviewing
  Go code, cmd/ entry points, internal packages, build scripts, or Go module setup.
---

# Go Code Quality — habibiahmada-terminal

Standar kode Go untuk proyek terminal portfolio ini.

## Arsitektur Package

```
cmd/portfolio/   → thin main, delegate ke internal/tui
cmd/ssh/         → thin main, Wish server, same internal/tui
internal/tui/    → Bubble Tea models & screens
internal/data/   → portfolio structs & data (no UI imports)
internal/components/ → reusable UI pieces
internal/styles/ → Lip Gloss only, no business logic
```

**Dependency rule:** `data` tidak import `tui`. `styles` tidak import `tui`.

## Entry Points

`cmd/*/main.go` hanya:

1. Parse flags (minimal)
2. Load data (dari `internal/data`)
3. Start TUI atau SSH server

Jangan taruh screen logic di `main.go`.

## Naming & Files

- Satu screen = satu file di `internal/tui/` (`projects.go`, `about.go`)
- Exported types hanya jika dipakai lintas package
- Package name = folder name (`tui`, `data`, `styles`)

## Error Handling

- Return error dari functions, handle di caller
- `main.go`: log fatal hanya untuk startup failure
- Jangan panic kecuali programmer error di init

## Data Layer

```go
// internal/data/projects.go
type Project struct {
    Name        string
    Description string
    Stack       []string
    GitHub      string
    Live        string
}
```

- Data sebagai variabel/func getter, bukan hard-code di View()
- v1: bundled static data, no HTTP client ke Supabase

## Build & Release

Cross-compile targets:

- linux/amd64, linux/arm64
- darwin/amd64, darwin/arm64
- windows/amd64

Makefile atau `scripts/build.sh` — satu command build all.

## Testing

- Unit test untuk `internal/data` (validasi struktur)
- Bubble Tea: pertimbangkan `teatest` untuk screen behavior
- Test file: `*_test.go` di package yang sama

## Checklist Review

- [ ] Satu TUI core dipakai portfolio + ssh entry
- [ ] Tidak ada duplikasi screen logic antar cmd
- [ ] Data terpisah dari View/Update
- [ ] Tidak ada dependency Supabase/API di v1
- [ ] Import cycle tidak ada
- [ ] `go vet ./...` dan `go test ./...` lulus

## Anti-Patterns

- Logic TUI di npm wrapper (npm hanya spawn binary)
- Dua implementasi TUI berbeda untuk npx vs SSH
- Hard-code portfolio text di `View()` methods
- Over-abstract helper 1-2 baris
