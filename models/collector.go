package models

// Collector defines a collector for a domain specific metric (e.g. CPU)
type Collector interface {
	Lookup()
	Collect()
	Print() Printable
}
