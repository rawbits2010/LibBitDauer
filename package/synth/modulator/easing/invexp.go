package easing

import "math"

type InverseExponential struct {
	Factor float64 // >5
}

// NewInverseExponential creates a new inverse exponential interpolation object
// for envelope processing. Factor must be [>5].
func NewInverseExponential() *InverseExponential {
	return &InverseExponential{
		Factor: 5,
	}
}

// GetValue does an inverse exponential interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (invexp InverseExponential) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	return math.Exp(currVal*invexp.Factor)/math.Exp(invexp.Factor) - 1
}
