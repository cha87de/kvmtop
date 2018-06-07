package collectors

import (
	"bytes"
	"encoding/gob"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"fmt"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
)

func cpuLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// get amount of cores
	vcpus, err := libvirtDomain.GetVcpus()
	if err != nil {
		return
	}
	cores := len(vcpus)
	newMeasurementCores := models.CreateMeasurement(uint64(cores))
	domain.AddMetricMeasurement("cpu_cores", newMeasurementCores)

	// get thread IDs
	tasksFolder := fmt.Sprint(config.Options.ProcFS, "/", domain.PID, "/task/*")
	files, err := filepath.Glob(tasksFolder)
	if err != nil {
		return
	}
	coreThreadIDs := make([]int, cores)
	otherThreadIDs := make([]int, len(files)-cores)
	i := 0
	j := 0
	for _, f := range files {
		taskID, _ := strconv.Atoi(path.Base(f))
		stat := util.GetProcStat(taskID)
		if strings.Contains(stat.Comm, "CPU") {
			// taskID is for a vCPU core
			coreThreadIDs[i] = taskID
			i++
		} else {
			// taskID is not for a vCPU core
			otherThreadIDs[j] = taskID
			j++
		}
	}
	// fmt.Printf("coreThreadIDs: %v \notherThreadIDs: %v", coreThreadIDs, otherThreadIDs)
	domain.AddMetricMeasurement("cpu_threadIDs", models.CreateMeasurement(coreThreadIDs))
	domain.AddMetricMeasurement("cpu_otherThreadIDs", models.CreateMeasurement(otherThreadIDs))
}

func cpuCollect(domain *models.Domain) {
	// PART A: stats for VCORES from threadIDs
	cpuCollectMeasurements(domain, "cpu_threadIDs", "cpu_")
	// PART B: stats for other threads (i/o or emulation)
	cpuCollectMeasurements(domain, "cpu_otherThreadIDs", "cpu_other_")
}

func cpuCollectMeasurements(domain *models.Domain, metricName string, measurementPrefix string) {
	threadIDs := domain.GetMetricIntArray(metricName)
	var cputimes []uint64
	var runqueues []uint64
	for _, threadID := range threadIDs {
		schedstat := util.GetProcSchedStat(threadID)
		cputimes = append(cputimes, schedstat.Cputime)
		runqueues = append(runqueues, schedstat.Runqueue)
	}
	domain.AddMetricMeasurement(fmt.Sprint(measurementPrefix, "times"), models.CreateMeasurement(cputimes))
	domain.AddMetricMeasurement(fmt.Sprint(measurementPrefix, "runqueues"), models.CreateMeasurement(runqueues))
}

func cpuPrint(domain *models.Domain) []string {
	cores := getMetricUint64(domain, "cpu_cores", 0)

	// cpu util for vcores
	cputimeAllCores := cpuPrintThreadMetric(domain, "cpu_times")
	queuetimeAllCores := cpuPrintThreadMetric(domain, "cpu_runqueues")

	// cpu util for for other threads (i/o or emulation)
	otherCputimeAllCores := cpuPrintThreadMetric(domain, "cpu_other_times")
	otherQueuetimeAllCores := cpuPrintThreadMetric(domain, "cpu_other_runqueues")

	// put results together
	result := append([]string{cores}, cputimeAllCores)
	result = append(result, queuetimeAllCores)
	result = append(result, otherCputimeAllCores)
	result = append(result, otherQueuetimeAllCores)

	return result
}

func cpuPrintThreadMetric(domain *models.Domain, metric string) string {
	var times []string
	var timeAllCores string
	if metric, ok := domain.GetMetric(metric); ok {
		if len(metric.Values) > 1 {
			byteValue1 := metric.Values[0].Value
			reader1 := bytes.NewReader(byteValue1)
			dec1 := gob.NewDecoder(reader1)

			byteValue2 := metric.Values[1].Value
			reader2 := bytes.NewReader(byteValue2)
			dec2 := gob.NewDecoder(reader2)

			var timesRaw1 []uint64
			var timesRaw2 []uint64
			dec1.Decode(&timesRaw1)
			dec2.Decode(&timesRaw2)

			timeDiff := metric.Values[0].Timestamp.Sub(metric.Values[1].Timestamp).Seconds()
			timeConversionFactor := 1000000000 / timeDiff

			// for each thread ...
			var timeSum float64
			for i, time1 := range timesRaw1 {
				if len(timesRaw2) <= i {
					continue
				}
				time2 := timesRaw2[i]
				if time1 < time2 {
					// unexpected case, since dealing with counters
					time2 = time1
				}
				time := float64(time1-time2) / timeConversionFactor
				timeSum = timeSum + time
				times = append(times, fmt.Sprintf("%.0f", time*100))
			}
			timeAllCores = fmt.Sprintf("%.0f", timeSum/float64(len(timesRaw1))*100)
		}
	}
	return timeAllCores
}
