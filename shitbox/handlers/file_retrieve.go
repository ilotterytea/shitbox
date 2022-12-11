package shitbox

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

// File request handler.
func HandleFileRetrieval(c *fiber.Ctx) error {
	path := fmt.Sprintf("./uploaded/%s", c.Params("id"))

	fmt.Println(path)
	_, err := os.Open(path)

	if os.IsNotExist(err) {
		c.SendStatus(404)
		return c.JSON(fiber.Map{
			"status":  404,
			"message": fmt.Sprintf("The file with ID %s does not exist, or its lifetime has expired, or it is a private file.", c.Params("id")),
		})
	}

	return c.SendFile(path, false)
}
