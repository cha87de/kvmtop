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

}

// Collect memory collector data
func (collector *CollectorMEM) Collect(domain *models.Domain, libvirtDomain libvirt.Domain) {

}

// CreateCollectorMEM creates a new memory collector
func CreateCollectorMEM() CollectorMEM {
	return CollectorMEM{}
}
