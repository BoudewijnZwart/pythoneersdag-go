package repository

import (
	"database/sql"
	"errors"

	"groteboek-api/internal/model"
)

var ErrCadeauNietGevonden = errors.New("cadeau niet gevonden")

type CadeauRepository struct {
	db *sql.DB
}

func NewCadeauRepository(db *sql.DB) *CadeauRepository {
	return &CadeauRepository{db: db}
}

// FindOrCreate zoekt een cadeau op naam (uniek) en maakt het aan als het nog
// niet bestaat, zodat er nooit twee cadeaus met dezelfde naam ontstaan.
func (r *CadeauRepository) FindOrCreate(c model.Cadeau) (model.Cadeau, error) {
	bestaand, err := r.FindByNaam(c.Naam)
	if err == nil {
		return bestaand, nil
	}
	if !errors.Is(err, ErrCadeauNietGevonden) {
		return model.Cadeau{}, err
	}

	res, err := r.db.Exec(
		`INSERT INTO cadeaus (naam, prijs, omschrijving) VALUES (?, ?, ?)`,
		c.Naam, c.Prijs, c.Omschrijving,
	)
	if err != nil {
		return model.Cadeau{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Cadeau{}, err
	}

	c.ID = id
	return c, nil
}

func (r *CadeauRepository) FindByNaam(naam string) (model.Cadeau, error) {
	row := r.db.QueryRow(
		`SELECT id, naam, prijs, omschrijving FROM cadeaus WHERE naam = ?`, naam,
	)

	var c model.Cadeau
	if err := row.Scan(&c.ID, &c.Naam, &c.Prijs, &c.Omschrijving); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Cadeau{}, ErrCadeauNietGevonden
		}
		return model.Cadeau{}, err
	}

	return c, nil
}

func (r *CadeauRepository) FindAll() ([]model.Cadeau, error) {
	rows, err := r.db.Query(`SELECT id, naam, prijs, omschrijving FROM cadeaus ORDER BY naam`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cadeaus := []model.Cadeau{}
	for rows.Next() {
		var c model.Cadeau
		if err := rows.Scan(&c.ID, &c.Naam, &c.Prijs, &c.Omschrijving); err != nil {
			return nil, err
		}
		cadeaus = append(cadeaus, c)
	}

	return cadeaus, rows.Err()
}
