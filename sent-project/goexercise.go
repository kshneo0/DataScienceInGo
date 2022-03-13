package main // Executable

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

// Create Slice/Struct to store values
type saveDetails struct {
	Sentence   string
	Label      string
	Vaderlabel float64
}

// entry point for execution
func main() {

	// Sentiment with Vader
	// Neg : words
	// Pos : word
	// Neutral :word
	// Compound: final sentiment

	// mytext := "I really hate apples badly"
	// // Parse Text & Compare with Dictionary/Lexicon
	// parsedtext := sentitext.Parse(mytext, lexicon.DefaultLexicon)
	// // Analysis
	// results := sentitext.PolarityScore(parsedtext)

	// fmt.Println(results)
	// fmt.Println("Pos:", results.Positive)
	// fmt.Println("Neg:", results.Negative)
	// fmt.Println("Neutral:", results.Neutral)
	// fmt.Println("Compound/Sentiment:", results.Compound)

	// Open File
	csvfile, err := os.Open("sent-project/data/amazondataset.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer csvfile.Close()

	// Read our CSV File with Gota/Pandas For Dataframe
	// df := dataframe.ReadCSV(csvfile)
	// fmt.Println(df)

	// // Column Names
	// fmt.Println(df.Names())
	// fmt.Println(df.Select("sentences").String())

	// Method For Reading CSV
	csvlines, err := csv.NewReader(csvfile).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Apply our Fxn
	var detailsList []saveDetails

	for _, line := range csvlines {
		mysent := line[0]
		mylabel := line[1]
		fmt.Println(mysent)
		myresult := analyze(mysent)
		// fmt.Println("Orig:", mylabel, "NewLabel:", myresult)
		detailsList = append(detailsList, saveDetails{
			Sentence: mysent, Label: mylabel, Vaderlabel: myresult,
		})
	}
	// fmt.Println(detailsList)

	// Results as A DataFrame
	// Create Slice/Struct to store values

	// Struct to Dataframe
	df := dataframe.LoadStructs(detailsList)
	fmt.Println(df)

	// Save using Gota
	f, err := os.Create("sent-project/data/newdata.csv")
	if err != nil {
		log.Fatal(err)
	}
	df.WriteCSV(f)

}

func analyze(text string) float64 {
	// Parse Text & Compare with Dictionary/Lexicon
	mytext := text
	parsedtext := sentitext.Parse(mytext, lexicon.DefaultLexicon)
	// Analysis
	results := sentitext.PolarityScore(parsedtext)
	return results.Compound

}
