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
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		memLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect memory collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		memCollect(domains[uuid])
	}
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{},
		DomainFields: []string{
			"ram_total",
			"ram_used",
			"ram_vsize",
			"ram_rss",
			"ram_minflt",
			"ram_cminflt",
			"ram_majflt",
			"ram_cmajflt",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = memPrint(domains[uuid])
	}

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new memory collector
func CreateCollector() Collector {
	return Collector{}
}
