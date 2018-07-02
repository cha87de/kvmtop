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
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	// lookup for each domain
	for uuid := range domains {
		cpuLookup(domains[uuid], libvirtDomains[uuid])
	}

	// lookup details for host
	cpuLookupHost(host)
}

// Collect cpu collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		cpuCollect(domains[uuid])
	}
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{
			"cpu_cores",
			"cpu_meanfreq",
		},
		DomainFields: []string{
			"cpu_cores",
			"cpu_total",
			"cpu_steal",
			"cpu_other_total",
			"cpu_other_steal",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = cpuPrint(domains[uuid])
	}

	// lookup for host
	printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
