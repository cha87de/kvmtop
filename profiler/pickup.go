package profiler

import (
	"strconv"

	"github.com/cha87de/kvmtop/models"

	"github.com/cha87de/kvmtop/collectors/cpucollector"
)

func pickupCPU(domain models.Domain) int {
	cputimeAllCores, _ := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_times"))
	queuetimeAllCores, _ := strconv.Atoi(cpucollector.CpuPrintThreadMetric(&domain, "cpu_runqueues"))
	cpuUtil := cputimeAllCores + queuetimeAllCores
	return cpuUtil
}

func pickupIO(domain models.Domain) int {
	// TODO
	return 0
}

func pickupNet(domain models.Domain) int {
	// TODO
	return 0
}
