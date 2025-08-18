package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect opens (and migrates) the SQLite database.
func Connect() (*sql.DB, error) {
	if DB != nil {
		return DB, nil
	}

	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "coffeeshop.db"
	}

	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := database.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return nil, err
	}

	if err := migrate(database); err != nil {
		return nil, err
	}

	if err := seedAdmin(database); err != nil {
		return nil, err
	}

	DB = database
	return DB, nil
}

func migrate(db *sql.DB) error {
	// menu_items table
	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS menu_items (
		id           INTEGER PRIMARY KEY AUTOINCREMENT,
		name         TEXT NOT NULL,
		price_cents  INTEGER NOT NULL,
		description  TEXT DEFAULT '',
		active       INTEGER NOT NULL DEFAULT 1,
		created_at   DATETIME DEFAULT CURRENT_TIMESTAMP
	);`); err != nil {
		return err
	}

	// users table
	if _, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id            INTEGER PRIMARY KEY AUTOINCREMENT,
		email         TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		role          TEXT NOT NULL CHECK (role IN ('admin','customer')),
		created_at    DATETIME DEFAULT CURRENT_TIMESTAMP
	);`); err != nil {
		return err
	}

	return seedIfEmpty(db)
}
