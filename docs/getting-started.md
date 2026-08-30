# Getting Started

Panduan memulai proyek **habibiahmada-terminal** — dari clone repo sampai TUI berjalan di terminal Anda.

## Apa Itu Proyek Ini?

Satu aplikasi TUI (Terminal User Interface) portfolio interaktif yang dibuat dengan **Go + Bubble Tea**. Proyek ini dapat diakses melalui dua jalur:

| Jalur | Perintah | Keterangan |
|-------|----------|------------|
| **Local** | `npx habibiahmada` | Binary Go dijalankan di laptop Anda |
| **Remote** | `ssh habibiahmada.dev` | TUI yang sama berjalan di server EC2 |

Kedua jalur memakai **TUI core yang sama** (`internal/tui/`). Perbedaannya hanya di transport layer.

## Prasyarat

### Untuk menjalankan TUI (pengguna)

| Jalur | Yang dibutuhkan |
|-------|-----------------|
| npx | Node.js ≥ 14, terminal modern |
| SSH | SSH client (sudah ada di Linux/macOS, Windows 10+) |

### Untuk development (kontributor)

| Tool | Versi minimum | Cek versi |
|------|---------------|-----------|
| Go | 1.22+ | `go version` |
| Git | — | `git --version` |
| Make | opsional | `make --version` |
| Node.js | ≥ 14 (untuk tes npm wrapper) | `node --version` |

## Clone & Setup

```bash
# 1. Clone repository
git clone https://github.com/habibiahmada/habibiahmada-terminal.git
cd habibiahmada-terminal

# 2. Unduh dependensi Go
go mod download

# 3. Verifikasi build berhasil
go build -o habibiahmada ./cmd/portfolio
```

## Jalankan TUI Pertama Kali

### Cara tercepat — development mode

```bash
go run ./cmd/portfolio
```

Atau jika Make tersedia:

```bash
make dev
```

TUI akan langsung terbuka di terminal Anda.

### Build binary lokal

```bash
go build -o habibiahmada ./cmd/portfolio
./habibiahmada
```

### Tes npm wrapper (lokal)

```bash
# Build binary untuk platform Anda
bash scripts/build.sh

# Salin ke npm/bin dengan nama yang benar
mkdir -p npm/bin
cp dist/habibiahmada-linux-amd64 npm/bin/habibiahmada-linux-x64   # Linux x64
# atau sesuaikan dengan OS/arch Anda

# Jalankan via npm
node npm/index.js
```

## Navigasi TUI

Setelah TUI terbuka:

```
┌───────────────────────────────────────────────┐
│ HABIBI AHMAD AZIZ                             │  ← Header
├──────────────┬────────────────────────────────┤
│ > About      │                                │
│   Projects   │       Konten halaman           │  ← Sidebar + Content
│   Skills     │                                │
│   Experience │                                │
│   Certificates                                │
│   Contact    │                                │
├──────────────┴────────────────────────────────┤
│ ↑↓ Navigate  Enter Select  ← Back  Q Quit    │  ← Footer
└───────────────────────────────────────────────┘
```

| Tombol | Aksi |
|--------|------|
| `↑` / `↓` atau `k` / `j` | Navigasi menu |
| `Enter` / `Space` | Pilih / masuk halaman |
| `←` / `Esc` | Kembali ke Home (dari Home = quit) |
| `Q` / `Ctrl+C` | Keluar |

## Struktur Proyek (Fase Saat Ini)

```
habibiahmada-terminal/
├── cmd/
│   ├── portfolio/main.go     # Entry point local / npx
│   └── ssh/main.go            # Entry point SSH server
├── internal/
│   ├── tui/
│   │   ├── app.go             # TUI core (semua screen)
│   │   └── keymap.go          # Keyboard shortcuts
│   ├── styles/styles.go       # Lip Gloss styles
│   └── data/portfolio.go      # Data portfolio (bundled)
├── npm/                       # npm distribution wrapper
├── scripts/                   # Build scripts
├── docs/                      # Dokumentasi
├── Makefile
└── go.mod
```

## Perintah Makefile

| Perintah | Fungsi |
|----------|--------|
| `make dev` | Jalankan TUI lokal (`go run ./cmd/portfolio`) |
| `make dev-ssh` | Jalankan SSH server lokal (port 2222) |
| `make build` | Cross-compile semua platform → `dist/` |
| `make build-ssh` | Build binary SSH server → `habibiahmada-ssh` |
| `make test` | Jalankan `go test ./...` |
| `make vet` | Static analysis |
| `make fmt` | Format kode Go |
| `make lint` | vet + fmt + test |
| `make clean` | Hapus artefak build |

## Langkah Berikutnya

| Tujuan | Dokumen |
|--------|---------|
| Akses portfolio sebagai pengguna | [user-guide.md](user-guide.md) |
| Development sehari-hari | [development-guide.md](development-guide.md) |
| Lihat task & progress | [task-list.md](task-list.md) |
| Arsitektur lengkap | [architecture.md](architecture.md) |
| Deploy ke npm / EC2 | [deployment.md](deployment.md) |
