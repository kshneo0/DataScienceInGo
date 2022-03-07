package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Open the file
	// Method 1: Using Ioutil
	content, err := ioutil.ReadFile("data/example_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Do Something

	// Parse/Process
	fmt.Println(string(content))

	fmt.Println("---------------------------")

	// Method 2: Using Os. Open
	f, err := os.Open("data/example_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Parser/Process
	// Create A Scanner
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
