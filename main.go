package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortUrl     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortUrl(OriginalUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalUrl))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println(hash[:8])
	return hash[:8]
}

func createURL(OriginalUrl string) string {
	shortUrl := generateShortUrl(OriginalUrl)
	id := shortUrl
	urlDB[id] = URL{
		ID:           id,
		OriginalUrl:  OriginalUrl,
		ShortUrl:     shortUrl,
		CreationDate: time.Now(),
	}
	return "http://localhost:8080/" + shortUrl
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")
}

func main() {
	fmt.Println("Hello, World!")
	OriginalUrl := "https://github.com/Srujangv05"
	generateShortUrl(OriginalUrl)

	// Register Handler
	http.HandleFunc("/", handler)

	// Start server
	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
