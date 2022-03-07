package main

import (
	"fmt"
	"log"

	"rsc.io/pdf"
)

func main() {
	// Open the File
	file, err := pdf.Open("data/my_example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file.NumPage()) // Number of pages

	fmt.Println(file.Page(1).Content())

	// Process/Parse
}
