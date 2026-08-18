package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS cadeaus (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	naam TEXT NOT NULL UNIQUE,
	prijs REAL NOT NULL,
	omschrijving TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS kinderen (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	naam TEXT NOT NULL UNIQUE,
	cadeau_id INTEGER NOT NULL REFERENCES cadeaus(id)
);
`

// New opent of maakt de sqlite database en zorgt
// dat het schema aanwezig is.
func New(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
