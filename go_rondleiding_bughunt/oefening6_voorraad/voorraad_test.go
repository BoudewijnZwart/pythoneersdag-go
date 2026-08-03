package oefening6

import (
	"errors"
	"testing"
)

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

	if !errors.Is(err, ErrOnvoldoendeVoorraad) {
		t.Errorf("Haal(5) = %v, verwacht ErrOnvoldoendeVoorraad", err)
	}
	if v.Aantal != 2 {
		t.Errorf("Aantal na mislukte Haal(5) = %d, verwacht ongewijzigd 2", v.Aantal)
	}
}

func TestHaalPreciesOp(t *testing.T) {
	v := &Voorraad{Aantal: 4}

	err := v.Haal(4)

	if err != nil {
		t.Errorf("Haal(4) gaf een onverwachte error: %v", err)
	}
	if v.Aantal != 0 {
		t.Errorf("Aantal na Haal(4) = %d, verwacht 0", v.Aantal)
	}
}
