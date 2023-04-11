package shortenurl

import (
	"github.com/gofiber/fiber/v2"
)

func RedirectURL(c *fiber.Ctx) error {
	url := c.Params("url")
	db := Connect()
	defer db.Close()

	originalUrl, err := db.Get(c.Context(), url).Result()
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.Redirect(originalUrl)
}
