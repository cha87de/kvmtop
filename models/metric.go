package models

// Metric contains a monitoring metric value with current and previous
type Metric struct {
	Name   string
	Values []Measurement
}
