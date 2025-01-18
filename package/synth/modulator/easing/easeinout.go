package easing

import "math"

type EaseInOut struct {
	Power float64
}

// NewEaseInOut creates a new ease-in-out object for envelope processing.
// The has a slow start and end, and a fast middle section.
// The power determines the steepness of the curve and needs to be [>=1].
// Values like 2-3 gives a smooth transition, while >5 gives a steep curve.
func NewEaseInOut() *EaseInOut {
	return &EaseInOut{
		Power: 2,
	}
}

// GetValue does a cubic ease-in-out interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (eio EaseInOut) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	/*
		if currVal < 0.5 {
			return 4 * currVal * currVal * currVal
		}
		return 1 - math.Pow(-2*currVal+2, 3)/2
	*/
	if currVal < 0.5 {
		return 0.5 * math.Pow(2*currVal, eio.Power)
	}
	return 1 - 0.5*math.Pow(2*(1-currVal), eio.Power)
}
