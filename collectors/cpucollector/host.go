package cpucollector

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"

	"kvmtop/config"
	"kvmtop/models"

	"kvmtop/util"
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

	cores, _ := strconv.Atoi(cpuCores)
	cpuUser := toPercent(host, "cpu_user", cores)
	cpuNice := toPercent(host, "cpu_nice", cores)
	cpuSystem := toPercent(host, "cpu_system", cores)
	cpuIdle := toPercent(host, "cpu_idle", cores)
	cpuIOWait := toPercent(host, "cpu_iowait", cores)
	cpuIRQ := toPercent(host, "cpu_irq", cores)
	cpuSoftIRQ := toPercent(host, "cpu_softirq", cores)
	cpuSteal := toPercent(host, "cpu_steal", cores)
	cpuGuest := toPercent(host, "cpu_guest", cores)
	cpuGuestNice := toPercent(host, "cpu_guestnice", cores)

	// put results together
	result := append([]string{cpuCores}, cpuCurfreq, cpuUser, cpuSystem, cpuIdle, cpuSteal)
	if config.Options.Verbose {
		result = append(result, cpuMinfreq, cpuMaxfreq, cpuNice, cpuIOWait, cpuIRQ, cpuSoftIRQ, cpuGuest, cpuGuestNice)
	}
	return result
}

func toPercent(host *models.Host, metricName string, cores int) string {
	perTime := true
	var output string
	var percent float64
	if metric, ok := host.GetMetric(metricName); ok {
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
				ts1 := metric.Values[0].Timestamp
				ts2 := metric.Values[1].Timestamp
				diffSeconds := ts1.Sub(ts2).Seconds()
				valuePerSecond := value / 100 // since value is in Hz
				ratio := valuePerSecond / diffSeconds
				ratio = ratio / float64(cores)
				percent = ratio * 100 // compute it as percent
			}
			output = fmt.Sprintf("%.0f", percent)
		}
	}
	return output

}
