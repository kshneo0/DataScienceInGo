package main

import (
	"fmt"
	"log"

	"github.com/jdkato/prose/v2"
)

func main() {
	// NER
	// Entity: Person/People/Org/Location/etc

	//var mytext string = "John Mark works in London as Golang Developer"
	var mytext string = "John Mark works in London as Golang Developer"

	// NLP Document Struct/Object
	doc, err := prose.NewDocument(mytext)

	if err != nil {
		log.Fatal(err)
	}

	// Do something
	for index, entity := range doc.Entities() {
		fmt.Println(index, entity.Text, entity.Label)
	}

	for _, entity := range doc.Entities() {
		fmt.Println(entity.Text, entity.Label)
	}

	// Tokenization:Find the entities alongside
	for index, token := range doc.Tokens() {
		fmt.Println(index, token.Text, token.Label)

	}

}
