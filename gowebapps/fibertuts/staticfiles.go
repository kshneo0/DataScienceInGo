package main

import "github.com/gofiber/fiber/v2"

func main() {

	// Init App
	app := fiber.New()

	// Serve Static File
	// localhost:3000/myimages.png
	app.Static("/","gowebapps/fibertuts/myfiles")
	
	//Listen
	app.Listen(":3000")

}
