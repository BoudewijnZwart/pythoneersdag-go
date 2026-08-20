package oefening4

import "testing"

func TestHaalGenoegVoorraad(t *testing.T) {
	v := &Voorraad{Aantal: 5}

	err := v.Haal(3)

	if err != nil {
		t.Errorf("Haal(3) gaf een onverwachte error: %v", err)
	}
	if v.Aantal != 2 {
		t.Errorf("Aantal na Haal(3) = %d, verwacht 2", v.Aantal)
	}
}

func TestHaalOnvoldoendeVoorraad(t *testing.T) {
	v := &Voorraad{Aantal: 2}

	err := v.Haal(5)

	if err != ErrOnvoldoendeVoorraad {
		t.Errorf("Haal(5) = %v, verwacht ErrOnvoldoendeVoorraad", err)
	}
	if v.Aantal != 2 {
		t.Errorf("Aantal na mislukte Haal(5) = %d, verwacht ongewijzigd 2", v.Aantal)
	}
}

func TestBezorgCadeauMetGenoegVoorraad(t *testing.T) {
	voorraad := &Voorraad{Aantal: 10}

	_, err := BezorgCadeau(voorraad, "Milan", 3)

	if err != nil {
		t.Errorf("onverwachte error: %v", err)
	}
	if voorraad.Aantal != 7 {
		t.Errorf("voorraad.Aantal = %d, verwacht 7", voorraad.Aantal)
	}
}

func TestBezorgCadeauGeeftErrorBijOnvoldoendeVoorraad(t *testing.T) {
	voorraad := &Voorraad{Aantal: 2}

	_, err := BezorgCadeau(voorraad, "Fenna", 5)

	if err == nil {
		t.Error("BezorgCadeau gaf geen error terug, terwijl er niet genoeg voorraad was")
	}
	if voorraad.Aantal != 2 {
		t.Errorf("voorraad.Aantal = %d, verwacht ongewijzigd 2 (er had niets afgehaald mogen worden)", voorraad.Aantal)
	}
}
