package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"groteboek-api/internal/model"
	"groteboek-api/internal/repository"
)

type KindHandler struct {
	kinderen repository.KindRepository
}

func NewKindHandler(kinderen repository.KindRepository) *KindHandler {
	return &KindHandler{kinderen: kinderen}
}

// List handelt GET /kinderen af.
func (h *KindHandler) List(w http.ResponseWriter, r *http.Request) {
	kinderen, err := h.kinderen.FindAll()
	if err != nil {
		log.Println("fout bij ophalen kinderen:", err)
		writeError(w, http.StatusInternalServerError, "kon kinderen niet ophalen")
		return
	}
	writeJSON(w, http.StatusOK, kinderen)
}

// Get handelt GET /kinderen/{naam} af.
func (h *KindHandler) Get(w http.ResponseWriter, r *http.Request) {
	naam := r.PathValue("naam")

	kind, err := h.kinderen.FindByNaam(naam)
	if err != nil {
		if errors.Is(err, repository.ErrKindNietGevonden) {
			writeError(w, http.StatusNotFound, "kind niet gevonden in het grote boek")
			return
		}
		log.Println("fout bij ophalen kind:", err)
		writeError(w, http.StatusInternalServerError, "kon kind niet ophalen")
		return
	}
	writeJSON(w, http.StatusOK, kind)
}

type createKindRequest struct {
	Naam string       `json:"naam"`
	Wens model.Cadeau `json:"wens"`
}

// Create handelt POST /kinderen af
// De wens wordt hergebruikt als er al een cadeau met
// diezelfde naam bestaat.
func (h *KindHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createKindRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "ongeldige JSON")
		return
	}

	req.Naam = strings.TrimSpace(req.Naam)
	req.Wens.Naam = strings.TrimSpace(req.Wens.Naam)

	if req.Naam == "" {
		writeError(w, http.StatusBadRequest, "naam van het kind is verplicht")
		return
	}
	if req.Wens.Naam == "" {
		writeError(w, http.StatusBadRequest, "naam van het cadeau (wens) is verplicht")
		return
	}
	if req.Wens.Prijs < 0 {
		writeError(w, http.StatusBadRequest, "prijs van het cadeau mag niet negatief zijn")
		return
	}

	kind, err := h.kinderen.Create(req.Naam, req.Wens)
	if err != nil {
		if errors.Is(err, repository.ErrKindBestaatAl) {
			writeError(w, http.StatusConflict, "kind staat al in het grote boek")
			return
		}
		log.Println("fout bij aanmaken kind:", err)
		writeError(w, http.StatusInternalServerError, "kon kind niet aanmaken")
		return
	}

	writeJSON(w, http.StatusCreated, kind)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
