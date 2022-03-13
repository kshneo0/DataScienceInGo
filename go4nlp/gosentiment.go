package main

import (
	"fmt"

	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func main() {

	mytext := "I hate onions so bad"
	parsetext := sentitext.Parse(mytext, lexicon.DefaultLexicon)
	results := sentitext.PolarityScore(parsetext)
	fmt.Println(results)
	fmt.Println("Positive:", results.Positive)
	fmt.Println("Negative:", results.Negative)
	fmt.Println("Neutral:", results.Neutral)
	fmt.Println("Sentiment/Compound:", results.Compound)
}
