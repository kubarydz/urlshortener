package shortenurl

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ShortenUrl(url string, length int) string {
	s := make([]rune, length)
	for i := range s {
		s[i] = rune(rand.Intn(26) + 97)
	}

	return string(s)
}

type shortenRequest struct {
	Original string `json:"original"`
	Custom   string `json:"custom"`
}

type shortenResponse struct {
	ShortUrl string `json:"short_url"`
}

const DEFAULT_SHORT_LENGTH = 5
const DEFAULT_SHORT_RETRIES = 20

func Shorten(c *fiber.Ctx) error {
	body := shortenRequest{}
	err := c.BodyParser(&body)
	if err != nil {
		return fiber.ErrBadRequest
	}
	db := Connect()
	defer db.Close()

	var shortLength int
	shortLengthSaved, err := db.Get(c.Context(), "shorturl:length").Result()
	if err == redis.Nil {
		shortLength = DEFAULT_SHORT_LENGTH
		db.Set(c.Context(), "shorturl:length", DEFAULT_SHORT_LENGTH, 0)
	} else {
		shortLength, _ = strconv.Atoi(shortLengthSaved)
	}

	shortUrl := body.Custom
	custom := false
	if shortUrl != "" {
		result, err := db.Get(c.Context(), shortUrl).Result()
		if err == nil && result != "" {
			return fiber.NewError(fiber.ErrBadRequest.Code, "request short url already exists")
		}
		custom = true
	} else {
		shortUrl = ShortenUrl(body.Original, shortLength)
	}

	// temporary: save only for 24h
	if !custom {
		for i := 0; i < 20; i++ {
			exists, err := db.Exists(c.Context(), shortUrl).Result()
			if err != nil {
				return err
			}
			if exists == 0 {
				break
			}
			shortUrl = ShortenUrl(body.Original, shortLength)
		}
		//TODO change shorting length on collisions
	}
	db.Set(c.Context(), shortUrl, body.Original, time.Hour*24)

	return c.JSON(&shortenResponse{ShortUrl: shortUrl})
}
