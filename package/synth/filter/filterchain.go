package filter

type FilterChain struct {
	filterChain []Filter
}

// NewFilterChain creates a filter chain object that chains any object that
// implements the Filter interface. The processing order is in the order
// of the AddFilter calls.
func NewFilterChain() *FilterChain {
	return &FilterChain{
		filterChain: make([]Filter, 0),
	}
}

// AddFilter adds a filter to the filter chain. The processing order is in
// the order of these calls.
func (fch *FilterChain) AddFilter(filter Filter) {
	fch.filterChain = append(fch.filterChain, filter)
}

// ClearFilters simply empties the filter slice.
func (fch *FilterChain) ClearFilters() {
	fch.filterChain = make([]Filter, 0)
}

// Filter takes a value and runs it through the filter chain.
func (fch *FilterChain) Filter(value float64) float64 {

	for _, filter := range fch.filterChain {
		value = filter.Filter(value)
	}

	return value
}

// Reset simply resets all the filters in the chain.
func (fch *FilterChain) Reset() {
	for _, filter := range fch.filterChain {
		filter.Reset()
	}
}
