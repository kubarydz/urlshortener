package main

import (
	"fmt"
	"net/http"

	"github.com/kubarydz/urlshortener/shortenurl"
)

func main() {
	db := shortenurl.Connect()

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.ShortenUrl(original)
		fmt.Println(shortened)

		db.Set(r.Context(), shortened, original, 0)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.RedirectURL(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
