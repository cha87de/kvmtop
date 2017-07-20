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

}

// Collect disk collector data
func (collector *CollectorDISK) Collect(domain *models.Domain, libvirtDomain libvirt.Domain) {

}

// CreateCollectorDISK creates a new cpu collector
func CreateCollectorDISK() CollectorDISK {
	return CollectorDISK{}
}
