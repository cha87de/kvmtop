package netcollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the network collector
type Collector struct {
	models.Collector
}

// Lookup network collector data
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		netLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect network collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		netCollect(domains[uuid])
	}
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{},
		DomainFields: []string{
			"net_tx",
			"net_rx",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = netPrint(domains[uuid])
	}

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new network collector
func CreateCollector() Collector {
	return Collector{}
}
