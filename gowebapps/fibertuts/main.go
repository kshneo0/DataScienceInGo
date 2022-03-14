package main

import "github.com/gofiber/fiber/v2"

// Method 2: Fxnal Approach
func home(c *fiber.Ctx) error {
	return c.SendString("Hello Go4DataScientist")
}

func main() {
	// Init App
	app := fiber.New()

	// Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Go4DataScientist")
	})

	// Method 2:
	app.Get("/home", home)

	// Params
	// localhost:3000/api/Jesse
	app.Get("/api/:name", func(c *fiber.Ctx) error {
		fname := c.Params("name")
		return c.SendString(fname)
	})

	// Lisen
	app.Listen(":3000")
}
