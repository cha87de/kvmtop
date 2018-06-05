package collectors

import (
	"github.com/cha87de/kvmtop/config"
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
	if config.Options.Verbose {
		return []string{
			"disk_stats_errs",
			"disk_stats_flushreq",
			"disk_stats_flushtotaltimes",
			"disk_stats_rdbytes",
			"disk_stats_rdreq",
			"disk_stats_rdtotaltimes",
			"disk_stats_wrbytes",
			"disk_stats_wrreq",
			"disk_stats_wrtotaltimes",
			"disk_delayblkio",
		}
	}
	return []string{
		"disk_stats_rdbytes",
		"disk_stats_wrbytes",
		"disk_delayblkio",
	}
}

// CreateCollectorDISK creates a new cpu collector
func CreateCollectorDISK() CollectorDISK {
	return CollectorDISK{}
}
