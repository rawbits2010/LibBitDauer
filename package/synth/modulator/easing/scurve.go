package easing

import "math"

type SCurve struct {
	Sharpness float64
	Midpoint  float64
}

// NewSCurve creates a new SCurve object for envelope processing.
//
// The sharpness determines the steepness of the S-curve. Low sharpness
// means a more gradual transition, high sharpness means a steeper curve.
// Use positive values here, like 1.0, 5.0, 10.0, ...
//
// The midpoint specifies the X-axis value where the curve inflection occurs.
// This value should be in the range [0-1]
func NewSCurve(sharpness float64, midpoint float64) *SCurve {
	return &SCurve{
		Sharpness: sharpness,
		Midpoint:  midpoint,
	}
}

// GetValue does a scurve (sigmoid) interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (sc SCurve) GetValue(startPos, endPos, currPos uint) float64 {
	currVal := NormalisePosition(startPos, endPos, currPos)
	// return currVal * currVal * (3 - 2*currVal)
	adjCurrVal := (currVal - sc.Midpoint) / (1 - sc.Midpoint)
	if currVal < sc.Midpoint {
		adjCurrVal = currVal / sc.Midpoint
	}
	return math.Pow(adjCurrVal, sc.Sharpness) / (math.Pow(adjCurrVal, sc.Sharpness) + math.Pow(1-adjCurrVal, sc.Sharpness))
}
