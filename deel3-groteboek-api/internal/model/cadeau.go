package model

// Cadeau is een cadeau uit het grote boek. De naam is uniek: elk cadeau
// bestaat maar één keer, ongeacht hoeveel kinderen het wensen.
type Cadeau struct {
	ID           int64   `json:"id"`
	Naam         string  `json:"naam"`
	Prijs        float64 `json:"prijs"`
	Omschrijving string  `json:"omschrijving"`
}
