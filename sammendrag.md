# Sammendrag av kurset

Kurset lærer deg å bruke KI-agenter effektivt i programvareutvikling, steg for steg.

**01 – Start:** Miljø settes opp i GitHub Codespaces og VS Code. Du verifiserer at Copilot Chat fungerer og gjør en første oppgave.

**02 – Chat:** Chat passer til raske spørsmål, små kodeendringer og utforsking. Du blir kjent med eksempelapplikasjonen *tidtaker* og dens teknologier (Go, PocketBase, htmx, Tailwind).

**03 – Kontekst:** Dårlige svar skyldes ofte manglende kontekst. Du lærer å legge ved filer, utvalgte linjer og informasjon fra internett for å styre modellen riktig.

**04 – Agenter:** Agentmodus lar deg gjøre komplette oppgaver i én instruks: kodeendring, test (TDD), dokumentasjon, sikkerhetsvurdering og commit-melding. Gjentatte instrukser lagres i `AGENTS.md` så de alltid er med.

**05 – Feilsøking:** Du bruker agenten direkte på feilmeldinger og stack traces. Gode beskrivelser av feilen gir treffsikre løsninger.

**06 – Utforskning med skills:** Skills er gjenbrukbare instrukser for oppgaver du ikke trenger hele tiden. *Grill me*-skillen brukes til å utforske løsninger gjennom dialog, for eksempel automatisk timeføring via nettleser.

**07 – Planer:** En god plan beskriver arbeidet så tydelig at agenten kan jobbe selvstendig. Du bruker Playwright MCP og OpenCode CLI for å samle nok detaljer til å implementere automatisk timeføring.

**08 – Veien videre:** Tips for å få gode resultater uten å sitte og passe på agenten: tydelig verifisering, ende-til-ende-tester, `AGENTS.md`, parallelt arbeid med git worktrees og subagenter, håndtering av fullt kontekstvindu, diktering, og bevissthet rundt kostnad og kodekvalitet.