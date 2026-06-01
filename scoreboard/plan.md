# Scoreboard – Plan

## Problem
Kurset trenger en scoreboard som vises på storskjerm under kursøkten. Den skal vise deltagerenes fremdrift i sanntid og spille en gamification-lyd når noen fullfører et steg.

## Arkitektur

```
scoreboard/
  server.js       – Node.js HTTP-server (kun innebygde moduler)
  public/
    index.html    – Scoreboard-side (HTML/CSS/JS)
    sounds.js     – Web Audio API lyd-generator (ingen ekstern lib)
```

### server.js
- Innebygd `node:http` – ingen Express eller andre biblioteker
- Serverer `public/index.html` og statiske filer
- Server-Sent Events (SSE) endpoint `GET /events` pusher tilstandsoppdateringer til nettleseren
- `GET /state` returnerer nåværende tilstand som JSON (for initial load)
- Bakgrunnsjobb hvert **60. sekund**: henter nye deltakere (forks) via `gh api`
- Bakgrunnsjobb hvert **10. sekund**: sjekker fremdrift for hver deltaker via `gh` CLI
- Når en deltaker fullfører nytt steg: sender SSE-event `progress` med deltaker og steg

### public/index.html
- Henter initial tilstand fra `/state`, deretter lytter på `/events` (SSE)
- CSS Grid som justerer kolonner/fontstørrelse basert på antall deltakere
- Hvert deltakerkort viser: brukernavn, avatar og fremdrifts-badges per modul
- Spiller lyd via `sounds.js` ved mottatt `progress`-event

### public/sounds.js
- Bruker **Web Audio API** (`AudioContext`, `OscillatorNode`, `GainNode`) – **ingen ekstern lib**
- Genererer en chiptune "power-up"-sekvens: en rekke stigende toner (f.eks. C4→E4→G4→C5) med kort varighet og fallende volum (decay)
- Eksponert som `playPowerUp()` funksjon

## Deltakere – oppdagelse
Bruker `gh api /repos/arve0/ki/forks --paginate` hvert minutt.
Returnerer liste med `{login, avatar_url, html_url}` per fork-eier.
Nye brukere legges til tilstanden; eksisterende berøres ikke (tilstand bevares).

## Fremdriftsdeteksjon per modul

Deteksjon er basert på signaler som allerede finnes i kursmaterialet – ingen ekstra script kreves av deltakere. All matching er **case-insensitive** (toLowerCase på begge sider). Alle API-kall gjøres via `gh api`.

### Deteksjonsmetoder (prioritert rekkefølge)

1. **Fileksistens** – `gh api /repos/{login}/ki/contents/{path}` → HTTP 200 = fullført, 404 = ikke gjort. Rask og stabil.
2. **Commit-melding fuzzy-match** – `gh api /repos/{login}/ki/commits?path={path}&per_page=20` → iterer over `commit.message`, lowercase, sjekk om noen av søkeordene finnes med `includes()`. Brukes der ingen naturlig fil finnes.
3. **Filinnhold fuzzy-match** – Hent fil via API, base64-decode, lowercase, sjekk at søkeord finnes. Brukes kun der innhold skiller mellom to moduler.

---

### Modul 01 – Start: Sjekk at det virker

**Kursoppgave:** `git add sammendrag.md && git commit -m "det virker!" && git push`

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Fileksistens | `sammendrag.md` finnes i rot |
| 2 | Commit fuzzy | Siste 20 commits i rot, melding inneholder `"virker"` |

---

### Modul 02 – Chat: Lagre resultatet

**Kursoppgave:**
```shell
git add tidtaker.md
git commit -m "utforske en kodebase og dens teknologier"
git push
```

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Fileksistens | `tidtaker.md` finnes i rot |
| 2 | Commit fuzzy | Commits på `tidtaker.md`, melding inneholder `"utforske"` eller `"kodebase"` eller `"teknologi"` |

---

### Modul 03 – Kontekst: Lagre resultater

**Kursoppgave:**
```shell
git add tidtaker/templating.md
git commit -m "legge til kontekst"
git push
```

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Fileksistens | `tidtaker/templating.md` finnes |
| 2 | Commit fuzzy | Commits på `tidtaker/templating.md`, melding inneholder `"kontekst"` eller `"templating"` |

---

### Modul 04 – Agenter: AGENTS.md + datoformat-endring

**Kursoppgave:** Deltaker oppretter `AGENTS.md` manuelt, deretter lar agenten gjøre en kodeendring (datoformat) og kjører agent-generert `git commit` + `git push`.

Agenten følger sin AGENTS.md-instruksjon og lager en commit med forklaring om datoformatendringen. Typiske ord i meldingen: "dato", "format", "norsk", "april", "juni", "dag".

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Fileksistens | `AGENTS.md` finnes i rot |
| 2 | Commit fuzzy | Commits på `tidtaker/`, melding inneholder `"dato"` eller `"format"` eller `"norsk"` eller `"april"` eller `"juni"` |

---

### Modul 05 – Feilsøking: Fiks timer-bug

**Kursoppgave:** Deltaker sender feilbeskrivelse til agenten. Agenten (via AGENTS.md-instruksjon) gjør kodeendring og lager commit. Ingen eksplisitt commit-melding er gitt i kursmaterialet – meldingen genereres av agenten.

Feilen: «Tiden teller ikke når en starter en tidtaker, men dersom en oppfrisker siden vises korrekt tid.»
Forventede ord i agent-generert commit-melding: "teller", "timer", "feil", "fiks", "live", "oppdater", "htmx", "sanntid", "tick", "count", "bug", "refresh".

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Commit fuzzy | Commits på `tidtaker/` (siste 20), melding inneholder `"teller"` eller `"timer"` eller `"feil"` eller `"fiks"` eller `"live"` eller `"htmx"` eller `"oppdater"` eller `"bug"` eller `"tick"` eller `"count"` |

> **Merk:** Modul 04 og 05 genererer begge commits på `tidtaker/`. Ordsettet overlapper minimalt: modul 04 handler om datoformat, modul 05 om timer/live-oppdatering. Er begge sett av ord funnet, regnes begge som fullført.

---

### Modul 06 – Utforskning: Grill-me skill + eksport.md

**Kursoppgave:**
```shell
git commit -m "la til grill-me skill"
git push
```
Deretter: "Lagre resultatet i git og push det" (eksport.md, ingen fast melding).

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Fileksistens | `.agents/skills/grill-me/SKILL.md` finnes i fork |
| 2 | Commit fuzzy | Commits, melding inneholder `"grill"` eller `"skill"` |
| 3 | Fileksistens | `eksport.md` eller `tidtaker/eksport.md` finnes |

---

### Modul 07 – Planer: Utvidet eksport.md med Playwright-detaljer

**Kursoppgave:**
```shell
git add tidtaker/eksport.md
git commit -m "flere detaljer for eksport/import"
git push
```

Filen `tidtaker/eksport.md` finnes fra modul 06, men innholdet utvides med playwright-interaksjoner i modul 07. Filinnhold-sjekk skiller de to modulene.

| Prioritet | Metode | Signal |
|-----------|--------|--------|
| 1 | Filinnhold fuzzy | `tidtaker/eksport.md` finnes OG innhold inneholder `"playwright"` |
| 2 | Commit fuzzy | Commits på `tidtaker/eksport.md`, melding inneholder `"eksport"` og `"import"` eller `"detaljer"` |

---

### Modul 08 – Veien videre

Modul 08 er referansemateriell uten konkrete commit-oppgaver. Ingen automatisk deteksjon – markeres som fullført manuelt av kursholder eller settes alltid som «bonus/ekstra».

## Tilstandsmodell (in-memory i server.js)

```js
{
  participants: {
    "githublogin": {
      login: "githublogin",
      avatarUrl: "https://...",
      completedModules: [1, 2, 3],   // array av fullførte modulnummer
      lastChecked: "2026-05-19T...", // ISO timestamp
    }
  }
}
```

Tilstanden beholdes i minnet – ingen database. Restartes serveren, bygges tilstanden opp på nytt (tar maks 10 sekunder).

## HTML-layout

- Bakgrunn: mørk (gaming-estetikk)
- CSS Grid: `grid-template-columns: repeat(auto-fill, minmax(Xrem, 1fr))` – tilpasser seg automatisk
- Fontstørrelse skaleres: `font-size: clamp(0.8rem, 2vw, 1.4rem)`
- Deltakerkort: avatar, brukernavn, progress-badges (emoji per modul: ✅ fullført, ⬜ ikke gjort)
- Liten animasjon (pulse/glow) på nylig fullført badge
- SSE reconnect-logikk i nettleseren

## Lyd

Implementeres med **Web Audio API** direkte i nettleseren, ingen server-kall, ingen ekstern lib:

```
playPowerUp():
  For note in [C4, E4, G4, C5]:
    OscillatorNode type=square, frequency=note
    GainNode envelope: attack 5ms, sustain, decay 150ms
    start ved t + offset (50ms mellom toner)
```

Alternativt kan du generere lydfila selv med [BeepBox](https://www.beepbox.co/) – nettleserbasert chiptune-editor, eksporter som .wav og legg i `public/`.

## Todos

1. Lag `scoreboard/server.js` med HTTP-server, SSE, polling av GitHub API
2. Lag `scoreboard/public/index.html` med grid-layout og SSE-klient
3. Lag `scoreboard/public/sounds.js` med Web Audio API power-up lyd
4. Test at ny deltaker dukker opp automatisk
5. Test at fullføring av modul trigger lyd og oppdaterer board
6. (Valgfritt) Legg til `scoreboard/README.md` med kjøre-instrukser

## Kjøring

```shell
cd scoreboard
node server.js
# Åpne http://localhost:3000 i nettleseren
```
