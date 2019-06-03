package cpucollector

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/util"
	"gonum.org/v1/gonum/stat"
)

func cpuLookupHost(host *models.Host) {

	// collect cpu freq
	cpufreqinfo := util.GetSysCPU()
	coreFreqMin := []float64{}
	coreFreqMax := []float64{}
	coreFreqCurrent := []float64{}
	for _, c := range cpufreqinfo {
		// convert kHz to MHz
		coreFreqMin = append(coreFreqMin, float64(c.MinFreq/1000))
		coreFreqMax = append(coreFreqMax, float64(c.MaxFreq/1000))
		coreFreqCurrent = append(coreFreqCurrent, float64(c.CurFreq/1000))
	}
	coreFreqMinMean := stat.Mean(coreFreqMin, nil)
	host.AddMetricMeasurement("cpu_minfreq", models.CreateMeasurement(coreFreqMinMean))
	coreFreqMaxMean := stat.Mean(coreFreqMax, nil)
	host.AddMetricMeasurement("cpu_maxfreq", models.CreateMeasurement(coreFreqMaxMean))
	coreFreqCurMean := stat.Mean(coreFreqCurrent, nil)
	host.AddMetricMeasurement("cpu_curfreq", models.CreateMeasurement(coreFreqCurMean))

	cpuinfos := util.GetProcCpuinfo()
	cores := len(cpuinfos)
	host.AddMetricMeasurement("cpu_cores", models.CreateMeasurement(uint64(cores)))
}

func cpuCollectHost(host *models.Host) {
	// TODO lookup cpu host utilisation, cf. #23
}

func cpuPrintHost(host *models.Host) []string {
	cpuMinfreq := host.GetMetricFloat64("cpu_minfreq", 0)
	cpuMaxfreq := host.GetMetricFloat64("cpu_maxfreq", 0)
	cpuCurfreq := host.GetMetricFloat64("cpu_curfreq", 0)
	cpuCores, _ := host.GetMetricUint64("cpu_cores", 0)

	// put results together
	result := append([]string{cpuCores}, cpuCurfreq)
	if config.Options.Verbose {
		result = append(result, cpuMinfreq, cpuMaxfreq)
	}
	return result
}
