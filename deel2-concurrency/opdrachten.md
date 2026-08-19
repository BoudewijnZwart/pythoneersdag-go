# Concurrency & parallelism opdrachten

De programmeertaal Go is vanaf het begin aan ontworpen met concurrency en parallellisme in gedachten. Deze opdrachten leren je de basis.

## Opdracht 1 - Normale synchrone code

Om het voordeel van concurrency te begrijpen beginnen we met een klein programma dat synchroon draait. Open de bestanden 01_sync/main.go, shared/utils/colors.go en shared/utils/toys.go en probeer te begrijpen wat deze code doet. Je kunt de code uitvoeren vanuit de map waarin dit markdown-bestand zich bevindt door het volgende uit te voeren:

```bash
go run ./01_sync/
```

De code draait behoorlijk traag. Zonder nog iets aan te passen, waar kan de grootste verbetering behaald worden?

## Opdracht 2 - Concurrency zonder parallellisme

Open het bestand 02/concurrent/main.go. Voor deze opdracht gebruiken we nog geen parallellisme. De regel

```go
runtime.GOMAXPROCS(1)
```

zet het aantal OS-threads dat werk mag uitvoeren op één, waardoor parallellisme voorlopig wordt uitgeschakeld. Probeer het aanmaken van toys concurrent te laten verlopen. Gebruik goroutines en wait groups. Pas alleen het bestand 02_concurrent_antwoorden/main.go aan.


## Opdracht 3 - Parallellisme met Mutex

Als een Go-programma over meerdere cores beschikt, worden de goroutines niet alleen concurrent maar ook parallel uitgevoerd, indien mogelijk. Je hoeft hier zelf niets extra's voor te doen, de Go-runtime regelt dit voor je.
Code parallel laten draaien brengt wel extra verantwoordelijkheid met zich mee voor de ontwikkelaar. Je moet ervoor zorgen dat goroutines niet tegelijkertijd hetzelfde geheugen proberen te gebruiken. Dit kan data races veroorzaken.

Go heeft een tool om data races te detecteren. Met het volgende commando kun je de mogelijke data race zien die we in opdracht 3 hebben gemaakt:

```bash
    go run -race /03_parallel_mutex
```

Waar vindt de data race plaats?

Probeer dit probleem op te lossen met een Mutex. Zie https://go.dev/tour/concurrency/9 voor meer informatie over Mutexen in Go. Pas alleen het bestand 03_parallel_mutex/main.go aan.

## Opdracht 4 - Parallellisme met Channels

Een gebruikelijkere manier om problemen met geheugentoegang in Go op te lossen is het gebruik van channels. Channels transporteren data tussen goroutines. De goroutine moet aangeven of hij van een channel leest of ernaar schrijft. Channels zijn zo ontworpen dat, zelfs als twee goroutines op precies hetzelfde moment proberen te schrijven naar of te lezen van een channel, er toch een volgorde wordt afgedwongen.

Zie https://go.dev/tour/concurrency/2 voor meer informatie over channels.

Probeer hetzelfde probleem als in opdracht 3 op te lossen, maar nu met channels. Pas alleen het bestand 04_parallel_channels/main.go aan.

**Hint**

Een manier om dit op te lossen is door het "fan in" patroon te gebruiken. Dit kan door de goroutines die het speelgoed maken het speelgoed in de channel te stoppen en een apparte goroutine de channel te laten leeghalen.

