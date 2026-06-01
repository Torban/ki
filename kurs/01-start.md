# Start

## Github Codespaces
Github Codespaces er en enkel måte å kjøre et prosjekt i en VM på Github, slik at vi slipper å bruke tid på oppsett. Vi skal koble til VM med [editoren VSCode](https://code.visualstudio.com), så det må du ha installert.

Gjennomfør disse stegene:

1. Gå til: https://github.com/arve0/ki/fork
2. Klikk på _Create fork_
   ![Oppretter en fork av dette repoet](fork.png)
3. Trykk på _Code_ → _Codespaces_ → _Create codespace on main_
   ![Lager en Codespace for repoet på ditt eget område](codespaces.png)
4. Opprettet codespace åpner i nettleser:
   ![Codespace åpent i nettleser](codespace.png)
5. Trykk helt nede i venstre hjørne der det står _Codespaces: et navn_
6. Velg _Open in VS Code Desktop_ i dialogen som kommer opp.

![åpner i lokal editor](vscode-desktop.png)


## Test at Copilot-chaten fungerer
### Du har Copilot-lisens?
1. Gå til https://github.com/settings/copilot/features
2. Se at du har *Business*, *Pro*, *Pro+*, *Enterprise* eller tilsvarende. *Free* er ikke tilstrekkelig.

![Copilot Business via Webstep](copilot-business.png)

Dersom OK, kan du hoppe over neste steg og gå rett til [Oppgave: Sjekk at det virker](#oppgave-sjekk-at-det-virker).

### Du har ikke Copilot lisens
Dersom du ikke har Copilot-lisens, kan du bruke en API-nøkkel som gir deg tilgang til
samme KI-modell via [OpenRouter](https://openrouter.ai).

1. Be kursholder om API-nøkkel
2. I VSCode, legg til OpenRouter API-nøkkel:
3. Ctrl/Cmd + Shift + P
4. `> Chat: Manage Language Models`
5. `+ Add Models...`
6. OpenRouter
7. Godta navnet OpenRouter
8. Lim inn API-nøkkelen

### Oppgave: Sjekk at det virker
1. Åpne chat-vinduet, om det ikke er åpent (Ctrl + Alt/Cmd + I)
2. Finn og velg modellen *Claude Sonnet 4.6* (vi bruker denne modellen for alle oppgaver inntil vi ser på ulike modeller senere i kurset)
3. Skriv inn denne instruksen:

> lag et sammendrag av kurset

4. Høyreklikk på responsen og kopier.
5. Lagre til filen sammendrag.md.
6. Commit og push: `git add sammendrag.md && git commit -m "det virker!" && git push` (jeg bruker historikken i din fork til å lage scoreboarden)

Neste steg er [02-chatting.md](02-chatting.md).
