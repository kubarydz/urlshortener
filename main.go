package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kubarydz/urlshortener/shortenurl"
)

func main() {

	app := fiber.New()
	app.Post("/shorten", shortenurl.Shorten)
	app.Get("/:url", shortenurl.RedirectURL)

	log.Fatal(app.Listen(":3000"), 301)
}
