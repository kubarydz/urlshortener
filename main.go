package main

import (
	"fmt"
	"net/http"

	"github.com/kubarydz/urlshortener/shortenurl"
)

func main() {
	db, err := shortenurl.Connect()
	if err != nil {
		panic("failed to connect to db")
	}
	db.AutoMigrate(&shortenurl.URL{})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.ShortenUrl(original)
		fmt.Printf(shortened)

		db.Create(&shortenurl.URL{Original: original, Shortened: shortened})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
