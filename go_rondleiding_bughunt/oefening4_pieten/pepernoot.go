// Package oefening4 gaat over structs en methods in Go.
package oefening4

// Pepernoot representeert een smaak pepernoten met een voorraad in het pakhuis.
type Pepernoot struct {
	Smaak    string
	Voorraad int
}

// GeefUit verlaagt de voorraad met 1 (er wordt één zakje uitgedeeld).
//
// BUG: na het aanroepen van deze method blijkt de voorraad niet aangepast te zijn.
func (p Pepernoot) GeefUit() {
	p.Voorraad--
}
