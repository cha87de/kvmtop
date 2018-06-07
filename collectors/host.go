package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorHOST describes the host collector
type CollectorHOST struct {
	models.Collector
}

// Lookup host collector data
func (collector *CollectorHOST) Lookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	hostLookup(domain, libvirtDomain)
}

// Collect host collector data
func (collector *CollectorHOST) Collect(domain *models.Domain) {
	hostCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *CollectorHOST) PrintValues(domain *models.Domain) []string {
	return hostPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *CollectorHOST) PrintFields() []string {
	return []string{
		"host_name",
	}
}

// CreateCollectorHOST creates a new host collector
func CreateCollectorHOST() CollectorHOST {
	return CollectorHOST{}
}
