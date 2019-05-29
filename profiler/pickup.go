package profiler

import (
	"fmt"
	"strconv"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/collectors/cpucollector"
)

func pickupCPU(domain models.Domain) (int, int, int) {
	cputimeAllCores, err := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_threadIDs", "cpu_times"))
	if err != nil {
		fmt.Printf("err Atiu cpu_times: %v\n", err)
	}
	queuetimeAllCores, err := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_threadIDs", "cpu_runqueues"))
	if err != nil {
		fmt.Printf("err Atiu cpu_runqueues: %v\n", err)
	}
	cpuUtil := cputimeAllCores + queuetimeAllCores
	min := 0
	max := 100
	return cpuUtil, min, max
}

func pickupIO(domain models.Domain) (int, int, int) {
	readBytes, err := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "io_read_bytes", true))
	if err != nil {
		fmt.Printf("err Atiu io_read_bytes: %v\n", err)
	}
	writtenbytes, err := strconv.Atoi(collectors.GetMetricDiffUint64(domain.Measurable, "io_write_bytes", true))
	if err != nil {
		fmt.Printf("err Atiu io_write_bytes: %v\n", err)
	}
	total := readBytes + writtenbytes
	min := 0
	// TODO get disk i/o max speed from system
	maxSata3 := 6                        // GBit/s
	max := maxSata3 * 1024 * 1024 * 1024 // Bit/s
	return total, min, max
}

func pickupNet(domain models.Domain) (int, int, int) {
	rawRx := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedBytes", true)
	if rawRx == "" {
		rawRx = "0"
	}
	receivedBytes, err := strconv.Atoi(rawRx)
	if err != nil {
		fmt.Printf("err Atiu net_ReceivedBytes: %v\n", err)
	}
	rawTx := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedBytes", true)
	if rawTx == "" {
		rawTx = "0"
	}
	transmittedBytes, err := strconv.Atoi(rawTx)
	if err != nil {
		fmt.Printf("err Atiu net_TransmittedBytes: %v\n", err)
	}
	total := receivedBytes + transmittedBytes
	min := 0
	rawNetSpeed := collectors.GetMetricUint64(models.Collection.Host.Measurable, "net_host_speed", 0)
	if rawNetSpeed == "0" {
		// set default to 1GBit/s
		defaultSpeed := 1 * 1024 * 1024 * 1024
		rawNetSpeed = fmt.Sprintf("%d", defaultSpeed)
		fmt.Printf("no netspeed given, set default to %s\n", rawNetSpeed)
	}
	max, _ := strconv.Atoi(rawNetSpeed) // MBit
	max = max * 1024 * 1024 / 8         // to Byte
	return total, min, max
}
