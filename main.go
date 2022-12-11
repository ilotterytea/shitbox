package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	handlers "github.com/ilotterytea/shitbox/shitbox/handlers"
)

func main() {
	_, err := os.Stat("./directory")
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir("./uploaded", os.FileMode(0777))
		} else {
			log.Fatal(err)
		}
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("yo")
	})

	app.Post("/upload", handlers.HandleFileUpload)

	log.Fatal(app.Listen(":8000"))
}
