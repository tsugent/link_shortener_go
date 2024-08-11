package main

import (
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
)

// definte methods URLRepositoryInterface must have
type URLRepositoryInterface interface {
	Create(url string) (URL, error)
	Get(urlUUID string) (URL, error)
}

// URLRepository for database interaction
type URLRepository struct {
	db *sql.DB
}

func (repo *URLRepository) Create(url string) (URL, error) {
	var createdURL URL
	uuidString := uuid.NewString()
	redirectToPath, _, _ := strings.Cut(uuidString, "-")

	sqlQuery := `INSERT INTO urls (original_url, uuid_string, created) VALUES (?, ?, ?)`
	_, err := repo.db.Exec(sqlQuery, url, redirectToPath, time.Now().String())
	if err != nil {
		return createdURL, err
	}
	return repo.Get(redirectToPath)
}

func (repo *URLRepository) Get(urlUUID string) (URL, error) {
	var url URL
	sqlQuery := `SELECT original_url, uuid_string, created FROM urls WHERE uuid_string = ?`
	row := repo.db.QueryRow(sqlQuery, urlUUID)
	err := row.Scan(&url.OriginalURL, &url.UUIDString, &url.Created)
	if err != nil {
		return url, err
	}
	return url, nil
}
