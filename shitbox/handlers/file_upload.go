package shitbox

import (
	"fmt"
	"mime"
	"time"

	"github.com/gofiber/fiber/v2"

	structures "github.com/ilotterytea/shitbox/shitbox/structures"
	utils "github.com/ilotterytea/shitbox/shitbox/utils"
)

// Handler for requests to upload a file.
func HandleFileUpload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	id, _ := utils.RandomHexToken(3)
	key, _ := utils.RandomHexToken(8)
	mime, err := mime.ExtensionsByType(file.Header.Get("Content-Type"))
	if err != nil {
		return err
	}
	if len(mime) == 0 {
		c.SendStatus(500)
		return c.JSON(fiber.Map{
			"status":  500,
			"message": "The MIME type of this file was not found by the server. The file upload was stopped.",
			"file":    nil,
		})
	}

	ifile := structures.File{
		LiteralId: id,
		MIME:      file.Header.Get("Content-Type"),
		Ext:       mime[len(mime)-1],
		Timestamp: time.Now().Unix(),
		SecretKey: key,
	}

	c.SaveFile(file, fmt.Sprintf("./uploaded/%s%s", id, mime[len(mime)-1]))

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Success!",
		"meta":    ifile,
		"get":     fmt.Sprintf("%s/%s%s", c.BaseURL(), ifile.LiteralId, ifile.Ext),
		"delete":  fmt.Sprintf("%s/%s%s?key=%s", c.BaseURL(), ifile.LiteralId, ifile.Ext, ifile.SecretKey),
	})
}
