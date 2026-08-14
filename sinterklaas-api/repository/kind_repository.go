package repository

import (
	"sinterklaas-api/models"

	"gorm.io/gorm"
)

// KindRepository beschrijft de databasehandelingen voor een Kind.
type KindRepository interface {
	Create(kind *models.Kind) error
	FindByNaam(naam string) (*models.Kind, error)
	FindAll() ([]models.Kind, error)
}

type gormKindRepository struct {
	db *gorm.DB
}

func NewKindRepository(db *gorm.DB) KindRepository {
	return &gormKindRepository{db: db}
}

func (r *gormKindRepository) Create(kind *models.Kind) error {
	return r.db.Create(kind).Error
}

func (r *gormKindRepository) FindByNaam(naam string) (*models.Kind, error) {
	var kind models.Kind
	if err := r.db.Preload("Wens").Where("naam = ?", naam).First(&kind).Error; err != nil {
		return nil, err
	}
	return &kind, nil
}

func (r *gormKindRepository) FindAll() ([]models.Kind, error) {
	var kinderen []models.Kind
	err := r.db.Preload("Wens").Find(&kinderen).Error
	return kinderen, err
}
