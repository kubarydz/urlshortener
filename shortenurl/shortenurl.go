package shortenurl

import (
	"fmt"
	"math/rand"
)

func ShortenUrl(url string) string {
	s := make([]rune, 6)
	for i := range s {
		s[i] = rune(rand.Intn(26) + 97)
	}

	shortened := fmt.Sprintf("http://localhost:8080/%s", string(s))
	return shortened
}
