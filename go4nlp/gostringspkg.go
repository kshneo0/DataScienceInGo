package main

import (
	"fmt"
	"strings"
)

func main() {
	// Working with strings & strings Packages
	// A String: collection[array/slice] of bytes or runes which characters
	var mystr string = "hello Go"
	fmt.Println(mystr)

	// Check fot Type
	fmt.Printf("%T \n", mystr)

	// String Manipulation
	// Length of string
	fmt.Println("Length: ", len(mystr))

	// To Uppercase
	fmt.Println("Uppercase:", strings.ToUpper(mystr))

	// To Lowercase
	fmt.Println("Lowercase:", strings.ToLower(mystr))

	// To Title Case
	fmt.Println("Titlecase:", strings.Title(mystr))

	//To Get the Count of Subset within our string
	fmt.Println("Count:", strings.Count(mystr, "l"))

	// Check if it contains something
	fmt.Println("Contains:", strings.Contains(mystr, "Go"))

	// Useful For NLP
	// Split: Tikenization
	tok := strings.Split(mystr, " ")
	fmt.Println(tok)
	fmt.Println(tok[0])

	// Split After
	// https://forteleaf.tistory.com/entry/golang-strings-패키지
	// func SplitAfter (s, sep string, n int) [] string
	// SplitAfter는 s에서 sep가 나타난 부분의 바로 슬라이스하고,
	// 이러한 부분 문자열 조각을 반환합니다. sep가 비어있을 때, SplitAfter UTF-8 시퀀스 단위로 분할합니다. /// n은 반환 된 부분 문자열의 수를 결정합니다.

	// n> 0 : 최대 n 개의 부분 문자열. 마지막 부분 문자열에는 분할되지 않은 나머지가 포함된다.
	// n == 0 : 결과는 ni​​l. (부분 문자열 없음)
	// n <0 : 모든 부분 문자열.
	// Title 함수
	// func Title (s string) string

	// Title 문자열 s 단어의 처음 Unicode 문자를 타이틀 케이스에지도 한 카피를 돌려줍니다.

	ex := "Go,4,Data,science"
	fmt.Println(strings.SplitAfter(ex, "4"))
	fmt.Println(len(strings.SplitAfter(ex, "4")))

	// Replace: Text Cleaning
	fmt.Println(strings.Replace(mystr, "Go", "Golang", 1))

	// Join
	f1 := []string{"N","L","P"}
	fmt.Println(strings.Join(f1, "."))

	// Subset/Indexing : Something
	mystem := "running"
	fmt.Println(len(mystem))
	fmt.Println("[0:3]", mystem[0:3])

	// Prefix/Suffix
	fmt.Println("Prefix:ru",strings.HasPrefix(mystem,"ru"))
	fmt.Println("Suffix:ing",strings.HasSuffix(mystem,"ing"))


}
