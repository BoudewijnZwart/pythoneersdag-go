// Package oefening6 gaat over error handling in Go.
package oefening6

import "errors"

// ErrOnvoldoendeVoorraad wordt teruggegeven als er niet genoeg pepernoten zijn.
var ErrOnvoldoendeVoorraad = errors.New("onvoldoende voorraad")

// Voorraad houdt de hoeveelheid pepernoten in het pakhuis bij.
type Voorraad struct {
	Aantal int
}

// Haal haalt n zakjes uit de voorraad. Als er niet genoeg voorraad is, wordt
// ErrOnvoldoendeVoorraad teruggegeven en verandert het Aantal niet.
//
// BUG: de errorcheck hieronder controleert de verkeerde conditie.
func (v *Voorraad) Haal(n int) error {
	if v.Aantal > n {
		return ErrOnvoldoendeVoorraad
	}
	v.Aantal -= n
	return nil
}
