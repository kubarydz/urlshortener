package shortenurl

import (
	"net/http"

	"gorm.io/gorm"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null"`
	Shortened string `gorm:"not null"`
}

func RedirectURL(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	var url URL
	short := "http://localhost:8080/" + id
	db.First(&url, "shortened = ?", short)
	http.Redirect(w, r, url.Original, http.StatusFound)
}
