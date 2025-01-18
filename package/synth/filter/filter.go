package filter

import "math"

const Tau = math.Pi * 2 // why would you only have half the pie?

// Filter is an interface for filter implementations.
type Filter interface {
	Filter(float64) float64
	Reset()
}
