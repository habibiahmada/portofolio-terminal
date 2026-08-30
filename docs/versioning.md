# Semantic Versioning

Panduan versi untuk **habibiahmada-terminal**.

## Skema

Proyek ini mengikuti [Semantic Versioning](https://semver.org/) / **`MAJOR.MINOR.PATCH`**.

```
MAJOR.MINOR.PATCH
  │      │      │
  │      │      └─ PATCH — bugfix, tidak merubah behavior
  │      └──────── MINOR — fitur baru, backward-compatible
  └─────────────── MAJOR — breaking change
```

## Aturan untuk Proyek Ini

| Kenaikan | Contoh | Kapan dipakai |
|----------|--------|---------------|
| `PATCH` | `v1.0.0 → v1.0.1` | Fix bug, update dokumentasi, perubahan kecil data portfolio (judul/deskripsi) |
| `MINOR` | `v1.0.0 → v1.1.0` | Fitur TUI baru, tambah project, tambah screen, tambah komponen |
| `MAJOR` | `v1.0.0 → v2.0.0` | Breaking change: ganti prompt/UX, ganti cara install binary, ganti nama npm package |

## Tag & npm

- Tag git dan versi npm **selalu sinkron** — tag `v1.0.1` → npm version `1.0.1`.
- Tag dipush untuk memicu release workflow (`v*`).
  ```bash
  git tag v1.1.0
  git push origin v1.1.0
  ```
- Workflow `release.yml` → build → `npm publish` + GitHub Release.

## Checklist Sebelum Release

- [ ] Working tree bersih (perubahan sudah commit)
- [ ] `make lint` lulus (fmt, vet, test)
- [ ] CI hijau (`.github/workflows/ci.yml`)
- [ ] Semua binary test lokal (`bash scripts/release.sh`)
- [ ] `CHANGELOG` / release notes berisi perubahan (generated otomatis oleh workflow)

## Perubahan Data → MINOR

Karena data v1 **dibundel di binary**, setiap perubahan portfolio (tambah/ubah project, skill, experience) adalah **fitur baru** bagi end user → `MINOR`.

> Deadline pembuatan data: perubahan yg sifatnya hanya typo/koreksi → `PATCH`.

## Referensi

- [deployment.md](deployment.md) — alur release end-to-end
- [task-list.md](task-list.md) — roadmap & status