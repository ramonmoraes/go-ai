package neural

import "fmt"

type perceptronTrainingData struct {
	Input  []float64
	Output float64
}

func OrExample() {
	var orDataset = []perceptronTrainingData{
		perceptronTrainingData{[]float64{0, 0}, 0}, // false
		perceptronTrainingData{[]float64{0, 1}, 1}, // true
		perceptronTrainingData{[]float64{1, 0}, 1}, // true
		perceptronTrainingData{[]float64{1, 1}, 1}, // true
	}

	perceptron := NewPerceptron(2)
	fmt.Println("Initial weights")
	fmt.Println(perceptron.Weights)

	trainCount := 100
	fmt.Printf("Training %d times\n", trainCount)

	for trainCount > 0 {
		trainCount--
		for _, data := range orDataset {
			perceptron.Train(data.Input, data.Output)
		}
	}

	fmt.Println("Final weights")
	fmt.Println(perceptron.Weights)
	fmt.Println("- Predictions -")
	predict := perceptron.Output([]float64{1, 1})
	fmt.Println(predict)

	predict = perceptron.Output([]float64{1, 0})
	fmt.Println(predict)

	predict = perceptron.Output([]float64{0, 0})
	fmt.Println(predict)
}
