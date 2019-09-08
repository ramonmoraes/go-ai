package neural

import "fmt"

type perceptronTrainingData struct {
	Input  []float64
	Output float64
}

var biggerThenFiveDataset = []perceptronTrainingData{
	perceptronTrainingData{[]float64{0, 0, 0}, 0},
	perceptronTrainingData{[]float64{0, 0, 1}, 0},
	perceptronTrainingData{[]float64{0, 1, 0}, 0},
	perceptronTrainingData{[]float64{0, 1, 1}, 0},
	perceptronTrainingData{[]float64{1, 0, 0}, 0},
	perceptronTrainingData{[]float64{1, 0, 1}, 0},
	perceptronTrainingData{[]float64{1, 1, 0}, 1},
	perceptronTrainingData{[]float64{1, 1, 1}, 1},
}

func BiggerThenFiveExample() {
	perceptron := NewPerceptron(3)
	trainCount := 100

	for trainCount > 0 {
		trainCount--
		for _, data := range biggerThenFiveDataset {
			expected := data.Output
			output := perceptron.Output(data.Input)
			perceptron.Train(expected, output)
		}
	}

	predict := perceptron.Output([]float64{1, 1, 1})
	fmt.Println(predict)

	predict = perceptron.Output([]float64{0, 0, 1})
	fmt.Println(predict)

	predict = perceptron.Output([]float64{0, 1, 1})
	fmt.Println(predict)
}
