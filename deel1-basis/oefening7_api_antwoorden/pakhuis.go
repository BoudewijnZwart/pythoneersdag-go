// Package oefening7 laat zien hoe weinig code er nodig is voor een klein
// HTTP-endpointje in Go: geen extern framework nodig, alleen net/http uit
// de standard library.
package oefening7

import (
	"fmt"
	"net/http"
	"strconv"
)

// NieuwPakhuis zet een kleine HTTP API op met twee endpoints:
//
//	GET /cadeaus/{naam}          - een kort welkomstbericht per cadeau
//	GET /cadeaus/{naam}/{aantal} - hoeveel er besteld zijn
//
// BUG: In een van deze handlers zit een zelfde soort fout als eerder ook
// langsgekomen is.
func NieuwPakhuis() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /cadeaus/{naam}", func(w http.ResponseWriter, r *http.Request) {
		naam := r.PathValue("naam")
		fmt.Fprintf(w, "Cadeau '%s' staat klaar in het pakhuis!\n", naam)
	})

	mux.HandleFunc("GET /cadeaus/{naam}/{aantal}", func(w http.ResponseWriter, r *http.Request) {
		naam := r.PathValue("naam")
		aantalStr := r.PathValue("aantal")

		// FIX: de error checken
		aantal, err := strconv.Atoi(aantalStr)
		if err != nil {
			http.Error(w, "ongeldig aantal", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%d keer '%s' besteld\n", aantal, naam)
	})

	return mux
}
