package buffer

import (
	"math"

	"github.com/rawbits2010/LibBitDauer/package/synth/generator"
)

// CalcSampleLength returns how many samples are in the specified duration
func CalcSampleLength(sampleRate uint, durationMS uint) uint {
	return uint(math.Ceil((float64(sampleRate) * 1 / 1000.0) * float64(durationMS)))
}

// Generate will create a buffer that's large enough to hold a sample of desired size
// and generates the sample into it.
func Generate(durationMS uint, gen generator.Generator) []float64 {

	buffSize := CalcSampleLength(gen.GetSampleRate(), durationMS)
	buffer := make([]float64, buffSize)

	return FillBuffer(buffer, gen)
}

// FillBuffer generate samples into the provided buffer to fill it up
func FillBuffer(buffer []float64, gen generator.Generator) []float64 {

	if buffer == nil { // just to be sure
		return []float64{}
	}

	gen.Reset()
	for i := 0; i < len(buffer); i++ {
		buffer[i] = gen.GetNextSample()
	}

	return buffer
}
