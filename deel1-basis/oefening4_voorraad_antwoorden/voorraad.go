// Package oefening4 gaat over error handling in Go. Specifiek de gewoonte
// / noodzaak om een teruggegeven error ook echt te checken.
package oefening4

import (
	"errors"
	"fmt"
)

// ErrOnvoldoendeVoorraad wordt teruggegeven als er niet genoeg pepernoten zijn.
var ErrOnvoldoendeVoorraad = errors.New("onvoldoende voorraad")

// Voorraad houdt de hoeveelheid pepernoten in het pakhuis bij.
type Voorraad struct {
	Aantal int
}

// Haal haalt n zakjes uit de voorraad. Als er niet genoeg voorraad is, wordt
// ErrOnvoldoendeVoorraad teruggegeven en verandert het Aantal niet.
func (v *Voorraad) Haal(n int) error {
	if v.Aantal < n {
		return ErrOnvoldoendeVoorraad
	}
	v.Aantal -= n
	return nil
}

// BezorgCadeau haalt pepernoten uit de voorraad om bij een cadeau te doen,
// en geeft een bevestigingsbericht terug.
//
// BUG: bij het resultaat van voorraad.Haal() wordt de eventuele error gewoon
// genegeerd. Go heeft geen exceptions, als je een teruggegeven error niet
// checkt, merkt niemand iets van een mislukte actie. Deze functie doet dus
// vrolijk alsof het cadeau klaar is, zelfs als er helemaal niet genoeg
// voorraad was.
func BezorgCadeau(voorraad *Voorraad, naam string, aantalPepernoten int) (string, error) {
	// FIX: de error wel checken en meteen teruggeven als er iets misging.
	err := voorraad.Haal(aantalPepernoten)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("cadeau voor %s is klaar!", naam), nil
}
