package easing

import "math"

type InverseLogarithmic struct {
	Base float64
}

// NewInverseLogarithmic creates a new inverse logarithmic interpolation object
// for envelope processing. The base should be [>1].
// Recommended default: 10
func NewInverseLogarithmic(base float64) *InverseLogarithmic {
	return &InverseLogarithmic{
		Base: base,
	}
}

// GetValue does a logarithmic interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (invlog InverseLogarithmic) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	return (math.Pow(invlog.Base, currVal)-1)/invlog.Base - 1
}
