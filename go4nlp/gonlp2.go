package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jdkato/prose/v2"
)

func main() {

	var mystr string = "Hello world this is Golang"

	fmt.Println(mystr)

	// NLP Document Struct
	// docx, _ := prose.NewDocument(mystr)
	// fmt.Printf("%T \n", docx)

	// Showing Errors
	docx, err := prose.NewDocument((mystr))
	if err != nil {
		log.Fatal(err)
	}

	// Do Something
	// Tokenization
	for i, tok := range docx.Tokens() {
		fmt.Println("Index:", i, "Tokens:", tok.Text)
	}

	// Reading From A Textfile
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Do Something
	docxfile, err := prose.NewDocument(string(content))
	if err != nil {
		log.Fatal(err)
	}

	// Word Tokens
	// for _, tok := range docxfile.Tokens() {
	// 	fmt.Println(tok.Text)
	// }
	// Sentence Tokenization
	// for i, sent := range docxfile.Sentences() {
	// 	fmt.Println("index:", i, ":", sent.Text)
	// }
	// Pos Tagging
	for _, tok := range docxfile.Tokens() {
		fmt.Println(tok.Text, tok.Tag)
	}

}
