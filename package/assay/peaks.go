package assay

// FindDominantFrequencyIdx returns the index of the highest magnitude.
func FindDominantFrequencyIdx(magnitudes []float64) int {

	maxMag := 0.0
	dominantIndex := 0
	for i, v := range magnitudes {

		if v > maxMag {
			maxMag = v
			dominantIndex = i
		}
	}

	return dominantIndex
}

// FindPeakIdxsWithThreshold returns the indexes of magnitudes which are above
// a given threshold.
func FindPeakIdxsWithThreshold(magnitudes []float64, threshold float64) []int {
	peaks := []int{}
	for i := 1; i < len(magnitudes)-1; i++ {
		if magnitudes[i] > magnitudes[i-1] && magnitudes[i] > magnitudes[i+1] && magnitudes[i] > threshold {
			peaks = append(peaks, i)
		}
	}
	return peaks
}

// FindPeakIdxsWithPercentageThreshold returns the indexes of magnitudes
// whose ratio to the maximum magnitude exceeds the given percentage threshold.
func FindPeakIdxsWithPercentageThreshold(magnitudes []float64, threshold float64) []int {

	percentageMagnitudes := GetMagnitudesAsPercentage(magnitudes)

	var peaks []int
	for i, percentage := range percentageMagnitudes {
		if percentage > threshold {
			peaks = append(peaks, i)
		}
	}

	return peaks
}

// GetMagnitudesAsPercentage converts magnitudes to percentage ratios
// relative to the highest magnitude.
func GetMagnitudesAsPercentage(magnitudes []float64) []float64 {

	maxMag := max(magnitudes)

	percentageMagnitudes := make([]float64, len(magnitudes))
	for i, mag := range magnitudes {
		percentageMagnitudes[i] = (mag / maxMag) * 100.0
	}

	return percentageMagnitudes
}

// GetMagnitudePercentageDistribution converts magnitudes to percentage ratios
// relative to the highest value and groups them into buckets.
func GetMagnitudePercentageDistribution(magnitudes []float64, bucketCount int) []int {
	percentageMagnitudes := GetMagnitudesAsPercentage(magnitudes)
	return CountValuesInBinsInRange(percentageMagnitudes, 0, 100, bucketCount)
}

// CalculateThreshold calculates a dynamic threshold based on the magnitude values
func CalculateThreshold(magnitudes []float64) float64 {

	stdDev, mean := standardDeviation(magnitudes)

	threshold := mean + 1.5*stdDev

	return threshold
}
