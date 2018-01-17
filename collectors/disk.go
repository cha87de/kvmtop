package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// CollectorDISK describes the disk collector
type CollectorDISK struct {
	models.Collector
}

// Lookup disk collector data
func (collector *CollectorDISK) Lookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	diskLookup(domain, libvirtDomain)
}

// Collect disk collector data
func (collector *CollectorDISK) Collect(domain *models.Domain) {
	diskCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *CollectorDISK) PrintValues(domain *models.Domain) []string {
	return diskPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *CollectorDISK) PrintFields() []string {
	return []string{
		"disk_read",
		"disk_write",
	}
}

// CreateCollectorDISK creates a new cpu collector
func CreateCollectorDISK() CollectorDISK {
	return CollectorDISK{}
}
