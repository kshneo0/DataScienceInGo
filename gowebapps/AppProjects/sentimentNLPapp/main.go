package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func main() {

	engine := html.New("gowebapps/AppProjects/sentimentNLPapp/views", ".html")
	// Reload
	engine.Reload(true)

	// instance of app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Route
	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello App")
		initMessage := "Sentiment Analysis App on the Go"
		return c.Render("index", fiber.Map{
			"initMessage": initMessage,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello App")
		initMessage := "Sentiment Analysis App on the Go"
		message := c.FormValue("message")
		sentimentResults := sentimentize(message)
		return c.Render("index", fiber.Map{
			"initMessage":      initMessage,
			"originalMsg":      message,
			"sentimentDetails": sentimentResults,
		})
	})

	// Listen Route
	app.Listen(":3000")
}

type SentimentDetails struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Neutral  float64 `json:"neutral"`
	Compound float64 `json:"compound"`
}

// Functions for processing
func sentimentize(docx string) SentimentDetails {
	// Parse
	parsedtext := sentitext.Parse(docx, lexicon.DefaultLexicon)
	// Process
	results := sentitext.PolarityScore(parsedtext)
	sentimentScores := SentimentDetails{Positive: results.Positive, Negative: results.Negative, Neutral: results.Neutral, Compound: results.Compound}
	return sentimentScores
}
