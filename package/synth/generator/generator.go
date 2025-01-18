package generator

import "math"

const Tau = math.Pi * 2 // why would you only have half the pie?

// Generator is an interface for all function generators and the like
// mainly for the oscillator.
type Generator interface {
	GetSampleRate() uint
	GetNextSample() float64
	Reset() // call this before a new playback
}
