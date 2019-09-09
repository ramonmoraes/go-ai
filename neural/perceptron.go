package neural

import (
	"math"
	"math/rand"
)

type Perceptron struct {
	Weights []float64
	Bias    float64
}

// NewPerceptron returns a Perceptron with initially randomized weights
func NewPerceptron(inputsAmount int8) Perceptron {
	weights := []float64{}

	for inputsAmount > 0 {
		inputsAmount--
		weights = append(weights, rand.Float64())
	}

	return Perceptron{
		Weights: weights,
		Bias:    1,
	}
}

func (p *Perceptron) activateFunction(sum float64) float64 {
	return sigmoid(sum)
}

// Predict is sum of the given inputs with it's weights
func (p *Perceptron) Predict(inputs []float64) float64 {
	var sum float64
	for _, input := range inputs {
		for _, weight := range p.Weights {
			sum += input * weight
		}
	}
	return p.activateFunction(sum + p.Bias)
}

func stepFunction(sum float64, treshold float64) float64 {
	if sum > treshold {
		return 1
	}
	return 0
}

func sigmoid(sum float64) float64 {
	return 1.0 / (1.0 + math.Exp(-sum))
}

func sigmoidDerivated(number float64) float64 {
	return number * (number - 1)
}

func (p *Perceptron) Train(inputs []float64, expectedOutput float64) {
	const learningRate = 0.1
	outputErr := expectedOutput - p.Predict(inputs)
	p.Bias += outputErr
	for i, _ := range p.Weights {
		p.Weights[i] -= outputErr
	}
}
