package collectors

import (
	"regexp"
	"strconv"

	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

func cpuLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// get amount of cores
	vcpus, err := libvirtDomain.GetVcpus()
	if err != nil {
		return
	}
	cores := len(vcpus)
	newMeasurementCores := models.CreateMeasurement(cores)
	domain.AddMetricMeasurement("cpu_cores", newMeasurementCores)

	// get core thread IDs
	vCPUThreads, err := libvirtDomain.QemuMonitorCommand("info cpus", libvirt.DOMAIN_QEMU_MONITOR_COMMAND_HMP)
	if err != nil {
		return
	}
	regThreadID := regexp.MustCompile("thread_id=([0-9]*)\\s")
	threadIDsRaw := regThreadID.FindAllStringSubmatch(vCPUThreads, -1)
	threadIDs := make([]int, len(threadIDsRaw))
	for i, thread := range threadIDsRaw {
		threadIDs[i], _ = strconv.Atoi(thread[1])
	}
	newMeasurementThreads := models.CreateMeasurement(threadIDs)
	domain.AddMetricMeasurement("cpu_threadIDs", newMeasurementThreads)
}

func cpuCollect(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// TODO
}
