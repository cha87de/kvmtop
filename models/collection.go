package models

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"
)

// Collection of domains and other stuff
var Collection struct {
	Host           *Host
	Domains        Domains
	Collectors     Collectors
	Printer        Printer
	LibvirtDomains LibvirtDomains
}

// Measurable holds collector metrics in a sync.Map
type Measurable struct {
	Metrics sync.Map
}

// Domain defines a domain in libvirt
type Domain struct {
	*Measurable
	Name string
	UUID string
	PID  int
}

// Host defines the local host libvirt runs on
type Host struct {
	*Measurable
}

// Printable represents a set of fields and values to be printed
type Printable struct {
	HostFields   []string
	DomainFields []string
	HostValues   []string
	DomainValues map[string][]string
}

// AddMetricMeasurement adds a metric measurement
func (measurable *Measurable) AddMetricMeasurement(metricName string, measurement Measurement) {
	// create empty metric if not existent
	if _, ok := measurable.Metrics.Load(metricName); !ok {
		measurable.Metrics.Store(metricName, &Metric{
			Name:   metricName,
			Values: []Measurement{},
		})
	}
	rawmetrics, _ := measurable.Metrics.Load(metricName)
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
func (measurable *Measurable) GetMetric(metricName string) (*Metric, bool) {
	rawmetric, exists := measurable.Metrics.Load(metricName)
	if !exists {
		return &Metric{}, false
	}
	metric := rawmetric.(*Metric)
	return metric, true
}

// GetMetricIntArray reads and returns a metric int array by metric name
func (measurable *Measurable) GetMetricIntArray(metricName string) []int {
	var array []int
	if metric, ok := measurable.GetMetric(metricName); ok {
		if len(metric.Values) > 0 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			dec := gob.NewDecoder(reader)
			dec.Decode(&array)
		}
	}
	return array
}

// GetMetricStringArray reads and returns a metric string array by metric name
func (measurable *Measurable) GetMetricStringArray(metricName string) []string {
	var array []string
	if metric, ok := measurable.GetMetric(metricName); ok {
		if len(metric.Values) > 0 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			dec := gob.NewDecoder(reader)
			dec.Decode(&array)
		}
	}
	return array
}

// Collector defines a collector for a domain specific metric (e.g. CPU)
type Collector interface {
	Lookup()
	Collect()
	Print() Printable
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
	Screen(Printable)
	Close()
}
