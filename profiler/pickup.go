package profiler

import (
	"strconv"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/collectors/cpucollector"
)

func pickupCPU(domain models.Domain) int {
	cputimeAllCores, _ := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_threadIDs", "cpu_times"))
	queuetimeAllCores, _ := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_threadIDs", "cpu_runqueues"))
	cpuUtil := cputimeAllCores + queuetimeAllCores
	return cpuUtil
}

func pickupIO(domain models.Domain) int {
	// TODO
	return 0
}

func pickupNet(domain models.Domain) int {
	receivedBytes, _ := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedBytes", true))
	transmittedBytes, _ := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedBytes", true))
	total := receivedBytes + transmittedBytes
	// max, _ := strconv.Atoi(collectors.GetMetricUint64(models.Collection.Host.Measurable, "net_host_speed", 0)) // MBit
	// max = max * 1024 * 1024 / 8                                                                                // to Byte
	return total
}
