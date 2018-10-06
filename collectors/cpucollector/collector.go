package cpucollector

import (
	"github.com/cha87de/kvmtop/models"
)

// Collector describes the cpu collector
type Collector struct {
	models.Collector
}

// Lookup cpu collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		cpuLookup(&domain, libvirtDomain)
		return true
	})

	// lookup details for host
	cpuLookupHost(models.Collection.Host)
}

// Collect cpu collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		cpuCollect(&domain)
		return true
	})
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	printable := models.Printable{
		HostFields: []string{
			"cpu_cores",
			"cpu_meanfreq",
		},
		DomainFields: []string{
			"cpu_cores",
			"cpu_total",
			"cpu_steal",
			// verbose:
			"cpu_other_total",
			"cpu_other_steal",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = cpuPrint(&domain)
		return true
	})

	// lookup for host
	printable.HostValues = cpuPrintHost(models.Collection.Host)

	return printable
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
