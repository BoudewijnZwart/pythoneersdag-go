package oefening7

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCadeauWelkomstbericht(t *testing.T) {
	pakhuis := NieuwPakhuis()
	req := httptest.NewRequest(http.MethodGet, "/cadeaus/lego", nil)
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, verwacht %d", rec.Code, http.StatusOK)
	}
	verwacht := "Cadeau 'lego' staat klaar in het pakhuis!\n"
	if rec.Body.String() != verwacht {
		t.Errorf("body = %q, verwacht %q", rec.Body.String(), verwacht)
	}
}

func TestBestelling(t *testing.T) {
	pakhuis := NieuwPakhuis()
	req := httptest.NewRequest(http.MethodGet, "/cadeaus/lego/3", nil)
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, verwacht %d", rec.Code, http.StatusOK)
	}
	verwacht := "3 keer 'lego' besteld\n"
	if rec.Body.String() != verwacht {
		t.Errorf("body = %q, verwacht %q", rec.Body.String(), verwacht)
	}
}

func TestOngeldigAantal(t *testing.T) {
	pakhuis := NieuwPakhuis()
	req := httptest.NewRequest(http.MethodGet, "/cadeaus/lego/veel", nil)
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, verwacht %d (Bad Request)", rec.Code, http.StatusBadRequest)
	}
}
