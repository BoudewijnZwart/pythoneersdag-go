// Package oefening7 laat zien hoe weinig code je nodig hebt voor een klein
// REST API'tje in Go: geen extern framework nodig, alleen net/http en
// encoding/json uit de standaardbibliotheek. Sinds Go 1.22 kan de ingebouwde
// http.ServeMux zelf al HTTP-methodes en pad-variabelen matchen.
package oefening7

import (
	"encoding/json"
	"net/http"
)

// Cadeau is een cadeau in het Sinterklaas-pakhuis.
type Cadeau struct {
	Naam string `json:"naam"`
	Voor string `json:"voor"`
}

// NieuwPakhuis zet een kleine HTTP API op met drie endpoints:
//
//	GET  /cadeaus         - lijst alle cadeaus
//	GET  /cadeaus/{naam}  - haal één cadeau op
//	POST /cadeaus         - voeg een nieuw cadeau toe
//
// BUG: één van de routes is verkeerd geregistreerd, waardoor die nooit
// matcht.
func NieuwPakhuis() http.Handler {
	cadeaus := map[string]Cadeau{
		"lego":    {Naam: "lego", Voor: "Fenna"},
		"voetbal": {Naam: "voetbal", Voor: "Milan"},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /cadeaus", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cadeaus)
	})

	mux.HandleFunc("GET /cadeau/{naam}", func(w http.ResponseWriter, r *http.Request) {
		naam := r.PathValue("naam")
		cadeau, gevonden := cadeaus[naam]
		if !gevonden {
			http.Error(w, "cadeau niet gevonden", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cadeau)
	})

	mux.HandleFunc("POST /cadeaus", func(w http.ResponseWriter, r *http.Request) {
		var nieuw Cadeau
		if err := json.NewDecoder(r.Body).Decode(&nieuw); err != nil {
			http.Error(w, "ongeldige JSON", http.StatusBadRequest)
			return
		}
		cadeaus[nieuw.Naam] = nieuw
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(nieuw)
	})

	return mux
}
