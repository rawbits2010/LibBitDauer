package modulator

type FlatModulation struct {
	sampleRate uint
}

// NewFlatModulation creates a flat modulator which will always returns 0.
// Used for the
func NewFlatModulation(sampleRate uint) *FlatModulation {
	return &FlatModulation{sampleRate: sampleRate}
}

func (fm FlatModulation) GetSampleRate() uint {
	return fm.sampleRate
}

func (fm FlatModulation) GetNextSample() float64 {
	return 0
}

func (fm *FlatModulation) Reset() {}
