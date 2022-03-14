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
		message := "Hello Data Scientist & Developers"
		return c.Render("index", fiber.Map{
			"coolMessage": message,
		})
	})

	// Lisen
	app.Listen(":3000")
}
