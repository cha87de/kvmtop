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
	stats := util.GetProcStatCPU()
	maincpuStat := util.ProcStatCPU{}
	for _, s := range stats {
		if s.Name == "cpu" {
			maincpuStat = s
			break
		}
	}
	host.AddMetricMeasurement("cpu_user", models.CreateMeasurement(maincpuStat.User))
	host.AddMetricMeasurement("cpu_nice", models.CreateMeasurement(maincpuStat.Nice))
	host.AddMetricMeasurement("cpu_system", models.CreateMeasurement(maincpuStat.System))
	host.AddMetricMeasurement("cpu_idle", models.CreateMeasurement(maincpuStat.Idle))
	host.AddMetricMeasurement("cpu_iowait", models.CreateMeasurement(maincpuStat.IOWait))
	host.AddMetricMeasurement("cpu_irq", models.CreateMeasurement(maincpuStat.IRQ))
	host.AddMetricMeasurement("cpu_softirq", models.CreateMeasurement(maincpuStat.SoftIRQ))
	host.AddMetricMeasurement("cpu_steal", models.CreateMeasurement(maincpuStat.Steal))
	host.AddMetricMeasurement("cpu_guest", models.CreateMeasurement(maincpuStat.Guest))
	host.AddMetricMeasurement("cpu_guestnice", models.CreateMeasurement(maincpuStat.GuestNice))
}

func cpuPrintHost(host *models.Host) []string {
	cpuMinfreq := host.GetMetricFloat64("cpu_minfreq", 0)
	cpuMaxfreq := host.GetMetricFloat64("cpu_maxfreq", 0)
	cpuCurfreq := host.GetMetricFloat64("cpu_curfreq", 0)
	cpuCores, _ := host.GetMetricUint64("cpu_cores", 0)

	cpuUser := host.GetMetricDiffUint64("cpu_user", true)
	cpuNice := host.GetMetricDiffUint64("cpu_nice", true)
	cpuSystem := host.GetMetricDiffUint64("cpu_system", true)
	cpuIdle := host.GetMetricDiffUint64("cpu_idle", true)
	cpuIOWait := host.GetMetricDiffUint64("cpu_iowait", true)
	cpuIRQ := host.GetMetricDiffUint64("cpu_irq", true)
	cpuSoftIRQ := host.GetMetricDiffUint64("cpu_softirq", true)
	cpuSteal := host.GetMetricDiffUint64("cpu_steal", true)
	cpuGuest := host.GetMetricDiffUint64("cpu_guest", true)
	cpuGuestNice := host.GetMetricDiffUint64("cpu_guestnice", true)

	// put results together
	result := append([]string{cpuCores}, cpuCurfreq, cpuUser, cpuSystem, cpuIdle, cpuSteal)
	if config.Options.Verbose {
		result = append(result, cpuMinfreq, cpuMaxfreq, cpuNice, cpuIOWait, cpuIRQ, cpuSoftIRQ, cpuGuest, cpuGuestNice)
	}
	return result
}
