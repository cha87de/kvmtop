package diskcollector

import (
	"strings"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the disk collector
type Collector struct {
	models.Collector
}

// Lookup disk collector data
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	hostDiskSources := ""
	for uuid := range domains {
		diskLookup(domains[uuid], libvirtDomains[uuid])

		// merge sourcedir metrics from domains to one metric for host
		disksources := strings.Split(collectors.GetMetricString(domains[uuid].Measurable, "disk_sources", 0), ",")
		for _, disksource := range disksources {
			if !strings.Contains(hostDiskSources, disksource) {
				if hostDiskSources != "" {
					hostDiskSources += ","
				}
				hostDiskSources += disksource
			}
		}
	}
	host.AddMetricMeasurement("disk_sources", models.CreateMeasurement(hostDiskSources))

	diskHostLookup(host)
}

// Collect disk collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		diskCollect(domains[uuid])
	}
	diskHostCollect(host)
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{
			"disk_device_reads",
			"disk_device_readsmerged",
			"disk_device_sectorsread",
			"disk_device_timereading",
			"disk_device_writes",
			"disk_device_writesmerged",
			"disk_device_sectorswritten",
			"disk_device_timewriting",
			"disk_device_currentops",
			"disk_device_timeforops",
			"disk_device_weightedtimeforops",
		},
		DomainFields: []string{
			"disk_size_capacity",
			"disk_size_allocation",
			"disk_size_physical",
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
	printable.HostValues = diskPrintHost(host)

	return printable
}

// CreateCollector creates a new disk collector
func CreateCollector() Collector {
	return Collector{}
}
