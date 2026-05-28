# Password Manager UI

Svelte 5 + TypeScript frontend for the local password manager. It talks to the Go API on `http://localhost:3000` and is built as a **single HTML file** so the backend can embed it with `go:embed`.

## Features

- Credential table with domain, username, and password columns
- Fuzzy search by domain
- Inline editing and save via the API
- One-click password copy to the clipboard
- Password change history per credential
- Modal to add new credentials (with show/hide password)
- Toast notifications for success and error states

## Stack

- [Svelte 5](https://svelte.dev/) (runes: `$state`, `$derived`, `$props`)
- [Vite](https://vite.dev/) for dev and production builds
- [vite-plugin-singlefile](https://github.com/richardtallent/vite-plugin-singlefile) — bundles JS/CSS into one `index.html` for embedding

## Prerequisites

- [Node.js](https://nodejs.org/) (LTS recommended)
- [pnpm](https://pnpm.io/) (lockfile is `pnpm-lock.yaml`)

The Go server must be running for API calls during normal use. For UI-only work you can run Vite alone, but list/create/update/history requests will fail without the backend.

## Development

From this directory:

```bash
pnpm install
pnpm dev
```

Vite serves the app (default port 5173). Start the Go server on port 3000 so fetches to `/credentials`, `/credential/*`, and `/credential-history/*` succeed.

Optional typecheck:

```bash
pnpm check
```

## Production build

```bash
pnpm build
```

Output goes to `ui/dist/` (gitignored). The root `main.go` embeds `ui/dist/index.html` and serves it at `/`.

After changing the UI, rebuild before running or packaging the Go binary:

```bash
cd ui && pnpm build && cd .. && go run .
```

## Project layout

| Path | Role |
|------|------|
| `src/App.svelte` | Main shell, API calls, search, toasts |
| `src/lib/CredentialTable.svelte` | Editable credential grid |
| `src/lib/PasswordCell.svelte` | Masked password display and copy |
| `src/lib/HistoryTable.svelte` | History drawer for a credential |
| `src/lib/AddCredentialModal.svelte` | Create-credential form |
| `src/lib/Toast.svelte` | Transient notifications |
| `vite.config.ts` | Svelte plugin + single-file bundle |

## IDE

[VS Code](https://code.visualstudio.com/) with the [Svelte extension](https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode) is recommended; see `.vscode/extensions.json`.
