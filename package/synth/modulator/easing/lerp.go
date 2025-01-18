package easing

type LERP struct {
}

// NewLERP creates a new linear interpolation object for envelope processing.
func NewLERP() *LERP {
	return &LERP{}
}

// GetValue does a linear interpolation between startPos and endPos.
// Returns the normalized value [0-1] at currPos.
func (lerp LERP) GetValue(startPos, endPos, currPos uint) float64 {
	return NormalisePosition(startPos, endPos, currPos)
}
