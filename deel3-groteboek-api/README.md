# Het Grote Boek van Sinterklaas — API

Een kleine Go-API die het grote boek van Sinterklaas vervangt: kinderen met hun wens (cadeau) worden opgeslagen in een sqlite-database.

Gebouwd met alleen de Go standaardbibliotheek en een pure-Go sqlite-driver ([modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)).

## Vereisten

- Go 1.22 of hoger (voor `net/http` path-patronen zoals `GET /kinderen/{naam}`)

Werkt zonder verdere setup op macOS, Windows en Linux: de sqlite-driver (`modernc.org/sqlite`) is pure Go, dus er is geen cgo en geen C-compiler nodig (geen Xcode command line tools, geen MSYS2/gcc).

> **Windows-tip:** in PowerShell is `curl` standaard een alias voor `Invoke-WebRequest`, dat de `-d`/`-i` flags uit de voorbeelden hieronder niet begrijpt. Gebruik `curl.exe` (expliciet), of voer de commando's uit in `cmd.exe`, Git Bash of WSL.

## Server starten

```bash
go run .
```

De server luistert op `:8080` en maakt bij de eerste start automatisch een `groteboek.db` bestand aan (inclusief tabellen) in de huidige map.

Voor een schone start kun je gewoon het databasebestand verwijderen.


## Endpoints

| Methode | Pad                  | Omschrijving                                  |
|---------|----------------------|-----------------------------------------------|
| GET     | `/kinderen`          | Vraag alle kinderen op                        |
| GET     | `/kinderen/{naam}`   | Eén kind opzoeken op naam                     |
| POST    | `/kinderen`          | Nieuw kind aanmaken met een wens (cadeau)     |

Namen van kinderen zijn uniek in het grote boek: een tweede kind met dezelfde naam aanmaken geeft `409 Conflict`. Cadeaus zijn uniek op naam: vraagt een volgend kind hetzelfde cadeau, dan wordt het bestaande cadeau hergebruikt in plaats van gedupliceerd.

## Voorbeelden met curl

### Kind aanmaken

```bash
curl -i -X POST localhost:8080/kinderen \
  -d '{"naam":"Pietje","wens":{"naam":"LEGO piratenset","prijs":49.99,"omschrijving":"Een leuke piratenset"}}'
```

```
HTTP/1.1 201 Created

{"id":1,"naam":"Pietje","wens":{"id":1,"naam":"LEGO piratenset","prijs":49.99,"omschrijving":"Een leuke piratenset"}}
```

### Nog een kind, met hetzelfde cadeau

Het cadeau bestaat al (uniek op naam), dus het wordt hergebruikt — prijs/omschrijving uit dit verzoek worden dan genegeerd.

```bash
curl -i -X POST localhost:8080/kinderen \
  -d '{"naam":"Klaasje","wens":{"naam":"LEGO piratenset","prijs":9999,"omschrijving":"andere waarde, wordt genegeerd"}}'
```

### Alle kinderen opvragen

```bash
curl -i localhost:8080/kinderen
```

### Eén kind opzoeken op naam

```bash
curl -i localhost:8080/kinderen/Pietje
```

Onbekende naam geeft `404 Not Found`:

```bash
curl -i localhost:8080/kinderen/Onbekend
```


## Projectstructuur

```
main.go                              startpunt van de applicatie, knoopt alles aan elkaar
internal/
  model/
    kind.go                          Kind{ID, Naam, Wens Cadeau}
    cadeau.go                        Cadeau{ID, Naam, Prijs, Omschrijving}
  database/
    database.go                      opent sqlite-db en zet het schema neer
  repository/
    kind_repository.go               KindRepository: Create, FindByNaam, FindAll
    cadeau_repository.go             CadeauRepository: FindOrCreate, FindByNaam, FindAll
  handler/
    kind_handler.go                  HTTP-handlers voor de 3 endpoints
```

`KindRepository` gebruikt intern `CadeauRepository` om de wens op te zoeken of aan te maken voordat het kind wordt weggeschreven — de handler-laag kent alleen `KindRepository`.

## Opdracht

Het is nu aan jullie om de API uit te breiden. Dit kan bijvoorbeeld door extra endpoints aan te maken, de modellen uit te breiden of een frontend toe te voegen.

Er is (nog) geen handler of endpoints voor cadeaus zelf. `CadeauRepository` heeft wel al `FindAll` en `FindByNaam` klaarstaan — een mogelijke uitbreiding is om zelf, naar het voorbeeld van `KindHandler`, een `CadeauHandler` te bouwen met bijvoorbeeld:

- `GET /cadeaus` — alle cadeaus tonen
- `GET /cadeaus/{naam}` — één cadeau opzoeken

en die te registreren in `main.go`.
