package filter

import "math"

// 2nd order Notch IIR as seen at https://www.youtube.com/@PhilsLab
type NotchIIR struct {
	sampleRate uint
	sampleTime float64

	centerFreq float64
	wc         float64
	notchwidth float64
	ww         float64

	alpha     float64
	beta      float64
	inValues  [3]float64
	outValues [3]float64
}

// NewNotchIIR creates a new 2nd order notch IIR filter object that
// implements the Filter interface. It filters a frequency spectrum around
// a center frequency. The bandwith can be set directly with SetNotchwidth
// or a quality factor with SetQualityFactor.
//
// This filter is tipically used to remove certain frequencies so there is
// no variable gain function. Use a PeakingIIR filter if you need that.
//
// The sample rate is in Hz and can't be changed later.
func NewNotchIIR(sampleRate uint) *NotchIIR {
	nfTmp := &NotchIIR{
		sampleRate: sampleRate,
		sampleTime: 1 / float64(sampleRate),
	}

	nfTmp.SetCenter(1000)
	nfTmp.SetNotchwidth(100)

	return nfTmp
}

// SetCenter sets the center frequency of the filter with freq in Hz.
func (iir *NotchIIR) SetCenter(freq float64) {
	iir.centerFreq = freq

	iir.wc = (2 * float64(iir.sampleRate)) * math.Tan(math.Pi*freq/float64(iir.sampleRate))
	iir.alpha = 4 + iir.wc*iir.wc/(float64(iir.sampleRate)*float64(iir.sampleRate))
}

// SetBandwidth sets how wide is the spectrum the filter processes in Hz.
func (iir *NotchIIR) SetNotchwidth(freq float64) {
	iir.notchwidth = freq

	iir.ww = Tau * freq
	iir.beta = 2 * iir.ww / float64(iir.sampleRate)
}

// SetQualityFactor sets how wide is the spectrum the filter processes
// using a quality factor.
func (iir *NotchIIR) SetQualityFactor(q float64) {
	iir.notchwidth = iir.centerFreq / q

	iir.ww = Tau * iir.notchwidth
	iir.beta = 2 * iir.ww / float64(iir.sampleRate)
}

// Filter takes a value and applies the notch filter to it.
func (iir *NotchIIR) Filter(value float64) float64 {

	iir.inValues[2] = iir.inValues[1]
	iir.inValues[1] = iir.inValues[0]
	iir.inValues[0] = value

	iir.outValues[2] = iir.outValues[1]
	iir.outValues[1] = iir.outValues[0]

	iir.outValues[0] = (iir.alpha*iir.inValues[0] + 2*(iir.alpha-8)*iir.inValues[1] + iir.alpha*iir.inValues[2] -
		(2*(iir.alpha-8)*iir.outValues[1] + (iir.alpha-iir.beta)*iir.outValues[2])) /
		(iir.alpha + iir.beta)

	return iir.outValues[0]
}

// Reset clears the rolling values but keeps the coefficients.
func (iir *NotchIIR) Reset() {

	for i := 0; i < len(iir.inValues); i++ {
		iir.inValues[i] = 0
	}
	for i := 0; i < len(iir.outValues); i++ {
		iir.outValues[i] = 0
	}
}
