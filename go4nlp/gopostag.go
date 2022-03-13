package main

import (
	"fmt"
	"strings"

	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
)

func main() {

	// mytext := "Jesse was going to fish a fish at the bank in London"
	mytext := " I am going to fish a fish at the bank"

	// Tokenize
	mytokens := tokenize.NewTreebankWordTokenizer().Tokenize(mytext)

	// Tag/Annotate
	postagger := tag.NewPerceptronTagger()
	for _, token := range postagger.Tag(mytokens) {
		fmt.Println(token.Text, token.Tag)
	}

	results := nounchunker(mytext)
	fmt.Println("Noun Chunks::", results)

	verbresults := verbchunker(mytext)
	fmt.Println("Verb Chunks::", verbresults)

}

func nounchunker(text string) []string {
	// Tokenize
	mytokens := tokenize.NewTreebankWordTokenizer().Tokenize(text)

	// Tag
	postagger := tag.NewPerceptronTagger()
	nounList := []string{}
	for _, token := range postagger.Tag(mytokens) {
		// If tag == NN
		// if token.Tag == "NN"
		if strings.HasPrefix(token.Tag, "N") {
			// Append to our list
			nounList = append(nounList, token.Text)
		}
	}
	return nounList

}

// Create VerbExtractor/Chunker

func verbchunker(text string) []string {
	// Tokenize
	mytokens := tokenize.NewTreebankWordTokenizer().Tokenize(text)

	// Tag
	postagger := tag.NewPerceptronTagger()
	verbList := []string{}
	for _, token := range postagger.Tag(mytokens) {
		// If tag == NN
		if strings.HasPrefix(token.Tag, "V") {
			// Append
			verbList = append(verbList, token.Text)
		}
	}
	return verbList

}
