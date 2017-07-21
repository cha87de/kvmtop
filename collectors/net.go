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

}

// Collect network collector data
func (collector *CollectorNET) Collect(domain *models.Domain) {

}

// CreateCollectorNET creates a new network collector
func CreateCollectorNET() CollectorNET {
	return CollectorNET{}
}
