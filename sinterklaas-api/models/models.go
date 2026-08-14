package models

import "gorm.io/gorm"

// Eén Cadeau kan de Wens zijn van meerdere Kinderen.
type Cadeau struct {
	gorm.Model
	Naam string `json:"naam" gorm:"unique"`
}

// Wens verwijst naar het Cadeau dat het kind graag wil hebben.
type Kind struct {
	gorm.Model
	Naam   string `json:"naam"`
	WensID uint   `json:"wens_id"`
	Wens   Cadeau `json:"wens" gorm:"foreignKey:WensID"`
}
