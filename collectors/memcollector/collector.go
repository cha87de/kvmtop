package memcollector

import (
	"github.com/cha87de/kvmtop/models"
)

// Collector describes the memory collector
type Collector struct {
	models.Collector
}

// Lookup memory collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		memLookup(&domain, libvirtDomain)
		return true
	})
}

// Collect memory collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		memCollect(&domain)
		return true
	})
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	printable := models.Printable{
		HostFields: []string{},
		DomainFields: []string{
			"ram_total",
			"ram_used",
			"ram_vsize",
			"ram_rss",
			"ram_minflt",
			"ram_cminflt",
			"ram_majflt",
			"ram_cmajflt",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = memPrint(&domain)
		return true
	})

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new memory collector
func CreateCollector() Collector {
	return Collector{}
}
