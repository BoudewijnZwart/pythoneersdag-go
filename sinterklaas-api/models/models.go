package models

// Eén Cadeau kan de Wens zijn van meerdere Kinderen.
type Cadeau struct {
	ID   uint   `json:"id"`
	Naam string `json:"naam"`
}

// Wens verwijst naar het Cadeau dat het kind graag wil hebben.
type Kind struct {
	ID     uint   `json:"id"`
	Naam   string `json:"naam"`
	WensID uint   `json:"wens_id"`
	Wens   Cadeau `json:"wens"`
}
