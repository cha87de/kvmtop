package hostcollector

import (
	"os"

	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the host collector
type Collector struct {
	models.Collector
}

// Lookup host collector data
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		hostLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect host collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		hostCollect(domains[uuid])
	}
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{
			"host_name",
		},
		DomainFields: []string{
			"host_name",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = hostPrint(domains[uuid])
	}

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)
	hostname, _ := os.Hostname()
	printable.HostValues = append(printable.HostValues, hostname)

	return printable
}

// CreateCollector creates a new host collector
func CreateCollector() Collector {
	return Collector{}
}
