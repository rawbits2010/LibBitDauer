package filter

// 1st order High-Pass IIR as seen at https://www.youtube.com/@PhilsLab
type HighPassIIR struct {
	sampleRate uint
	coeff      float64
	lastIn     float64
	lastOut    float64
}

// NewHighPassIIR creates a new 1st order high-pass IIR filter object that
// implements the Filter interface. It lets frequencies through above the
// cutoff frequency.
// The sample rate is in Hz and can't be changed later.
func NewHighPassIIR(sampleRate uint) *HighPassIIR {
	iirTmp := &HighPassIIR{
		sampleRate: sampleRate,
	}
	return iirTmp
}

// SetCutoff Calculates coefficients for the cutoff frequency.
// 0 <= freq <= sampleRate/2
func (iir *HighPassIIR) SetCutoff(freq float64) {

	alpha := Tau * freq / float64(iir.sampleRate)

	iir.coeff = 1 / (1 + alpha)
}

// Filter takes a value and applies the high-pass filter to it.
func (iir *HighPassIIR) Filter(value float64) float64 {

	out := iir.coeff * ((value - iir.lastIn) + iir.lastOut)

	if out > 1 {
		out = 1
	} else if out < -1 {
		out = -1
	}

	iir.lastIn = value
	iir.lastOut = out
	return iir.lastOut
}

// Reset clears the rolling values but keeps the coefficients.
func (iir *HighPassIIR) Reset() {
	iir.lastIn = 0
	iir.lastOut = 0
}
