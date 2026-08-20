// Package oefening6 gaat over interfaces in Go. Specifiek hoe method sets
// werken met pointer-receivers vs value-receivers.
package oefening6

import "fmt"

// Bezorgmethode is de interface die elke manier van bezorgen moet implementeren.
type Bezorgmethode interface {
	Bezorg(pakket string) string
}

// Stoomboot is de traditionele manier waarop Sinterklaas zelf aankomt. Een
// stoomboot heeft geen eigen state nodig, dus een value receiver volstaat.
type Stoomboot struct{}

func (s Stoomboot) Bezorg(pakket string) string {
	return fmt.Sprintf("%s wordt per stoomboot bezorgd", pakket)
}

// Paard is hoe Amerigo de pakjes langs de daken brengt. Een paard onthoudt
// hoeveel pakketten het al bezorgd heeft, dus Bezorg gebruikt een pointer
// receiver om die teller daadwerkelijk te kunnen aanpassen.
type Paard struct {
	Naam          string
	AantalBezorgd int
}

func (p *Paard) Bezorg(pakket string) string {
	p.AantalBezorgd++
	return fmt.Sprintf("%s wordt door %s over de daken bezorgd (bezorging nr. %d)", pakket, p.Naam, p.AantalBezorgd)
}

// NieuwPaard maakt een nieuwe Bezorgmethode aan in de vorm van een opgegeven paard.
//
// BUG: dit compileert niet. Paard{...} is een waarde, geen pointer, en
// Bezorg heeft een pointer receiver - een waarde van het type Paard zit
// daarom niet in de method set van de Bezorgmethode interface.
func NieuwPaard(naam string) Bezorgmethode {
	return &Paard{Naam: naam} // FIX: een pointer teruggeven, die zit wel in de method set
}
