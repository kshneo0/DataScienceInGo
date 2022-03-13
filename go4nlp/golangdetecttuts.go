package main

import (
	"fmt"

	"github.com/rylans/getlang"
)

func main() {
	var mystr string = "Hello world everyone"
	var mystrfr = "Bonjour a tous"
	var mystrko = "안녕하세요"

	lang := getlang.FromString(mystr)
	langfr := getlang.FromString(mystrfr)
	langko := getlang.FromString(mystrko)

	fmt.Println("Text:", mystr)
	fmt.Println(lang.LanguageCode()) // Language as Code
	fmt.Println(lang.Confidence())   // Confidence/Accuracy of Prediction

	fmt.Println("Text:", mystrfr)
	fmt.Println(langfr.LanguageCode()) // Language as Code
	fmt.Println(langfr.Confidence())   // Confidence/Accuracy of Prediction

	fmt.Println("Text:", mystrko)
	fmt.Println(langko.LanguageCode()) // Language as Code
	fmt.Println(langko.Confidence())   // Confidence/Accuracy of Prediction
}
