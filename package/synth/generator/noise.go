package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rawbits2010/LibBitDauer/package/synth/filter"
)

type NoiseType int

const (
	RedNoise NoiseType = iota
	PinkNoise
	WhiteNoise
	BlueNoise
	VioletNoise
)

type NoiseGenerator struct {
	noiseType  NoiseType
	rndSeed    int64
	rndFunc    rand.Rand
	sampleRate uint
	filter     filter.Filter
}

// NewNoiseGenerator creates a noise generator object the implements the
// Generator interface. You can select the noise type with SetNoiseType.
// The sample rate is in Hz and can't be changed later.
func NewNoiseGenerator(sampleRate uint) *NoiseGenerator {
	ngTmp := &NoiseGenerator{
		sampleRate: sampleRate,
		noiseType:  WhiteNoise,
		filter:     nil,
	}
	ngTmp.SetSeed(time.Now().UnixNano())

	return ngTmp
}

// SetSeed sets the seed for the random generator for consistency.
func (ng *NoiseGenerator) SetSeed(seed int64) {
	ng.rndSeed = seed
	ng.rndFunc = *rand.New(rand.NewSource(seed))
}

// SetNoiseType selects the noise type to use. The different types are
// implemented by filtering white noise and are calculation heavy.
func (ng *NoiseGenerator) SetNoiseType(noiseType NoiseType) error {

	// setup filters as seen at https://github.com/DIDAVA/dNoise
	switch noiseType {
	case RedNoise:
		ng.filter = generateFilterChain(15, -6, ng.sampleRate)

	case PinkNoise:
		ng.filter = generateFilterChain(15, -3, ng.sampleRate)

	case WhiteNoise:
		ng.filter = nil

	case BlueNoise:
		ng.filter = generateFilterChain(15, 3, ng.sampleRate)

	case VioletNoise:
		ng.filter = generateFilterChain(51, 6, ng.sampleRate)

	default:
		return fmt.Errorf("invalid noise type: %d", noiseType)
	}

	ng.noiseType = noiseType

	return nil
}

func generateFilterChain(start, step int, sampleRate uint) *filter.FilterChain {

	fch := *filter.NewFilterChain()

	lp := filter.NewLowPassIIR(sampleRate)
	lp.SetCutoff(22050)
	fch.AddFilter(lp)

	hp := filter.NewHighPassIIR(sampleRate)
	hp.SetCutoff(0)
	fch.AddFilter(hp)

	// 31 lowshelf, 22050 highshelf
	octaves := []float64{31, 62, 125, 250, 500, 1000, 2000, 4000, 8000, 16000}
	for freqIdx, freq := range octaves {

		filter := filter.NewPeakingIIR(sampleRate)
		filter.SetCenter(freq)
		filter.SetQualityFactor(0.45)
		filter.SetGaindB(float64(start + (freqIdx * step)))

		fch.AddFilter(filter)
	}

	return &fch
}

// GetNextSample returns the next value for the set type of noise.
func (ng NoiseGenerator) GetNextSample() float64 {

	value := ng.getNextValue()

	if ng.noiseType == WhiteNoise {
		return value
	}

	value = ng.filter.Filter(value) / 4.0

	return value
}

func (ng NoiseGenerator) getNextValue() float64 {
	return ng.rndFunc.Float64()*2 - 1
}

// Reset re-initializes the random generator with the set seed.
func (ng *NoiseGenerator) Reset() {
	ng.SetSeed(ng.rndSeed)
}
