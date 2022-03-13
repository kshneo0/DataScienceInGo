package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	cregex "github.com/mingrammer/commonregex"
)

func main() {
	// Textcleaning using Regex & More
	var mystr string = "Hello godev my email is kshneo@gmail.com"
	// Multi Large Text string literal ``
	// 	var docx string = `
	// Golang was designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Ken called Rob on 519-555-7765 which was redirected to +44 22 777 555.
	// Jesse sent an email to jc.@gmail.com which he found on the website http://jcharistech.com.
	// Golang was publicly announced in November 2009 and version 1.0 was released in March 2012.
	// Go is widely used in production at Google USA and in many other organizations and open-source projects.
	// In November 2016, the Go and Go Mono fonts were released by type designers Charles Bigelow and Kris Holmes specifically for use by the Go project. Go is a humanist sans-serif which resembles Lucida Grande and Go Mono is monospaced. Each of the fonts adhere to the WGL4 character set and were designed to be legible with a large x-height and distinct letterforms. Both Go and Go Mono adhere to the DIN 1450 standard by having a slashed zero, lowercase l with a tail, and an uppercase I with serifs.
	// I have been coding since 4:00 AM this morning.Accra is big but not bigger as London.
	// john.smith@yahoo.com
	// 		`
	// Reading Text From A File
	// os,ioutil,bufio
	content, err := ioutil.ReadFile("go4nlp/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Do something
	docx2 := string(content)
	fmt.Println("Sample:::", docx2)

	// Text Preprocessing
	// Normalizing : case, unicode
	// Remove noise[special char, emails, phones]
	// Lemma/Stemming
	// Tokenization

	// Normalize
	fmt.Println(strings.ToLower(mystr))
	// fmt.Println(docx)

	// Extract Emails
	// Method 1: Standard Library Regexp
	// Pattern
	p := regexp.MustCompile(`godev`)
	// Find
	fmt.Println(p.FindAllString(mystr, 1))

	//Replace
	fmt.Println(p.ReplaceAllString(mystr, "REPLACED"))

	// Method 2: External Library CommonRegex : NeatText.py
	fmt.Println(cregex.Emails(mystr))

	// Exercise 1.: Extraction
	fmt.Println(cregex.Emails(docx2))
	fmt.Println(cregex.Date(docx2))
	fmt.Println(cregex.Links(docx2))

	// Remove/Replace Document Redaction/Text Cleaning
	m := regexp.MustCompile(cregex.EmailPattern)
	fmt.Println(m.ReplaceAllString(docx2, "REPLACED"))
}
