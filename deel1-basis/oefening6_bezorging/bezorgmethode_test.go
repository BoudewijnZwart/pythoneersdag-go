package oefening6

import "testing"

func TestStoombootImplementeertInterface(t *testing.T) {
	var methode Bezorgmethode = Stoomboot{}

	resultaat := methode.Bezorg("een zak pepernoten")
	verwacht := "een zak pepernoten wordt per stoomboot bezorgd"

	if resultaat != verwacht {
		t.Errorf("Bezorg(...) = %q, verwacht %q", resultaat, verwacht)
	}
}

func TestPaardBezorgtEnTeltMee(t *testing.T) {
	methode := NieuwPaard("Amerigo")

	eerste := methode.Bezorg("een lego doos")
	tweede := methode.Bezorg("een puzzel")

	verwachtEerste := "een lego doos wordt door Amerigo over de daken bezorgd (bezorging nr. 1)"
	verwachtTweede := "een puzzel wordt door Amerigo over de daken bezorgd (bezorging nr. 2)"

	if eerste != verwachtEerste {
		t.Errorf("eerste Bezorg(...) = %q, verwacht %q", eerste, verwachtEerste)
	}
	if tweede != verwachtTweede {
		t.Errorf("tweede Bezorg(...) = %q, verwacht %q", tweede, verwachtTweede)
	}
}
