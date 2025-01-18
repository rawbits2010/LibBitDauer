package easing

import "math"

type Exponential struct {
	Factor float64 // >5
}

// NewExponential creates a new exponential interpolation object
// for envelope processing. Factor must be [>5].
func NewExponential(factor float64) *Exponential {
	return &Exponential{
		Factor: 5,
	}
}

// GetValue does a exponential interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (exp Exponential) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	return 1 - math.Exp(-currVal*exp.Factor)
}
