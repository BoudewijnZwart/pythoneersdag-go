package repository

import (
	"database/sql"
	"errors"
	"strings"

	"groteboek-api/internal/model"
)

var (
	ErrKindNietGevonden = errors.New("kind niet gevonden")
	ErrKindBestaatAl    = errors.New("kind staat al in het grote boek")
)

// KindRepository beschrijft de opslag en opvraging van kinderen in het grote boek.
type KindRepository interface {
	Create(naam string, wens model.Cadeau) (model.Kind, error)
	FindByNaam(naam string) (model.Kind, error)
	FindAll() ([]model.Kind, error)
}

type SqliteKindRepository struct {
	db      *sql.DB
	cadeaus CadeauRepository
}

func NewKindRepository(db *sql.DB, cadeaus CadeauRepository) *SqliteKindRepository {
	return &SqliteKindRepository{db: db, cadeaus: cadeaus}
}

// Create zoekt de gewenste wens op (of maakt hem aan) en schrijft het kind
// met die wens in het grote boek.
func (r *SqliteKindRepository) Create(naam string, wens model.Cadeau) (model.Kind, error) {
	cadeau, err := r.cadeaus.FindOrCreate(wens)
	if err != nil {
		return model.Kind{}, err
	}

	res, err := r.db.Exec(`INSERT INTO kinderen (naam, cadeau_id) VALUES (?, ?)`, naam, cadeau.ID)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return model.Kind{}, ErrKindBestaatAl
		}
		return model.Kind{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return model.Kind{}, err
	}

	return model.Kind{ID: id, Naam: naam, Wens: cadeau}, nil
}

func (r *SqliteKindRepository) FindByNaam(naam string) (model.Kind, error) {
	row := r.db.QueryRow(`
		SELECT k.id, k.naam, c.id, c.naam, c.prijs, c.omschrijving
		FROM kinderen k
		JOIN cadeaus c ON c.id = k.cadeau_id
		WHERE k.naam = ?`, naam)

	var k model.Kind
	if err := row.Scan(&k.ID, &k.Naam, &k.Wens.ID, &k.Wens.Naam, &k.Wens.Prijs, &k.Wens.Omschrijving); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Kind{}, ErrKindNietGevonden
		}
		return model.Kind{}, err
	}

	return k, nil
}

func (r *SqliteKindRepository) FindAll() ([]model.Kind, error) {
	rows, err := r.db.Query(`
		SELECT k.id, k.naam, c.id, c.naam, c.prijs, c.omschrijving
		FROM kinderen k
		JOIN cadeaus c ON c.id = k.cadeau_id
		ORDER BY k.naam`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kinderen := []model.Kind{}
	for rows.Next() {
		var k model.Kind
		if err := rows.Scan(&k.ID, &k.Naam, &k.Wens.ID, &k.Wens.Naam, &k.Wens.Prijs, &k.Wens.Omschrijving); err != nil {
			return nil, err
		}
		kinderen = append(kinderen, k)
	}

	return kinderen, rows.Err()
}
