---
name: graphify-workflow
description: >-
  Menjalankan dan memakai Graphify knowledge graph sebelum explore codebase.
  Use when starting work on this repo, exploring architecture, finding file
  relationships, or when the user mentions graphify, knowledge graph, or
  codebase navigation.
---

# Graphify Workflow

Proyek ini memakai [Graphify](https://github.com/Graphify-Labs/graphify) (PyPI: `graphifyy`, CLI: `graphify`).

## Setup (sudah dikonfigurasi)

- Rule Cursor: `.cursor/rules/graphify.mdc` (`alwaysApply: true`)
- Output graph: `graphify-out/graph.json`, `GRAPH_REPORT.md`

## Sebelum Explore Codebase

1. Baca `graphify-out/GRAPH_REPORT.md` jika ada
2. Query graph, jangan langsung grep banyak file:

```bash
graphify query "what connects cmd/portfolio to internal/tui?" --graph graphify-out/graph.json
graphify query "show navigation flow" --dfs --graph graphify-out/graph.json
graphify explain "internal/tui/app.go" --graph graphify-out/graph.json
```

3. Gunakan `graphify path "A" "B"` untuk shortest path antar node

## Update Graph

Setelah perubahan kode signifikan:

```bash
graphify update .
```

Untuk perubahan dokumen/non-code, rebuild penuh via assistant: `/graphify .`

## Hook (opsional)

```bash
graphify hook install
```

Post-commit hook rebuild graph incrementally.

## Cursor Install/Uinstall

```bash
graphify cursor install    # tulis .cursor/rules/graphify.mdc
graphify cursor uninstall  # hapus rule
```

## Prioritas

**Graphify query → targeted file read → grep**

Jangan baca puluhan file untuk memetakan struktur jika graph sudah tersedia.

## Catatan

- Package resmi: `graphifyy` (double-y), bukan `graphify`
- Install: `uv tool install graphifyy`
- Tidak ada command `graphify init`; gunakan `graphify update .` + `graphify cursor install`
