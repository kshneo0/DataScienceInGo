package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Data Analysis in Go
	// Open CSV
	csvfile, err := os.Open("goeda/data/diamonds.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	// Read CSV
	df := dataframe.ReadCSV(csvfile)
	// fmt.Println(df)

	// EDA

	// Shape of Dataset
	row, col := df.Dims()
	fmt.Println("Shape of DF:", row, col)

	// Get Only Row Size
	fmt.Println(df.Nrow())

	// Get only Columns Size
	fmt.Println(df.Ncol())

	// Get Column Names
	fmt.Println(df.Names())

	// Get DataTypes
	fmt.Println(df.Types())

	// Describe/Summary
	fmt.Println("Summary", df.Describe())

	// Selection of Columns & Rows
	// Select  columns by Column name
	fmt.Println(df.Select("carat"))

	// Select column by index
	fmt.Println(df.Select(1))

	// Multiple Columns Selection
	//df[["carat","cut"]] //pandas
	//[]string{"carat","cut"} // slice/array
	fmt.Println(df.Select([]string{"carat", "cut"}))

	// Selection of Rows
	// Subset : iloc
	fmt.Println(df.Subset(0))

	// Series and Apply Functions
	ds := df.Col("carat") // A Series
	// fmt.Println(ds)
	// fmt.Printf("%T \n", ds)
	// Apply A Function
	// Get the mean
	dsmean := ds.Mean()
	fmt.Println("Mean of Series:", dsmean)

	// Check For Missing Values
	// fmt.Println(ds.IsNaN())
	gmean := stat.Mean(ds.Float(), nil)
	fmt.Println("Go Num Mean for series:", gmean)

	// Apply Conditions/Filter
	// fmt.Println(df.Select("cut"))
	// fmt.Println(df)
	ispremium := df.Filter(dataframe.F{Colname: "cut", Comparator: "==", Comparando: "Premium"})
	fmt.Println(ispremium.Dims())

}
