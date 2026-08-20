// Package oefening1 gaat over basis types en variabelen in Go.
package oefening1

// BerekenGemiddeldGewicht berekent het gemiddelde gewicht (in grammen) van een
// lijst zakjes pepernoten. Sint is in augustus al druk met proefzakjes wegen,
// lang voor de intocht.
//
// BUG: er gaat hier iets mis met de precisie van de berekening.
func BerekenGemiddeldGewicht(gewichten []int) float64 {
	totaal := 0
	for _, g := range gewichten {
		totaal += g
	}

	gemiddelde := float64(totaal) / float64(len(gewichten)) // FIX: expliciete conversie naar float64 voor de deling

	return float64(gemiddelde)
}
