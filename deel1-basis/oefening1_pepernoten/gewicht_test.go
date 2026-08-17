package oefening1

import (
	"math"
	"testing"
)

func TestBerekenGemiddeldGewicht(t *testing.T) {
	gewichten := []int{100, 150, 175} // in grammen
	verwacht := 141.6666666666667

	resultaat := BerekenGemiddeldGewicht(gewichten)

	if math.Abs(resultaat-verwacht) > 0.0001 {
		t.Errorf("BerekenGemiddeldGewicht(%v) = %v, verwacht %v", gewichten, resultaat, verwacht)
	}
}

func TestBerekenGemiddeldGewichtEenZakje(t *testing.T) {
	gewichten := []int{300}
	verwacht := 300.0

	resultaat := BerekenGemiddeldGewicht(gewichten)

	if resultaat != verwacht {
		t.Errorf("BerekenGemiddeldGewicht(%v) = %v, verwacht %v", gewichten, resultaat, verwacht)
	}
}
