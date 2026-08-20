### Welkom bij de Sinterklaas Go Workshop! 

Sinterklaas is in augustus al druk bezig met voorbereidingen voor het einde van het jaar. Sint heeft een hoop werk te doen, daarom maakt hij gebruik van de taal Go. Go helpt hem razendsnelle code te schrijven zonder al te veel complexiteit. Helaas hebben enkele freelance-pieten slordige code geschreven waardoor de codebase enkele bugs heeft. Kun je de sint helpen om de boel weer draaiende te krijgen?

### Hoe te starten?

De oefeningen zijn genummerd van **1 tot en met 7**. Het is belangrijk dat je de volgorde aanhoudt, omdat elk concept voortbouwt op het vorige. 

1. Open de map van de oefening waar je aan wilt werken (bijv. oefening1_pepernoten).
2. Open het .go bestand (bijv. gewicht.go) en lees de instructies/opdracht in de comments.
3. Fix de code of los de bug op!

###  Hoe draai je de tests?

In Go zit de test-runner ingebouwd in de compiler. Je hebt geen extern framework zoals pytest nodig. 

### 1. Alle tests van de hele workshop draaien

Wil je zien hoe ver je bent met de hele workshop? Draai dit commando vanuit de **root-map** (waar deze README staat): 

```bash
go test ./...
```


### 2. De tests van één specifieke oefening draaien

Wil je gefocust aan één opdracht werken? Navigeer naar die map en draai go test: 

```bash
cd oefening1_pepernoten
go test
```


### 3. Uitgebreide (verbose) output zien

Wil je extra informatie of fmt.Println logs zien tijdens het testen? Voeg de -v flag toe: 

```bash
go test -v
```


### Handige Links (Go by Example)

In tegenstelling tot de sint heeft Go geen groot boek nodig. Je kan [Go by Example](https://gobyexample.com/) gebruiken als spiekbrief tijdens deze opdrachten. 


****1. Pepernoten****
Basis Types & Variabelen int, float, str: [Variables](https://gobyexample.com/variables) & [Values](https://gobyexample.com/values)

****2. Kwaliteitscontrole****
Loops & Conditionals if/else, for-loops: [If/Else](https://gobyexample.com/if-else) & [For](https://gobyexample.com/for)

****3. Pieten****
Slices & Arrays list: [Arrays](https://gobyexample.com/arrays) & [Slices](https://gobyexample.com/slices)

****4. Voorraad****
Mapsdict (dictionaries)[Maps](https://gobyexample.com/maps)

****5. Pakhuis****
Structs & Methods: [Structs](https://gobyexample.com/structs) & [Methods](https://gobyexample.com/methods)

****6. Bezorging****
Pointers & Interfaces: [Pointers](https://gobyexample.com/pointers) & [Interfaces](https://gobyexample.com/interfaces)

****7. API****
Error Handling: [Errors](https://gobyexample.com/errors) & [HTTP Server](https://gobyexample.com/http-server)

