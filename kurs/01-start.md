# Start

## Github Codespaces
Github Codespaces er en enkel måte å kjøre [editoren VSCode](https://code.visualstudio.com) direkte i nettleseren, slik at vi slipper å bruke tid på oppsett.

Gjennomfør disse stegene:

1. Gå til: https://github.com/arve0/ki/fork
2. Klikk på _Create fork_
   ![Oppretter en fork av dette repoet](fork.png)
3. Trykk på _Code_ → _Codespaces_ → _Create codespace on main_
   ![Lager en Codespace for repoet på ditt eget område](codespaces.png)
4. Åpne opprettet codespace:
   ![Codespace åpent i nettleser](codespace.png)

## Test at Copilot-chaten fungerer
### Du har Copilot-lisens?
1. Gå til https://github.com/settings/copilot/features
2. Se at du har *Pro*, *Pro+*, *Enterprise* eller tilsvarende. *Free* er ikke tilstrekkelig.

Dersom OK, kan du hoppe over neste steg og gå rett til [Oppgave: Sjekk at det virker](#oppgave-sjekk-at-det-virker).

### Du har ikke Copilot lisens
1. Gå til https://arve.dev/ki
2. Kopier API-nøkkelen
2. I VSCode, legg til OpenRouter API-nøkkel:
   1. Ctrl/Cmd + Shift + P
   2. > Chat: Manage Language Models
   3. + Add Models...
   4. OpenRouter
   5. Godta navnet OpenRouter
   6. Lim inn API-nøkkelen

### Oppgave: Sjekk at det virker
1. Åpne chat-vinduet, om det ikke er åpent (Ctrl + Alt/Cmd + I)
2. Finn og velg modellen *Claude Sonnet 4.6* (vi bruker denne modellen for alle oppgaver inntil vi ser på ulike modeller senere i kurset)
3. Skriv inn denne instruksen:

> lag et sammendrag av kurset

4. Høyreklikk på responsen og kopier.
5. Lagre til filen sammendrag.md.
6. Commit og push: `git add sammendrag.md && git commit -m "det virker!" && git push` (jeg bruker historikken i din fork til å lage scoreboarden)

Neste steg er [02-chatting.md](02-chatting.md).
