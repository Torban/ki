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

Å finne ut hvordan agenten skal klare å verifisere sitt arbeid er hva jeg bruker mest hjernekapasitet på i mitt arbeid nå. Hvilke ting trenger den tilgang til? Hvordan kan jeg gjøre det sikkert? Hvordan kan jeg unngå at agenten overser instruks om verifikasjon?


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

> Start ved å skrive en test som sjekker at endepunkt /a inneholder teksten "ny funksjonalitet". Sørg for at testen feiler (TDD rød). Implementer funksjonaliteten. Kjør testen på ny. Sørg for at testen er OK (TDD grønn). Verifiser at alle andre tester også kjører og er grønn.


# Hva er gode tester for en agent?
Agenten er veldig flink til å skrive tester, kanskje for flink. Eksempelvis, dersom du har kode som er vanskelig å teste på grunn av sterkt koblede avhengigheter, så mocker agenten gjerne bort avhengighetene for å teste en bit av koden. Den er sykt god på mocking, men tester med mye mocking representerer ofte dårlig kode og sårbare tester.

Jeg ber ikke ageten skrive slike tester, jeg fokuserer kun på ende-til-ende tester. Det er utfallet, hensikten og formålet som er viktig. Selfølgelig ingen regler uten unntak, det gir mening å skrive tester for moduler du skal skrive om, men tenk deg om for hvilket grensesnitt du legger testene på. Kommer du deg lenger unna koden og nærmere produktet er det mye enklere å verifisere at testene er riktig.

I den grad jeg gjør review av kode, er det review av testkode jeg bruker mest tid på. Ble dette slik jeg forestilte meg? Beskriver det formålet med den nye funksjonaliteten? Vil testen fungere likevel om jeg bytter ut implementasjonen?


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
Dersom en kombinerer [subagents](#subagents) med [resultatet virker ikke](#resultatet-virker-ikke), da har du essensielt en [Ralph loop](https://ghuntley.com/loop/). Kort sagt er det å la en hovedagent styre subagenter helt til hovedagenten er fornøyd.

Dette kommer i mange farger og former, her er noen ressurser du kan sjekke ut:

- [ralphx](https://github.com/iyaki/ralphex)
- [claude goals](https://code.claude.com/docs/en/goal)


# Kostnad og ytelse
Å lage og gjennomføre dette kurset kostet meg omtrent 60 USD ved å kjøpe tokens direkte fra [openrouter.ai](https://openrouter.ai). Det er det kanskje verdt, men det er en ny kostnad på toppen av lønnskostnader, lisenser og hardware. Per nå vet vi ikke hva kostnaden kommer til å være, så en gjør riktig i å være kritisk til ukritisk KI-forbruk.

Samtidig er det viktig å ikke sette unaturlige begrensinger i hva en kan utforske. Derfor vil jeg anbefale at en bruker KI fritt en periode, før en gjør seg opp en mening om kostnaden svarer til verdien det skaper.

Rent hardware-messig kan vi også regne på det, stemmer det at tokens er på billigsalg? I følge estimater og gjetning fra internett trenger Claude Sonnet 4.6 en NVIDIA DGX H100 med 640GB minne som koster omtrent 4,5 millioner kroner. Den kan generere opp til 80 tokens i sekundet, det vil si håndere en bruker om gangen. Gitt en nedbetalingstid på 2 år og perfekt 100% utnyttelse sitter vi da med en kostnad på omtrent 0,26 kroner / time.

```
4 500 000 kroner / (2 år x 365 dager/år x 24 timer/dag) ~ 0,26 kroner / time
```

Det er mye rimeligere enn et menneske, men helt klart en forenklet model. Eksempelvis har vi ikke tatt med treningskosten, hvor stor utnyttelse en klarer å oppnå på serverparken, kostnad for strøm, leie og vedlikehold.


# Kjøre modeller lokalt
Hvis du har en ny macbook med mye RAM, er det relativt enkelt å kjøre en modell lokalt. Modellene er ikke like gode som de du kan kjøpe på abonnement, men de klarer mye som en skulle tro en måtte ha en leid modell på.

Du kan prøve å starte GPT-OSS som fungerer til mye:
```shell
brew install llama.cpp
llama-server -hf ggml-org/gpt-oss-20b-GGUF --jinja -c 0 --host 127.0.0.1 --port 8080
```

Se om den klarer å trekke ut informasjon fra JSON:

> Se innlimt JSON. Hent ut x y og z. La output være strukturert JSON på formen {"x": "...", "y": "...", "z": "..."}

Advarsel: Å gjøre seg kjent med de tekniske begrepene for KI-modeller er et kaninhull. Antall parametre, kvantifisert, mode of expert, gguf, osv. Her kan du synke mye tid. Tips, bruk gemini til å forklare hva de ulike tingene er.


# Skal vi bry oss om kodekvalitet nå?
Ja, men ikke på samme måte som før. Detaljer er mindre viktige nå, mener jeg, men oppførsel er fortsatt like viktig. Uten agenter var det ikke uvanlig at en endring på et sted ga en stor og alvorlig feil et helt annet sted. Ofte et signal på dårlig kodekvalitet. Sannsynligheten for slike feil går ikke ned med agenter.

Naivt kan en tenke seg at utfallet er:
```
sannsynlighet * antall endringer = antall feil
```

Gitt at sannsynligheten er den samme eller omtrent lik, med agenter er `antall endringer` økende, derfor må en gjøre noe med `sannsynlighet`. Hvordan gjør en det?

Gode abstraksjoner (interfaces, modularitet). Bra testing (ende til ende, beskriver hensikt). Solid rekkverk (blue green deployments, automatiserte reviews). Du skjønner tegningen.


# Diktering
Siden KI-modellene er så gode på språk og tastatur er tregere enn å snakke, prøv diktering. Det er effektivt spesielt i en planleggingsfase der du ber agenten om å lage en plan.

På MacOS er diktering innebygd, du finner det under _Systeminnstillinger_, søk etter _Diktering_:

![innstillinger for diktering på macos](diktering.png)

> dette er skrevet med diktering. Veldig nyttig når du har mange detaljer om en ting du skal lage men ikke orker å skrive ned alle detaljene som fin prosa. Ofte oppdager jeg at jeg tenker saktere enn jeg snakker og jeg må stoppe opp for å tenke hva som er neste naturlige setning. I motsetning til når jeg skriver der teksten alltid er klar for meg fordi jeg skriver så sakte.


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
