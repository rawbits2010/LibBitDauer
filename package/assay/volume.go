package assay

import "math"

// CalculateVolumeInBand calculates the volume within a specific frequency band
// by summing the magnitudes of frequencies around a target frequency and its
// bandwidth.
func CalculateVolumeInBand(magnitudes []float64, freqs []float64, targetFreq float64, bandwidth float64) float64 {

	totalVolume := 0.0
	count := 0

	for i, freq := range freqs {
		if freq >= targetFreq-bandwidth/2 && freq <= targetFreq+bandwidth/2 {
			totalVolume += magnitudes[i]
			count++
		}
	}

	if count == 0 {
		return 0
	}

	return totalVolume / float64(count)
}

// CalculateTimeVaryingVolume calculates the time-varying volume (envelope) by
// computing the amplitude energy within a sliding window (e.g., 1024 samples)
// and moving the window across time, updating the volume value at each step.
func CalculateTimeVaryingVolume(samples []float64, windowSize int) []float64 {

	numWindows := len(samples) / windowSize

	windowSums := make([]float64, numWindows)
	for i, v := range samples {
		windowSums[i/windowSize] += v * v
	}

	volumes := make([]float64, numWindows)
	for i, v := range windowSums {
		volumes[i] = math.Sqrt(v / float64(windowSize))
	}

	return volumes
}

// GetAttackTime calculates the attack time of a sound, which indicates how
// quickly the sound reaches its maximum volume from silence.
func GetAttackTime(samples []float64, sampleRate int) float64 {

	maxAmplitude := max(samples)
	threshold := maxAmplitude * 0.9 // e.g. 90% of max volume

	for i, sample := range samples {
		if sample >= threshold {
			return float64(i) / float64(sampleRate) // time in sec
		}
	}

	return 0
}
