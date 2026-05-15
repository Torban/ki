# Planer
Utforskningen fra forrige steg er et fint utgangspunkt for å lage en plan. En plan er et skriftlig selvstendig dokument som beskriver hva som skal lages, og som kan agenten kan jobbe med uten at du trenger å gi input. Her sparer en seg mye tid, for en har beskrevet et større stykke arbeid som agenten kan jobbe uavhengig med.

Hvis du ikke gjorde det i forrige steg, lagre oppsummeringen nå:

>  Lagre oppsummeringen til eksport.md.

## Oppgave: Vurder om det er nok informasjon
Skumles eksport.md. Hvis du selv skulle implementert denne funksjonaliteten, ville du hatt nok informasjon til å lage den 100% uten å bruke andre kilder?

Hvis du ikke holdt på lenge og var nøye med å gi den elementer fra timeføringsprogrammet, tipper jeg den ikke inneholder nok informasjon. Spesielt slike detaljer som hvordan en skal klikke mellom uker og hvilke felt som må fylles ut.

Hvis svaret ditt er "ja", da har du en plan 🙌 Hvis ikke, gjør oppgavene under.

## Playwright MCP-server og OpenCode CLI
Her skal vi bruke to verktøy for å utvide planen:

1. [Playwright MCP-server](https://playwright.dev/docs/getting-started-mcp)
2. [OpenCode CLI](https://opencode.ai)

Playwright MCP-server gir oss muligheten til å styre en nettleser med naturlig språk, "gå til den siden", "trykk på det", osv. OpenCode CLI er en AI-agent, tilsvarende Copilot.

Grunnen til at vi bruker OpenCode CLI er fordi organisasjonen som har ditt Copilot-abonnement kan velge å skru av MCP-støtte i VSCode dersom du har enterprise-abonnementet. Dette omgår vi ved å bruke OpenCode, og får samtidig erfaring med å bruke agenten fra kommandolinje. Det er ingen stor forskjell fra det vi har gjort til nå, det er fortsatt en chat, en gir instrukser og en kan legge til filer og nettsted som kontekst, den støtter AGENTS.md og skills, men grensesnittet er noe ulikt. Eksempelvis legger en til filer med `@filnavn` istedenfor å trykke på knappene i VSCode.

## Oppgave: Logge på timeføringsprogrammet på egen PC
Agenten kan hente nettsider selv, men da har den ikke dine cookies og innlogginger. For å la agenten se timeføringsprogrammet og klikke i det, skal vi bruke en [Playwright](https://playwright.dev/) først lokalt for å logge på.

Siden vi kjører oppgavene i Codespaces som er en VM på Github, må vi logge på timeføringsprogrammet på egen PC, så overføre cookies til Codespaces. Codespaces kan åpne nettlesere via VNC eller remote X11, men det er knotete. Dette er enklere:

1. Åpne en terminal på egen PC.
2. Installer playwright: `npm install --global playwright && npx playwright install chrome`
3. Start playwright i opptaksmodus: `npx playwright codegen --save-storage=auth.json`
4. Naviger til timeføringssiden.
5. Logg inn.
6. Lukk nettleseren.
7. Verifiser at du har en auth.json med cookies.
8. Kopier innholdet i filen og opprett en fil med samme navn i Codespaces.

Merk: For å gjennomføre timeføringen kunne vi benyttet denne metoden for å finne riktige ting å trykke på, men vi skal heller la agenten finne ut av det selv.

## Oppgave: Koble OpenCode til AI-leverandør
1. I Codespaces, åpne terminalen
2. Start `opencode`
3. Skriv `/connect` og trykk enter
4. Følg anvisningen for å logge på Github eller OpenRouter

## Oppgave: Verifisere at Playwright fungerer fra OpenCode i Codespaces
I OpenCode, skriv denne prompten:

> åpne <url> med playwright, bruk cookies fra auth.json, la skjermstørrelse være 1024x768, lagre skjermbilde til timeføring1.png

Åpne skjermbildet og sjekk at du er innlogget. Fortsett med noe som dette:

> Naviger til Hjem / Timelister / Timeregistrering. Lagre skjermbilde til timeføring2.png

Hvis OK, prøv:

> Legg inn en rad med timeføring. Trykk på feltene og velg første alternativ på valg som blir synlig. Legg inn teksten "AI-kurs" på beskrivelse og 1.5 timer tirsdag. Lagre skjermbilde til timeføring3.png

Dette er informasjonen som planen trenger, be agenten legge det til planen:

> @tidtaker/eksport.md utvid planen med aktuelle ting en må trykke på for å gjennomføre timeføringen. Bruk selektorer som sannsynligvis ikke endrer seg, slik som tekst eller titler.

Lagre endringene i eksport.md:

```shell
git add tidtaker/eksport.md
git commit -m "flere detaljer for eksport/import"
git push
```

Neste steg er [08-gode-afk-resultat.md](08-gode-afk-resultat.md).
