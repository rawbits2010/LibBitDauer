package easing

import "math"

type EaseOut struct {
	Power float64
}

// NewEaseOut creates a new ease-out object for envelope processing.
// The curve has a fast start with a slow end.
// The power determines the steepness of the curve and needs to be [>=1].
// Values like 2-3 gives a smooth transition, while >5 gives a steep curve.
func NewEaseOut(power float64) *EaseOut {
	return &EaseOut{
		Power: power,
	}
}

// GetValue does an ease-out interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (eo EaseOut) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	return 1 - math.Pow(1-currVal, eo.Power)
}
