---
marp: true
paginate: true
lang: nl
size: 16:9
style: |
  section {
    font-size: 26px;
  }
---

<!--
Notities voor de presentator:
Dit is een Marp-presentatie. Open het bestand met de Marp VS Code-extensie
voor een preview en presenter view, of exporteer met Marp CLI:
  npx @marp-team/marp-cli@latest go-introductie.md -o go-introductie.pptx
  npx @marp-team/marp-cli@latest go-introductie.md -o go-introductie.pdf

HTML-comments zoals dit blok staan niet op de slide zelf, maar wel in de
presenter view (toets 'p'). Elke slide hieronder heeft zo'n blok met het
bijbehorende spreekscript. Een losse, leesbare versie van hetzelfde
script staat in go-introductie-spreekscript.md.

Richttijd: 35 tot 40 minuten.
-->



<!-- _class: lead -->
<!-- _paginate: false -->

# Pythoneersdag
### GO!


---


# Deel 1: waarom bestaat Go?

<!--
We beginnen bij het begin. Waarom is deze taal er eigenlijk gekomen? Dat
verhaal verklaart een hoop van de keuzes die Go later maakt.
-->

---

## Het probleem bij Google, 2007

Google liep tegen een paar hardnekkige problemen aan:

- C++ compileerde traag bij grote codebases, en het build-systeem was zelf al een probleem
- Java werkte, maar bracht veel boilerplate en zware OOP-structuren mee
- Python was fijn om snel iets te bouwen, maar schaalde niet goed en hield weinig vast in types

<!--
Google had tienduizenden engineers die aan enorme, gedeelde
C++-codebases werkten. 1 kleine wijziging kon betekenen dat je
tientallen minuten wachtte tot de build klaar was. Java loste het
snelheidsprobleem deels op, maar bracht weer veel ceremonie mee. Python
was geweldig voor scripts en snelle tools, maar minder geschikt voor
grote, langlopende systemen met honderden engineers eraan werkend.

Geen van die talen deed alles wat Google nodig had: snel compileren,
snel draaien, onderhoudbaar op schaal, en prettig voor concurrency.
-->

---

## De makers

Drie mensen van Google, met behoorlijk wat geschiedenis:

- Robert Griesemer werkte eerder mee aan de V8 JavaScript-engine en de Java HotSpot VM
- Rob Pike hielp UTF-8 bedenken en werkte aan Unix en Plan 9
- Ken Thompson is medebedenker van Unix zelf, en van de taal B, de voorloper van C

<!--
Geen junior engineers die even een taaltje in elkaar knutselden. Ken
Thompson heeft mee aan de wieg gestaan van Unix, het besturingssysteem
waar macOS en Linux nog steeds op voortbouwen. Rob Pike was betrokken bij
UTF-8, de tekstcodering die je dagelijks gebruikt zonder het te weten.
Decennia ervaring met systeemprogrammeren, en een goed beeld van welke
pijnpunten ze wilden oplossen.
-->

---

## Tijdlijn

- 2007: intern gestart bij Google, tijdens het wachten op een lange C++-compilatie
- 2009: open source uitgebracht
- 2012: Go 1.0, de eerste stabiele versie met een compatibiliteitsbelofte
- Nu: een van de belangrijkste talen in cloud-infrastructuur

<!--
Er gaat een bekend verhaal rond dat Rob Pike het idee kreeg terwijl hij
zat te wachten op een grote C++-build. Sinds Go 1.0 in 2012 belooft het
Go-team backwards compatibility: code die toen compileerde, compileert
nu nog steeds. Voor een taal die zo snel gegroeid is, is dat best uniek.
-->

---

## Het ontwerpdoel: eenvoud

> "Go is een poging om de complexiteit terug te dringen."

- Zo'n 25 keywords, tegenover 35 in Python en 50 in Java
- Bij voorkeur 1 manier om iets te doen
- Geen classes, geen exceptions
- Bewust eenvoudig, soms zelfs bewust saai

<!--
Onthoud dit vandaag: Go is expliciet ontworpen om voorspelbaar te zijn.
Veel talen gaan prat op expressiviteit, op meerdere manieren om
hetzelfde te doen. Go kiest het tegenovergestelde: 1 duidelijke
manier, ook al is die soms verbose. Dat merk je zo bij de error handling
- dat voelt in het begin misschien primitief na Python, maar het is een
bewuste keuze.
-->

---

# Deel 2: wat is Go eigenlijk?

---

## Kenmerken

- Compiled naar native machinecode, geen VM of interpreter nodig
- Statically typed: types worden gecheckt voor het draaien
- Garbage collected, geen handmatig geheugenbeheer zoals in C
- Concurrency zit ingebouwd in de taal zelf, geen externe library nodig
- Compileert naar 1 binary, geen runtime nodig op de server
- Formatter, linter en testrunner zitten er al in

<!--
Dit verklaart al veel van waarom Go zo populair is voor
infrastructuur-tools. Het compileert naar 1 simpel bestand dat je
overal naartoe kopieert en draait, zonder dat er een Python-interpreter,
virtualenv of JVM op de doelmachine hoeft te staan. Dat is een groot
verschil met hoe je een Python-applicatie normaal deployt.
-->

---

## Wie gebruikt Go?

Grote kans dat je er vandaag al mee werkte, zonder het te weten:

- Docker en Kubernetes
- Terraform
- Delen van GitHub, Cloudflare, Uber
- CLI-tools zoals gh, kubectl, hugo, caddy

<!--
Als je ooit Docker, Kubernetes of Terraform gebruikte, werkte je al de
hele tijd met Go. Het is de facto de taal geworden voor
cloud-infrastructuur en command-line tools, precies omdat het
compileert naar snelle, makkelijk te distribueren binaries met
ingebouwde concurrency.
-->

---

## Voor wie is Go bedoeld?

Sterk in: backend services en API's, cloud-infrastructuur, command-line
tools, en alles met veel concurrency zoals netwerkverkeer of queues.

Minder voor de hand liggend: data science en ML, waar Python de standaard
blijft, snelle wegwerp-scripts, en frontend.

<!--
Belangrijk om eerlijk te zijn: Go gaat Python niet vervangen voor data
science of snelle scriptjes - daar blijft Python sterk. Maar voor de
backend-services en infrastructuur die veel van jullie bouwen, is Go een
serieuze en vaak betere fit. Dat is ook waarom we voor deze workshop een
backend-achtergrond aannemen.
-->

---

---

## Voordelen en nadelen, eerlijk gezegd

Voordelen:
- Makkelijk te lezen, weinig verborgen magie
- Snel compileren, snel draaien
- Concurrency is ingebakken en prettig om te gebruiken
- 1 binary, simpele deployment
- Sterke ingebouwde tooling, inclusief een race detector

Nadelen:
- Foutafhandeling is repetitief: `if err != nil` overal
- Geen list comprehensions, geen ternary operator - soms meer regels voor iets simpels
- Nil kan verrassen: een "typed nil" gedraagt zich net anders dan je zou verwachten
- Geen ingebouwde unions

<!--
Even eerlijk zijn, want anders wordt dit een verkooppraatje. Go lost veel
problemen goed op, maar heeft ook duidelijke nadelen. De herhaling van if
err != nil went sommige mensen nooit, en dat is een terechte klacht -
straks bij de error handling-slide zie je precies waarom dat zo is
opgezet, maar het blijft omslachtig vergeleken met try/except. Generics
zijn relatief nieuw in Go, sinds versie 1.18 in 2022, en de standaardbieb
gebruikt ze nog niet overal - je merkt af en toe dat de taal er niet mee
geboren is. En het typed-nil probleem - een interface die nil lijkt maar
het technisch niet is - is een van de meest verwarrende dingen in Go voor
nieuwkomers. Niks van dit alles is dealbreakers, maar het is goed om
vooraf te weten waar je tegenaan kan lopen.
-->

---

# Deel 3: Go versus Python

---

## Het overzicht

| | Python | Go |
|---|---|---|
| Typing | Dynamisch | Statisch |
| Uitvoering | Interpreted | Compiled |
| Concurrency | GIL / asyncio | Goroutines + channels |
| Error handling | Exceptions | Return values |
| OOP | Classes + inheritance | Structs + interfaces |
| Deployment | Runtime + dependencies | 1 binary |
| Package tool | pip/poetry/uv + venv | go mod, ingebouwd |

<!--
Het spoorboekje voor de rest van de talk. We lopen de belangrijkste
rijen hierna een voor een langs, en elk verschil komt straks terug in de
code-voorbeelden.
-->

---

## Typing: dynamisch versus statisch

```python
# Python: compileert (want dat bestaat niet), en crasht pas
# als regel 3 daadwerkelijk wordt uitgevoerd
def bereken(a, b):
    return a + b

bereken("3", 4)  # TypeError, maar pas tijdens runtime
```

```go
// Go: dit compileert al niet
func bereken(a, b int) int {
    return a + b
}

bereken("3", 4) // compile error: cannot use "3" as int
```

<!--
Waarschijnlijk het meest voelbare verschil met Python. In Python ontdek
je een type-fout pas als die regel code daadwerkelijk wordt uitgevoerd -
misschien pas in productie, in een edge case die niemand testte. In Go
ontdek je diezelfde fout al bij het compileren, voordat de code ooit
draait. Voelt in het begin als extra gedoe, maar vangt een hele
categorie bugs af voordat ze een probleem worden.
-->

---

## Snelheid en compilatie

Python interpreteert code elke keer opnieuw, of vertaalt het naar
bytecode. Go compileert naar native machinecode, 1 keer, vooraf. Het
resultaat is doorgaans een factor 10 tot 100 sneller in ruwe rekenkracht.
En de compilatie zelf is opvallend snel, ook bij grote projecten.

<!--
Nuance: dit gaat niet altijd om "Go is beter". Voor veel backend-taken
maakt ruwe snelheid weinig uit, want je wacht toch op de database of een
netwerkcall. Maar voor CPU-intensieve taken, of gewoon voor het gevoel
van een razendsnelle build-cyclus tijdens development, is het verschil
goed merkbaar.
-->

---

## Concurrency: het GIL-probleem

Python heeft de Global Interpreter Lock: maar 1 thread voert
Python-bytecode tegelijk uit. De workarounds zijn asyncio, dat
single-threaded blijft, of multiprocessing, met zware losse processen.

Go heeft goroutines: extreem lichte, door de taal zelf beheerde threads.
Duizenden tegelijk starten is geen probleem.

<!--
Een van de grootste redenen dat mensen voor backend-werk naar Go
overstappen. Python's GIL betekent dat echte parallelle CPU-uitvoering
binnen 1 proces niet kan met standaard threads. Go heeft dat probleem
niet. Dit is het onderwerp van ronde 2 vanmiddag, dus we gaan er nu niet
dieper op in - onthoud alleen dit moment.
-->

---

## Deployment: venv-narigheid versus 1 bestand

```bash
# Python
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
python app.py
```

```bash
# Go
go build -o app
scp app server:/app
ssh server ./app
```

<!--
Herkenbare pijn: wie heeft er weleens "werkt op mijn machine" gehad door
een Python-versie- of dependency-conflict? Go compileert naar 1
zelfstandig binary bestand met alle dependencies erin. Geen runtime
nodig op de server, geen versie-mismatches. Je kopieert 1 bestand en
het draait.
-->

---

# Deel 4: een snelle tour van Go syntax

<!--
Nu praktisch: genoeg syntax om zo dadelijk de oefeningen te kunnen doen.
Geen complete taalcursus - dat leer je vanmiddag al doende. Zie dit als
een naslagwerkje.
-->

---

## Hello, World

```python
# Python
print("Hallo, Sinterklaas!")
```

```go
// Go
package main

import "fmt"

func main() {
    fmt.Println("Hallo, Sinterklaas!")
}
```

<!--
Elk Go-bestand hoort bij een package - het uitvoerbare startpunt is
altijd package main met een func main(). Er is geen impliciete
top-level scriptuitvoering zoals in Python; alles zit in een functie.
fmt is de standaard-package voor eenvoudige in- en uitvoer.
-->

---

## Variabelen en types

```go
var naam string = "Amerigo"  // expliciet, met type
var leeftijd = 12             // expliciet, type wordt afgeleid
snelheid := 42                // kort, alleen binnen functies

const Max = 100               // constante
```

De `:=`-vorm gebruik je bijna altijd binnen functies. `var` gebruik je
vooral zonder startwaarde, of buiten functies. En Go is statisch
getypeerd: eenmaal een int, altijd een int.

<!--
De korte variabele-declaratie met := is waar je het merendeel van de
tijd gebruik van maakt binnen functies - het leidt zowel het type af als
declareert en vult de variabele in 1 stap. var gebruik je typisch op
package-niveau, of wanneer je een variabele wil declareren zonder hem
meteen een waarde te geven.
-->

---

## Zero values

```go
var i int        // 0
var f float64     // 0.0
var s string      // "" (lege string, geen nil)
var b bool         // false
var p *int         // nil
```

Elke variabele krijgt automatisch een zinnige startwaarde. Geen
NameError, geen undefined.

<!--
Geen directe Python-equivalent. In Python bestaat een variabele pas
zodra je hem een waarde geeft; probeer je iets niet-bestaands te
gebruiken, krijg je een NameError. In Go bestaat elke gedeclareerde
variabele altijd, met een voorspelbare startwaarde. Voorkomt een hele
categorie bugs, maar betekent ook dat je soms niet kan zien of een
waarde expres leeg is of nog nooit gezet. Kwam trouwens letterlijk terug
in opdracht 1 van vandaag, waar een berekening stilletjes het verkeerde
type gebruikte.
-->

---

## Functies met meerdere return values

```go
func haalOp(m map[string]int, sleutel string) (int, bool) {
    waarde, gevonden := m[sleutel]
    return waarde, gevonden
}

waarde, ok := haalOp(voorraad, "pepernoten")
if !ok {
    fmt.Println("niet gevonden")
}
```

Functies kunnen meerdere waarden teruggeven. Dit patroon zie je overal
in Go, vooral bij error handling zo dadelijk.

<!--
Waar je in Python vaak een tuple zou teruggeven en die zou uitpakken, is
dit in Go een taalfeature op zichzelf. Het patroon waarde-en-ok, of
waarde-en-error zoals we zo zien, is zo fundamenteel dat je het in bijna
elk Go-bestand tegenkomt.
-->

---

## Control flow

```go
if leeftijd >= 18 {
    fmt.Println("volwassen")
} else {
    fmt.Println("kind")
}

for i := 0; i < 5; i++ {
    fmt.Println(i)
}

for _, cadeau := range cadeaus {
    fmt.Println(cadeau)
}
```

Geen haakjes om de conditie, wel altijd accolades. En maar 1
loop-keyword: for doet alles, er is geen while.

<!--
De grote structuur is herkenbaar als je van Python komt. Het
belangrijkste verschil: for vervult in Go alle rollen die for + while
in andere talen spelen. Geen haakjes rond de conditie, maar de
accolades zijn wel verplicht - geen indentation-based blocks zoals in
Python.
-->

---

## Structs: geen classes

```go
type Cadeau struct {
    Voor   string
    Inhoud string
}

c := Cadeau{Voor: "Fenna", Inhoud: "Lego"}
fmt.Println(c.Voor)
```

Geen class, geen `__init__`, geen inheritance. Een struct is puur een
verzameling velden. Gedrag komt er apart bij, via methods.

<!--
Een van de grotere mentale omschakelingen komend vanuit Python's
class-gebaseerde OOP. Go heeft geen classes en dus ook geen inheritance.
Een struct is puur data. Gedrag voeg je apart toe via methods, en in
plaats van inheritance gebruikt Go composition: je stopt de ene struct in
de andere, in plaats van ervan te erven.
-->

---

## Methods: gedrag toevoegen aan een struct

```go
type Pepernoot struct {
    Smaak    string
    Voorraad int
}

func (p *Pepernoot) GeefUit() {
    p.Voorraad--
}
```

Dat `(p *Pepernoot)` heet de receiver, en er zijn er twee soorten. Een
value receiver, `(p Pepernoot)`, werkt op een kopie. Een pointer
receiver, `(p *Pepernoot)`, werkt op het origineel.

<!--
Belangrijk stuk voor vanmiddag - een van de oefeningen draait hier
letterlijk om. Met een value receiver krijgt je method een kopie van de
struct, en elke wijziging daaraan verdwijnt zodra de method klaar is.
Met een pointer receiver werkt de method op het origineel, en blijven
wijzigingen behouden. Lijkt een detail, maar het is een van de meest
voorkomende beginnersfouten in Go.
-->

---

## Slices: dynamische lijsten, met een addertje

```go
cadeaus := []string{"Lego", "Voetbal", "Puzzel"}
cadeaus = append(cadeaus, "Knuffel")

eersteTwee := cadeaus[:2]  // let op: geen kopie
```

Een slice is Go's meest gebruikte lijst-achtige structuur, en anders dan
`list[:n]` in Python is het geen kopie van de data. Het is een klein
venster op een gedeelde onderliggende array.

<!--
Waarschijnlijk het meest verrassende verschil met Python vandaag, en het
onderwerp van een van de lastigere oefeningen. In Python maakt een
list-slice een echte, onafhankelijke kopie. In Go doet slicing dat niet:
een slice is intern een klein structuurtje met een pointer, een lengte
en een capaciteit, wijzend naar een gedeelde array. Twee slices kunnen
dus dezelfde data delen, en append kan zelfs stiekem in die gedeelde
array schrijven als er nog ruimte over is. Onthoud deze slide als je
straks vastloopt.
-->

---

## Interfaces: structural typing

```go
type Bezorgmethode interface {
    Bezorg(pakket string) string
}

type Stoomboot struct{}

func (s Stoomboot) Bezorg(pakket string) string {
    return pakket + " wordt per stoomboot bezorgd"
}
```

Geen implements-keyword nodig. Zodra een type alle methods van een
interface heeft, voldoet het er automatisch aan.

<!--
Anders dan Java of TypeScript, waar je expliciet zou schrijven "class
Stoomboot implements Bezorgmethode". In Go bestaat die koppeling niet -
een type voldoet impliciet aan een interface zodra het toevallig alle
vereiste methods heeft. Er zit wel een addertje onder het gras met de
receivers van net: of een type aan een interface voldoet, hangt af van
of je pointer- of value-receivers gebruikte. Komt terug in een van de
oefeningen vanmiddag, die zelfs helemaal niet compileert totdat je het
doorhebt.
-->

---

## Error handling: geen exceptions

```python
# Python
try:
    resultaat = riskante_actie()
except ValueError as e:
    print(f"Er ging iets mis: {e}")
```

```go
// Go
resultaat, err := riskanteActie()
if err != nil {
    fmt.Println("Er ging iets mis:", err)
    return
}
```

<!--
Waarschijnlijk de grootste omschakeling van vandaag. Go heeft geen
exceptions, geen try/except. Een functie die kan falen geeft een extra
error-waarde terug als laatste return value. Is die nil, dan ging het
goed, anders niet - en jij bent verantwoordelijk om dat te checken. Doe
je dat niet, dan verdwijnt de fout stilletjes in het niets. Geen vangnet
zoals een onafgehandelde exception die crasht met een duidelijke stack
trace. Dwingt je om bij elke risicovolle aanroep na te denken over het
faalpad. Onderwerp van een van de oefeningen vanmiddag.
-->

---

## Tooling

```bash
go run main.go     # compileer en draai in 1 stap
go build             # compileer naar een binary
go test ./...         # draai alle tests
go fmt ./...            # formatteer je code
go vet ./...              # zoek veelgemaakte foutpatronen
go mod init naam            # start een nieuwe module
```

Geen Black, geen Flake8, geen losse testrunner. Zit er al in, en
iedereen gebruikt dezelfde tools.

<!--
Onderschat voordeel van Go: alle tooling die je in de Python-wereld apart
kiest en installeert - formatter, linter, testrunner - zit standaard in
de go-CLI. Omdat gofmt maar 1 manier van formatteren kent, zonder
configureerbare opties zoals bij Black, ziet Go-code er in de praktijk
verrassend consistent uit, ongeacht wie het schreef.
-->

---

<!-- _class: lead -->

# Klaar om te bug-hunten

<!--
Dat was de theorie, tijd voor de praktijk. Zo meteen gaan jullie in
tweetallen of drietallen aan de slag met een Go-repo vol opzettelijke
bugs, verdeeld in oplopende moeilijkheid.
-->

---

## Hoe werkt de bug hunt?

1. Clone de repo, elke oefening staat in zijn eigen mapje
2. Elke oefening heeft tests die nu falen
3. Los de bug op in de code, niet de tests, tot alles slaagt: `go test ./oefening1...`
4. Oefeningen zijn enigszins oplopend in moeilijkheid

```bash
git clone <repo-url> sinterklaas-workshop
cd sinterklaas-workshop
```

<!--
Clone de repo, en per oefening-map staat er een falende test. Pas de
productiecode aan, niet de tests, tot alles slaagt. De oefeningen zijn
met opzet genummerd op moeilijkheid, dus begin bij 1 en werk omhoog.
Wij lopen rond voor vragen.
-->

---

<!-- _class: lead -->

# Succes!

<!--
Veel plezier. Ik loop rond, roep gerust als je vastloopt of als iets uit
de talk nog onduidelijk is. Na ronde 1 komen we weer bij elkaar om te
bespreken wat iedereen tegenkwam.
-->
