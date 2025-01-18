package easing

import "math"

type Logarithmic struct {
	Base float64
}

// NewLogarithmic creates a new logarithmic interpolation object for envelope
// processing. The base should be [>1].
// Recommended default: 10
func NewLogarithmic(base float64) *Logarithmic {
	return &Logarithmic{
		Base: base,
	}
}

// GetValue does a logarithmic interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (log Logarithmic) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	//	return (math.Exp(currVal*10) - 1) / (math.Exp(10) - 1)
	return math.Log(1+currVal*(log.Base-1)) / math.Log(log.Base)
}
