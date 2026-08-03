# 🎁 Sinterklaas' Pakhuis,Go Bug Hunt

Het is dan wel augustus, maar bij Sinterklaas begint de voorbereiding vroeg —
de eerste pepernoten worden al gebakken en de brieven stromen binnen. Het
pakhuis draait op een klein Go-systeempje voor cadeaus, voorraad en
bezorging. Er zitten alleen wat kleine bugs in. Aan jullie de taak om ze op
te sporen en te repareren, één voor één.

## Aan de slag

Zorg dat je Go geïnstalleerd hebt (`go version` in je terminal,versie 1.22
of hoger). Clone deze repo en ga naar de root van het project:

```bash
git clone <deze-repo> sinterklaas-workshop
cd sinterklaas-workshop
```

Elke oefening staat in zijn eigen mapje en heeft een `_test.go` bestand met
tests die momenteel **falen**. Jullie doel: pas de code aan (niet de tests!)
zodat alle tests slagen.

Test één oefening:

```bash
go test ./oefening1_pepernoten/...
```

Test alles in één keer:

```bash
go test ./...
```

Gebruik `-v` voor meer detail over welke (sub)tests slagen of falen:

```bash
go test -v ./oefening3_pakhuis/...
```

## De oefeningen

De oefeningen zijn verdeeld in niveaus. Je hoeft ze niet per se op volgorde
te doen, maar makkelijk → gemiddeld → moeilijk is wel de bedoeling.

### 🟢 Makkelijk,basis syntax & types

- **`oefening1_pepernoten`**,een gewichtsberekening die net niet klopt.
  Iets met types en precisie.
- **`oefening2_verlanglijstjes`**,een regel ("elke 5de brief levert een
  verrassing op") die de verkeerde brieven als verrassing aanmerkt.

### 🟡 Gemiddeld,slices & structs

- **`oefening3_pakhuis`**,een functie die een cadeau uit de inpakwachtrij
  moet verwijderen, maar er ontstaat een raar resultaat. Klassieke
  slice-valkuil.
- **`oefening4_pieten`**,een Piet die een zakje pepernoten uitdeelt, maar de
  voorraad lijkt in het niets te verdwijnen.

### 🔴 Moeilijk,interfaces & errors

- **`oefening5_bezorging`**,deze package **compileert zelfs niet**. De
  stoomboot voldoet niet aan de `Bezorgmethode` interface. Waarom niet?
  (Tip: let goed op hoofdletters in Go,dit raakt aan hoe export/zichtbaarheid
  werkt.)
- **`oefening6_voorraad`**,pepernoten die uit de pakhuisvoorraad gehaald
  worden, maar de errorcheck redeneert precies verkeerd om.

### 🎁 Backend-highlight,een echte HTTP API

- **`oefening7_api`**,een klein REST API'tje voor cadeaus, met alleen de
  standaardbibliotheek (`net/http` + `encoding/json`), geen extern framework.
  Sinds Go 1.22 kan `http.ServeMux` zelf HTTP-methodes en pad-variabelen
  matchen (`"GET /cadeaus/{naam}"`),dit laat mooi zien hoe weinig code een
  eenvoudige API in Go nodig heeft. Er zit hier één kleine typefout in een
  route. De tests draaien tegen de handlers via `net/http/httptest`, dus je
  hoeft niets zelf op te starten.

  Gezien de meesten van jullie backend-achtergrond hebben, is dit misschien
  wel de leukste om even goed naar te kijken.

## Tips

- De comments in de code (in het Nederlands) geven een hint over waar de bug
  zit, maar niet wat de fix precies is.
- Bij compile-errors: lees de foutmelding van de Go-compiler goed, die is vaak
  verrassend precies.
- Los het per twee of drie op,hardop redeneren over waarom iets niet werkt
  is de helft van het leerproces.
- Klaar met alles? Kijk eens rond in de code van de andere groepjes, of
  probeer in `oefening7_api` zelf een `DELETE /cadeaus/{naam}` endpoint toe
  te voegen.

Veel plezier, en mogen jullie voorraad nooit negatief worden. 🎅
