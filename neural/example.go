package neural

type perceptronTrainingData struct {
	Input  []float32
	Output bool
}

var biggerThenFiveDataset = []perceptronTrainingData{
	perceptronTrainingData{[]float32{0, 0, 0}, false},
	perceptronTrainingData{[]float32{0, 0, 1}, false},
	perceptronTrainingData{[]float32{0, 1, 0}, false},
	perceptronTrainingData{[]float32{0, 1, 1}, false},
	perceptronTrainingData{[]float32{1, 0, 0}, false},
	perceptronTrainingData{[]float32{1, 0, 1}, false},
	perceptronTrainingData{[]float32{1, 1, 0}, true},
	perceptronTrainingData{[]float32{1, 1, 1}, true},
}
