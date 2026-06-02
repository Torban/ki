# Tidtaker

En tidsregistreringsapplikasjon bygget med PocketBase og HTMX.

## Funksjoner

- **Autentisering** – Registrering og innlogging med e-post/passord
- **Timer** – Start/stopp tidtaking med ett klikk
- **Beskrivelser og tags** – Legg til beskrivelse og tagger på registreringer
- **Redigering** – Endre start-/stopptid og beskrivelse i etterkant
- **Uendelig scroll** – Bla gjennom alle registreringer med automatisk lasting
- **Søk** – Fulltekstsøk på beskrivelse og tagger (AND-basert)
- **Filter og sortering** – Filtrer på tagger, sorter på tid eller beskrivelse
- **Norsk datoformat** – Datoer vises på formen «21. april 2026»

## Stack

| Komponent | Teknologi |
|-----------|-----------|
| Backend   | [PocketBase](https://pocketbase.io) (Go) |
| Frontend  | [HTMX](https://htmx.org) + HTML |
| Styling   | [Tailwind CSS](https://tailwindcss.com) (CDN) |
| Testing   | [Playwright](https://playwright.dev) (E2E) |

## Kom i gang

### Forutsetninger

- Go 1.21+
- Node.js 18+ (for E2E-tester)

### Kjør utviklingsserver

```bash
# Klon og gå inn i prosjektet
cd tidtaker

# Start utviklingsserver
./scripts/dev.sh

# Eller manuelt:
go run . serve --http=127.0.0.1:8090 --dev
```

Åpne [http://localhost:8090](http://localhost:8090) i nettleseren.

### Bygg

```bash
# Bygg for gjeldende plattform
./scripts/build.sh 1.0.0

# Bygg for Linux (deployment)
./scripts/build-linux.sh 1.0.0
```

## Prosjektstruktur

```
tidtaker/
├── main.go                  # Inngangspunkt, PocketBase-oppsett
├── routes.go                # Rute-registrering
├── handlers_auth.go         # Auth-handlere (register, login, logout)
├── handlers_timer.go        # Timer-handlere (start, stop, edit, søk)
├── templates.go             # Template-initialisering og -hjelpefunksjoner
├── pb_migrations/           # Database-migrasjoner
│   └── timings.go           # Timings-collection
├── templates/               # HTML-templates (embedded)
│   ├── base.html            # Felles layout
│   ├── auth/
│   │   ├── login.html
│   │   └── register.html
│   ├── timer/
│   │   ├── timer_page.html
│   │   ├── timer_controls.html
│   │   ├── timing_item.html
│   │   ├── timing_edit.html
│   │   └── timings_list.html
│   └── search/
│       ├── search_bar.html
│       └── filter_bar.html
├── static/                  # Statiske filer (embedded)
│   └── js/htmx.min.js
├── scripts/                 # Hjelpeskript
│   ├── dev.sh               # Start utviklingsserver
│   ├── build.sh             # Bygg lokal binær
│   ├── build-linux.sh       # Bygg for Linux
│   └── deploy.sh            # Deploy til Hetzner
└── tests/e2e/               # E2E-tester (Playwright)
```

## API-endepunkter

| Metode | Sti | Beskrivelse |
|--------|-----|-------------|
| POST | `/api/timings/start` | Start ny timing |
| POST | `/api/timings/:id/stop` | Stopp timing |
| POST | `/api/timings/:id/add-tag` | Legg til tag |
| POST | `/api/timings/:id/remove-tag` | Fjern tag |
| PATCH | `/api/timings/:id/edit-time` | Endre start-/stopptid |
| PATCH | `/api/timings/:id/edit-description` | Endre beskrivelse |
| GET | `/api/timings/list` | Hent liste (paginert) |
| GET | `/api/timings/search` | Søk med filter |
| DELETE | `/api/timings/:id` | Slett timing |

## Deployment

Se `scripts/deploy.sh` for automatisert deployment til Hetzner med systemd.

```bash
./scripts/deploy.sh 1.0.0
```

## E2E-tester

```bash
cd tidtaker
npm install
npx playwright test
```

## Feilsøking

- **Server starter ikke:** Sjekk at port 8090 er ledig
- **Migrasjoner feiler:** Slett `pb_data/`-mappen og start på nytt
- **Templates oppdateres ikke:** Templates er embedded – bygg på nytt med `go build`

## Utvikling

### Legge til nye features

1. Legg til nye felter i migrasjonen (`pb_migrations/`)
2. Legg til handlers i `handlers_*.go`
3. Registrer ruter i `routes.go`
4. Lag HTML-templates i `templates/`
5. Bygg og test: `go build && go run . serve --dev`

### Søk-implementering

Søk er server-side og AND-basert: hvert mellomrom-separert ord må matche enten beskrivelse eller tags. Eksempel: "arbeid sprint" matcher registreringer med både "arbeid" OG "sprint".
