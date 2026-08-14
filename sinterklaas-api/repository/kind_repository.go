package repository

import (
	"database/sql"

	"sinterklaas-api/models"
)

// KindRepository beschrijft de databasehandelingen voor een Kind.
type KindRepository interface {
	Create(kind *models.Kind) error
	FindByNaam(naam string) (*models.Kind, error)
	FindAll() ([]models.Kind, error)
}

type sqliteKindRepository struct {
	db *sql.DB
}

func NewKindRepository(db *sql.DB) KindRepository {
	return &sqliteKindRepository{db: db}
}

func (r *sqliteKindRepository) Create(kind *models.Kind) error {
	res, err := r.db.Exec(`INSERT INTO kinderen (naam, wens_id) VALUES (?, ?)`, kind.Naam, kind.WensID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	kind.ID = uint(id)
	return nil
}

const kindSelect = `
	SELECT k.id, k.naam, k.wens_id, c.id, c.naam
	FROM kinderen k
	JOIN cadeaus c ON c.id = k.wens_id
`

func (r *sqliteKindRepository) FindByNaam(naam string) (*models.Kind, error) {
	var kind models.Kind
	row := r.db.QueryRow(kindSelect+` WHERE k.naam = ?`, naam)
	if err := row.Scan(&kind.ID, &kind.Naam, &kind.WensID, &kind.Wens.ID, &kind.Wens.Naam); err != nil {
		return nil, err
	}
	return &kind, nil
}

func (r *sqliteKindRepository) FindAll() ([]models.Kind, error) {
	rows, err := r.db.Query(kindSelect + ` ORDER BY k.id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kinderen := []models.Kind{}
	for rows.Next() {
		var kind models.Kind
		if err := rows.Scan(&kind.ID, &kind.Naam, &kind.WensID, &kind.Wens.ID, &kind.Wens.Naam); err != nil {
			return nil, err
		}
		kinderen = append(kinderen, kind)
	}
	return kinderen, rows.Err()
}
