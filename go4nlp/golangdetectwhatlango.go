package main

import (
	"fmt"

	"github.com/abadojack/whatlanggo"
)

func main() {
	var mydocx string = "Hello world This is John from London"
	lang := whatlanggo.Detect(mydocx)
	fmt.Println(lang.Lang.String())
	fmt.Println(lang.Confidence)

	mydocx2 := `привет друзья как дела?
	Учиться новому - весело`
	lang2 := whatlanggo.Detect((mydocx2))
	fmt.Println(lang2.Lang.String())
	fmt.Println(lang2.Confidence)
	fmt.Println(whatlanggo.Scripts[lang2.Script])

	mydocx3 := `안녕하세요?
	누구세요?`
	lang3 := whatlanggo.Detect((mydocx3))
	fmt.Println(lang3.Lang.String())
	fmt.Println(lang3.Confidence)
	fmt.Println(whatlanggo.Scripts[lang3.Script])

}
