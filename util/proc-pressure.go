package util

import (
	"bufio"
	"fmt"
	"os"

	"kvmtop/config"
)

// some avg10=0.00 avg60=0.00 avg300=0.00 total=109155294
// full avg10=0.00 avg60=0.00 avg300=0.00 total=71768841

// ProcPressureMetric describes the type for Metric in a ProcPressure element
type ProcPressureMetric string

const (
	// ProcPressureMetricSome defines the metric type "some" for a ProcPressure element
	ProcPressureMetricSome ProcPressureMetric = "some"
	// ProcPressureMetricFull defines the metric type "full" for a ProcPressure element
	ProcPressureMetricFull ProcPressureMetric = "Full"
)

// ProcPressureResource describes the resource (cpu,io,mem)
type ProcPressureResource string

const (
	// ProcPressureResourceCPU defines the resource type "cpu" for a ProcPressure element
	ProcPressureResourceCPU ProcPressureResource = "cpu"
	// ProcPressureResourceIO defines the metric type "io" for a ProcPressure element
	ProcPressureResourceIO ProcPressureResource = "io"
	// ProcPressureResourceMemory defines the metric type "memory" for a ProcPressure element
	ProcPressureResourceMemory ProcPressureResource = "memory"
)

// ProcPressure describes one row in /proc/pressure/{io,cpu,mem}
type ProcPressure struct {
	Metric ProcPressureMetric // some or full
	Avg10  float64
	Avg60  float64
	Avg300 float64
	Total  uint64
}

// GetProcPressure reads and returns the pressures for the given resource
func GetProcPressure(resource ProcPressureResource) []ProcPressure {
	pressures := []ProcPressure{}
	filepath := fmt.Sprint(config.Options.ProcFS, "/pressure/", resource)

	file, err := os.Open(filepath)
	if err != nil {
		// cannot open file ...
		return pressures
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	format := "%s avg10=%f avg60=%f avg300=%f total=%d"

	for scanner.Scan() {
		row := scanner.Text()
		pressure := ProcPressure{}

		_, err := fmt.Sscanf(
			string(row), format,
			&pressure.Metric,
			&pressure.Avg10,
			&pressure.Avg60,
			&pressure.Avg300,
			&pressure.Total,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse row in proc pressure %s: %s\n", resource, err)
			continue
		}
		pressures = append(pressures, pressure)
	}
	return pressures
}
