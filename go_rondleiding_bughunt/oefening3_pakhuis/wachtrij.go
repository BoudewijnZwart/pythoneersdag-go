// Package oefening3 gaat over slices in Go.
package oefening3

// Cadeau representeert een cadeau in de inpakwachtrij.
type Cadeau struct {
	Voor   string
	Inhoud string
}

// VerwijderUitWachtrij verwijdert het cadeau op de gegeven index uit de
// inpakwachtrij (bijvoorbeeld omdat het al ingepakt is) en geeft de nieuwe
// (kortere) wachtrij terug.
//
// BUG: de slice-truc hieronder doet niet wat je zou verwachten.
func VerwijderUitWachtrij(wachtrij []Cadeau, index int) []Cadeau {
	return append(wachtrij[:index], wachtrij[index:]...)
}
