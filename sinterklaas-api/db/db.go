package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS cadeaus (
	id   INTEGER PRIMARY KEY AUTOINCREMENT,
	naam TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS kinderen (
	id      INTEGER PRIMARY KEY AUTOINCREMENT,
	naam    TEXT NOT NULL,
	wens_id INTEGER NOT NULL REFERENCES cadeaus(id)
);
`

// New opent (of maakt) het SQLite-databasebestand en voert de migraties uit.
func New(path string) (*sql.DB, error) {
	database, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if _, err := database.Exec(schema); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}
