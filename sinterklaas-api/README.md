# Sinterklaas API

Met het toenemende aantal mensen op de wereld is het grote boek van Sinterklaas niet meer
een handige manier om bij te houden wie er cadeautjes moeten krijgen op 5 december. Daarom
hebben de Pieten het een API gemaakt met Go.

## De API starten

```bash
go run .
```

De server luistert op `http://localhost:8081`. Bij de eerste start wordt
automatisch een `sinterklaas.db`-bestand aangemaakt met de juiste tabellen.

## Endpoints

| Methode | Pad                | Omschrijving                     |
|---------|---------------------|-----------------------------------|
| GET     | `/kinderen`         | Alle kinderen ophalen             |
| GET     | `/kinderen/{naam}`  | Eén kind ophalen op naam          |
| POST    | `/kinderen`         | Nieuw kind aanmaken met een wens  |

## Uitproberen met curl

### Kind aanmaken

```bash
curl -X POST http://localhost:8081/kinderen \
  -H "Content-Type: application/json" \
  -d '{"naam": "Pietje", "wens": "Fiets"}'
```

Maak nog een kind aan met dezelfde wens, om te zien dat het cadeau hergebruikt
wordt in plaats van dubbel aangemaakt:

```bash
curl -X POST http://localhost:8081/kinderen \
  -H "Content-Type: application/json" \
  -d '{"naam": "Marieke", "wens": "Fiets"}'
```

### Eén kind opzoeken op naam

```bash
curl http://localhost:8081/kinderen/Pietje
```

### Alle kinderen ophalen

```bash
curl http://localhost:8081/kinderen
```

## Projectstructuur

```
main.go                 opzetten van db, repositories, handlers en routes
models/                 Kind en Cadeau structs
repository/             interfaces + repositories
handlers/               HTTP-handlers die met de repository-interfaces praten
db/                     opzetten van de databaseverbinding en migraties
```

## De opdracht

Probeer er achter te komen hoe de API werkt. Daarna mag je de API op een originele manier uitbreiden. Denk hierbij aan het toevoegen van nieuwe modellen, nieuwe endpoints, een frontend, etc.

De beste uitbreiding van deze API krijgt een prijs. 
