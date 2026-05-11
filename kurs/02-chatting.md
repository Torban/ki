# Chat
Chat er en enkel måte å gjøre små oppgaver, der du jobber interaktivt med agenten.

Passer bra til:
- Spørre om ting. Alt fra ting som vil finnes i dokumentasjon på et bibliotek til konkrete spørsmål om kodebasen. Eksempler:

  > Hvilken verktøy brukes i go for å formatere kode?
  > Hva gjør denne kodebasen?
  > Hvordan kan jeg teste denne koden?

- Gjøre små endringer som er kontekstavhengig. Eksempler:

  > Omdøp alle ..
  > Flytt denne seksjonen av koden.
  > Koden er repetiv, gjør koden DRY.
  > Endre URL for https://tjeneste.apps.test.firma.no til http://tjeneste.namespace.cluster.local.

- Utforske hvordan du skal løse et problem. Eksempler:

  > Jeg har denne feilen her, hva kan være problemet?
  > Jeg skal lage en ny ting, hvordan kan jeg gjøre det?

## Norsk eller engelsk?
Min erfaring er at norsk er like godt som engelsk, samt at min hjerne er flinkere å produsere fullstendige setninger på norsk og at det er enklere å diktere. Altså foretrekker jeg norsk når jeg chatter. Har du fått et dårlig svar og ønsker å prøve på engelsk, kan du enkelt ta den lange instruks og oversette den, før en prøver igjen på engelsk:

> Oversett dette til engelsk: ....

## Kodebase
Kurset baserer seg på kodebasen [tidtaker](../tidtaker/), som er et enkelt program for å holde oversikt over arbeidstimene sine. _Tidtaker_ er grunnlaget for oppgavene som kommer.

## Oppgave: Bli kjent med kodebasen
Start en ny chat og prøv disse instruksene:

> I mappen tidtaker ligger det et program. Forklar meg hva det gjør, hvilken teknologi det bruker og hvordan jeg kommer i gang.

## Oppgave: Bli kjent med teknologien
Var det noe nytt her som du ikke kjenner til? Prøv noe som dette:

> Kan du forklare meg mer om teknologien x, y og z?

## Oppgave: Lagre resultatet
Høyreklikk, velg _Copy all_, lagre til tidtaker.md, commit og push:

```shell
git add tidtaker.md
git commit -m "utforske en kodebase og dens teknologier"
git push
```

Neste steg er [03-kontekst.md](03-kontekst.md).

