package shortenurl

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

type URL struct {
	ID        uint   `gorm:"primary_key"`
	Original  string `gorm:"not null"`
	Shortened string `gorm:"not null"`
}

func RedirectURL(db *redis.Client, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	short := "http://localhost:8080/" + id
	original, err := db.Get(r.Context(), short).Result()
	if err != nil {
		panic(err) //TODO
	}
	http.Redirect(w, r, original, http.StatusFound)
}
