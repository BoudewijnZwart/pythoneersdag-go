package oefening7

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLijstCadeaus(t *testing.T) {
	pakhuis := NieuwPakhuis()

	req := httptest.NewRequest(http.MethodGet, "/cadeaus", nil)
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /cadeaus gaf status %d, verwacht %d", rec.Code, http.StatusOK)
	}
}

func TestHaalCadeauOp(t *testing.T) {
	pakhuis := NieuwPakhuis()

	req := httptest.NewRequest(http.MethodGet, "/cadeaus/lego", nil)
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("GET /cadeaus/lego gaf status %d, verwacht %d", rec.Code, http.StatusOK)
	}

	var cadeau Cadeau
	if err := json.Unmarshal(rec.Body.Bytes(), &cadeau); err != nil {
		t.Fatalf("kon response niet parsen: %v", err)
	}
	if cadeau.Voor != "Fenna" {
		t.Errorf("cadeau.Voor = %q, verwacht %q", cadeau.Voor, "Fenna")
	}
}

func TestHaalOnbekendCadeauOp(t *testing.T) {
	pakhuis := NieuwPakhuis()

	req := httptest.NewRequest(http.MethodGet, "/cadeaus/nietbestaand", nil)
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("GET /cadeaus/nietbestaand gaf status %d, verwacht %d", rec.Code, http.StatusNotFound)
	}
}

func TestVoegCadeauToe(t *testing.T) {
	pakhuis := NieuwPakhuis()

	nieuw := Cadeau{Naam: "puzzel", Voor: "Timo"}
	body, _ := json.Marshal(nieuw)

	req := httptest.NewRequest(http.MethodPost, "/cadeaus", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("POST /cadeaus gaf status %d, verwacht %d", rec.Code, http.StatusCreated)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/cadeaus/puzzel", nil)
	getRec := httptest.NewRecorder()
	pakhuis.ServeHTTP(getRec, getReq)

	if getRec.Code != http.StatusOK {
		t.Fatalf("GET /cadeaus/puzzel na toevoegen gaf status %d, verwacht %d", getRec.Code, http.StatusOK)
	}
}

func TestVoegOngeldigCadeauToe(t *testing.T) {
	pakhuis := NieuwPakhuis()

	req := httptest.NewRequest(http.MethodPost, "/cadeaus", bytes.NewReader([]byte("dit is geen json")))
	rec := httptest.NewRecorder()
	pakhuis.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("POST /cadeaus met ongeldige JSON gaf status %d, verwacht %d", rec.Code, http.StatusBadRequest)
	}
}
