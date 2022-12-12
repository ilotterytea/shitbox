package shitbox

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ilotterytea/shitbox/db"
	shitbox "github.com/ilotterytea/shitbox/shitbox"
)

// File deletion request handler.
func HandleFileDeletion(c *fiber.Ctx) error {
	path := fmt.Sprintf("./uploaded/%s", c.Params("id"))

	fmt.Println(path)

	fmt.Println(c.Params("id"))
	ifile, err := shitbox.DBClient.File.FindFirst(
		db.File.ID.Equals(strings.Split(c.Params("id"), ".")[0]),
	).Exec(shitbox.DBCtx)

	if err != nil {
		return err
	}

	if c.Query("key") != ifile.Secret {
		c.SendStatus(401)
		return c.JSON(fiber.Map{
			"status":  401,
			"message": "Your secret key not match with file's secret key.",
		})
	}

	err = os.Remove(path)

	if err != nil {
		c.SendStatus(404)
		return c.JSON(fiber.Map{
			"status":  404,
			"message": fmt.Sprintf("The file with ID %s does not exist, or its lifetime has expired, or it is a private file.", c.Params("id")),
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Successfully deleted the file!",
	})
}
