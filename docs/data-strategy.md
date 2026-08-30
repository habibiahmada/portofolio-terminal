# Strategi Data

Keputusan arsitektur data untuk website vs terminal portfolio.

## Konteks

Ekosistem portfolio mencakup:

- **Website** (`habibiahmada.dev`) — Next.js di Vercel, data dari Supabase
- **CLI** (`npx habibiahmada`) — Go TUI, data lokal
- **SSH** (`ssh ssh.habibiahmada.dev`) — Go TUI yang sama, data di server

## Opsi yang Dibahas

### Opsi A — Data Dibundel di Go (Dipilih untuk v1)

```
Go binary
├── TUI
└── Portfolio data
```

| Kelebihan | Kekurangan |
|-----------|------------|
| Sangat sederhana | Setiap perubahan portfolio perlu release binary baru |
| Cepat | |
| Tidak butuh API | |
| Bisa offline setelah binary tersedia | |
| SSH tidak perlu akses Supabase | |

### Opsi B — TUI Fetch dari API (Tidak dipilih v1)

```
Go TUI → HTTPS → Vercel API → Supabase
```

| Kelebihan | Kekurangan |
|-----------|------------|
| Update portfolio langsung muncul | Butuh internet |
| Tidak perlu rebuild untuk setiap perubahan konten | API bisa down |
| | Arsitektur lebih kompleks |
| | SSH TUI bergantung API |

**Keputusan v1:** Opsi A. Portfolio tidak berubah setiap jam; release baru saat menambah project bukan masalah.

## Data Flow per Interface

```
Website ───────► Supabase ◄──── Telegram Agent
                    ▲
                    │
              (blog automation)

CLI ───────────► Bundled portfolio data (internal/data/)
SSH ───────────► Bundled portfolio data (sama)
```

CLI tidak perlu melalui Supabase.

## Alternatif Source of Truth (Dibahas, Bukan v1 Terminal)

Konsep awal single source dengan JSON files:

```
portfolio/
├── data/
│   ├── profile.json
│   ├── projects.json
│   ├── skills.json
│   └── experience.json
├── web/
└── cli/
```

Website dan CLI membaca `data/*.json`. Untuk terminal v1, data di-embed sebagai Go structs/constants di `internal/data/`.

## Model Data Go (Contoh)

```go
type Project struct {
    Name        string
    Description string
    Stack       []string
    GitHub      string
    Live        string
}

var Projects = []Project{
    {
        Name:        "Renshuu",
        Description: "Japanese learning platform",
        Stack:       []string{"Laravel", "React", "Inertia"},
    },
    // SmartFarm AI, CultureConnect, Spacelab, ...
}
```

## Supabase Schema (Website — Referensi)

| Tabel | Contoh field |
|-------|--------------|
| projects | id, slug, name, description, thumbnail, github_url, live_url, featured, created_at |
| skills | — |
| experiences | — |
| certificates | — |
| articles | — |
| social_links | — |
| profile | — |

## Failure Isolation

| Supabase down | Website mungkin terdampak; CLI & SSH tetap jalan |
| Vercel down | Website down; CLI & SSH tetap jalan |
| EC2 down | Website & CLI tetap jalan; SSH down |

## Evolusi Masa Depan

Sinkronisasi penuh via Portfolio API ke Supabase dimungkinkan:

```
                    Supabase
                       │
                  Portfolio API
                       │
          ┌────────────┼────────────┐
          │            │            │
       Next.js        CLI          SSH
```

Belum diimplementasi di v1 terminal.
