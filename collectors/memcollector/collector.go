package memcollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the memory collector
type Collector struct {
	models.Collector
}

// Lookup memory collector data
func (collector *Collector) Lookup(domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		memLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect memory collector data
func (collector *Collector) Collect(domain *models.Domain) {
	memCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *Collector) PrintValues(domain *models.Domain) []string {
	return memPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *Collector) PrintFields() []string {
	return []string{
		"ram_total",
		"ram_used",
		"ram_vsize",
		"ram_rss",
		"ram_minflt",
		"ram_cminflt",
		"ram_majflt",
		"ram_cmajflt",
	}
}

// CreateCollector creates a new memory collector
func CreateCollector() Collector {
	return Collector{}
}
