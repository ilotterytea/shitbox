package shitbox

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Handler for requests to upload a file.
func HandleFileUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	return c.SaveFile(file, fmt.Sprintf("./uploaded/%s", file.Filename))
}
