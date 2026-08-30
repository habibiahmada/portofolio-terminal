# Panduan Pengguna

Cara mengakses portfolio **Habibi Ahmad Aziz** di terminal Anda — tanpa perlu clone repo atau install Go.

## Ringkasan Cepat

```bash
# Jalankan di laptop Anda (local)
npx habibiahmada

# Atau hubungkan ke server (remote)
ssh habibiahmada.dev
```

Kedua perintah di atas menampilkan **TUI portfolio interaktif yang sama**.

---

## Opsi 1: npx habibiahmada (Local Experience)

Cara termudah untuk mayoritas developer. Binary Go diunduh dan dijalankan langsung di laptop Anda.

### Persyaratan

- Terminal modern (iTerm2, Windows Terminal, GNOME Terminal, dll.)
- Node.js ≥ 14 (untuk menjalankan `npx`)
- Koneksi internet (hanya saat pertama kali unduh)

### Cara Menjalankan

```bash
npx habibiahmada
```

npm akan:
1. Mengunduh package `habibiahmada` (jika belum ada di cache)
2. Mendeteksi OS dan arsitektur Anda
3. Menjalankan binary Go yang sesuai
4. Membuka TUI portfolio

### Platform yang Didukung

| OS | Arsitektur |
|----|------------|
| Linux | x64, arm64 |
| macOS | x64 (Intel), arm64 (Apple Silicon) |
| Windows | x64 |

### Pertanyaan Umum

**Apakah perlu install Go?**
Tidak. Binary sudah dikompilasi. Anda hanya perlu Node.js untuk `npx`.

**Apakah perlu internet setiap kali?**
Tidak selalu. Setelah package di-cache oleh npm, bisa dijalankan tanpa unduh ulang.

**Apakah EC2/server harus hidup?**
Tidak. npx berjalan 100% di laptop Anda.

**Berapa ukuran download?**
Sekitar ~10–20 MB (tergantung platform).

---

## Opsi 2: ssh habibiahmada.dev (Remote Experience)

TUI berjalan di server AWS EC2. Laptop Anda hanya mengirim input keyboard dan menerima output tampilan.

### Persyaratan

- SSH client (sudah tersedia di Linux/macOS; Windows 10+ punya OpenSSH bawaan)
- Koneksi internet
- **Server EC2 harus dalam keadaan ON**

### Cara Menjalankan

```bash
ssh habibiahmada.dev
```

Setelah terhubung, TUI portfolio langsung terbuka — tidak perlu mengetik perintah tambahan.

### Kapan Memilih SSH?

| Situasi | Rekomendasi |
|---------|-------------|
| Tidak punya Node.js | SSH |
| Ingin demo "server-side TUI" | SSH |
| EC2 sedang mati | Gunakan npx |
| Butuh offline | Gunakan npx (setelah cache) |

### Catatan Ketersediaan

SSH portfolio menumpang pada EC2 yang juga menjalankan Telegram Agent. EC2 dinyalakan **3× sehari** via scheduler. Jika SSH gagal connect, kemungkinan EC2 sedang STOPPED — gunakan `npx habibiahmada` sebagai alternatif.

---

## Opsi 3: Website

```bash
# Buka di browser
https://habibiahmada.dev
```

Website dan terminal portfolio menampilkan informasi yang sama, dengan presentation layer berbeda.

---

## Navigasi di TUI

```
┌───────────────────────────────────────────────┐
│ HABIBI AHMAD AZIZ                             │
│ Full-Stack Web Developer                      │
├──────────────┬────────────────────────────────┤
│ > About      │  Konten halaman yang dipilih   │
│   Projects   │                                │
│   Skills     │                                │
│   Experience │                                │
│   Certificates                                │
│   Contact    │                                │
├──────────────┴────────────────────────────────┤
│ ↑↓ Navigate  Enter Select  ← Back  Q Quit    │
└───────────────────────────────────────────────┘
```

### Kontrol Keyboard & Navigasi

| Tombol / Shortcut | Aksi |
|---|---|
| `↑` / `↓` / `k` / `j` | Navigasi menu sidebar / scroll konten |
| `→` / `Enter` | Masuk ke area konten / buka detail project |
| `←` / `Esc` | Kembali ke sidebar navigasi / keluar dari modal |
| `P` | *[Shortcut]* Langsung lompat ke halaman **Projects** |
| `C` | *[Shortcut]* Langsung lompat ke halaman **Contact** |
| `V` | *[Shortcut]* Buka **CV Modal Viewer** (dari Home) |
| `M` | *[Easter Egg]* Interaksi dengan Maskot CRT Robot |
| `s` / `S` | Toggle mode seleksi teks native terminal |
| `?` / `F1` | Buka / tutup overlay panduan bantuan |
| `q` / `Ctrl+C` | Keluar dari aplikasi TUI |

#### Navigasi di Halaman Project Detail:
- `←` / `h`: Pindah ke project sebelumnya
- `→` / `l`: Pindah ke project berikutnya
- `Esc`: Kembali ke daftar project

### Dukungan Mouse

- **Wheel Up / Down**: Scroll konten atas dan bawah
- **Klik Menu Sidebar**: Langsung berpindah ke halaman tersebut
- **Klik Card Project**: Membuka halaman detail project
- **Scrollbar Drag/Click**: Loncat cepat ke posisi konten yang diinginkan
- **Klik Maskot**: Memancing reaksi kedip mata (`─ ─`, `★ ★`) atau reaksi marah (`╬`) jika diklik berkali-kali!

### Menu yang Tersedia

| Menu | Isi |
|---|---|
| **Home** | Hero banner `HABIBI`, featured project cards interaktif, trusted partners |
| **About** | Profil, bio, filosofi rekayasa software, statistik |
| **Projects** | Daftar seluruh proyek dengan tag pills, badge featured, dan case study detail |
| **Skills** | Keahlian teknis per domain (Frontend, Backend, DevOps, AI) |
| **Experience** | Riwayat pekerjaan dan pendidikan formal |
| **Certificates** | Sertifikasi profesional terverifikasi |
| **Services** | Layanan dan solusi rekayasa web yang ditawarkan |
| **Contact** | Email, GitHub, LinkedIn, Website, dan ketersediaan |

---

## Perbandingan Jalur Akses

| | npx | SSH | Website |
|---|-----|-----|---------|
| Perintah | `npx habibiahmada` | `ssh habibiahmada.dev` | Browser → `habibiahmada.dev` |
| Butuh Node.js | Ya | Tidak | Tidak |
| Butuh SSH client | Tidak | Ya | Tidak |
| Butuh browser | Tidak | Tidak | Ya |
| TUI interaktif | Ya | Ya | Tidak |
| Offline | Bisa (setelah cache) | Tidak | Tidak |
| Bergantung EC2 | Tidak | Ya | Tidak |

Detail teknis perbandingan npx vs SSH: [npx-vs-ssh.md](npx-vs-ssh.md).

---

## Troubleshooting

### npx: command not found

Install Node.js dari [nodejs.org](https://nodejs.org/) atau via package manager:

```bash
# Ubuntu/Debian
sudo apt install nodejs npm

# macOS (Homebrew)
brew install node
```

### npx: Binary not found

Package mungkin belum ter-publish atau versi lokal belum di-build. Untuk development, lihat [getting-started.md](getting-started.md).

### ssh: Connection refused / timed out

EC2 kemungkinan sedang STOPPED. Alternatif:

```bash
npx habibiahmada
```

### TUI tampilan rusak / karakter aneh

Pastikan terminal Anda mendukung:
- UTF-8 encoding
- Minimal 80×24 karakter
- True color (opsional, untuk styling terbaik)

### TUI tidak merespons keyboard

Pastikan fokus ada di jendela terminal (bukan di panel lain). Tekan `Ctrl+C` untuk keluar paksa.

---

## Untuk Recruiter / Hiring Manager

Tiga cara cepat melihat portfolio:

```bash
# Terminal (local) — paling mudah
npx habibiahmada

# Terminal (remote) — demo server-side
ssh habibiahmada.dev

# Web
open https://habibiahmada.dev
```

Proyek terminal ini sendiri adalah bagian dari portfolio: *"Interactive CLI Portfolio built with Go and Bubble Tea"*.
