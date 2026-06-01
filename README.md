# Kunstig intelligens for utviklere
Jeg har gått gjennom standard løype for å bli kjent med ki-assistert utvikling;

1. Starter med chatting.
2. Tar i bruk agenter.
3. Bedre kontrollerer hva jeg ønsker.

Dette er et kurs som går gjennom hva jeg har lært, steg for steg, slik at andre kan komme raskere i gang.

## Hva trenger du for å gjennomføre kurset?
1. En Github-bruker.
2. Abonnement for Copilot hos Github.
3. [VSCode](https://code.visualstudio.com) installert på din PC.

> Kurset legger opp til å bruke Github 100%, så vi slipper bruke tid på oppsett, men du kan gjennomføre kurset lokalt på din egen PC med eget KI-abonnement også, på egen risiko.

## Hvem er kurset for?
### Utvikler uten erfaring med KI-agenter
Få en rask innføring i hva som er mulig og hvordan du gjør det.

### Utvikler med erfaring med AI-agenter
Del erfaringer med andre, snakk med likesinnede, hjelp andre som er ny.

### Ikke-utvikler som lurer på hva det handler om
Følg anvisningene og erfar hvordan utvikling kan gjøres med vanlig norsk, uten detaljert kjennskap til programmeringsspråk.

## Kurset
[Start her!](kurs/01-start.md)

### Innhold

#### Del 1 – Oppsett og kom i gang
Sett opp GitHub Codespaces og koble til Copilot eller OpenRouter.
- Opprett en fork og åpne kodebasen i nettleseren med Codespaces
- Koble til Copilot (Pro/Pro+/Enterprise) eller OpenRouter med API-nøkkel
- Sjekk at chat fungerer ved å lage et sammendrag av kurset

#### Del 2 – Chat som verktøy
Chat er raskere enn å google fordi den forstår din kontekst.
- Spør om ukjent teknologi og kodebasen med naturlig norsk
- Gjøre små kontekstsensitive endringer i koden
- Utforske og forstå eksempelapplikasjonen _tidtaker_

#### Del 3 – Kontekst
Dårlige svar skyldes ofte manglende kontekst – her lærer du å styre modellen.
- Legg til filer som kontekst for mer presise svar
- Marker enkeltlinjer for å holde modellen fokusert
- Hent oppdatert informasjon fra internett for å unngå utdaterte svar

#### Del 4 – Agenter
Agenter kan gjøre mer enn å svare – de kan endre koden din.
- La agenten endre implementasjon og kjøre eksisterende tester
- Gi agenten en full arbeidsliste i én instruks (TDD, dokumentasjon, sikkerhetsvurdering, commit)
- Bruk `AGENTS.md` for å unngå å gjenta de samme instruksene om og om igjen

#### Del 5 – Feilsøking
KI er svært god på feilmeldinger og stack traces.
- Gi agenten feilbeskrivelsen og la den finne og fikse årsaken

#### Del 6 – Utforskning og skills
Bruk KI som sparringspartner og utvid agenten med gjenbrukbare skills.
- La agenten utforske løsninger du ikke hadde tenkt på selv
- Installer og bruk _grill-me_ skill for strukturert problemutforskning
- Legg til HTML-elementer fra nettleseren som kontekst
- Lag og del egne skills for gjenbruk på tvers av prosjekter

#### Del 7 – Planer
En god plan lar agenten jobbe selvstendig over lengre tid.
- Lag et detaljert plandokument agenten kan jobbe fra uten avbrudd
- Bruk Playwright MCP-server for å styre nettleseren med naturlig språk
- Bruk OpenCode CLI som alternativ til Copilot Chat i terminalen
- La agenten implementere planen fra start til slutt

#### Del 8 – Veien videre
Forslag til vei videre, tips og triks som er kjekt å kjenne til.
