package main

import "fmt"

func main() {
	// Strings
	// String : a collection of characters
	// String : an array/slice of bytes/runes

	var char byte = 'A'
	var char2 rune = 'A'
	fmt.Println("Char as Byte", char)
	fmt.Println("Char as Rune", char2)

	// Type
	fmt.Printf("%T \n", char)
	fmt.Printf("%T \n", char2)
	// Method2: Create Charaters with Method
	charA := byte('A')
	charB := rune('A')
	fmt.Println("Char as Byte:Fxn", charA)
	fmt.Println("Char as Rune:Fxn", charB)

	// Actual Representation
	str1 := string(char)
	str2 := string(char2)
	fmt.Println(str1)
	fmt.Println(str2)
	
	// DataType
	fmt.Printf("%T \n", str1)
	fmt.Printf("%T \n", str2)

	// Representing : method2 %c + Printf()
	fmt.Printf("%c\n", char)
	fmt.Printf("%c\n", char2)

	
}
