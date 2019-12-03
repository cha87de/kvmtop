package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	drModels "github.com/disresc/lib/models"
)

// NewMeasurable instantiates and returns a new Measurable
func NewMeasurable() *Measurable {
	return &Measurable{
		//Metrics: sync.Map{},
		metrics: make(map[string]Metric),
		access:  sync.Mutex{},
	}
}

// Measurable holds collector metrics in a map
type Measurable struct {
	//Metrics sync.Map
	access  sync.Mutex
	metrics map[string]Metric
}

// AddMetricMeasurement adds a metric measurement
func (measurable *Measurable) AddMetricMeasurement(metricName string, measurement Measurement) {
	measurable.access.Lock()
	defer measurable.access.Unlock()

	// create empty metric if not existent
	if _, ok := measurable.metrics[metricName]; !ok {
		measurable.metrics[metricName] = Metric{
			Name:   metricName,
			Values: []Measurement{},
		}
	}

	metric := measurable.metrics[metricName]

	// add measurement as value to metric
	end := len(metric.Values) - 1
	if end <= 0 {
		end = len(metric.Values)
	}
	metric.Values = append([]Measurement{measurement}, metric.Values[0:end]...)

	// store back
	measurable.metrics[metricName] = metric

	/*items := make(map[string]*drModels.Item)
	items["blabla"] = &drModels.Item{
		Type:  "blabla",
		Value: "Blubb",
	}*/
	SendUpdateEvent(drModels.Event{
		Name:  "das ist die quelle",
		Value: "blablabla",
		//Items:  items,
	})
}

// DelMetricMeasurement removes a metric
func (measurable *Measurable) DelMetricMeasurement(metricName string) {
	measurable.access.Lock()
	defer measurable.access.Unlock()
	delete(measurable.metrics, metricName)
}

// GetMetric reads and returns the metric values by metric name
func (measurable *Measurable) GetMetric(metricName string) (*Metric, bool) {
	measurable.access.Lock()
	defer measurable.access.Unlock()
	var metric Metric
	metric, exists := measurable.metrics[metricName]
	if !exists {
		return &Metric{}, false
	}
	return &metric, true
}

// GetMetricString returns the given metric value at measurement index as string
func (measurable *Measurable) GetMetricString(metricName string, measurementIndex int) string {
	var output string
	if metric, ok := measurable.GetMetric(metricName); ok {
		if len(metric.Values) > measurementIndex {
			byteValue := metric.Values[measurementIndex].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var valuetype string
			decoder.Decode(&valuetype)
			output = fmt.Sprintf("%s", valuetype)
		}
	}
	return output
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

// Dump returns all the measurements as map
func (measurable *Measurable) Dump() map[string]Metric {
	return measurable.metrics
}

// GetMetricUint64 returns the given metric value at measurement index as string
func (measurable *Measurable) GetMetricUint64(metricName string, measurementIndex int) (string, error) {
	var output string
	if metric, ok := measurable.GetMetric(metricName); ok {
		if len(metric.Values) > measurementIndex {
			byteValue := metric.Values[measurementIndex].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var valuetype uint64
			decoder.Decode(&valuetype)
			output = fmt.Sprintf("%d", valuetype)
		} else {
			return "", fmt.Errorf("metric %s has no index %d", metricName, measurementIndex)
		}
	} else {
		// fmt.Printf("metric %s not found, dump: %+v\n", metricName, measurable.Dump())
		return "", fmt.Errorf("metric %s not found", metricName)
	}
	return output, nil
}

// GetMetricFloat64 computes the diff as Float64 between the two measurements for given metric and returns it as string
func (measurable *Measurable) GetMetricFloat64(metricName string, measurementIndex int) string {
	var output string
	if metric, ok := measurable.GetMetric(metricName); ok {
		if len(metric.Values) > measurementIndex {
			byteValue := metric.Values[measurementIndex].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var valuetype float64
			decoder.Decode(&valuetype)
			output = fmt.Sprintf("%f", valuetype)
		}
	}
	return output
}

// GetMetricDiffUint64AsFloat computes the diff as Uint64 between the two measurements for given metric and returns it as string
func (measurable *Measurable) GetMetricDiffUint64AsFloat(metricName string, perTime bool) float64 {
	var output float64
	if metric, ok := measurable.GetMetric(metricName); ok {
		if len(metric.Values) >= 2 {
			// get first value
			byteValue1 := metric.Values[0].Value
			reader1 := bytes.NewReader(byteValue1)
			decoder1 := gob.NewDecoder(reader1)
			var value1 uint64
			decoder1.Decode(&value1)

			// get second value
			byteValue2 := metric.Values[1].Value
			reader2 := bytes.NewReader(byteValue2)
			decoder2 := gob.NewDecoder(reader2)
			var value2 uint64
			decoder2.Decode(&value2)

			// calculate value diff per time
			value := float64(value1 - value2)

			// get time diff
			if perTime {
				timeDiff := metric.Values[0].Timestamp.Sub(metric.Values[1].Timestamp).Seconds()
				value = value / timeDiff
			}
			output = value
		}
	}
	return output
}

// GetMetricDiffUint64 computes the diff as Uint64 between the two measurements for given metric and returns it as string
func (measurable *Measurable) GetMetricDiffUint64(metricName string, perTime bool) string {
	var output string
	diff := measurable.GetMetricDiffUint64AsFloat(metricName, perTime)
	output = fmt.Sprintf("%.0f", diff)
	return output
}
