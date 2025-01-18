package assay

// GetZeroCrossingRate computes the Zero Crossing Rate (ZCR) of the given
// audio samples.
// The ZCR represents the rate at which the signal changes sign
// (crosses the zero amplitude level).
// Higher ZCR suggests noisier or higher-frequency sounds, while lower ZCR
// indicates deeper, clearer sounds.
func GetZeroCrossingRate(samples []float64) float64 {
	crossings := 0
	for i := 1; i < len(samples); i++ {
		if samples[i-1] >= 0 && samples[i] < 0 || samples[i-1] < 0 && samples[i] >= 0 {
			crossings++
		}
	}

	return float64(crossings) / float64(len(samples)-1)
}
