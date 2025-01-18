package easing

// Easing is an interface for easing functions for envelope processing.
type Easing interface {
	GetValue(startPos, endPos, currPos uint) float64
}
