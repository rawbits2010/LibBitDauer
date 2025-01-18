package filter

// 1st order Low-Pass IIR as seen at https://www.youtube.com/@PhilsLab
type LowPassIIR struct {
	sampleRate uint
	coeff      [2]float64
	lastOut    float64
}

// NewLowPassIIR creates a new 1st order low-pass IIR filter object that
// implements the Filter interface. It lets frequencies through below the
// cutoff frequency.
// The sample rate is in Hz and can't be changed later.
func NewLowPassIIR(sampleRate uint) *LowPassIIR {
	iirTmp := &LowPassIIR{
		sampleRate: sampleRate,
	}
	return iirTmp
}

// SetCutoff Calculates coefficients for the cutoff frequency.
// 0 <= freq <= sampleRate/2
func (iir *LowPassIIR) SetCutoff(freq float64) {

	alpha := Tau * freq / float64(iir.sampleRate)

	iir.coeff[0] = alpha / (1 + alpha)
	iir.coeff[1] = 1 / (1 + alpha)
}

// Filter takes a value and applies the low-pass filter to it.
func (iir *LowPassIIR) Filter(value float64) float64 {

	out := iir.coeff[0]*value + iir.coeff[1]*iir.lastOut

	if out > 1 {
		out = 1
	} else if out < -1 {
		out = -1
	}

	iir.lastOut = out
	return iir.lastOut
}

// Reset clears the rolling value but keeps the coefficients.
func (iir *LowPassIIR) Reset() {
	iir.lastOut = 0
}
