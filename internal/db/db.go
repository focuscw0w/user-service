package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(datasource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", datasource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createSchema(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        );
    `
	_, err := db.Exec(query)
	return err
}
