package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"sinterklaas-api/models"
	"sinterklaas-api/repository"

	"gorm.io/gorm"
)

// KindHandler bevat de HTTP-handlers voor het Kind-endpoint. Hij praat alleen
// met de repository-interfaces, niet met GORM zelf.
type KindHandler struct {
	kindRepo   repository.KindRepository
	cadeauRepo repository.CadeauRepository
}

func NewKindHandler(kindRepo repository.KindRepository, cadeauRepo repository.CadeauRepository) *KindHandler {
	return &KindHandler{kindRepo: kindRepo, cadeauRepo: cadeauRepo}
}

type createKindRequest struct {
	Naam string `json:"naam"`
	Wens string `json:"wens"`
}

// Create maakt een nieuw Kind aan. De Wens wordt opgezocht of aangemaakt aan
// de hand van de naam van het cadeau, zodat meerdere kinderen dezelfde wens
// kunnen delen.
func (h *KindHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createKindRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "ongeldige request body", http.StatusBadRequest)
		return
	}

	req.Naam = strings.TrimSpace(req.Naam)
	req.Wens = strings.TrimSpace(req.Wens)
	if req.Naam == "" || req.Wens == "" {
		http.Error(w, "naam en wens zijn verplicht", http.StatusBadRequest)
		return
	}

	cadeau, err := h.cadeauRepo.FindOrCreateByNaam(req.Wens)
	if err != nil {
		http.Error(w, "kon cadeau niet aanmaken", http.StatusInternalServerError)
		return
	}

	kind := models.Kind{Naam: req.Naam, WensID: cadeau.ID, Wens: *cadeau}
	if err := h.kindRepo.Create(&kind); err != nil {
		http.Error(w, "kon kind niet aanmaken", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, kind)
}

// GetByNaam haalt één Kind op aan de hand van zijn naam.
func (h *KindHandler) GetByNaam(w http.ResponseWriter, r *http.Request) {
	naam := r.PathValue("naam")

	kind, err := h.kindRepo.FindByNaam(naam)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "kind niet gevonden", http.StatusNotFound)
			return
		}
		http.Error(w, "er ging iets mis", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, kind)
}

// List geeft alle kinderen terug, inclusief hun wens.
func (h *KindHandler) List(w http.ResponseWriter, r *http.Request) {
	kinderen, err := h.kindRepo.FindAll()
	if err != nil {
		http.Error(w, "er ging iets mis", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, kinderen)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
