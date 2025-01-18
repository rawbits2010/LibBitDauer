package main

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
	"github.com/rawbits2010/LibBitDauer/package/assay"
)

func main() {

	file, err := os.Open("out.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	audioBuf, err := decoder.FullPCMBuffer()
	if err != nil {
		panic(err)
	}

	sampleRate := float64(audioBuf.Format.SampleRate)
	samples := audioBuf.AsFloatBuffer().Data

	windowedSignal := assay.ApplyHammingWindow(samples)
	spectrum := assay.GetFFTForFullSample(windowedSignal)

	frequencies, magnitudes, phases := assay.ExtractFFTResults(spectrum, sampleRate)

	dominantFreqIdx := assay.FindDominantFrequencyIdx(magnitudes)
	fmt.Printf("Dominant Frequency: %.2f Hz\n", frequencies[dominantFreqIdx])

	err = assay.ExportFFTResultsToCSV(frequencies, magnitudes, phases, "fft_spectrum.csv")
	if err != nil {
		fmt.Println("Error writing file:", err)
	} else {
		fmt.Println("FFT Spectrum exported to fft_spectrum.csv")
	}

	err = assay.PlotMagnitudeAndPhase(frequencies, magnitudes, phases, "fft_spectrum.png")
	if err != nil {
		fmt.Println("Error writing file:", err)
	} else {
		fmt.Println("FFT Spectrum exported to fft_spectrum.png")
	}

	threshold := 10.0
	peaks := assay.FindPeakIdxsWithPercentageThreshold(magnitudes, threshold)

	fmt.Println("Found Peaks:")
	for _, index := range peaks {
		fmt.Printf("Frequency: %.2f Hz, Magnitude: %.5f\n", frequencies[index], magnitudes[index])
	}

	counts := assay.GetMagnitudePercentageDistribution(magnitudes, 100)
	err = assay.PlotBins(counts, 0, 100, 100, "fft_percdist.png")
	if err != nil {
		fmt.Println("Error writing file:", err)
	} else {
		fmt.Println("Magnitude distribution exported to fft_percdist.png")
	}

	waveform := IdentifyByHarmonicRatios(peaks, magnitudes)
	fmt.Printf("Identified waveform based on harmonics ratios: %s\n", waveform)

	waveform = IdentifyWaveformWithMetrics(peaks, magnitudes, frequencies)
	fmt.Printf("Identified waveform with metrics: %s\n", waveform)
}

// An example way of the identification process is to check the magnitude changes
// in the harmonics.
// This is only an example, peaks need to belong to a single signal. The ratios
// and number of peaks are all tunable values, and they depend on the peak finding
// method.
func IdentifyByHarmonicRatios(peaks []int, magnitudes []float64) string {

	harmonicRatios := make([]float64, len(peaks))
	for i, peakIdx := range peaks {
		harmonicRatios[i] = magnitudes[peakIdx] / magnitudes[0]
	}

	// values needs to be tuned
	if len(peaks) == 1 {
		// only base frequency
		return "Sine wave"
	} else if len(peaks) > 5 && (len(harmonicRatios) > 1 && harmonicRatios[1] > 0.5) {
		// base frequency + odd harmonics (3f, 5f, 7f...).
		return "Square wave"
	} else if len(peaks) <= 5 && (len(harmonicRatios) > 1 && harmonicRatios[1] < 0.5) {
		// base frequency + all odd harmonics, but amplitude shows heavily decay (e.g., 1/9, 1/25).
		return "Triangle wave"
	} else {
		// base frequency + all harmonics (odd and even)
		// or we have a complex waveform
		return "Sawtooth or Complex waveform"
	}
}

// An example way of the identification process is to check the bandwidth and
// the spectral centroid of the signal.
// This is only an example, peaks need to belong to a single signal. The ratios
// and number of peaks are all tunable values, and they depend on the peak finding
// method.
func IdentifyWaveformWithMetrics(peaks []int, magnitudes, frequencies []float64) string {

	bandwidth := assay.CalculateBandwidth(peaks, frequencies)
	spectralCentroid := assay.CalculateSpectralCentroid(magnitudes, frequencies)

	// values needs to be tuned
	if bandwidth > 1000 && spectralCentroid > 1500 {
		// Wider bandwidth due to odd harmonics, with the spectral centroid shifted to higher frequencies.
		// High bandwidth and spectral centroid suggest a square wave.
		return "Square wave"
	} else if bandwidth < 300 && spectralCentroid < 1000 {
		// A sine wave has no harmonics, so its bandwidth is very small or zero.
		// The spectral centroid is at the fundamental frequency, as there are no higher harmonics.
		// If the bandwidth is small and the spectral centroid is at the fundamental, it's likely a sine wave.
		return "Sine wave"
	} else if bandwidth >= 300 && bandwidth <= 1000 {
		// Typically has a narrower bandwidth with harmonics decreasing with frequency.
		// The spectral centroid is lower than for square waves, as harmonics decrease relative to the fundamental.
		// If the bandwidth is medium and the spectral centroid is lower, it's likely a triangle wave.
		return "Triangle wave"
	} else {
		return "Complex waveform"
	}
}
