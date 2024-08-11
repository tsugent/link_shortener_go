package main

import "fmt"

type URL struct {
	OriginalURL string
	Created     string
	UUIDString  string
}

func (url *URL) getShortenedLink() string {
	return fmt.Sprintf(`{"shortened_link": "http://localhost:8080/short/?r=%s"}`, url.UUIDString)
}

func (url *URL) getOriginalLink() string {
	return fmt.Sprintf(`{"original_url": "%s"}`, url.OriginalURL)
}
