# npx vs SSH

Perbandingan dua jalur akses terminal portfolio.

## Ringkasan

| | `npx habibiahmada` | `ssh habibiahmada.dev` |
|---|-------------------|------------------------|
| Label UX | **LOCAL EXPERIENCE** | **REMOTE EXPERIENCE** |
| TUI berjalan di | Laptop user | AWS EC2 |
| Perlu Node.js | Ya (untuk npx) | Tidak |
| Perlu SSH client | Tidak | Ya |
| Download aplikasi | Ya (biasanya pertama kali) | Tidak |
| EC2 diperlukan | Tidak | Ya |
| EC2 harus hidup | Tidak | Ya |
| CPU/RAM | Laptop user | EC2 |
| Internet | Untuk mengambil package | Untuk koneksi SSH |
| Offline setelah tersedia | Bisa, secara teknis | Tidak |
| Deployment | npm | EC2 |
| Maintenance server | Tidak | Ya |
| Cocok untuk | Developer umum | Technical users |

## npx — Local Experience

### Alur

```
Laptop pengguna
│ npx habibiahmada
├── Node.js / npm
├── download package (jika belum cache)
├── extract package
└── menjalankan Go binary
          │
          ▼
       Bubble Tea → TUI
```

### Internet

- **Pertama kali:** ya — npm registry → download package → laptop → run
- **Setelah cache:** bisa tanpa download ulang, tergantung npm/npx
- **Jangan desain** dengan asumsi npx selalu offline

### Resource Server

Setelah download, **EC2 tidak digunakan**. EC2 bisa STOPPED dan npx tetap jalan.

```
Recruiter → npm Registry (download sekali) → Laptop → Go TUI
```

Bukan: Recruiter → Vercel → EC2 → TUI.

### npx Bukan Runtime TUI

```
npx → Distribution → Go binary → Bubble Tea → TUI
```

## SSH — Remote Experience

### Alur

```
Laptop (SSH client only)
  │ keyboard input
  ▼
SSH connection
  ▼
AWS EC2
  ├── SSH Server (Wish)
  └── Go TUI → Bubble Tea
```

Input (`↑`, `↓`, `Enter`, `q`) dikirim ke EC2. EC2 memproses dan kirim balik perubahan tampilan.

### Kelebihan

- Tidak perlu Node.js, npm, go install, atau git clone
- Hanya butuh SSH client (Linux/macOS/Windows modern)
- Binary (~20 MB) menggunakan CPU/RAM EC2, bukan laptop user

### Kelemahan

- EC2 harus hidup — jika STOPPED, SSH gagal
- Perlu maintenance server

## Konsep Bersama

npx dan SSH hanya dua jalur berbeda untuk menjalankan **TUI yang sama**:

```
             npx                    SSH
              │                      │
              ▼                      ▼
        Distribution          Connection layer
              │                      │
              └──────────┬───────────┘
                         ▼
                    Go binary
                         │
                         ▼
                    Bubble Tea
                         │
                         ▼
                        TUI
```

## Strategi EC2

SSH portfolio menumpang pada EC2 yang sudah ada untuk Telegram Agent:

```
AWS EC2
├── Telegram Agent (3× daily scheduler)
└── SSH Portfolio (available when EC2 ON)
```

npx tetap tersedia tanpa bergantung EC2.

## Tampilan Website (CTA)

```
INTERACTIVE PORTFOLIO

  Try my portfolio in your terminal

  $ npx habibiahmada

  Or, connect remotely:

  $ ssh habibiahmada.dev
```
