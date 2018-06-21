package iocollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorIO describes the io collector
type Collector struct {
	models.Collector
}

// Lookup io collector data
func (collector *Collector) Lookup(domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		ioLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect io collector data
func (collector *Collector) Collect(domain *models.Domain) {
	ioCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *Collector) PrintValues(domain *models.Domain) []string {
	return ioPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *Collector) PrintFields() []string {
	return []string{
		"io_rchar",
		"io_wchar",
		"io_syscr",
		"io_syscw",
		"io_read_bytes",
		"io_write_bytes",
		"io_cancelled_write_bytes",
	}
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
