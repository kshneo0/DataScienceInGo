package main

import (
	"fmt"
	"math"

	"github.com/montanaflynn/stats"
	gstat "gonum.org/v1/gonum/stat"
)

func main() {

	fmt.Println("Statistics in Go")
	// Array/Slice
	even := []float64{2, 4, 6, 8, 10, 8, 8}

	var odd = []float64{3, 5, 7, 9}

	fmt.Println("Even:", even)
	fmt.Println("Odd:", odd)

	// Basic Maths
	// Mean
	evenmean, _ := stats.Mean(even)
	fmt.Println("Mean for Even::", evenmean)
	evenmin, _ := stats.Min(even)
	fmt.Println("Minimum:", evenmin)

	evenmax, _ := stats.Max(even)
	fmt.Println("Maximum:", evenmax)

	// Mode
	// Mode gets the mode [most frequent value(s)] of a slice of float64s
	evenmode, _ := stats.Mode(even)
	fmt.Println(evenmode)

	// Arithmetic,Harm,Geo
	// GeometricMean returns the median of the data
	gmean, _ := stats.GeometricMean(even)
	fmt.Println("GeoMean", gmean)

	// HarmonicMean returns the mode of the data
	harmean, _ := stats.HarmonicMean(even)
	fmt.Println("HarmonicMean", harmean)

	// STD
	oddstd, _ := stats.StandardDeviation(odd)
	fmt.Println("Standard Dev", oddstd)

	// Variance
	oddvar, _ := stats.PopulationVariance(odd)
	fmt.Println("Pop Variance", oddvar)

	// Math Packages
	fmt.Println("STDEV From Math:", math.Sqrt(oddvar))

	// Using Gonum For Stats

	oddmean := gstat.Mean(odd, nil)
	fmt.Println("Mean for Odd via Gonum::", oddmean)

}
