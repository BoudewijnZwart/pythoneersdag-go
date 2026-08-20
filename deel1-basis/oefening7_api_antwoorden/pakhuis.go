// Package oefening7 laat zien hoe weinig code er nodig is voor een klein
// HTTP-endpointje in Go: geen extern framework nodig, alleen net/http uit
// de standard library.
package oefening7

import (
	"fmt"
	"net/http"
	"strconv"
)

// BestelCadeaus krijgt een cadeaunaam en een aantal en returnt dat als string en mogelijk een error.
// In deze handler zit een bug, probeer deze op te lossen.
func BestelCadeaus(naam string, aantalStr string) (string, error) {
	aantal, err := strconv.Atoi(aantalStr)

	if err != nil {
		return "", err
	}

	boodschap := fmt.Sprintf("%d keer '%s' besteld\n", aantal, naam)
	return boodschap, nil
}

// NieuwPakhuis zet een kleine HTTP API op met twee endpoints:
//
//	GET /cadeaus/{naam}          - een kort welkomstbericht per cadeau
//	GET /cadeaus/{naam}/{aantal} - hoeveel er besteld zijn
//
// De bug zit niet hier verscholen, los deze op in BestelCadeaus
func NieuwPakhuis() http.Handler {
	// mux kun je zien als een router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /cadeaus/{naam}", func(w http.ResponseWriter, r *http.Request) {
		naam := r.PathValue("naam")
		fmt.Fprintf(w, "Cadeau '%s' staat klaar in het pakhuis!\n", naam)
	})

	mux.HandleFunc("GET /cadeaus/{naam}/{aantal}", func(w http.ResponseWriter, r *http.Request) {
		naam := r.PathValue("naam")
		aantalStr := r.PathValue("aantal")
		boodschap, err := BestelCadeaus(naam, aantalStr)

		if err != nil {
			http.Error(w, "Ongeldig aantal ingevoerd!", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, boodschap)

	})

	return mux
}
