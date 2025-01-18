package assay

import "math"

// CalculateBandwidth calculates the frequency range (bandwidth) of the signal
// based on the provided peaks and their corresponding frequencies.
func CalculateBandwidth(peaks []int, frequencies []float64) float64 {
	return frequencies[peaks[len(peaks)-1]] - frequencies[peaks[0]]
}

// CalculateSpectralCentroid computes the spectral centroid, which is the
// weighted average of the frequencies based on their magnitudes.
func CalculateSpectralCentroid(magnitudes, frequences []float64) float64 {

	var weightedSum float64
	var totalMagnitude float64
	for i, mag := range magnitudes {
		weightedSum += frequences[i] * mag
		totalMagnitude += mag
	}

	return weightedSum / totalMagnitude
}

// CalculateSpectralSlope calculates the spectral slope, which indicates
// how higher frequencies relate to lower frequencies. It helps identify
// whether a sound is "bright" or "dull".
// Dull sounds: Typically have a steeper spectral slope, with less energy
// at higher frequencies.
// Bright sounds: Have a flatter spectral slope, with more energy at higher
// frequencies.
func CalculateSpectralSlope(magnitudes, frequences []float64) float64 {
	meanFreq := 0.0
	meanMag := 0.0
	sumXY := 0.0
	sumXX := 0.0
	totalMagnitudes := 0.0

	// mean freq and mag
	for _, mag := range magnitudes {
		totalMagnitudes += mag
	}
	for i, mag := range magnitudes {
		meanFreq += frequences[i] * mag / totalMagnitudes
		meanMag += mag / totalMagnitudes
	}

	// linear regression
	for i, mag := range magnitudes {
		freq := frequences[i]
		sumXY += (freq - meanFreq) * (mag - meanMag)
		sumXX += (freq - meanFreq) * (freq - meanFreq)
	}

	if sumXX == 0 {
		return 0
	}

	return sumXY / sumXX
}

// CalculateSpectralContrast calculates the spectral contrast, measuring energy
// differences across frequency bands. It helps distinguish uniform energy
// sounds from those with prominent frequency differences.
// Low contrast: Uniform energy, as in white noise or synthetic sounds.
// High contrast: Significant energy differences, as in drums, where low
// frequencies dominate.
func CalculateSpectralContrast(magnitudes []float64, frequences []float64, numBands int) []float64 {

	minFreq := frequences[0]
	maxFreq := frequences[len(frequences)-1]
	bandWidth := (maxFreq - minFreq) / float64(numBands)

	minValues := make([]float64, numBands)
	maxValues := make([]float64, numBands)
	for i := 0; i < numBands; i++ {
		minValues[i] = math.MaxFloat64
		maxValues[i] = -math.MaxFloat64
	}

	for i, mag := range magnitudes {
		bandIndex := int((frequences[i] - minFreq) / bandWidth)

		if bandIndex < numBands {
			if mag < minValues[bandIndex] {
				minValues[bandIndex] = magnitudes[i]
			}
			if mag > maxValues[bandIndex] {
				maxValues[bandIndex] = magnitudes[i]
			}
		}
	}

	contrasts := make([]float64, numBands)
	for i := 0; i < numBands; i++ {
		contrasts[i] = maxValues[i] - minValues[i]
	}

	return contrasts
}

// CalculateTotalEnergyDensity calculates the total energy density of a sound.
func CalculateTotalEnergyDensity(magnitudes []float64) float64 {
	totalEnergy := 0.0
	for _, mag := range magnitudes {
		totalEnergy += mag * mag
	}
	return totalEnergy
}

// CalculateEnergyDensityByBand calculates the energy density for each
// frequency band, showing how energy is distributed across frequency ranges.
// It helps identify where the energy is concentrated: low, mid, or high
// frequencies.
// Example: A bass sound has high energy in low frequencies, while a cymbal or
// synthetic drum sound carries more energy in high frequencies.
func CalculateEnergyDensityByBand(magnitudes, frequences []float64, numBands int) []float64 {

	minFreq := frequences[0]
	maxFreq := frequences[len(frequences)-1]
	bandWidth := (maxFreq - minFreq) / float64(numBands)

	energyByBand := make([]float64, numBands)

	for i, mag := range magnitudes {
		bandIndex := int((frequences[i] - minFreq) / float64(bandWidth))
		energyByBand[bandIndex] += mag * mag
	}

	return energyByBand
}
