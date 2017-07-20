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
func (collector *CollectorCPU) Collect(domain *models.Domain, libvirtDomain libvirt.Domain) {
	cpuCollect(domain, libvirtDomain)
}

// CreateCollectorCPU creates a new cpu collector
func CreateCollectorCPU() CollectorCPU {
	return CollectorCPU{}
}
