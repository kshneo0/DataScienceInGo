package main

import (
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

	// Route
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

	app.Post("/", func(c *fiber.Ctx) error {
		initMessage := "Posted New Message Here"
		fname := c.FormValue("firstname")
		message := c.FormValue("message")

		return c.Render("index", fiber.Map{
			"coolMessage": initMessage,
			"firstName":   fname,
			"newMessage":  message,
		})
	})

	// Lisen
	app.Listen(":3000")
}
