package modulator

import (
	"github.com/rawbits2010/LibBitDauer/package/synth/buffer"
	"github.com/rawbits2010/LibBitDauer/package/synth/modulator/easing"
)

type ADSRPhase int

const (
	Attack ADSRPhase = iota
	Decay
	Sustain
	Release
)

// TODO: need to rethink this so GetNextSample won't modify the object.
// For example make ADSR into a separate Envelope object and pass the
// currSample to the GetNextSample function.
type ADSR struct {
	attackS      uint
	AttackCurve  easing.Easing
	decayS       uint
	DecayCurve   easing.Easing
	sustainS     uint
	sustain      float64
	releaseS     uint
	ReleaseCurve easing.Easing

	sampleRate       uint
	currPhase        ADSRPhase
	releaseTriggered bool
	ManualSustain    bool // infinite sustain time when true
	currSample       uint
}

// NewADSR creates a new ADSR envelope. The sample rate is in Hz and can't
// be changed afterwards.
func NewADSR(sampleRate uint) *ADSR {
	adsrTmp := &ADSR{
		sampleRate:   sampleRate,
		sustain:      1,
		AttackCurve:  *easing.NewLERP(),
		DecayCurve:   *easing.NewLERP(),
		ReleaseCurve: *easing.NewLERP(),
	}
	return adsrTmp
}

// SetAttackLength sets the length of the attack phase in milliseconds.
func (adsr *ADSR) SetAttackLength(attackMS uint) {
	adsr.attackS = buffer.CalcSampleLength(adsr.sampleRate, attackMS)
}

// SetDecayLength sets the length of the decay phase in milliseconds.
func (adsr *ADSR) SetDecayLength(decayMS uint) {
	adsr.decayS = buffer.CalcSampleLength(adsr.sampleRate, decayMS)
}

// SetSustain sets the sustain value directly.
func (adsr *ADSR) SetSustain(value float64) {
	adsr.sustain = value
}

// SetSustainLength Sets the exact sustain time. Make sure you also set
// the release time! Use this when you know the length of the sound.
func (adsr *ADSR) SetSustainLength(sustainMS uint) {
	adsr.sustainS = buffer.CalcSampleLength(adsr.sampleRate, sustainMS)
}

// TriggerRelease Start the release phase instatly when in sustain phase.
// Use SetReleaseLength in this case. Useful for triggering on MIDI key release.
func (adsr *ADSR) TriggerRelease() {
	adsr.releaseTriggered = true
}

// SetReleaseLength Sets the release time. Use this if you want to trigger
// the release specifically with TriggerRelease.
func (adsr *ADSR) SetReleaseLength(releaseMS uint) {
	adsr.releaseS = buffer.CalcSampleLength(adsr.sampleRate, releaseMS)
}

// GetSampleRate returns the sample rate with which the envelope
// was created.
func (adsr ADSR) GetSampleRate() uint {
	return adsr.sampleRate
}

// GetNextSample returns the next envelope value based on the ADSR settings.
func (adsr *ADSR) GetNextSample() float64 {

	curr := adsr.currSample
	adsr.currSample++

	switch adsr.currPhase {

	case Attack:
		if curr == adsr.attackS {
			adsr.currPhase = Decay
			adsr.currSample = 0
		}
		return adsr.AttackCurve.GetValue(0, adsr.attackS, curr)

	case Decay:
		if curr == adsr.decayS {
			adsr.currPhase = Sustain
			adsr.currSample = 0
		}
		return adsr.sustain + ((1 - adsr.sustain) * (1 - adsr.DecayCurve.GetValue(0, adsr.decayS, curr)))

	case Sustain:
		if !adsr.releaseTriggered {
			if !adsr.ManualSustain && curr >= adsr.sustainS {
				adsr.currPhase = Release
				adsr.currSample = 0
			}
			return adsr.sustain
		}
		fallthrough

	default:
		return adsr.sustain * (1 - adsr.ReleaseCurve.GetValue(0, adsr.releaseS, curr))
	}
}

// Reset will set the
func (adsr *ADSR) Reset() {
	adsr.releaseTriggered = false
}
