package oefening5

import "testing"

func TestVoegSpoedbestellingToeLaatPloegBOngemoeid(t *testing.T) {
	wachtrij := []Cadeau{
		{Voor: "Fenna", Inhoud: "Lego"},
		{Voor: "Milan", Inhoud: "Voetbal"},
		{Voor: "Timo", Inhoud: "Puzzel"},
		{Voor: "Sanne", Inhoud: "Knuffel"},
	}

	ploegA, ploegB := VerdeelWachtrij(wachtrij)

	spoedbestelling := Cadeau{Voor: "Bram", Inhoud: "Boek"}
	ploegA = VoegSpoedbestellingToe(ploegA, spoedbestelling)

	if len(ploegA) != 3 {
		t.Fatalf("len(ploegA) = %d, verwacht 3", len(ploegA))
	}
	if ploegA[2] != spoedbestelling {
		t.Errorf("ploegA[2] = %v, verwacht de spoedbestelling %v", ploegA[2], spoedbestelling)
	}

	verwachtEersteVanB := Cadeau{Voor: "Timo", Inhoud: "Puzzel"}
	if ploegB[0] != verwachtEersteVanB {
		t.Errorf(
			"ploegB[0] = %v, verwacht %v - de spoedbestelling voor ploeg A heeft blijkbaar ploeg B's cadeau overschreven",
			ploegB[0], verwachtEersteVanB,
		)
	}
}
