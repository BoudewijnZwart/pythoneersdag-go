package oefening4

import "testing"

func TestGeefUit(t *testing.T) {
	p := Pepernoot{Smaak: "Kaneel", Voorraad: 10}

	p.GeefUit()

	if p.Voorraad != 9 {
		t.Errorf("Voorraad na GeefUit() = %d, verwacht 9", p.Voorraad)
	}
}

func TestGeefUitMeerdereKeren(t *testing.T) {
	p := Pepernoot{Smaak: "Honing", Voorraad: 5}

	p.GeefUit()
	p.GeefUit()
	p.GeefUit()

	if p.Voorraad != 2 {
		t.Errorf("Voorraad na 3x GeefUit() = %d, verwacht 2", p.Voorraad)
	}
}
