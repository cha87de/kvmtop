package cpucollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/util"
	"gonum.org/v1/gonum/stat"
)

func cpuLookupHost(host *models.Host) {
	cpuinfo := util.GetProcCpuinfo()

	// calculate average MHz
	coreFrequencies := []float64{}
	for _, c := range cpuinfo {
		coreFrequencies = append(coreFrequencies, float64(c.CPUMhz))
	}
	freqMean := stat.Mean(coreFrequencies, nil)
	host.AddMetricMeasurement("cpu_meanfreq", models.CreateMeasurement(freqMean))

	cores := len(coreFrequencies)
	host.AddMetricMeasurement("cpu_cores", models.CreateMeasurement(uint64(cores)))
}

func cpuPrintHost(host *models.Host) []string {
	cpuMeanfreq := collectors.GetMetricFloat64(host.Measurable, "cpu_meanfreq", 0)
	cpuCores := collectors.GetMetricUint64(host.Measurable, "cpu_cores", 0)

	// put results together
	result := append([]string{cpuCores}, cpuMeanfreq)
	return result
}
