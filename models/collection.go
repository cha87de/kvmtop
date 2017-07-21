package models

import (
	"bytes"
	"encoding/gob"
	"time"

	libvirt "github.com/libvirt/libvirt-go"
)

// Collection of domains and other stuff
var Collection struct {
	Domains    map[string]*Domain
	Collectors map[string]Collector
}

// Domain defines a domain in libvirt
type Domain struct {
	Name    string
	UUID    string
	Metrics map[string]*Metric
}

// AddMetricMeasurement adds a metric measurement to the domain
func (domain *Domain) AddMetricMeasurement(metricName string, measurement Measurement) {
	// create empty metrics map if not existent
	if domain.Metrics == nil {
		domain.Metrics = make(map[string]*Metric)
	}

	// create empty metric if not existent
	if _, ok := domain.Metrics[metricName]; !ok {

		domain.Metrics[metricName] = &Metric{
			Name:   metricName,
			Values: []Measurement{},
		}
	}
	metrics := domain.Metrics[metricName]

	// add measurement as value to metric
	end := len(metrics.Values) - 1
	if end <= 0 {
		end = len(metrics.Values)
	}
	//domain.Metrics[metricName].Values =
	metrics.Values = append([]Measurement{measurement}, metrics.Values[0:end]...)
}

// Collector defines a collector for a metric (e.g. CPU)
type Collector interface {
	Lookup(domain *Domain, libvirtDomain libvirt.Domain)
	Collect(domain *Domain)
	Print(domain *Domain) []string
}

// Metric contains a monitoring metric value with current and previous
type Metric struct {
	Name   string
	Values []Measurement
}

// Measurement represents one measurement value at a timestamp
type Measurement struct {
	Value     []byte
	Timestamp time.Time
}

// CreateMeasurement creates a new measurement with recent time and bytes from provided value
func CreateMeasurement(value interface{}) Measurement {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(value)
	if err != nil {
		return Measurement{}
	}
	return Measurement{
		Value:     buffer.Bytes(),
		Timestamp: time.Now(),
	}
}
