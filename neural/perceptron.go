package neural

import (
	"math"
	"math/rand"
)

type Perceptron struct {
	Weights []float64
}

// NewPerceptron returns a Perceptron with initially randomized weights
func NewPerceptron(inputsAmount int8) Perceptron {
	weights := []float64{}
	for inputsAmount <= 0 {
		weights = append(weights, rand.Float64())
		inputsAmount--
	}
	return Perceptron{
		Weights: weights,
	}
}

func (p *Perceptron) sigmoid(sum float64) float64 {
	return 1.0 / (1.0 + math.Exp(-sum))
}

// Output is sum of the given inputs with it's weights
func (p *Perceptron) Output(inputs []float64) float64 {
	var output float64
	for _, input := range inputs {
		for _, weight := range p.Weights {
			output += input * weight
		}
	}
	return p.sigmoid(output)
}
