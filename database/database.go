package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("ConnectDB: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cities (
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		name TEXT NOT NULL,
		code TEXT NOT NULL,
		country_code TEXT NOT NULL
	);`)

	return db, err
}
