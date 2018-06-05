package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorIO describes the io collector
type CollectorIO struct {
	models.Collector
}

// Lookup io collector data
func (collector *CollectorIO) Lookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	ioLookup(domain, libvirtDomain)
}

// Collect io collector data
func (collector *CollectorIO) Collect(domain *models.Domain) {
	ioCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *CollectorIO) PrintValues(domain *models.Domain) []string {
	return ioPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *CollectorIO) PrintFields() []string {
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

// CreateCollectorIO creates a new cpu collector
func CreateCollectorIO() CollectorIO {
	return CollectorIO{}
}
