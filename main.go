package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
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

func shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortUrl := createURL(data.URL)
	fmt.Fprintf(w, shortUrl)
}

func main() {
	// Register Handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", shortUrlHandler)

	// Start server
	fmt.Println("Starting server on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
