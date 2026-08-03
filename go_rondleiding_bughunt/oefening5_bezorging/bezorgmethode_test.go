package oefening5

import "testing"

func TestStoombootImplementeertInterface(t *testing.T) {
	var methode Bezorgmethode = Stoomboot{}

	resultaat := methode.Bezorg("een zak pepernoten")
	verwacht := "een zak pepernoten wordt per stoomboot bezorgd"

	if resultaat != verwacht {
		t.Errorf("Bezorg(...) = %q, verwacht %q", resultaat, verwacht)
	}
}

func TestPaardImplementeertInterface(t *testing.T) {
	var methode Bezorgmethode = Paard{}

	resultaat := methode.Bezorg("een lego doos")
	verwacht := "een lego doos wordt per paard over de daken bezorgd"

	if resultaat != verwacht {
		t.Errorf("Bezorg(...) = %q, verwacht %q", resultaat, verwacht)
	}
}
