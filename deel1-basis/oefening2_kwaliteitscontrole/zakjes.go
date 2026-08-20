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
//
// BUG: het lijkt erop dat de staat van Gecontroleerd op de zakjes niet
// echt worden gewijzigd. Kun je het probleem oplossen?
func KeurAlleZakjesGoed(zakjes []Zakje) {
	for _, zakje := range zakjes {
		zakje.Gecontroleerd = true
	}
}
