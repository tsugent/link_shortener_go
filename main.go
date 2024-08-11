package main

import (
	"fmt"
	"net/http"
)

func main() {
	db, err := connectDatabase("./links.db")
	if err != nil {
		fmt.Printf("Failed to connect to the database: %s\n", err)
		return
	}
	defer db.Close()

	if err := checkAndInitializeDatabase(db); err != nil {
		fmt.Printf("Failed to initialize the database: %s\n", err)
		return
	}
	urlRepo := NewURLRepository(db)
	http.HandleFunc("/shorten", urlRepo.HandleShorten)
	http.HandleFunc("/short/", urlRepo.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
