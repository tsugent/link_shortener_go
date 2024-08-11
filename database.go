package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// Connect to database
func connectDatabase(filePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Check that the table exists and create it if it doesnt
// TODO: move to dockerfile
func checkAndInitializeDatabase(db *sql.DB) error {
	_, err := db.Query("SELECT 1 FROM urls LIMIT 1")
	if err != nil {
		createTableQuery := `CREATE TABLE urls (original_url TEXT, uuid_string TEXT UNIQUE, created TEXT)`
		_, err := db.Exec(createTableQuery)
		if err != nil {
			return err
		}
	}
	return nil
}
