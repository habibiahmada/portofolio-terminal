# habibiahmada

Interactive terminal portfolio for **Habibi Ahmad Aziz** — full-stack web developer.

Run the Bubble Tea TUI locally without cloning the Go repo:

```bash
npx habibiahmada
```

Remote SSH experience (same UI, runs on EC2):

```bash
ssh ssh.habibiahmada.dev
```

Website: [habibiahmada.dev](https://habibiahmada.dev)

## What you get

- **Home** — featured projects, availability, quick actions
- **Projects** — full archive with case-study detail (live data when online)
- **Skills, Experience, Certificates, Services, Contact**

Portfolio project data syncs from the website API when internet is available; a bundled offline copy is used as fallback.

## Requirements

- **Node.js** ≥ 14 (wrapper only — the TUI is a native Go binary)
- Supported platforms: Linux (x64/arm64), macOS (Intel/Apple Silicon), Windows (x64)

## Navigation

| Keys | Action |
|------|--------|
| `↑` / `↓` / `j` / `k` | Navigate / scroll |
| `→` / `Enter` | Focus content / open detail |
| `←` / `Esc` | Back to nav |
| `P` | Jump to Projects (from Home) |
| `C` | Contact |
| `V` | CV viewer |
| `R` | Retry live portfolio sync (when offline) |
| `?` | Help |
| `q` / `Ctrl+C` | Quit |

Mouse wheel scroll and click navigation are supported. Press `s` to toggle native text selection.

## Environment

| Variable | Description |
|----------|-------------|
| `PORTFOLIO_API_URL` | Override API base (default `https://www.habibiahmada.dev`) |

## Install globally (optional)

```bash
npm install -g habibiahmada
habibiahmada
```

## License

MIT © [Habibi Ahmad Aziz](https://habibiahmada.dev)

## Repository

[github.com/habibiahmada/portofolio-terminal](https://github.com/habibiahmada/portofolio-terminal)
