---
name: bubble-tea-tui
description: >-
  Bubble Tea and Lip Gloss TUI design for habibiahmada-terminal. Use when
  building terminal UI, screens, components, styles, keyboard navigation,
  or layout in internal/tui and internal/components.
---

# Bubble Tea TUI Design

Panduan UI terminal untuk portfolio Habibi Ahmad Aziz.

## Stack

- **Bubble Tea** — TUI framework (Init, Update, View)
- **Lip Gloss** — styling & layout
- **Wish** — SSH layer (cmd/ssh only)

## Visual Language

Styles terpusat di `internal/styles/styles.go`:

| Style | Penggunaan |
|-------|------------|
| TitleStyle | Header "HABIBI AHMAD AZIZ" |
| SubtitleStyle | "Full-Stack Web Developer" |
| SelectedStyle | Menu item aktif (`> About`) |
| NormalStyle | Menu item biasa |
| BorderStyle | Box/border layout |
| MutedStyle | Hint, timestamp |
| SuccessStyle | Status positif |
| LinkStyle | GitHub, Live Demo (OSC 8 jika didukung) |

Definisikan sekali di package level. Jangan raw ANSI di screen files.

## Layout Pattern

```
┌───────────────────────────────────────────────┐
│ HABIBI AHMAD AZIZ                             │ ← components/header
├──────────────┬────────────────────────────────┤
│ > About      │                                │
│   Projects   │       Content                  │ ← sidebar + content
│   Skills     │                                │
├──────────────┴────────────────────────────────┤
│ ↑↓ Navigate  Enter Select  ← Back  Q Quit     │ ← components/footer
└───────────────────────────────────────────────┘
```

Komponen reusable: `header`, `sidebar`, `footer`, `card`, `list`, `modal`.

## Navigation

| Key | Action |
|-----|--------|
| ↑ ↓ | Navigate menu |
| Enter | Select |
| ← | Back |
| q | Quit |

Definisikan di `internal/tui/keymap.go`.

## Screens

| Screen | Konten |
|--------|--------|
| Home | Menu utama |
| About | Bio developer |
| Projects | Renshuu, SmartFarm AI, CultureConnect, Spacelab |
| Skills | Laravel, Next.js, React, Node.js, PostgreSQL, Supabase, Docker, GCP |
| Experience | Riwayat kerja |
| Certificates | Sertifikasi |
| Contact | Website, GitHub, LinkedIn |

Projects detail: nama, deskripsi, stack, [GitHub] [Live Demo].

## Bubble Tea Architecture

Top-level model = message router + compositor:

```
App (top-level)
 ├── Home model
 ├── Projects model
 ├── Skills model
 └── ...
```

- Route messages ke child model yang aktif
- Compose View dari child `View()` outputs
- Jangan simpan side effects di `View()`

## Responsive Layout

1. Handle `tea.WindowSizeMsg` di Update
2. Pass width/height ke child components
3. Kirim `WindowSizeMsg` ke child saat screen switch (child perlu resize saat pertama tampil)

```go
case tea.WindowSizeMsg:
    m.width = msg.Width
    m.height = msg.Height
    m.sidebar.SetSize(m.sidebarWidth, m.height-footerHeight)
```

## Lip Gloss Layout

- `lipgloss.JoinVertical` / `JoinHorizontal` untuk compose
- `lipgloss.Place` untuk center content
- `MaxWidth`/`MaxHeight` untuk terminal resize
- `AdaptiveColor` / `LightDark` untuk light/dark terminal

Biarkan Bubble Tea handle color profile detection — jangan force TrueColor.

## SSH (Wish)

TUI di `internal/tui` identik untuk local dan SSH. Wish hanya transport — tidak duplikasi UI.

## Best Practices (Charm / Community)

1. **Styles once** — package-level Lip Gloss variables
2. **Views are pure** — no mutation in View()
3. **Window size** — always react to resize
4. **Layout arithmetic** — test dengan terminal kecil & besar
5. **Receiver methods** — gunakan judiciously; prefer explicit Update on child
6. **Cache expensive renders** jika View rebuild berat

## Anti-Patterns

- Hard-code ASCII art di setiap screen (satu header component)
- Inline styles di 10 file berbeda
- Ignore WindowSizeMsg → broken layout
- Web box-model mental model — Lip Gloss sizing berbeda
- Dua TUI codebase untuk npx vs SSH

## Referensi Proyek

- Ilustrasi & signature art: `docs/tui-illustration.md`
- Struktur folder: `docs/folder-structure.md`
- UX local vs remote: `docs/npx-vs-ssh.md`
