package main

import (
	"encoding/json"
	"net/http"
)

// handler to create and return shortened url
func (repo *URLRepository) HandleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var data URL
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	shortenedURL, err := repo.Create(data.OriginalURL)
	if err != nil {
		http.Error(w, "Server Error Occurred", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(shortenedURL.getShortenedLink()); err != nil {
		http.Error(w, "Server Error Occurred", http.StatusInternalServerError)
	}
}

// Handler to redirect from url
func (repo *URLRepository) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortArg := r.URL.Query().Get("r")
	if shortArg == "" {
		http.Error(w, "Missing Shortened URL ID", http.StatusBadRequest)
		return
	}
	shortenedURL, err := repo.Get(shortArg)
	if err != nil {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(shortenedURL.getOriginalLink()); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
