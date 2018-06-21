package hostcollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the host collector
type Collector struct {
	models.Collector
}

// Lookup host collector data
func (collector *Collector) Lookup(domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		hostLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect host collector data
func (collector *Collector) Collect(domain *models.Domain) {
	hostCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *Collector) PrintValues(domain *models.Domain) []string {
	return hostPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *Collector) PrintFields() []string {
	return []string{
		"host_name",
	}
}

// CreateCollector creates a new host collector
func CreateCollector() Collector {
	return Collector{}
}
