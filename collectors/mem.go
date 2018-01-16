package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorMEM describes the memory collector
type CollectorMEM struct {
	models.Collector
}

// Lookup memory collector data
func (collector *CollectorMEM) Lookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	memLookup(domain, libvirtDomain)
}

// Collect memory collector data
func (collector *CollectorMEM) Collect(domain *models.Domain) {
	memCollect(domain)
}

// PrintValues the collected cpu data for a domain
func (collector *CollectorMEM) PrintValues(domain *models.Domain) []string {
	return memPrint(domain)
}

// PrintFields the collected cpu data for a domain
func (collector *CollectorMEM) PrintFields() []string {
	return []string{
		"ram_total",
		"ram_used",
	}
}

// CreateCollectorMEM creates a new memory collector
func CreateCollectorMEM() CollectorMEM {
	return CollectorMEM{}
}
