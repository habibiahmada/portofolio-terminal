# Ilustrasi TUI

Panduan visual identity, ilustrasi ASCII, dan animasi untuk **habibiahmada-terminal**.

Dokumen ini harus selesai diimplementasikan **sebelum Fase 4 — Deployment** (setelah Fase 3 Portfolio Parity), agar pengalaman `npx habibiahmada` dan `ssh habibiahmada.dev` sudah terasa seperti produk yang sengaja didesain — bukan website yang dipindahkan ke terminal.

---

## Prinsip Desain

### Hybrid: library + signature custom

Jangan memilih antara "pakai library" atau "buat sendiri". Gunakan keduanya:

```
                TERMINAL PORTFOLIO
                        │
           ┌────────────┴────────────┐
           │                         │
      Existing Assets             Custom
           │                         │
           ▼                         ▼
      FIGlet fonts              Habibi logo
      Bubbles components        Signature art
      Spinners                  Animasi splash
      Lip Gloss borders         Identitas visual
           │                         │
           └────────────┬────────────┘
                        ▼
                   Bubble Tea
                        │
                    Lip Gloss
                        │
                        ▼
                       TUI
```

| Layer | Pakai library | Buat sendiri |
|-------|---------------|--------------|
| Infrastruktur | Bubble Tea, Bubbles, Lip Gloss, FIGlet | — |
| Typography besar | `go-figure` / FIGlet fonts | — |
| Layout & border | Lip Gloss `BorderStyle`, `CardStyle` | — |
| Identitas | — | Signature illustration, logo terminal, micro-art per screen |
| Animasi | Spinner Bubbles (loading) | Splash sequence custom (≤ 2 detik) |

### Terminal illustration, bukan foto ASCII

Terminal punya keterbatasan: font monospace, lebar kolom, dukungan Unicode/warna, dan variasi client SSH. **Jangan** jadikan foto realistis atau ASCII art ultra-detail sebagai fondasi.

Yang direkomendasikan:

- Geometric ASCII art (box-drawing, simbol sederhana)
- Logo/mascot minimal (`>_` prompt, terminal window)
- Hierarchy visual via Lip Gloss (warna, padding, border)
- Ilustrasi yang **menyusut dan menyederhanakan diri** saat terminal kecil

Yang dihindari di v1:

- Sixel / Kitty graphics / iTerm inline images sebagai fondasi (tidak kompatibel lintas SSH)
- ASCII portrait detail
- Animasi berat di setiap screen switch

---

## Signature Illustration

Identitas visual utama portfolio terminal — **"Habibi Terminal"**.

### Konsep

Kotak terminal dengan prompt `>_`, nama singkat, dan label "TERMINAL". Ini menjadi watermark visual yang konsisten di splash, home, dan about.

```
              ╭─────────────────╮
              │                 │
              │      >_         │
              │                 │
              │    HABIBI       │
              │    TERMINAL     │
              │                 │
              ╰─────────────────╯
```

Variasi compact (sidebar/header kecil):

```
╭──────────╮
│   >_     │
│  HABIBI  │
╰──────────╯
```

Variasi inline (footer/status):

```
>_ habibiahmada · online
```

### Elemen signature

| Elemen | Makna | Warna (Lip Gloss) |
|--------|-------|-------------------|
| `>_` prompt | Developer, CLI-first identity | `ColorAccent` (#FFE66D) |
| Kotak rounded | Terminal window | `ColorBorder` |
| "HABIBI TERMINAL" | Branding | `ColorPrimary` + `ColorSecondary` |
| Cursor blink (animasi) | Hidup, interaktif | Toggle `_` / ` ` di splash saja |

Signature ini **tidak** diambil dari library random — desain manual, disimpan sebagai string template di repo.

---

## Hierarki Visual per Screen

Setiap screen mengikuti urutan yang sama agar terasa kohesif:

```
Header (nama + status)
    ↓
Illustration / hero (opsional, responsive)
    ↓
Konten utama (data-driven)
    ↓
Footer (keymap hint)
```

### Home — Welcome Hero

Layar pertama setelah splash. Fokus: first impression + CTA navigasi.

```
╭────────────────────────────────────────────────────╮
│                                                    │
│              ╭─────────────────╮                   │
│              │      >_         │                   │
│              │    HABIBI       │                   │
│              │    TERMINAL     │                   │
│              ╰─────────────────╯                   │
│                                                    │
│              HABIBI AHMAD AZIZ                     │
│              Full-Stack Web Developer              │
│                                                    │
│         [ ENTER ] Explore Portfolio                │
│                                                    │
╰────────────────────────────────────────────────────╯
```

Setelah Enter → layout sidebar + content (screen About default).

### About — Developer Story

Micro-illustration: terminal kecil di samping bio, bukan hero penuh.

```
╭──────────╮   I build web applications,
│   >_     │   cloud infrastructure, and
│  HABIBI  │   developer tools.
╰──────────╯
```

### Projects — Stack Tags Visual

Bukan ilustrasi besar; gunakan **card + tag** sebagai visual rhythm:

```
┌─ Renshuu ─────────────────────┐
│ Language learning platform    │
│ [Laravel] [Vue] [PostgreSQL]    │
└─────────────────────────────────┘
```

Ikon ASCII opsional per kategori project:

| Tipe | Ikon |
|------|------|
| Web app | `🌐` atau `[www]` |
| AI/ML | `◇` atau `[ai]` |
| Dev tools | `>_` |
| Mobile | `📱` atau `[app]` |

(Gunakan ASCII fallback jika Unicode tidak aman — lihat [Responsive Behavior](#responsive-behavior).)

### Skills — Radar / Bar Visual

Ilustrasi data-driven, bukan art statis:

```
Go          ████████░░  80%
Laravel     █████████░  90%
Next.js     ███████░░░  70%
```

Bar disusun ulang lebar terminal; jumlah skill visible menyesuaikan tinggi window.

### Experience — Timeline

Garis vertikal ASCII sebagai ilustrasi struktural:

```
2024 ──●── Senior Developer @ Company
         │
2022 ──●── Full-Stack @ Startup
         │
2020 ──●── Intern @ Agency
```

Panjang garis menyesuaikan lebar; di terminal sempit, timeline jadi list datar tanpa garis.

### Certificates — Badge Grid

```
╭─────╮  ╭─────╮  ╭─────╮
│ GCP │  │ AWS │  │ ... │
╰─────╯  ╰─────╯  ╰─────╯
```

Grid kolom = `floor((width - padding) / badgeWidth)`.

### Contact — Connection Diagram

Ilustrasi hubungan sederhana:

```
        [ GitHub ]
            │
    [ Website ]─── YOU ───[ LinkedIn ]
            │
        [ Email  ]
```

Di terminal kecil: list vertikal biasa, diagram disembunyikan.

---

## Splash & Animasi

Animasi **hanya** di startup — tidak di setiap navigasi screen.

### Sequence (target ≤ 2 detik)

```
Frame 1 (0ms)     Initializing portfolio...
Frame 2 (400ms)   [████████░░░░░░░░] 50%
Frame 3 (800ms)   [████████████████] 100%
Frame 4 (1000ms)  Signature illustration + cursor blink
Frame 5 (1800ms)  Transition → Home hero
```

Implementasi: Bubble Tea `tea.TickMsg` + state machine di model splash terpisah. Jangan block SSH handshake terlalu lama — skip splash jika `TERM` tidak support warna atau lebar < 40.

### Referensi inspirasi

| Proyek | Pelajaran |
|--------|-----------|
| [ascii-animations](https://github.com/cyperx84/ascii-animations) | Pola splash, banner, color demo |
| [firework](https://github.com/erik-adelbert/firework) | Animasi frame-based dengan Bubble Tea |
| Bubbles spinner | Loading indicator standar |

---

## Responsive Behavior

Ilustrasi **tidak kaku** — setiap hero/adaptif punya beberapa **variant** berdasarkan lebar terminal.

### Breakpoints

| Lebar | Mode | Perilaku |
|-------|------|----------|
| ≥ 100 | `wide` | Full signature art + sidebar + content side-by-side |
| 80–99 | `comfortable` | Signature compact, sidebar sempit |
| 60–79 | `narrow` | Signature mini atau hanya FIGlet nama |
| 40–59 | `minimal` | Text-only hero, stack sidebar di atas content |
| < 40 | `tiny` | Sembunyikan ilustrasi; konten & navigasi prioritas |

Tinggi terminal:

| Tinggi | Perilaku |
|--------|----------|
| ≥ 30 | Tampilkan hero + full footer |
| 20–29 | Kurangi padding hero |
| < 20 | Single-column; splash dilewati |

### Implementasi responsif

```
internal/
├── components/
│   └── illustration.go    # RenderIllustration(width, variant)
├── assets/
│   └── art/
│       ├── signature_wide.txt
│       ├── signature_compact.txt
│       ├── signature_mini.txt
│       └── splash_frames.go   # optional: generated frames
```

Pseudocode:

```go
func IllustrationVariant(width int) Variant {
    switch {
    case width >= 100:
        return VariantWide
    case width >= 80:
        return VariantCompact
    case width >= 60:
        return VariantMini
    default:
        return VariantHidden
    }
}

func RenderSignature(w int) string {
    switch IllustrationVariant(w) {
    case VariantWide:
        return assets.SignatureWide
    case VariantCompact:
        return assets.SignatureCompact
    case VariantMini:
        return assets.SignatureMini
    default:
        return ""
    }
}
```

### Aturan agar tidak kaku

1. **Multi-variant**, bukan satu ASCII art yang di-crop paksa
2. **Center dengan `lipgloss.Place`**, bukan hard-code kolom
3. **`MaxWidth` pada art**, bukan asumsi 80 kolom
4. **Graceful degradation** — hilangkan art sebelum layout rusak
5. **Test matrix**: 40×12, 80×24, 120×40, 200×50, SSH via Windows Terminal & macOS Terminal

### Unicode & SSH safety

| Karakter | Wide mode | Narrow mode |
|----------|-----------|-------------|
| Box-drawing (`╭╮│`) | ✅ | ✅ |
| Block elements (`██`) | ✅ | FIGlet fallback |
| Emoji | ⚠️ opsional | ❌ ASCII fallback |

Deteksi: jika `runewidth` atau lebar efektif tidak cocok, fallback ke variant ASCII-only.

---

## FIGlet & Typography

Untuk teks besar (nama, judul section):

```
HABIBI  →  ██╗  ██╗ █████╗ ...
           (font: standard / slant / small)
```

Library: [`github.com/common-nighthawk/go-figure`](https://github.com/common-nighthawk/go-figure)

Aturan:

- Font **small** atau **standard** untuk lebar < 80
- Font **slant** atau **big** hanya di `wide` mode
- Cache render FIGlet per `(text, font, width)` — jangan regenerate setiap frame

---

## Struktur File (Target)

```
internal/
├── assets/
│   └── art/
│       ├── signature_wide.txt
│       ├── signature_compact.txt
│       ├── signature_mini.txt
│       └── about_terminal.txt
├── components/
│   ├── illustration.go      # Signature + variant picker
│   ├── hero.go              # Home hero block
│   └── progress_bar.go      # Splash progress
├── styles/
│   └── illustration.go      # ArtStyle, PromptStyle, HeroStyle
└── tui/
    ├── splash.go              # Splash screen model
    └── home.go                # Hero + welcome
```

Styles ilustrasi terpisah dari styles umum:

```go
// internal/styles/illustration.go
var (
    PromptStyle = lipgloss.NewStyle().Foreground(ColorAccent).Bold(true)
    ArtBorderStyle = lipgloss.NewStyle().Foreground(ColorSecondary)
    HeroTitleStyle = lipgloss.NewStyle().Foreground(ColorPrimary).Bold(true)
)
```

---

## Checklist Implementasi

Urutan sebelum deploy ke production:

| # | Task | Prioritas |
|---|------|-----------|
| I.1 | Buat signature art (wide / compact / mini) | Tinggi |
| I.2 | Komponen `RenderSignature(width)` | Tinggi |
| I.3 | Home hero dengan ilustrasi responsif | Tinggi |
| I.4 | Splash screen + progress bar | Sedang |
| I.5 | Integrasi FIGlet untuk nama/judul | Sedang |
| I.6 | Micro-illustration About | Rendah |
| I.7 | Skill bar / timeline visual | Rendah |
| I.8 | Test matrix terminal sizes + SSH | Tinggi |
| I.9 | Dokumentasi ini ✅ | Selesai |

---

## Referensi

| Resource | URL |
|----------|-----|
| Bubble Tea | https://github.com/charmbracelet/bubbletea |
| Lip Gloss | https://github.com/charmbracelet/lipgloss |
| Bubbles (spinner, list, viewport) | https://github.com/charmbracelet/bubbles |
| Bubble Tea examples | https://github.com/charmbracelet/bubbletea/tree/main/examples |
| go-figure (FIGlet) | https://github.com/common-nighthawk/go-figure |
| ASCII Animations (inspirasi) | https://github.com/cyperx84/ascii-animations |
| Firework (animasi) | https://github.com/erik-adelbert/firework |

## Dokumen Terkait

- [architecture.md](architecture.md) — tiga presentation layer, identitas sama
- [tech-stack.md](tech-stack.md) — Bubble Tea, Lip Gloss
- [folder-structure.md](folder-structure.md) — `assets/`, `components/`
- [development-guide.md](development-guide.md) — workflow edit TUI
- [task-list.md](task-list.md) — Fase 2.5 Ilustrasi
- [npx-vs-ssh.md](npx-vs-ssh.md) — kompatibilitas terminal remote
