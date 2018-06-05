package collectors

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cha87de/kvmtop/models"
)

func getMetricUint64(domain *models.Domain, metric string, measurementIndex int) string {
	var output string
	if metric, ok := domain.GetMetric(metric); ok {
		if len(metric.Values) > measurementIndex {
			byteValue := metric.Values[measurementIndex].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var valuetype uint64
			decoder.Decode(&valuetype)
			output = fmt.Sprintf("%d", valuetype)
		}
	}
	return output
}

func getMetricDiffUint64(domain *models.Domain, metric string, perTime bool) string {
	var output string
	if metric, ok := domain.GetMetric(metric); ok {
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

			output = fmt.Sprintf("%.0f", value)
		}
	}
	return output
}
