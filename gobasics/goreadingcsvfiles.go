package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	// Open the File
	csvfile, err := os.Open("data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	// Do Something

	// Parse/Process
	// Method 1: Long using csv
	// r := csv.NewReader(csvfile)
	// for {
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// fmt.Println(record)
	// 	fmt.Println(record[3])
	// }

	// Method 2: Using A Package (gota, gframes, dataframe-go)
	df := dataframe.ReadCSV(csvfile)
	fmt.Println(df)

}
