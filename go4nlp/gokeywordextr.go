package main

import (
	"fmt"
	"io/ioutil"
	"log"

	rake "github.com/Obaied/RAKE.Go"
)

func main() {

	// var mytext string = `Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data. The result is a computer capable of "understanding" the contents of documents, including the contextual nuances of the language within them. The technology can then accurately extract information and insights contained in the documents as well as categorize and organize the documents themselves.

	// Challenges in natural language processing frequently involve speech recognition, natural language understanding, and natural-language generation.
	// Natural language processing (NLP) is a subfield of linguistics, computer science, and artificial intelligence concerned with the interactions between computers and human language, in particular how to program computers to process and analyze large amounts of natural language data. The result is a computer capable of "understanding" the contents of documents, including the contextual nuances of the language within them. The technology can then accurately extract information and insights contained in the documents as well as categorize and organize the documents themselves.

	// Challenges in natural language processing frequently involve speech recognition, natural language understanding, and natural-language generation.`

	// Opening and Reading From A File
	content, err := ioutil.ReadFile("go4nlp/exampletext.txt")
	if err != nil {
		log.Fatal(err)
	}

	candidates := rake.RunRake(string(content))

	// Store in a Map [K:V]/Dictionary
	keywordMap := make(map[string]float64)
	for _, word := range candidates {
		// fmt.Println("Keyword: ", word.Key, "Score", word.Value)
		keywordMap[word.Key] = word.Value
	}

	fmt.Println(keywordMap)

}
