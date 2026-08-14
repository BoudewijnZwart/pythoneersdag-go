package repository

import (
	"database/sql"

	"sinterklaas-api/models"
)

// CadeauRepository beschrijft de databasehandelingen voor een Cadeau.
type CadeauRepository interface {
	FindOrCreateByNaam(naam string) (*models.Cadeau, error)
}

type sqliteCadeauRepository struct {
	db *sql.DB
}

func NewCadeauRepository(db *sql.DB) CadeauRepository {
	return &sqliteCadeauRepository{db: db}
}

// FindOrCreateByNaam zorgt dat meerdere kinderen naar hetzelfde Cadeau kunnen
// verwijzen zonder dubbele rijen aan te maken voor eenzelfde wens.
func (r *sqliteCadeauRepository) FindOrCreateByNaam(naam string) (*models.Cadeau, error) {
	if _, err := r.db.Exec(`INSERT INTO cadeaus (naam) VALUES (?) ON CONFLICT(naam) DO NOTHING`, naam); err != nil {
		return nil, err
	}

	var cadeau models.Cadeau
	err := r.db.QueryRow(`SELECT id, naam FROM cadeaus WHERE naam = ?`, naam).Scan(&cadeau.ID, &cadeau.Naam)
	if err != nil {
		return nil, err
	}
	return &cadeau, nil
}
