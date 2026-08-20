// Package oefening3 gaat over structs en methods in Go.
package oefening3

// Pepernoot representeert een smaak pepernoten met een voorraad in het pakhuis.
type Pepernoot struct {
	Smaak    string
	Voorraad int
}

// GeefUit verlaagt de voorraad met 1 (er wordt 1 zakje uitgedeeld).
//
// BUG: na het aanroepen van deze method blijkt de voorraad niet aangepast te zijn.
// FIX: pointer receiver zodat de wijziging blijft bestaan
func (p *Pepernoot) GeefUit() {
	p.Voorraad--
}
