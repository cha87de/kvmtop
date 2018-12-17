package profiler

import (
	"fmt"
	"strconv"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/collectors/cpucollector"
)

func pickupCPU(domain models.Domain) int {
	cputimeAllCores, err := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_threadIDs", "cpu_times"))
	if err != nil {
		fmt.Printf("err Atiu cpu_times: %v\n", err)
	}
	queuetimeAllCores, err := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_threadIDs", "cpu_runqueues"))
	if err != nil {
		fmt.Printf("err Atiu cpu_runqueues: %v\n", err)
	}
	cpuUtil := cputimeAllCores + queuetimeAllCores
	return cpuUtil
}

func pickupIO(domain models.Domain) int {
	readBytes, err := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "io_read_bytes", true))
	if err != nil {
		fmt.Printf("err Atiu io_read_bytes: %v\n", err)
	}
	writtenbytes, err := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "io_write_bytes", true))
	if err != nil {
		fmt.Printf("err Atiu io_write_bytes: %v\n", err)
	}
	total := readBytes + writtenbytes
	return total
}

func pickupNet(domain models.Domain) int {
	receivedBytes, err := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedBytes", true))
	if err != nil {
		fmt.Printf("err Atiu net_ReceivedBytes: %v\n", err)
	}
	transmittedBytes, err := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedBytes", true))
	if err != nil {
		fmt.Printf("err Atiu net_TransmittedBytes: %v\n", err)
	}
	total := receivedBytes + transmittedBytes
	// max, _ := strconv.Atoi(collectors.GetMetricUint64(models.Collection.Host.Measurable, "net_host_speed", 0)) // MBit
	// max = max * 1024 * 1024 / 8                                                                                // to Byte
	return total
}
