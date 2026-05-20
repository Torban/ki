# Veien videre
Nå er den strukturelle delen av kurset ferdig. Under er noen forslag til vei videre, tips og triks som er kjekt å kjenne til.


# Gode AFK resultater
Nå "koker" agenten vår med den nye funksjonaliteten, og vi kan ta en kaffe, jobbe med noe annet eller gå hjem. Det vi ønsker å optimalisere i dette steget er at når agenten er ferdig, så skal resulatet være bra. Vi har allerede gjort noen grep her, ved å lage [AGENTS.md](../AGENTS.md).

Her er noen ting jeg bruker å tenke på:

1. Bruker jeg en agent som for lov å jobbe lenge, uten å bli avbrutt? Chat: nei, CLI: ja.
2. Kan agenten kjøre kommandoer uten å be om lov? `--yolo` i en sandboks.
3. Vet agenten hvordan arbeidet skal verifiseres? Har en instruks om hvordan det verifiseres. Jeg har guidet agenten til å lage gode tester som beskriver hensikt (e2e).
4. Kunne jeg gitt flere instrukser, slik at agenten gjør mest mulig arbeid? Kan den også kjøre sikkerhetsvurdering, benchmark, forberede en commit-melding, foreslå arbeid videre, osv.

For denne planen, tipper jeg at det ikke fungerer ut av boksen. Det er fordi vi har ikke laget noe oppsett for at agenten kan sjekke selv om det fungerer. Du kan nå velge om du chatter videre, for å korrigere:

> når jeg gjør x, så skjer det ingenting, kan du finne feilen og korrigere?

...eller starte på en ny chat for å utforske hvordan det kan testes:

> er det mulig å teste chrome plugins med playwright?

Akkurat her er det vanskelig på grunn av at jeg ikke ønsker å dele innloggingen med KI, men en kan kanskje gjøre noe som dette?

> kan du lage meg et playwright-script jeg kan kjøre for å holde sesjonen i auth.json aktivt, slik at jeg ikke blir logget ut?

Dette er hva jeg bruker mest hjernekapasitet på i mitt arbeid nå. Hvordan kan jeg hjelpe agenten til å verifisere sitt eget arbeid?


# Jobbe med flere agenter samtidig
Klør det i fingrene når du venter på agenten som jobber? Da er det riktig tidspunkt å starte med en sideoppgave.

Det er vanskelig å gjøre mange oppgaver samtidig i samme repo uten at agentene går i bena på hverandre, så du trenger en måte å isolere arbeidet fra hverandre. [Git worktrees](https://git-scm.com/docs/git-worktree) løser akkurat dette problemet.

Kort sagt er git worktrees en lokal kopi av en git branch. Altså cirka dette:

```shell
git checkout -b ny-branch
mkdir ../ny-branch
cp -r . ../ny-branch
```

Jeg anbefaler å sjekke ut [_Agents View_  i VSCode](https://www.youtube.com/watch?v=DC_z7VjJCJM), men dersom du foretrekker terminal og gjøre det selv kan du gjøre noe slik som dette:

```shell
# lag en ny branch med worktree og sjekk ut i en mappe
git worktree add -b en-fiks ../en-fiks

# jobb selvstendig i mappen
code ../en-fiks

# lagre endringene på branchen i worktree
cd ../en-fiks
echo "" > hypotetisk-fiks
git add --all .
git commit -m "jeg fikset det"

# branch er tilgjengelig fra hovedmappen Lovisa
cd - # gå tilbake
git push origin en-fiks      # kan også gjøres fra worktree, dette er bare et eksempel som viser at branch er tilgjengelig i repo-mappen
gh pr create --head en-fiks

# fjern lokal utsjekking av branchen når du er ferdig (fra repo-mappen)
git worktree remove ../en-fiks
```


# Agenten gjør ikke som jeg vil
Dette er de tre vanligste variantene av at KI-modellen feiler på noe vis:

1. Den gjetter og gjør dårlig arbeid.
2. Den kjører seg fast og kommer ikke videre.
3. Resultatet virker ikke.

## Dårlige antakelser og dårlige utfall
Antakelser og overraskende løsninger er vanlig når instruksen er for dårlig. Går det helt i feil retning er det vanskelig å korrigere, fordi konteksten til forespørselen fylles med det dårlige forslaget. Eksempelvis fungerer det sjeldent å korrigere etterpå, altså dette fungerer ikke:

> Du implementerte med React, men jeg foretrekker at man bruker vanilla js.

Da er det bedre å stoppe agenten, gå tilbake til forrige sjekkpunkt og skrive instruksen på ny. Eksempelvis legge til det som var viktig for deg, eller be agenten skrive en plan, gjennomgå/korrigere den og så be agenten jobbe med implementasjonen.

## Klarer ikke løse problemet
En gang i blant vet ikke KI-modellen svaret "direkte". Altså klarer den ikke å bruke treningssettet sitt til å umiddelbart løse problemet. Istedenfor å la den koke lenge er det bedre å styre agenten frem til løsningen.

Her hjelper det med reel erfaring om hvordan problemet løses, men en kan også få hjelp til gode strategier for å feilsøke:

> Vi sitter fast og finner ikke en løsning som fungerer. Ta et steg tilbake og tenk på hvordan vi kan feilsøke og finne ut av hva som er feil. Kan vi bruke noen verktøy eller metoder for å utelukke løsninger som er feil? Foreslå tre veier videre.

## Resultatet virker ikke
Agenten er ferdig og den sier "nå er det ferdig og alt fungerer". En lystløgner med andre ord.

Nesten alltid er dette på grunn av at agenten ikke har noen god måte å verifisere at hensikten eller målet med oppgaven er nådd. Ofte må en være ganske presis i instruksen sin, mer rom for tolkning, mindre presisjon.

Sammenligne disse:

> Implementer funksjonaliteten.

> Implementer funksjonaliteten. Sørg for at alt virker.

> Implementer funksjonaliteten. Når du er ferdig skal du kjøre alle testene.

> Implementer funksjonaliteten. Når du er ferdig skal du kjøre alle testene. Kjør opp server og sjekk at endepunkt /a nå inneholder teksten "ny funksjonalitet", først da er du ferdig.

> Start ved å skrive en test som sjekker at endepunkt /a inneholder teksten "ny funksjonalitet". Sørg for at testen feiler (TDD red). Implementer funksjonaliteten. Kjør testen på ny. Sørg for at testen er OK (TDD grønn). Verifiser at alle andre tester også kjører og er grønn.


# Problemer med fyllt kontekst
Hvis du jobber lenge i en og samme chat, eller du jobber i en stor kodebase der agenten leser mye kode, kan en havne i en situasjon der instrukser ikke lenger fungerer. KI-modeller har en øvre grense på hvor mye informasjon og hvor mange instrukser en og samme sesjon kan inneholde.

I VSCode kan du se hvor mye av kontekstvinduet som er i bruk, her 9% av 200K tokens:

![9% av kontekstvindu er i bruk](kontekstvindu.png)

De fleste agentene prøver å fikse fullt kontekstvindu for deg ved å kjøre en automatisk komprimering. Dette kan du også selv trigge ved å skrive `/compact`. Kompringering er essensielt "lag et sammendrag av denne sesjonen", der kall etter komprimering kun har med sammendrag + nye chats. Altså sendes ikke tidligere dialog med, slik som til vanlig.

Et problem med komprimeringen er at sammendraget fjerner informasjon som du anser som viktig. Det ligner litt på å gi en kort instruks, der du ikke selv går gjennom den lengre planen som KI-modellen vil lage selv.

Uansett, når du havner i en slik situasjon, er svaret alltid å starte på ny med tomt kontekstvindu. Start en ny chat-sesjon, enten ved å bruke `/clear` eller ved å bruke plusstegnet.

Tenk også på hvordan oppgaven kan løses opp i flere deler. Du kan godt be agenten hjelpe deg med det:

> Del denne oppgaven opp i flere deloppgaver. Hver deloppgave skal være mulig å verifisere, slik at en er sikker på om deloppgaven er gjort.


# Subagents
Subagenter er agenter som er spunnet opp fra agenten du snakker med. Det er en effektiv måte å håndtere store oppgaver på, der subagenten for kontekst for en deloppgave, og hovedagenten orkestrerer.

Det høres avansert ut, men fra KI-modellens side er det å spinne opp en subagent er det samme som å kjøre en kommando. Altså kan vi aktivere det med vanlig norsk, så lenge agenten vår støtter subagenter.

Dersom du har en stor plan med mange deloppgaver kan du prøve noe som dette:

> implementer planen. bruk subagenter til å implementere delene, gi de tilstrekkelig informasjon om oppgaven og hvordan den skal testes. du tar ansvar som orkestrator og kvalitetsjekk etter subagenten har gjort seg ferdig.

Subagenter er også en fin måte å utforske mange ulike ideer:

> Lag git worktrees for alle alternative løsningene som du har foreslått. Bruk subagenter til å lage alle de ulike løsningene. Kjør subagentene parallelt. Når de er ferdig, lag et sammendrag som vekter fordeler og ulemper med de ulike løsningene, og presenter det i en tabell. Ta også med en kommando jeg kan kjøre for å inspisere hver av løsningene.


# Agenten stopper før den er ferdig
Ralph loop.


# Kostnad og ytelse


# Kjøre modeller lokalt


# Rammeverk
https://github.com/iyaki/ralphex


# Hva er gode tester for en agent?


# Skal vi bry oss om kodekvalitet nå?


# Diktering


# Ulike typer modeller


# Ting KI kan gjøre
Prøv disse tingene:
- lag et grafana dashboard, agenten produserer JSON basert instruks + curl av metrikk-endepunktet
- gjøre review av kode, dokumentasjon eller planer: Gjennomgå denne x og finn blindsoner og uklarheter
- oversette mellom norsk og engelsk
- lage arkitekturskisser med mermaid som du kan legge i markdown
- skrive forklarende commit meldinger, prøv dette når du har kodeendringer fra en chat du har jobbet i:
  > lag en passende git commit melding, den skal beskrive hva og hvorfor, ha både tittel og body. linjene i body skal ikke være lengre enn 80 chars. la output være en git kommando jeg kan kjøre selv
- opprette pull requests:
  > lag en pr med gh kommandolinje. tittel på pr skal inneholde saksnummeret fra jira, beskrivelsen skal inneholde hva som er gjort og hvorfor. det skal være en liste med commits i kronologisk rekkefølge, prefikset med kort sha1, slik at det kan trykkes på i github, en forklaring på hensikten/hvorfor endringen gjøres, samt en notis om hvordan det kan testes av den som skal gjøre review. la output være en gh-kommando jeg kan kjøre selv


# Ingen generert dokumentasjon


# Remote tilgang


# Tips fra folk på internett
https://www.aihero.dev/ways-ai-coding-has-rewired-my-brain
