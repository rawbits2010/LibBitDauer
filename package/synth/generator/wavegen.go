package generator

import (
	"math"
)

type FunctionGenerator struct {
	Frequency float64

	sampleRate      uint
	angleStep       float64
	currentAngle    float64
	phaseShiftAngle float64
	waveFunc        WaveFunction
}

// NewFunctionGenerator creates a new function generator that implements the
// Generator interface. Noise and pulse width modulation is implemented separately.
// The actual function is set with SetFunction. It defaults
// to a flat line with a constant 0 value.
// The sample rate is in Hz and can't be changed later.
func NewFunctionGenerator(sampleRate uint) *FunctionGenerator {

	fgTmp := &FunctionGenerator{
		sampleRate: sampleRate,
		angleStep:  Tau / float64(sampleRate),
		waveFunc:   FlatlineFunction,
	}

	return fgTmp
}

// GetSampleRate returns the sample rate with which the oscillator
// was created.
func (fg FunctionGenerator) GetSampleRate() uint {
	return fg.sampleRate
}

// ShiftPhase sets the phase shift for the generator in degrees.
func (fg *FunctionGenerator) ShiftPhase(angleDeg float64) {

	// normalize to 0-360
	roundedAngle := int(angleDeg) % 360
	angleDeg = angleDeg - math.Floor(angleDeg) + float64(roundedAngle)

	fg.currentAngle += angleDeg / 360.0 * Tau
	if fg.currentAngle >= Tau {
		fg.currentAngle -= Tau
	}

	fg.phaseShiftAngle = fg.currentAngle
}

// SetFunction sets the actual generator function to use.
func (fg *FunctionGenerator) SetFunction(wf WaveFunction) {
	fg.waveFunc = wf
}

// GetNextSample advances the angle and returns the next sample from the
// function generator.
func (fg *FunctionGenerator) GetNextSample() float64 {

	sample := fg.waveFunc(fg.currentAngle)

	fg.currentAngle += fg.angleStep * fg.Frequency
	if fg.currentAngle >= Tau {
		fg.currentAngle -= Tau
	}

	return sample
}

// Reset sets the generator back to it's starting state.
func (fg *FunctionGenerator) Reset() {
	fg.currentAngle = fg.phaseShiftAngle
}
