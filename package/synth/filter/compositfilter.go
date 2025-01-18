package filter

type CompositFilter struct {
	chainList []FilterChain
	gainList  []float64
}

// NewCompositFilter creates an object that can hold multiple filter chains,
// each having an a gain value. It implements the Filter interface. Used to
// run the value through the chains separately and returns an average value.
// The gain can be used as a weight.
func NewCompositFilter() *CompositFilter {
	cfTmp := &CompositFilter{}
	cfTmp.Reset()

	return cfTmp
}

// AddFilterChain adds a filter chain with a gain value. The gain can be used
// as a weight and best to have it [0-1].
func (cf *CompositFilter) AddFilterChain(chain FilterChain, gain float64) {
	cf.chainList = append(cf.chainList, chain)
	cf.gainList = append(cf.gainList, gain)
}

// ClearFilters simply empties the filter and gain slices.
func (cf *CompositFilter) ClearFilterChains() {
	cf.chainList = make([]FilterChain, 0)
	cf.gainList = make([]float64, 0)
}

// Filter runs the value through the filter chains separately, taking the gains
// into account, and averages the results.
func (cf *CompositFilter) Filter(value float64) float64 {

	tmpOutValues := make([]float64, len(cf.chainList))

	for chainIdx, chain := range cf.chainList {
		tmpOutValues[chainIdx] = chain.Filter(value)
	}

	var out float64
	for outIdx, outValue := range tmpOutValues {
		out += outValue * cf.gainList[outIdx]
		out /= 2
	}

	return out
}

// Reset simply resets all the filters in the chain.
func (cf *CompositFilter) Reset() {
	for _, filteChain := range cf.chainList {
		filteChain.Reset()
	}
}
