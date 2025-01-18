package easing

import "math"

type EaseIn struct {
	Power float64
}

// NewEaseIn creates a new ease-in object for envelope processing.
// The curve has a slow start with a fast end.
// The power determines the steepness of the curve and needs to be [>=1].
// Values like 2-3 gives a smooth transition, while >5 gives a steep curve.
func NewEaseIn(power float64) *EaseIn {
	return &EaseIn{
		Power: power,
	}
}

// GetValue does an ease-in interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (ei EaseIn) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	return math.Pow(currVal, ei.Power)
}
