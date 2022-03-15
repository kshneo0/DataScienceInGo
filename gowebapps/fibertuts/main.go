package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// Render HTML
	// Templating Engine
	engine := html.New("gowebapps/fibertuts/views", ".html")

	// Reload For Changes :For Dev
	engine.Reload(true)

	// Init App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Route : Using Get Method
	app.Get("/", func(c *fiber.Ctx) error {
		initMessage := "Hello Data Scientist & Developers"
		// fname := c.FormValue("firstname")
		// message := c.FormValue("message")

		return c.Render("index", fiber.Map{
			"coolMessage": initMessage,
			// "firstName": fname,
			// "newMessage": message,
		})
	})

	// Route : Using Post Method
	app.Post("/", func(c *fiber.Ctx) error {
		initMessage := "Posted New Message Here"
		fname := c.FormValue("firstname")
		message := c.FormValue("message")

		// File Uploads
		file, err := c.FormFile("filename")
		if err != nil {
			return err
		}
		fileName := file.Filename
		fileSize := file.Size
		// Saving File
		c.SaveFile(file, fmt.Sprintf("gowebapps/fibertuts/myfiles/%s", file.Filename))

		return c.Render("index", fiber.Map{
			"coolMessage":   initMessage,
			"firstName":     fname,
			"newMessage":    message,
			"SavedFileName": fileName,
			"SavedFileSize": fileSize,
		})
	})

	// Lisen
	app.Listen(":3000")
}
