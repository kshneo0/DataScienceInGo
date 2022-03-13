package main

import (
	"fmt"

	"github.com/cdipaolo/sentiment"
)

func main() {
	mytext := "I hate eating apples and coding"

	// Model : restore or train(project directory)
	sentimentModel, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}
	results := sentimentModel.SentimentAnalysis(mytext, sentiment.English)
	fmt.Println(results)
	// Sentiment For the Entire Sentence
	fmt.Println("Sentiment Score:", results.Score)

}
