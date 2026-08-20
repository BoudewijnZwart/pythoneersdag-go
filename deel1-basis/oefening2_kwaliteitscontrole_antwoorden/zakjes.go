// Package oefening2 gaat over slices en range in Go.
package oefening2

// Zakje representeert een zakje pepernoten dat door een Piet gecontroleerd
// wordt voordat het de deur uit mag.
type Zakje struct {
	Smaak         string
	Gecontroleerd bool
}

// KeurAlleZakjesGoed loopt door alle zakjes en markeert ze als
// gecontroleerd.
func KeurAlleZakjesGoed(zakjes []Zakje) {
	// FIX: itereren over de index en via zakjes[i] het echte element in de
	// slice aanpassen, in plaats van de kopie die range teruggeeft.
	for i := range zakjes {
		zakjes[i].Gecontroleerd = true
	}
}
