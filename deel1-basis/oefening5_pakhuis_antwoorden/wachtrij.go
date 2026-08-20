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

	// FIX: maak voor beide ploegen een nieuwe slice.
	//
	// Een simpele slice-expressie zoals wachtrij[:helft] maakt geen kopie.
	// De nieuwe slice verwijst dan nog steeds naar dezelfde onderliggende
	// array als wachtrij (en daarmee ook als de andere slice).
	//
	// Door make() te gebruiken maken we voor iedere ploeg een nieuwe
	// onderliggende array. Met copy() kopiëren we vervolgens de cadeaus
	// naar die nieuwe arrays.
	//
	// Hierdoor zijn ploegA en ploegB volledig onafhankelijk van elkaar:
	// een append of wijziging aan ploegA kan de data van ploegB niet meer
	// overschrijven.
	ploegA = make([]Cadeau, helft)
	ploegB = make([]Cadeau, len(wachtrij)-helft)

	copy(ploegA, wachtrij[:helft])
	copy(ploegB, wachtrij[helft:])

	return ploegA, ploegB
}

// VoegSpoedbestellingToe komt later op de dag binnen: een extra cadeau moet
// nog snel bij ploeg A's lijstje.
//
// Omdat ploegA nu een eigen onderliggende array heeft, kan append hier
// veilig gebruikt worden. Als er voldoende capaciteit is, wordt het nieuwe
// element in ploegA's eigen array geplaatst. Als er niet genoeg capaciteit
// is, maakt Go automatisch een nieuwe array.
//
// In beide gevallen blijft ploegB volledig onaangetast.
func VoegSpoedbestellingToe(ploegA []Cadeau, extra Cadeau) []Cadeau {
	return append(ploegA, extra)
}
