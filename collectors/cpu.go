package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorCPU describes the cpu collector
type CollectorCPU struct {
	models.Collector
}

// Lookup cpu collector data
func (collector *CollectorCPU) Lookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	cpuLookup(domain, libvirtDomain)
}

// Collect cpu collector data
func (collector *CollectorCPU) Collect(domain *models.Domain) {
	cpuCollect(domain)
}

// PrintValues the collected cpu data for a domain
func (collector *CollectorCPU) PrintValues(domain *models.Domain) []string {
	return cpuPrint(domain)
}

// PrintFields the collected cpu data for a domain
func (collector *CollectorCPU) PrintFields() []string {
	return []string{
		"cpu_cores",
		"cpu_total",
		"cpu_steal",
	}
}

// CreateCollectorCPU creates a new cpu collector
func CreateCollectorCPU() CollectorCPU {
	return CollectorCPU{}
}
