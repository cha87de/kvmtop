package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorNET describes the network collector
type CollectorNET struct {
	models.Collector
}

// Lookup network collector data
func (collector *CollectorNET) Lookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	netLookup(domain, libvirtDomain)
}

// Collect network collector data
func (collector *CollectorNET) Collect(domain *models.Domain) {
	netCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *CollectorNET) PrintValues(domain *models.Domain) []string {
	return netPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *CollectorNET) PrintFields() []string {
	return []string{
		"net_tx",
		"net_rx",
	}
}

// CreateCollectorNET creates a new network collector
func CreateCollectorNET() CollectorNET {
	return CollectorNET{}
}
