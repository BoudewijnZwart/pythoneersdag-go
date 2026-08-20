package model

// Kind staat in het grote boek van Sinterklaas met de wens die het heeft.
type Kind struct {
	ID   int64  `json:"id"`
	Naam string `json:"naam"`
	Wens Cadeau `json:"wens"`
}
