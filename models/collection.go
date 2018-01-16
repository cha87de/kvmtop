package models

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"

	libvirt "github.com/libvirt/libvirt-go"
)

// Collection of domains and other stuff
var Collection struct {
	Domains    map[string]*Domain
	Collectors map[string]Collector
	Printer    Printer
}

// Domain defines a domain in libvirt
type Domain struct {
	Name string
	UUID string
	//Metrics map[string]*Metric
	Metrics sync.Map
}

// AddMetricMeasurement adds a metric measurement to the domain
func (domain *Domain) AddMetricMeasurement(metricName string, measurement Measurement) {

	// create empty metric if not existent
	if _, ok := domain.Metrics.Load(metricName); !ok {
		domain.Metrics.Store(metricName, &Metric{
			Name:   metricName,
			Values: []Measurement{},
		})
	}
	rawmetrics, _ := domain.Metrics.Load(metricName)
	metrics := rawmetrics.(*Metric)

	// add measurement as value to metric
	end := len(metrics.Values) - 1
	if end <= 0 {
		end = len(metrics.Values)
	}
	//domain.Metrics[metricName].Values =
	metrics.Values = append([]Measurement{measurement}, metrics.Values[0:end]...)

	// store back
	//domain.Metrics.Store(metricName, metrics)
}

// GetMetric reads and returns the metric values by metric name
func (domain *Domain) GetMetric(metricName string) (*Metric, bool) {
	rawmetric, exists := domain.Metrics.Load(metricName)
	if !exists {
		return &Metric{}, false
	}
	metric := rawmetric.(*Metric)
	return metric, true
}

// Collector defines a collector for a metric (e.g. CPU)
type Collector interface {
	Lookup(domain *Domain, libvirtDomain libvirt.Domain)
	Collect(domain *Domain)
	PrintFields() []string
	PrintValues(domain *Domain) []string
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

// Printer defines a printer for output
type Printer interface {
	Open()
	Screen(fields []string, values [][]string)
	Close()
}
