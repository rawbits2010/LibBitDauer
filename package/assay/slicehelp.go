package assay

import "math"

func max(values []float64) float64 {

	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}

func min(values []float64) float64 {

	minVal := values[0]
	for _, v := range values {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}

func mean(values []float64) float64 {

	sum := 0.0
	for _, mag := range values {
		sum += mag
	}

	mean := sum / float64(len(values))

	return mean
}

func normalize(values []float64) []float64 {

	maxVal := max(values)
	for i := range values {
		values[i] /= maxVal
	}

	return values
}

func standardDeviation(values []float64) (float64, float64) {

	mean := mean(values)

	varianceSum := 0.0
	for _, mag := range values {
		varianceSum += (mag - mean) * (mag - mean)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(values)))

	return stdDev, mean
}

// CountValuesInBinsInRange counts how many values fall into each bin within
// the specified range and returns the number of values in each bin.
func CountValuesInBinsInRange(values []float64, minValue, maxValue float64, numBins int) []int {

	bins := make([]int, numBins)

	rangeWidth := float64(maxValue - minValue)
	binWidth := rangeWidth / float64(numBins)

	for _, value := range values {

		if value < minValue || value > maxValue {
			continue
		}

		binIndex := int(math.Floor((value - minValue) / binWidth))

		if binIndex == numBins {
			binIndex = numBins - 1
		}

		bins[binIndex]++
	}

	return bins
}
