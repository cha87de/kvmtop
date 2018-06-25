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
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		ioLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect io collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		ioCollect(domains[uuid])
	}
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{},
		DomainFields: []string{
			"io_rchar",
			"io_wchar",
			"io_syscr",
			"io_syscw",
			"io_read_bytes",
			"io_write_bytes",
			"io_cancelled_write_bytes",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = ioPrint(domains[uuid])
	}

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
