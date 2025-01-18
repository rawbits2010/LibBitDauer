package filter

import (
	"math"
)

// 2nd order Peaking IIR as seen at https://www.youtube.com/@PhilsLab
type PeakingIIR struct {
	sampleRate uint
	centerFreq float64
	bandwidth  float64

	inValues  [3]float64 // 0 current
	outValues [3]float64 // 0 current
	inCoeff   [3]float64
	outCoeff  [3]float64

	sampleTime float64
	wcT        float64
	q          float64
	gain       float64
}

// NewPeakingIIR creates a new 2nd order peaking IIR filter object that
// implements the Filter interface. It processes a frequency spectrum around
// a center frequency. The bandwith can be set directly with SetBandwidth
// or a quality factor with SetQualityFactor.
// The sample rate is in Hz and can't be changed later.
func NewPeakingIIR(sampleRate uint) *PeakingIIR {
	pfTmp := &PeakingIIR{
		sampleRate: sampleRate,
		sampleTime: 1 / float64(sampleRate),
	}

	pfTmp.bandwidth = 1
	pfTmp.gain = 1
	pfTmp.SetCenter(1)

	return pfTmp
}

// SetCenter sets the center frequency of the filter with freq in Hz.
func (iir *PeakingIIR) SetCenter(freq float64) {
	iir.centerFreq = freq

	// Hz to rad/s, pre-warp, multiply by sample time
	iir.wcT = 2 * math.Tan(math.Pi*freq*iir.sampleTime)

	iir.SetBandwidth(iir.bandwidth)
}

// SetBandwidth sets how wide is the spectrum the filter processes in Hz.
func (iir *PeakingIIR) SetBandwidth(freq float64) {
	iir.bandwidth = freq

	iir.q = iir.centerFreq / iir.bandwidth

	iir.SetGainLinear(iir.gain)
}

// SetQualityFactor sets how wide is the spectrum the filter processes
// using a quality factor.
func (iir *PeakingIIR) SetQualityFactor(q float64) {
	iir.q = q

	iir.bandwidth = iir.centerFreq / q

	iir.SetGainLinear(iir.gain)
}

// SetGainLinear sets the gain as a dircets value as follows:
// gain > 1 -> boost, gain < 1 -> cut
func (iir *PeakingIIR) SetGainLinear(gain float64) {
	iir.gain = gain

	iir.calcCoefficients()
}

// SetGaindB sets the gain using decibels
func (iir *PeakingIIR) SetGaindB(gaindB float64) {
	iir.gain = math.Pow(10, gaindB/20.0)

	iir.calcCoefficients()
}

func (iir *PeakingIIR) calcCoefficients() {

	iir.inCoeff[0] = 4 + 2*(iir.gain/iir.q)*iir.wcT + iir.wcT*iir.wcT
	iir.inCoeff[1] = 2*iir.wcT*iir.wcT - 8
	iir.inCoeff[2] = 4 - 2*(iir.gain/iir.q)*iir.wcT + iir.wcT*iir.wcT

	iir.outCoeff[0] = 1 / (4 + 2/iir.q*iir.wcT + iir.wcT*iir.wcT)
	iir.outCoeff[1] = -(2*iir.wcT*iir.wcT - 8)
	iir.outCoeff[2] = -(4 - 2/iir.q*iir.wcT + iir.wcT*iir.wcT)
}

// Filter takes a value and applies the peaking filter to it.
func (iir *PeakingIIR) Filter(value float64) float64 {

	iir.inValues[2] = iir.inValues[1]
	iir.inValues[1] = iir.inValues[0]
	iir.inValues[0] = value

	iir.outValues[2] = iir.outValues[1]
	iir.outValues[1] = iir.outValues[0]

	iir.outValues[0] = (iir.inCoeff[0]*iir.inValues[0] + iir.inCoeff[1]*iir.inValues[1] + iir.inCoeff[2]*iir.inValues[2] +
		(iir.outCoeff[1]*iir.outValues[1] + iir.outCoeff[2]*iir.outValues[2])) * iir.outCoeff[0]

	return iir.outValues[0]
}

// Reset clears the rolling values but keeps the coefficients.
func (iir *PeakingIIR) Reset() {

	for i := 0; i < len(iir.inValues); i++ {
		iir.inValues[i] = 0
	}
	for i := 0; i < len(iir.outValues); i++ {
		iir.outValues[i] = 0
	}
}
