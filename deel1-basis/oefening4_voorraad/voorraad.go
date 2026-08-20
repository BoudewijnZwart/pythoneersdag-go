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
// BUG: volgens de output lijkt het altijd goed te gaan, ook al de voorraad niet
// minder wordt.
// los de bug hier onder op
func BezorgCadeau(voorraad *Voorraad, naam string, aantalPepernoten int) (string, error) {
	voorraad.Haal(aantalPepernoten)
	return fmt.Sprintf("cadeau voor %s is klaar!", naam), nil
}
