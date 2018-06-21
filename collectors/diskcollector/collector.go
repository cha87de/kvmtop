package diskcollector

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the disk collector
type Collector struct {
	models.Collector
}

// Lookup disk collector data
func (collector *Collector) Lookup(domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		diskLookup(domains[uuid], libvirtDomains[uuid])
	}
}

// Collect disk collector data
func (collector *Collector) Collect(domain *models.Domain) {
	diskCollect(domain)
}

// PrintValues the collected data for a domain
func (collector *Collector) PrintValues(domain *models.Domain) []string {
	return diskPrint(domain)
}

// PrintFields the collected data for a domain
func (collector *Collector) PrintFields() []string {
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

// CreateCollector creates a new disk collector
func CreateCollector() Collector {
	return Collector{}
}
