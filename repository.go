package main

import "database/sql"

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
	return
}

func (repo *URLRepository) Get(urlUUID string) (URL, error) {
	return
}
