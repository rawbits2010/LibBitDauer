package modulator

import "github.com/rawbits2010/LibBitDauer/package/synth/generator"

type LFO struct {
	Generator generator.FunctionGenerator
	Deviation float64 // a multiplyer for the sample
}

// NewLFO creates a new function generator that can be used as a low frequency
// oscillator. The sample rate is in Hz and can't be changed afterwards.
func NewLFO(sampleRate uint) *LFO {
	return &LFO{
		Generator: *generator.NewFunctionGenerator(sampleRate),
	}
}

// GetSampleRate returns the sample rate the function generator is set to.
func (lfo LFO) GetSampleRate() uint {
	return lfo.Generator.GetSampleRate()
}

// GetNextSample returns the next sample from the generator multiplied by
// the Deviation.
func (lfo LFO) GetNextSample() float64 {
	return lfo.Deviation * lfo.Generator.GetNextSample()
}

// Reset simply resets the generator for the LFO.
func (lfo *LFO) Reset() {
	lfo.Generator.Reset()
}
