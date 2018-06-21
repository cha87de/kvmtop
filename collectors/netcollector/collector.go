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
func (collector *Collector) Lookup(domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		netLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect network collector data
func (collector *Collector) Collect(domain *models.Domain) {
	netCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *Collector) PrintValues(domain *models.Domain) []string {
	return netPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *Collector) PrintFields() []string {
	return []string{
		"net_tx",
		"net_rx",
	}
}

// CreateCollector creates a new network collector
func CreateCollector() Collector {
	return Collector{}
}
