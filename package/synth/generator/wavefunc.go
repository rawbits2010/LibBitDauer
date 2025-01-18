package generator

import "math"

type WaveFunction func(float64) float64

// SineFunction returns the value [0-1] for the given angle in radians,
// using a sine function. Use this in a FunctionGenerator.
func SineFunction(angle float64) float64 {
	samp := math.Sin(angle)
	return samp
}

// SquareFunction returns the value [0-1] for the given angle in radians,
// using a square function. Use this in a FunctionGenerator.
func SquareFunction(angle float64) float64 {
	if angle <= math.Pi {
		return 1
	}
	return -1
}

// SawtoothFunction returns the value [0-1] for the given angle in radians,
// using a sawtooth function. Use this in a FunctionGenerator.
func SawtoothFunction(angle float64) float64 {
	samp := angle / Tau
	return 2*samp - 1
}

// RevSawtoothFunction returns the value [0-1] for the given angle in radians,
// using a reverse sawtooth function. Use this in a FunctionGenerator.
func RevSawtoothFunction(angle float64) float64 {
	samp := SawtoothFunction(angle) * -1
	return samp
}

// TriangleFunction returns the value [0-1] for the given angle in radians,
// using a triangle function. Use this in a FunctionGenerator.
func TriangleFunction(angle float64) float64 {
	samp := SawtoothFunction(angle)
	if samp < 0 {
		samp = -samp
	}
	return 2*samp - 1
}

// FlatlineFunction gives a constant 0 value
func FlatlineFunction(angle float64) float64 {
	return 0
}
