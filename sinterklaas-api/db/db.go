package db

import (
	"sinterklaas-api/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// New opent (of maakt) het SQLite-databasebestand en voert de migraties uit.
func New(path string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := database.AutoMigrate(&models.Cadeau{}, &models.Kind{}); err != nil {
		return nil, err
	}

	return database, nil
}
