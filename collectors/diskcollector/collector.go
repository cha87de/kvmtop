package diskcollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the disk collector
type Collector struct {
	models.Collector
}

// Lookup disk collector data
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		diskLookup(domains[uuid], libvirtDomains[uuid])
	}
	diskHostLookup()
}

// Collect disk collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		diskCollect(domains[uuid])
	}
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{},
		DomainFields: []string{
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
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = diskPrint(domains[uuid])
	}

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new disk collector
func CreateCollector() Collector {
	return Collector{}
}
