package filter

type BandPassIIR struct {
	hp HighPassIIR
	lp LowPassIIR
}

// NewBandPassIIR creates a new band pass filter object that implements the
// Filter interface. It composed of a high-pass and a low-pass filter with a
// common cutoff frequency.
// The sample rate is in Hz and can't be changed later.
func NewBandPassIIR(sampleRate uint) *BandPassIIR {
	return &BandPassIIR{
		hp: *NewHighPassIIR(sampleRate),
		lp: *NewLowPassIIR(sampleRate),
	}
}

// SetCutoff sets the same cutoff frequency for the high-pass and the low-pass
// filters. The freq value is in Hz.
func (irr *BandPassIIR) SetCutoff(freq float64) {
	irr.hp.SetCutoff(freq)
	irr.lp.SetCutoff(freq)
}

// Filter takes a value and applies the band-pass filter to it.
func (irr *BandPassIIR) Filter(value float64) float64 {
	out := irr.hp.Filter(value)
	out = irr.lp.Filter(out)
	return out
}

// Reset simply resets the high-pass and low-pass filters.
func (irr *BandPassIIR) Reset() {
	irr.hp.Reset()
	irr.lp.Reset()
}
