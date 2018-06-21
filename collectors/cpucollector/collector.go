package cpucollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the cpu collector
type Collector struct {
	models.Collector
}

// Lookup cpu collector data
func (collector *Collector) Lookup(domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		cpuLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect cpu collector data
func (collector *Collector) Collect(domain *models.Domain) {
	cpuCollect(domain)
}

// PrintValues the collected cpu data for a domain
func (collector *Collector) PrintValues(domain *models.Domain) []string {
	return cpuPrint(domain)
}

// PrintFields the collected cpu data for a domain
func (collector *Collector) PrintFields() []string {
	return []string{
		"cpu_cores",
		"cpu_total",
		"cpu_steal",
		"cpu_other_total",
		"cpu_other_steal",
	}
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
