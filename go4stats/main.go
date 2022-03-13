package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Scalar: A number
	var a int = 44
	fmt.Println("Scalar:", a)
	// Vectors
	// 1 Dim
	// Row Vector/Column Vector[list of same datatype]
	// Method 1
	myvector := []float64{1.2, 3.4, 4.5, 3.5, 4.4}
	fmt.Println(myvector)
	fmt.Printf("%T \n", myvector)

	// Method 2: Using Gonum
	// Syntax: mat.NewVecDense(size,[]type{v1,v2})
	myvectorA := mat.NewVecDense(2, []float64{1.2, 3.4})
	myvectorB := mat.NewVecDense(2, []float64{3.2, 4.4})
	fmt.Printf("%T \n", myvectorA)
	fmt.Printf("%T \n", myvectorB)

	// Dot Product
	dp := mat.Dot(myvectorA, myvectorB)
	fmt.Println(dp)

	// Sum of Vector
	fmt.Println("Sum :", mat.Sum(myvectorA))

	// Max of Vector
	fmt.Println("Max :", mat.Max(myvectorA))

	// Matrix
	// Creating A Matrix
	data := []float64{1.7, 3.4, 3.3, 4.5, 3.3, 1.4}
	// Syntax: (row,col,data)
	// Syntax: (dim/shape,data)
	mymatrix := mat.NewDense(2, 3, data)
	fmt.Println(mymatrix)
	fm := mat.Formatted(mymatrix, mat.Prefix(" "))
	fmt.Println(fm)
	matPrint(mymatrix)

	// Matrix of Zeros
	allzeros := mat.NewDense(3, 2, nil)
	matPrint(allzeros)

	// Shape/Dimension
	fmt.Println(mymatrix.Dims())
	// Get Min/Max
	fmt.Println("Min", mat.Min(mymatrix))
	fmt.Println("Max", mat.Max(mymatrix))

	// Selection
	// arr[0] :python
	// Will not work
	// fmt.Println(mymatrix[0])
	// Get All Values of  Row1
	row1 := mat.Row(nil, 0, mymatrix)
	fmt.Println(row1)
	// Method 2 : Get All Values of A Row
	fmt.Println(mymatrix.RowView(0))

	// Get All values of Col2
	col1 := mat.Col(nil, 0, mymatrix)
	fmt.Println(col1)

	// Method 2 : Get All Values of A Column
	fmt.Println(mymatrix.ColView(0))

	// Select A single value : At
	fmt.Println(mymatrix.At(0, 0))

	// Modify/Change
	// arr[0][0] = 34 : numpy
	// SetCol.Set.SetRow()
	mymatrix.Set(0, 0, 7.1)
	matPrint(mymatrix)

	// Change An Entire Row
	mymatrix.SetRow(1, []float64{3.0, 3.0, 3.0})
	matPrint(mymatrix)

	// Change An Entire Column
	mymatrix.SetCol(1, []float64{0.0, 0.0})
	matPrint(mymatrix)

	// Transpose
	matPrint(mymatrix.T())

}

// To Display Matrix
func matPrint(X mat.Matrix) {
	fm := mat.Formatted(X, mat.Prefix(" "), mat.Squeeze())
	fmt.Printf("%v\n", fm)
}
