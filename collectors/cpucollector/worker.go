package cpucollector

import (
	"path"
	"path/filepath"
	"regexp"
	"strconv"

	"fmt"

	"github.com/cha87de/kvmtop/collectors"
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

	// cache old thread IDs for cleanup
	var oldThreadIds []int
	oldThreadIds = append(oldThreadIds, domain.GetMetricIntArray("cpu_threadIDs")...)
	oldThreadIds = append(oldThreadIds, domain.GetMetricIntArray("cpu_otherThreadIDs")...)

	// get core thread IDs
	vCPUThreads, err := libvirtDomain.QemuMonitorCommand("info cpus", libvirt.DOMAIN_QEMU_MONITOR_COMMAND_HMP)
	if err != nil {
		return
	}
	regThreadID := regexp.MustCompile("thread_id=([0-9]*)\\s")
	threadIDsRaw := regThreadID.FindAllStringSubmatch(vCPUThreads, -1)
	coreThreadIDs := make([]int, len(threadIDsRaw))
	for i, thread := range threadIDsRaw {
		threadID, _ := strconv.Atoi(thread[1])
		coreThreadIDs[i] = threadID
		oldThreadIds = removeFromArray(oldThreadIds, threadID)
	}
	newMeasurementThreads := models.CreateMeasurement(coreThreadIDs)
	domain.AddMetricMeasurement("cpu_threadIDs", newMeasurementThreads)

	// get thread IDs
	tasksFolder := fmt.Sprint(config.Options.ProcFS, "/", domain.PID, "/task/*")
	files, err := filepath.Glob(tasksFolder)
	if err != nil {
		return
	}
	otherThreadIDs := make([]int, 0)
	i := 0
	for _, f := range files {
		taskID, _ := strconv.Atoi(path.Base(f))
		found := false
		for _, n := range coreThreadIDs {
			if taskID == n {
				// taskID is for vCPU core. skip.
				found = true
				break
			}
		}
		if found {
			// taskID is for vCPU core. skip.
			continue
		}
		// taskID is not for a vCPU core
		otherThreadIDs = append(otherThreadIDs, taskID)
		oldThreadIds = removeFromArray(oldThreadIds, taskID)
		i++
	}
	domain.AddMetricMeasurement("cpu_otherThreadIDs", models.CreateMeasurement(otherThreadIDs))

	// remove cached but not existent thread IDs
	for _, id := range oldThreadIds {
		domain.DelMetricMeasurement(fmt.Sprint("cpu_times_", id))
		domain.DelMetricMeasurement(fmt.Sprint("cpu_runqueues_", id))
		domain.DelMetricMeasurement(fmt.Sprint("cpu_other_times_", id))
		domain.DelMetricMeasurement(fmt.Sprint("cpu_other_runqueues_", id))
	}
}

func cpuCollect(domain *models.Domain) {
	// PART A: stats for VCORES from threadIDs
	cpuCollectMeasurements(domain, "cpu_threadIDs", "cpu_")
	// PART B: stats for other threads (i/o or emulation)
	cpuCollectMeasurements(domain, "cpu_otherThreadIDs", "cpu_other_")

	// PART C: collect frequencies

}

func cpuCollectMeasurements(domain *models.Domain, metricName string, measurementPrefix string) {
	threadIDs := domain.GetMetricIntArray(metricName)
	for _, threadID := range threadIDs {
		schedstat := util.GetProcSchedStat(threadID)
		domain.AddMetricMeasurement(fmt.Sprint(measurementPrefix, "times_", threadID), models.CreateMeasurement(schedstat.Cputime))
		domain.AddMetricMeasurement(fmt.Sprint(measurementPrefix, "runqueues_", threadID), models.CreateMeasurement(schedstat.Runqueue))
	}
}

func cpuPrint(domain *models.Domain) []string {
	cores := collectors.GetMetricUint64(domain.Measurable, "cpu_cores", 0)

	// cpu util for vcores
	cputimeAllCores := cpuPrintThreadMetric(domain, "cpu_threadIDs", "cpu_times")
	queuetimeAllCores := cpuPrintThreadMetric(domain, "cpu_threadIDs", "cpu_runqueues")

	// cpu util for for other threads (i/o or emulation)
	otherCputimeAllCores := cpuPrintThreadMetric(domain, "cpu_otherThreadIDs", "cpu_other_times")
	otherQueuetimeAllCores := cpuPrintThreadMetric(domain, "cpu_otherThreadIDs", "cpu_other_runqueues")

	// put results together
	result := append([]string{cores}, cputimeAllCores, queuetimeAllCores)
	if config.Options.Verbose {
		result = append(result, otherCputimeAllCores, otherQueuetimeAllCores)
	}
	return result
}

func cpuPrintThreadMetric(domain *models.Domain, lookupMetric string, metric string) string {
	threadIDs := domain.GetMetricIntArray(lookupMetric)
	var measurementSum float64
	var measurementCount int
	for _, threadID := range threadIDs {
		metricName := fmt.Sprint(metric, "_", threadID)
		measurementStr := collectors.GetMetricDiffUint64(domain.Measurable, metricName, true)
		if measurementStr == "" {
			continue
		}
		measurement, err := strconv.ParseUint(measurementStr, 10, 64)
		if err != nil {
			continue
		}
		measurementSeconds := float64(measurement) / 1000000000 // since counters are nanoseconds
		measurementSum += measurementSeconds
		measurementCount++
	}

	avg := float64(measurementSum) / float64(measurementCount)
	percent := avg * 100
	return fmt.Sprintf("%.0f", percent)
}

func removeFromArray(s []int, r int) []int {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
