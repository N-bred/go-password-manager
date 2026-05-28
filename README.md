# Password Manager

A small, **local-first** password manager: a Go HTTP server with a SQLite database and a Svelte web UI embedded in the binary. Credentials stay on your machine; there is no cloud sync or remote auth layer.

## What it does

- Store credentials (domain, username, password) in SQLite (`db.db` at the project root when running)
- Serve a single-page UI at `http://127.0.0.1:3000`
- REST API for listing, creating, updating credentials and viewing password history
- Bootstrap the database from a text or JSON file on first launch via `--source`

Password history is recorded when credentials are created or updated, so you can see previous password values for a given entry.

## Architecture

```
┌─────────────────────────────────────────────────────────┐
│  Browser  →  http://127.0.0.1:3000                      │
│              ├── GET /          (embedded Svelte UI)    │
│              ├── GET /credentials/                      │
│              ├── GET /credential/{id}                   │
│              ├── POST /credential/create                │
│              ├── PUT /credential/update                 │
│              └── GET /credential-history/{id}           │
└─────────────────────────────────────────────────────────┘
                          │
                          ▼
                   SQLite (db.db)
```

| Area | Technology |
|------|------------|
| Backend | Go 1.26+, `net/http`, [go-sqlite3](https://github.com/ncruces/go-sqlite3) |
| Frontend | Svelte 5, Vite, bundled into one HTML file — see [ui/README.md](ui/README.md) |
| Data | SQLite tables `Credentials` and `Credentials_History` |

## Quick start

### 1. Build the UI (required before `go run` / `go build`)

```bash
cd ui
pnpm install
pnpm build
cd ..
```

### 2. First run — populate the database

On the **first** run, pass `--source` with a `.txt` or `.json` import file so SQLite is created and filled. Without this step the database has no credentials.

```bash
go run . --source path/to/credentials.txt
# or
go run . --source path/to/credentials.json
```

See [Import file formats](#import-file-formats) below for the expected shape of those files.

### 3. Later runs

Once the database exists, start the server **without** `--source`:

```bash
go run .
```

Open [http://127.0.0.1:3000](http://127.0.0.1:3000). Add or edit entries in the UI anytime; you only need `--source` again if you want to re-import from a file (that replaces/populates via `CreateAndPopulateDB`).

## Import file formats

### JSON (`.json`)

A JSON array of credential objects. Each entry must include `id`, `domain`, `username`, and `password` (the importer unmarshals into `models.Credential`):

```json
[
  {
    "id": 1,
    "domain": "github.com",
    "username": "octocat",
    "password": "example-password-1"
  },
  {
    "id": 2,
    "domain": "mail.example.org",
    "username": "user@example.com",
    "password": "example-password-2"
  }
]
```

### Text (`.txt`)

Line-oriented blocks: **domain**, then **username**, then **password**, then a blank separator line. Repeat for each credential. Files are usually saved with Windows-style line endings (`CRLF`); the parser treats a line that is only a carriage return as the separator between entries.

```text
github.com
octocat
example-password-1

mail.example.org
user@example.com
example-password-2

```

(Each blank line between blocks is the separator the importer expects—not arbitrary empty lines inside a single entry.)

## API overview

| Method | Path | Description |
|--------|------|-------------|
| GET | `/credentials/` | All credentials |
| GET | `/credential/{id}` | One credential by ID |
| POST | `/credential/create` | Create (JSON body) |
| PUT | `/credential/update` | Update (JSON body with `id`) |
| GET | `/credential-history/{id}` | Password history for credential `id` |

CORS is enabled for local development. The server binds to **127.0.0.1:3000** only.

## Repository layout

```
password_manager/
├── main.go              # Server, routing, optional --source import, UI embed
├── db/                  # SQLite schema and queries
├── models/              # Credential and history structs
├── routes/              # HTTP handlers
├── helpers/             # Import parsers, JSON helpers, CLI helpers
├── ui/                  # Svelte frontend (see ui/README.md)
└── db.db                # SQLite file (created at runtime)
```

## Development notes

- After UI changes: `cd ui && pnpm build`, then restart the Go process.
- UI development with hot reload: run `go run .` and `cd ui && pnpm dev` (see [ui/README.md](ui/README.md) for how the dev server relates to the API).

## License

Add a license file here if you plan to publish or share the project.
