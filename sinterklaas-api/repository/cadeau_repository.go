package repository

import (
	"sinterklaas-api/models"

	"gorm.io/gorm"
)

// CadeauRepository beschrijft de databasehandelingen voor een Cadeau.
type CadeauRepository interface {
	FindOrCreateByNaam(naam string) (*models.Cadeau, error)
}

type gormCadeauRepository struct {
	db *gorm.DB
}

func NewCadeauRepository(db *gorm.DB) CadeauRepository {
	return &gormCadeauRepository{db: db}
}

// FindOrCreateByNaam zorgt dat meerdere kinderen naar hetzelfde Cadeau kunnen
// verwijzen zonder dubbele rijen aan te maken voor eenzelfde wens.
func (r *gormCadeauRepository) FindOrCreateByNaam(naam string) (*models.Cadeau, error) {
	var cadeau models.Cadeau
	if err := r.db.Where(models.Cadeau{Naam: naam}).FirstOrCreate(&cadeau).Error; err != nil {
		return nil, err
	}
	return &cadeau, nil
}
