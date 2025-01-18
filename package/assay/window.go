package assay

import (
	"math"

	"gonum.org/v1/gonum/dsp/window"
)

/*
Windowing is used to reduce spectral leakage when performing a Fourier Transform,
particularly for non-periodic signals or finite-length samples.

Common Window Types:
- Rectangular (no window): Best for periodic signals but causes the most leakage for
  non-periodic signals.
- Hann window: Reduces leakage well but slightly widens the main lobe.
- Hamming window: Similar to Hann but with slightly better sidelobe attenuation.
- Blackman window: Strong attenuation of side lobes but increases the width of the
  main lobe.

Potential Drawbacks:
- Frequency resolution loss: The windowing function can widen the main lobe in the
  frequency spectrum, reducing the precision with which you can identify frequency
  components.
- Amplitude distortion: The window may reduce the signal's overall amplitude or
  alter the relative amplitudes of different frequency components.
- Introduces bias: The window shape can introduce a bias towards certain frequency
  components depending on the window type.
*/

// Moderate side-lobe suppression with good frequency resolution. Suitable for
// general-purpose use when no specific windowing requirements exist.
func GetHannWindow(size int) []float64 {

	window := make([]float64, size)
	for i := 0; i < size; i++ {
		window[i] = 0.5 * (1 - math.Cos(2*math.Pi*float64(i)/float64(size-1)))
	}

	return window
}

// Better side-lobe suppression than Hann, but slightly worse frequency
// resolution. Useful when minimizing spectral leakage is more important
// than precise frequency localization.
func GetHammingWindow(size int) []float64 {

	window := make([]float64, size)
	for i := 0; i < size; i++ {
		window[i] = 0.54 - 0.46*math.Cos(2*math.Pi*float64(i)/float64(size-1))
	}

	return window
}

// Strongest side-lobe suppression of the three, but the main lobe is wider,
// reducing frequency resolution.
// Ideal for applications where minimizing side-lobes is critical, even at the
// expense of less sharp frequency peaks.
func GetBlackmanWindow(size int) []float64 {

	window := make([]float64, size)
	for i := 0; i < size; i++ {
		window[i] = 0.42 - 0.5*math.Cos(2*math.Pi*float64(i)/float64(size-1)) + 0.08*math.Cos(4*math.Pi*float64(i)/float64(size-1))
	}

	return window
}

func ApplyHannWindow(samples []float64) []float64 {
	return window.Hann(samples)
}

func ApplyHammingWindow(samples []float64) []float64 {
	return window.Hamming(samples)
}

func ApplyBlackmanWindow(samples []float64) []float64 {
	return window.Blackman(samples)
}
