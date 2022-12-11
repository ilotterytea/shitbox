package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	handlers "github.com/ilotterytea/shitbox/shitbox/handlers"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("yo")
	})

	app.Post("/upload", handlers.HandleFileUpload)

	log.Fatal(app.Listen(":8000"))
}
