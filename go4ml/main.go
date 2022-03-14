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
	// fmt.Println("Save Model")
	// model.PersistToFile("go4ml/logisticHCVmodel.json")

	// Evaluate
	fmt.Println("Evalute Model")
	res := evaluatemodel(xtrain, ytrain, s1, ytest)
	fmt.Printf("%+v\n", res)
}

// Confusion Matrix Struct
type ConfusionMatrix struct {
	positive      int     `json:"positive"`
	negative      int     `json:"negative"`
	truePositive  int     `json:"truePositive"`
	trueNegative  int     `json:"trueNegative"`
	falsePositive int     `json:"falsePositive"`
	falseNegative int     `json:"falseNegative"`
	accuracy      float64 `json:"accuracy"`
	precision     float64 `json:"precision"`
	recall        float64 `json:"recall"`
}

func evaluatemodel(xtrain [][]float64, ytrain []float64, xTest, yTest []float64) ConfusionMatrix {
	model := linear.NewLogistic(base.BatchGA, 0.00001, 0, 1000, xtrain, ytrain)
	err := model.Learn()
	if err != nil {
		log.Fatal(err)
	}
	// c:= color.New(color.FgCyan, color.Bold)
	fmt.Println("Finishing Training")
	// c.Println("Finish Training")
	myprediction, err := model.Predict(xTest)
	if err != nil {
		panic(err)
	}
	//color.Magnta("Prediction:")
	fmt.Println(myprediction)
	fmt.Println(math.Round(myprediction[0]))
	cm := ConfusionMatrix{}
	for _, y := range yTest {
		if y == 1.0 {
			cm.positive++
		}
		if y == 0.0 {
			cm.negative++
		}
	}
	// Evaluate the Model on the Test data
	var decisionBoundary = 0.5
	for i := range xTest {
		prediction, err := model.Predict(xTest)
		if err != nil {
			panic(err)
		}
		y := int(yTest[i])
		positive := prediction[0] >= decisionBoundary

		if y == 1 && positive {
			cm.truePositive++
		}
		if y == 1 && !positive {
			cm.falseNegative++
		}
		if y == 0 && positive {
			cm.falsePositive++
		}
		if y == 0 && !positive {
			cm.trueNegative++
		}
	}

	// c.Println("Calculating Confusion Metrics")
	// Calculate Evaluation Metrics
	cm.accuracy = (float64(cm.truePositive) + float64(cm.trueNegative)) / 
		(float64(cm.positive) + float64(cm.negative))

	cm.precision = float64(cm.truePositive) / (float64(cm.truePositive) + float64(cm.falsePositive))
	cm.recall = float64(cm.truePositive) / (float64(cm.truePositive) + float64(cm.falseNegative))

	return cm
}