package synth

import (
	"github.com/rawbits2010/LibBitDauer/package/synth/buffer"
	"github.com/rawbits2010/LibBitDauer/package/synth/generator"
	"github.com/rawbits2010/LibBitDauer/package/synth/modulator"
)

type OscGeneratorType uint

const (
	OscTypeWave OscGeneratorType = iota
	OscTypePulse
	OscTypeNoise
)

// TODO: extract delay, so GetNextSample don't need to modify the object
type Oscillator struct {
	Wave          *generator.FunctionGenerator
	Pulse         *generator.Generator //TODO
	Noise         *generator.NoiseGenerator
	generatorType OscGeneratorType // select one of the above

	Frequency    float64
	FrequencyMod generator.Generator // modulates Frequency

	Volume    float64
	VolumeMod generator.Generator // modulates Volume

	Envelope    *modulator.ADSR
	UseEnvelope bool // turn on the envelope

	sampleRate uint

	delayS     uint // delays the generator start
	currDelayS uint
}

// NewOscillator creates a new oscillator which has a function generator,
// an optional envelope, a starting volume and frequency value with
// modulation options. The sample rate is in Hz and can't be changed later.
func NewOscillator(sampleRate uint) *Oscillator {
	oscTmp := &Oscillator{
		Wave: generator.NewFunctionGenerator(sampleRate),
		//TODO: Pulse: generator.NewPulseGenerator(sampleRate),
		Noise:         generator.NewNoiseGenerator(sampleRate),
		generatorType: OscTypeWave,
		FrequencyMod:  modulator.NewFlatModulation(0),
		Volume:        1,
		VolumeMod:     modulator.NewFlatModulation(0),
		Envelope:      modulator.NewADSR(sampleRate),
		UseEnvelope:   false,
		sampleRate:    sampleRate,
	}

	return oscTmp
}

// GetSampleRate returns the sample rate with which the oscillator
// was created.
func (osc Oscillator) GetSampleRate() uint {
	return osc.sampleRate
}

// SetDelay sets a delay in milliseconds for the oscillator start.
func (osc *Oscillator) SetDelay(durationMS uint) {
	osc.delayS = buffer.CalcSampleLength(osc.sampleRate, durationMS)
}

// SwitchGeneratorType will instantly switch to another type of function
// generator. Won't change any values for them.
func (osc *Oscillator) SwitchGeneratorType(oscType OscGeneratorType) {
	osc.generatorType = oscType
}

// GetNextSample returns the next sample using the set oscillator state.
func (osc *Oscillator) GetNextSample() float64 {

	if osc.currDelayS < osc.delayS {
		osc.currDelayS++
		return 0
	}

	var sample float64
	switch osc.generatorType {

	case OscTypeWave:
		osc.Wave.Frequency = osc.Frequency + osc.FrequencyMod.GetNextSample()
		sample = osc.Wave.GetNextSample()

	case OscTypePulse:
		// TODO

	case OscTypeNoise:
		sample = osc.Noise.GetNextSample()
	}

	sample *= osc.Volume + osc.VolumeMod.GetNextSample()

	if osc.UseEnvelope {
		sample *= osc.Envelope.GetNextSample()
	}

	return sample
}

// Reset resets ALL the oscillator values to their default.
func (osc *Oscillator) Reset() {
	osc.currDelayS = 0
	osc.Wave.Reset()
	//TODO: osc.Pluse.Reset()
	osc.Noise.Reset()
	osc.FrequencyMod.Reset()
	osc.VolumeMod.Reset()
}
