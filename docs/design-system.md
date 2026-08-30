# Design System TUI

Panduan visual **habibiahmada-terminal** yang diselaraskan dengan identitas website portfolio — "Red & Blue Glitch" — diadaptasi untuk terminal monospace, Lip Gloss, dan Bubble Tea.

Dokumen ini adalah referensi implementasi untuk `internal/styles/`, komponen TUI, dan ilustrasi ASCII. Lihat juga [tui-illustration.md](tui-illustration.md) untuk detail animasi dan signature art.

---

## Prinsip

| Prinsip | Website | Terminal |
|---------|---------|----------|
| Identitas | Red accent, glitch, geometric | Merah brand + biru sekunder, border Lip Gloss |
| Typography | Space Grotesk (heading) + Geist (body) | FIGlet/hero bold + monospace body |
| Hierarchy | Section labels `//`, zinc contrast | Label muted + title primary + body normal |
| Motion | Glitch, marquee, view transitions | Splash ≤2s, spinner, scroll indicator |
| Aksesibilitas | Light/dark theme | Dark-first terminal; degradasi tiny mode |

**Aturan emas:** Terminal illustration, bukan foto ASCII. Ilustrasi menyusut di terminal kecil (≥40×12 minimum, optimal ≥80×24).

---

## Palet Warna

### Brand (dari website)

| Token | Light web | Dark web | TUI (target) | Penggunaan |
|-------|-----------|----------|--------------|------------|
| Brand accent | `#ef4444` | `#f87171` | `#ef4444` | Logo dot, H1 accent, selected menu |
| Glitch blue | `#3b82f6` | `#3b82f6` | `#3b82f6` | Secondary accent, links, tags |
| Background | `#fafafa` | `#030303` | `#1a1a1a` | Area konten (SSH dark bg) |
| Foreground | `#09090b` | `#fafafa` | `#f5f5f5` | Body text |
| Muted | zinc-500 | zinc-400 | `#636e72` | Label, footer hints, metadata |
| Border | zinc-200 | zinc-800 | `#3f3f46` | Card, sidebar divider |
| Success | — | — | `#00B894` | Badge "Current", availability |

### Migrasi dari palette lama

Palette saat ini di `internal/styles/styles.go` menggunakan `#FF6B6B` (primary) dan `#4ECDC4` (secondary). **Target parity:**

```
ColorPrimary   → #ef4444  (brand red)
ColorSecondary → #3b82f6  (glitch blue)
ColorAccent    → #f87171  (soft red, selection highlight)
ColorLink      → #3b82f6
ColorText      → #f5f5f5  (dark terminal)
ColorBackground→ #1a1a1a
ColorBorder    → #3f3f46
ColorMuted     → #71717a  (zinc-500)
```

---

## Typography

### Mapping font web → terminal

| Web | Terminal | Implementasi |
|-----|----------|--------------|
| Space Grotesk (heading) | FIGlet large | `components/figlet.go`, font adaptif |
| Geist Sans (body) | Default terminal font | Lip Gloss tanpa override font |
| Geist Mono (labels, CTAs) | Monospace uppercase | `styles.LabelStyle`, tracking via spacing |

### Skala teks

| Level | Style | Contoh |
|-------|-------|--------|
| Hero H1 | FIGlet + `HeroTitleStyle` | `Building digital experiences...` (potong di tiny) |
| Page H1 | `TitleStyle` bold primary | `All Projects` |
| Section label | `MutedStyle` + prefix `//` | `// About` |
| Section H2 | `TitleStyle` | `Tools & Technologies` |
| Body | `NormalStyle` / `ContentStyle` | Paragraf bio |
| Meta | `MutedStyle` italic | `May 2026 – Now · Karawang` |
| CTA hint | `MutedStyle` + bracket | `[ ENTER ] Explore Portfolio` |
| Tag | `TagStyle` border secondary | `Laravel` |

### Responsive typography

| Lebar terminal | Perilaku |
|----------------|----------|
| < 60 cols | FIGlet hidden, plain bold title |
| 60–89 cols | FIGlet compact font |
| ≥ 90 cols | FIGlet full + signature wide |
| < 40×20 | Skip splash, `VariantHidden` ilustrasi |

---

## Spacing & Layout

### Grid website → TUI

| Web token | Nilai | TUI equivalent |
|-----------|-------|----------------|
| `max-w-7xl` | 90rem | `contentWidth()` max 76 cols, sidebar + divider + margin |
| Nav | vertical rail kiri | `NavRail` — `[Screen]` aktif + hints, di-center vertikal |
| Border nav↔konten | — | `VerticalRule` — satu garis `│` kontinu |
| Footer | bottom bar | `height - footerLines` body zone, footer baris terakhir |

### Layout struktur

```
│ [Home]  │ habibiahmada. · title          │
│  About  │ page content                   │
│  ...    │                                │
│ ↑↓ hint │                                │
├──────────────────────────────────────────┤
│ habibiahmada.          ↑↓ Screens · j/k · … │
└──────────────────────────────────────────┘
```

| Konstanta | Nilai | File |
|-----------|-------|------|
| `NavRailWidth` | 16 | `internal/components/sidebar.go` |
| `footerHeight` | 2 | `internal/tui/app.go` |
| `maxContentWidth` | 76 | `internal/tui/layout.go` |

---

## Komponen Visual

### Logo & brand mark

| Elemen | Web | TUI |
|--------|-----|-----|
| Wordmark | `habibiahmada.` (dot merah) | Header: `habibiahmada` + styled `.` |
| Signature | — | `internal/assets/art/signature_*.txt` |
| Terminal mini | CPU illustration | `AboutTerminal()` ASCII |

### Card

**Web:** rounded-2xl, border zinc, hover glitch  
**TUI:** `CardStyle` — `RoundedBorder`, `BorderForeground(ColorBorder)`, padding 1×2

```
┌─────────────────────────┐
│ 01 / Design             │
│ Web Design & Mobile-First│
│ Translating ideas into… │
└─────────────────────────┘
```

### Tag / Badge

| Variant | Style |
|---------|-------|
| Tech tag | `TagStyle` — border `#3b82f6` |
| Status badge | `SuccessStyle` — "Current", "Open to freelance" |
| Award badge | `ColorAccent` — "Top 15 Capstone" |

### List & selection

| State | Style |
|-------|-------|
| Normal item | `NormalStyle` |
| Selected (↑↓) | `ListSelectedStyle` / `SidebarItemSelectedStyle` |
| Active screen | Sidebar highlight `activeKey` |

### Modal

Help overlay (`?`) dan CV viewer (`V` dari Home) memakai `components.Modal` — centered, border primary, max 80% width.

---

## Identitas "Glitch" di Terminal

Website memakai chromatic aberration dan glitch text. Di terminal:

| Efek web | Adaptasi TUI |
|----------|--------------|
| Glitch H1 | Dual-color render: char offset merah/biru (opsional, tiny off) |
| Red dot brand | `.` dengan `ColorPrimary` |
| Node network bg | ASCII dot grid muted di hero (optional, ≥100 cols) |
| Marquee tech logos | Horizontal scroll text di Skills (Tea.Tick) |
| CRT scanline | **Tidak** — tidak kompatibel lintas SSH client |

Implementasi glitch ringan (Fase 3):

```go
// Contoh: alternating foreground on accent words
glitchStyle := lipgloss.NewStyle().Foreground(ColorPrimary)
altStyle   := lipgloss.NewStyle().Foreground(ColorSecondary)
```

---

## Screen-specific Design

### Home

| Section | Visual treatment |
|---------|------------------|
| Hero | Signature + FIGlet name + availability badge |
| Companies | Single line horizontal: `Neskar · PPLG · …` |
| Featured | 5 compact cards, number prefix |
| Press | 2 bordered cards, quote indent |
| Process | Numbered list 1–4, title bold |
| CTA | Primary border card, email as link style |

### About

| Section | Visual treatment |
|---------|------------------|
| Hero stats | 3-column inline: `3+ Years · 10+ Projects · 2 Awards` |
| Intro | Text + `AboutTerminal()` side-by-side (≥100 cols) |
| Journey CTA | Link style hint |

### Projects

| Section | Visual treatment |
|---------|------------------|
| Archive grid | List dengan year + tags inline |
| Detail | Section headers `── Where it started ──` |
| Prev/Next | Footer hints `← prev · next →` |

### Skills

Marquee simulation: rotating single line of 16 tools, or static wrapped grid.

### Experience

Timeline: `│` vertical line + `●` node, period muted left.

### Certificates

Pinned: `★` prefix; grid scroll dengan load indicator.

### Services

Numbered cards `01 / Design` — match web card hierarchy.

### Contact

Email prominent `LinkStyle`; social list dengan icon prefix `[GH]` `[LI]` `[IG]`.

---

## Ilustrasi & Assets

| Asset | Variant | Breakpoint |
|-------|---------|------------|
| Signature wide | ≥90 cols | `signature_wide.txt` |
| Signature compact | 60–89 | `signature_compact.txt` |
| Signature mini | 40–59 | `signature_mini.txt` |
| Hidden | <40 | `VariantHidden` |
| About terminal | ≥80 cols | `about_terminal.txt` |
| Splash | ≥40×20 | `splash.go` + progress bar |

Detail lengkap: [tui-illustration.md](tui-illustration.md)

---

## Keyboard & Interaction

Selaras dengan footer hints website CTAs:

| Key | Aksi global |
|-----|-------------|
| ↑↓ / j k | Navigate / scroll |
| Enter | Select / open detail |
| ← / Esc | Back |
| ? / F1 | Help modal |
| q / Ctrl+C | Quit |
| V | CV viewer (Home) |
| P | Jump Projects (Home, opsional) |

---

## Layout & Performance

### Center content

Konten pendek di-center vertikal via newline padding (`components.CenterInViewport`) — **bukan** `lipgloss.Place` full grid (mahal di RAM/CPU).

### Layout cache

`renderBodyCached()` menyimpan header + sidebar + content. Key **tidak** memasukkan `footerFrame` — animasi footer tidak rebuild seluruh Home/About/projects.

### Navigation

| Key | Aksi |
|-----|------|
| `↑↓` | Ganti screen langsung |
| `j/k` | Scroll / pilih item list |
| `Enter` | Buka detail |

### Footer

Satu baris: `>_ habibiahmada.` kiri · hints kanan. Prompt `>_` selalu berkedip (blinking `▊` cursor, opencode-style). Tick animasi **300ms**. Equalizer hanya jika `width ≥ 90`.

---

## Checklist Parity Desain

| # | Item | Status |
|---|------|--------|
| 1 | Brand red `#ef4444` menggantikan `#FF6B6B` | ⬜ |
| 2 | Glitch blue `#3b82f6` untuk links/tags | ⬜ |
| 3 | Header wordmark `habibiahmada.` | ⬜ |
| 4 | Section label pattern `// Label` | ⬜ |
| 5 | Availability badge di hero | ⬜ |
| 6 | Footer tagline match website | ⬜ |
| 7 | Card style konsisten di Services/Press | ⬜ |
| 8 | Timeline visual Experience | ✅ (partial) |
| 9 | Skill bar / marquee | ✅ (partial) |
| 10 | Dark-first palette | ⬜ |

---

## Referensi Internal

| Topik | File |
|-------|------|
| Styles | `internal/styles/styles.go` |
| Illustration styles | `internal/styles/illustration.go` |
| Hero | `internal/components/hero.go` |
| FIGlet | `internal/components/figlet.go` |
| Signature | `internal/components/illustration.go` |
| Layout | `internal/tui/app.go` |
| Halaman & copy | [pages.md](pages.md) |
| Task implementasi | [task-list.md](task-list.md) |

---

Terakhir diperbarui: Agustus 2026
