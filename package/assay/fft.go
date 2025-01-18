package assay

import (
	"math/cmplx"

	"gonum.org/v1/gonum/dsp/fourier"
)

// Advantage: Provides the highest frequency resolution, which is ideal for
// identifying fine frequency details in stationary signals.
// Use case: Use this method when the signal doesn't change much over time and
// you're primarily interested in frequency precision.
func GetFFTForFullSample(sample []float64) []complex128 {
	fft := fourier.NewFFT(len(sample))
	return fft.Coefficients(nil, sample)
}

// Advantage: Better time resolution; allows you to capture how the frequencies
// change over time, though at the cost of frequency resolution.
// Use case: Ideal for signals where you want to observe frequency changes
// over time.
func GetFFTWithShorterWindows(sample []float64, windowSize int) [][]complex128 {

	fft := fourier.NewFFT(windowSize)

	numWindows := len(sample) / windowSize
	windowSpectrums := make([][]complex128, numWindows)

	for i := 0; i < numWindows; i++ {
		start := i * windowSize

		window := sample[start : start+windowSize]
		spectrum := fft.Coefficients(nil, window)

		windowSpectrums[i] = spectrum
	}

	return windowSpectrums
}

// Advantage: Provides a balance between frequency and time resolution. By
// averaging over multiple windows with overlap, you get smoother spectral
// analysis.
// Use case: Useful for analyzing non-stationary signals while maintaining good
// frequency and time resolution.
func GetFFTWithOverlappingWindows(sample []float64, windowSize int, overlap int) [][]complex128 {

	fft := fourier.NewFFT(windowSize)

	stepSize := windowSize - overlap
	numWindows := (len(sample)-windowSize)/stepSize + 1
	windowSpectrums := make([][]complex128, numWindows)

	for i := 0; i < numWindows; i++ {
		start := i * stepSize

		window := sample[start : start+windowSize]
		spectrum := fft.Coefficients(nil, window)

		windowSpectrums[i] = spectrum
	}

	return windowSpectrums
}

// Convert the FFT result to frequencies, magnitude, and phase (rad).
// Will only do the left side - 0 to Nyquist-freq.
func ExtractFFTResults(spectrum []complex128, sampleRate float64) ([]float64, []float64, []float64) {
	n := len(spectrum) / 2
	frequencies := make([]float64, n)
	magnitudes := make([]float64, n)
	phases := make([]float64, n)

	for i := 0; i < n; i++ {
		//mag, phase := cmplx.Polar(spectrum[i])
		frequencies[i] = float64(i) * (float64(n) / sampleRate)
		magnitudes[i] = cmplx.Abs(spectrum[i])
		//phases[i] = math.Atan2(imag(spectrum[i]), real(spectrum[i]))
		phases[i] = cmplx.Phase(spectrum[i])
	}

	return frequencies, magnitudes, phases
}
