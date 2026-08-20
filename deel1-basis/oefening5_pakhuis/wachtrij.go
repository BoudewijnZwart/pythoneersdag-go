// Package oefening5 gaat over slices in Go en over de
// onderliggende array die een slice deelt met andere slices.
package oefening5

// Cadeau representeert een cadeau in de ochtendwachtrij.
type Cadeau struct {
	Voor   string
	Inhoud string
}

// VerdeelWachtrij verdeelt de ochtendwachtrij in tweeen: de eerste helft
// gaat naar ploeg A, de tweede helft naar ploeg B.
func VerdeelWachtrij(wachtrij []Cadeau) (ploegA, ploegB []Cadeau) {
	helft := len(wachtrij) / 2

	// BUG: bij het maken van ploegA en ploegB gaat er iets niet goed.
	// Het lijkt er wel op dat latere wijzigingen in ploegA ook terecht komen in
	// ploegB!
	return wachtrij[:helft], wachtrij[helft:]
}

// VoegSpoedbestellingToe komt later op de dag binnen: een extra cadeau moet
// nog snel bij ploeg A's lijstje.
//
// De code hieronder lijkt op het eerste gezicht helemaal correct.
// Toch kan het toevoegen van een cadeau ervoor zorgen dat ploegB onverwacht
// verandert.
//
// TIP: denk na over wat append eigenlijk doet.
// Heeft ploegA een eigen onderliggende array, of deelt hij die met
// een andere slice? Los de bug in de functie VerdeelWachtrij op.
func VoegSpoedbestellingToe(ploegA []Cadeau, extra Cadeau) []Cadeau {
	return append(ploegA, extra)
}
