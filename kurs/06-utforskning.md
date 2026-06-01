# Utforskning
En ting jeg oppdaget tidlig er at agenten ofte har andre tilnærminger enn jeg har selv. Min første tanke var "idiot, du kan ikke", men en av gangene oppdaget jeg at agenten brukte noe funksjonalitet jeg ikke var klar over på et felt jeg anser meg selv som ekspert. En naturlig tanke er "joda, men jeg kan jo ikke vite om all funksjonalitet i alle produkter, grunnkunnskapen er den viktige", men en bedre holdning er å tenke "KI-modellen har hele verden som kunnskap, den vil alltid være mer kunnskapsrik enn meg".

Med denne innstillingen går ditt arbeid fra "å vite alt selv" til "jeg bruker hele verden som kunnskap til å ta mine valg". Dette har vært mulig før også, ved å tilegne seg verdens kunnskap i bøker, fra Google, på Stack Overflow og YouTube. Men nå er kostnaden gått fra timer til minutter for å bruke hele verdens kunnskap.

Til dere som sier: Du kan ikke stole på en KI-modell! Nei, du kan heller ikke stole på informasjon fra bøker og internett, du er bare vant til at innholdet ofte stemmer.

Så, la oss utforske!

## Oppgave: Hvordan bruke tidtakingene i ditt timeføringsprogram?
Tidtakingen sparer oss for noe manuelt vedlikehold av timelister på en lapp eller i et regneark, men den største tidstyven i timeføring er å legge inn timene i Enterprise Economy Ultra fra din favorittleverandør av programvare.

Her skal vi bruke en [skill](https://agentskills.io) for å hjelpe oss å finne ut hvordan problemet løses. En skill er tilsvarende AGENTS.md, men kun navnet og beskrivelsen legges til i kontekst. Agenten bestemmer selv når skillen skal tas i bruk. Å ta i bruk er her å legge hele innholdet til skillen i konteksten, altså utvide instruksen.

Vi starter med å legge til [grill me skill](https://github.com/mattpocock/skills/blob/main/skills/productivity/grill-me/SKILL.md):

```shell
mkdir -p "$(git rev-parse --show-toplevel)/.agents/skills/grill-me"
curl --location https://raw.githubusercontent.com/mattpocock/skills/refs/heads/main/skills/productivity/grill-me/SKILL.md --output "$(git rev-parse --show-toplevel)/.agents/skills/grill-me/SKILL.md"
```

Prøv nå denne instruksen:

> /grill-me Jeg ønsker å benytte tidtakingene til å automatisk fylle inn mine timelister i en webapplikasjon. Jeg ønsker at timene skal rundes til nærmeste halvtime. Nedlasting av tidtaking skal laste ned en måned. Jeg ønsker enten en plugin for nettleser eller et script for playwright. Det skal være mulig for meg å logge inn i timeføringssystemet selv, og deretter skal programmet gjøre timeføringen. For plugin kan timeføringen trigges av at jeg drar over filen jeg lastet ned. For playwright kan du foreslå hvordan en gir filen. Beskriv fordeler og ulemper med plugin vs playwright og hjelp meg til å utforske hvordan dette kan løses.

Merk: _/grill-me_ aktiverer skillen. Aktivering kan også skje ved at beskrivelsen til skill inneholder "bruk denne instruksen hver gang du lager git commits" eller at din instruks er "bruk grill-me skill".

Hvis den spør om hvilket timeføringssystem du bruker, se neste oppgave.

## Oppgave: Legge til del av nettside som kontekst
De fleste kjipe timeføringssystemer har ingen API-er, derfor har jeg valgt å automatisere via en plugin i nettleser. Grill-me er ikke deterministisk, men sannsynligvis vil den spørre om hvordan timeføringsprogram du bruker, samt om hvordan timene skal legges inn. Da kan du mate det spesifikke HTML-elementet fra timeføringsprogrammet som kontekst:

1. _Ctrl/Cmd + Shift + P_
2. _Browser: Open Integrated Browser_
3. Gå til og logg inn på timeføringssiden.
4. Trykk på knappen bak URL _Share with agent_ og del med agenten.
5. Prøv denne instruksen.

> jeg bruker timeføringssystemet som du ser på websiden

Dersom du får spørsmål om spesifikke element, lar neste knapp _Add element to chat_ et spesifikt HTML-element.

## Oppgave: Avslutte
Grill-me er omstendig og en kan sitte i en time og svare på spørsmål i noen tilfeller. Når du har gitt informasjon om hvordan timene skal føres inn, be om å avslutte:

> Det er nok nå. Lag en oppsummering av det du har så langt, og gjør antakelser for det du ikke vet. Skriv oppsummeringen her, men lagre resultatet også til eksport.md.

Lagre resultatet i git og push det.

## Når bruker en skills?
Tenk "jeg kommer til å bruke denne instruksen igjen, men den trengs ikke alltid". Jeg har brukt skills til å skrive gode git-commits, opprette pull requests basert på commits, oppgradere java, standardisere oppsett, sørge for at siste versjon benyttes for java-biblioteker, oppgradere github workflows, osv.

## Tips for skills
1. En skill er bare en vanlig instruks, tekst du kan skrive selv.
2. KI-modellen kan veldig mye, så du kan være unøyaktig i beskrivelsene.
3. En kan be om input fra brukeren underveis.
4. Skills kan installeres i din hjemmekatalog, slik at du kan benytte de på tvers av prosjekter.
5. Brukere som vet input up-front kan gi de i initiell instruks: "upgrade java, case number is GLAD-491"
6. Det finnes mange kataloger over skills på nett du kan prøve, eksempelvis [skills.sh](https://skills.sh).
7. KI-modeller er flinke til å skrive skills, her er en [skill for å skrive skills](https://github.com/mattpocock/skills/blob/main/skills/productivity/write-a-skill/SKILL.md) 🤓


Neste steg er [07-planer.md](07-planer.md).
