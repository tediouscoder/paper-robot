package model

// Filter is the interface used for paper filter.
type Filter interface {
	Satisfy(source *Paper) bool
}

// NewFilterByYear create a new FilterByYear.
func NewFilterByYear(year int) *FilterByYear {
	return &FilterByYear{year}
}

// FilterByYear will filter papers by year.
type FilterByYear struct {
	year int
}

// Satisfy implements Filter's Satisfy.
func (f *FilterByYear) Satisfy(source *Paper) bool {
	return f.year == source.Year
}
