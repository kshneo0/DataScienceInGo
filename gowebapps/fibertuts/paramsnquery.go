package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init App
	app := fiber.New()

	// Route For Params
	// localhost:3000/api/Jesse

	app.Get("/api/:name", func(c *fiber.Ctx) error {
		fname := c.Params("name")
		return c.SendString(fname)
	})

	// Foute For Query
	// localhost:3000/api/?text\"my wonderful app"

	app.Get("/api/:text?", func(c *fiber.Ctx) error {
		msg := c.Query("text")
		// return c.SendString(msg)
		newmsg := strings.ToUpper(msg)

		return c.JSON(fiber.Map{
			"original":msg,
			"modified":newmsg,
		})
	})

	// Lisen
	app.Listen(":3000")
}
