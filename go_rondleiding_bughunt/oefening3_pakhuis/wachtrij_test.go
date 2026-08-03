package oefening3

import (
	"reflect"
	"testing"
)

func TestVerwijderUitWachtrij(t *testing.T) {
	wachtrij := []Cadeau{
		{Voor: "Fenna", Inhoud: "Lego"},
		{Voor: "Milan", Inhoud: "Voetbal"},
		{Voor: "Timo", Inhoud: "Puzzel"},
	}

	resultaat := VerwijderUitWachtrij(wachtrij, 1)

	verwacht := []Cadeau{
		{Voor: "Fenna", Inhoud: "Lego"},
		{Voor: "Timo", Inhoud: "Puzzel"},
	}

	if !reflect.DeepEqual(resultaat, verwacht) {
		t.Errorf("VerwijderUitWachtrij(wachtrij, 1) = %v, verwacht %v", resultaat, verwacht)
	}

	if len(resultaat) != 2 {
		t.Errorf("lengte van resultaat = %d, verwacht 2", len(resultaat))
	}
}

func TestVerwijderUitWachtrijLaatsteItem(t *testing.T) {
	wachtrij := []Cadeau{
		{Voor: "Fenna", Inhoud: "Lego"},
		{Voor: "Milan", Inhoud: "Voetbal"},
	}

	resultaat := VerwijderUitWachtrij(wachtrij, 1)

	verwacht := []Cadeau{
		{Voor: "Fenna", Inhoud: "Lego"},
	}

	if !reflect.DeepEqual(resultaat, verwacht) {
		t.Errorf("VerwijderUitWachtrij(wachtrij, 1) = %v, verwacht %v", resultaat, verwacht)
	}
}
