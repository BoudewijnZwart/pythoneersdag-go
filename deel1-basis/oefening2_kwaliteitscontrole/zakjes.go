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
// BUG: range geeft bij elke iteratie een kopie van het element terug, niet
// het originele element in de slice. De aanpassing aan zakje in de loop
// verdwijnt dus zodra de volgende iteratie begint, de originele zakjes in
// de slice blijven ongewijzigd.
func KeurAlleZakjesGoed(zakjes []Zakje) {
	for _, zakje := range zakjes {
		zakje.Gecontroleerd = true
	}
}
