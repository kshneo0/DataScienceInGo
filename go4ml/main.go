package main

import (
	"fmt"
	"log"
	"math"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
)

func main() {

	// Python Pandas
	// Requirement For Dataset
	// Numerical
	// No Missing value & Header
	// Train/Test CSV
	// [xfeatures]<target>

	// Load our dataset
	xtrain, ytrain, err := base.LoadDataFromCSV("go4ml/data/prepdata/trainhcvData.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Testing Data
	xtest, ytest, err := base.LoadDataFromCSV("go4ml/data/prepdata/testhcvData.csv")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Training Dataset")
	fmt.Printf("%T \n", xtrain)
	fmt.Printf("%T \n", ytrain)

	fmt.Println("Testing Dataset")
	fmt.Printf("%T \n", xtest)
	fmt.Printf("%T \n", ytest)

	// Initialize Model
	// Optimization Method ()
	// Learning Rate
	// Regulation : for overfitting
	// Dataset (Xfeatures)
	// Class(binary 0 and 1)
	model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xtrain, ytrain)

	// Train
	err = model.Learn()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trainging Done")
	fmt.Println("Prediction")

	// Prediction
	s1 := []float64{32.0, 1.0, 38.5, 52.5, 7.7, 22.1, 7.5, 6.93, 3.23, 106.0, 12.1, 69.0}
	// Negative No Hep
	s2 := []float64{62.0, 0.0, 32.0, 416.6, 5.9, 110.3, 50.0, 5.57, 6.3, 55.7, 650.9, 68.5}
	// Yes Hep

	mypred1, err := model.Predict(s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mypred1)
	fmt.Println("Rounded Prediction")
	fmt.Println(math.Round(mypred1[0]))

	mypred2, err := model.Predict(s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mypred2)
	fmt.Println("Rounded Prediction")
	fmt.Println(math.Round(mypred2[0]))

	// Save Model
	fmt.Println("Save Model")
	model.PersistToFile("go4ml/logisticHCVmodel.json")

	// Evaluate
}
