package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	shitbox "github.com/ilotterytea/shitbox/shitbox"
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

	if err := shitbox.DBClient.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shitbox.DBClient.Prisma.Disconnect(); err != nil {
			panic(fmt.Errorf("could not disconnect: %w", err))
		}
	}()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("yo")
	})

	app.Post("/upload", handlers.HandleFileUpload)
	app.Get("/:id", handlers.HandleFileRetrieval)
	app.Delete("/:id", handlers.HandleFileDeletion)

	log.Fatal(app.Listen(":8000"))
}
