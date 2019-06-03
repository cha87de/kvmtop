package models

// Printable represents a set of fields and values to be printed
type Printable struct {
	HostFields   []string
	DomainFields []string
	HostValues   []string
	DomainValues map[string][]string
}
