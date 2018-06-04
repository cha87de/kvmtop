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

// PrintValues the collected data for a domain
func (collector *CollectorMEM) PrintValues(domain *models.Domain) []string {
	return memPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *CollectorMEM) PrintFields() []string {
	return []string{
		"ram_total",
		"ram_used",
		"ram_vsize",
		"ram_rss",
		"ram_minflt",
		"ram_cminflt",
		"ram_majflt",
		"ram_cmajflt",
	}
}

// CreateCollectorMEM creates a new memory collector
func CreateCollectorMEM() CollectorMEM {
	return CollectorMEM{}
}
