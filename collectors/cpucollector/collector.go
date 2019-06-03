package cpucollector

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

// Collector describes the cpu collector
type Collector struct {
	models.Collector
}

// Lookup cpu collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		cpuLookup(&domain, libvirtDomain)
		return true
	})

	// lookup details for host
	cpuLookupHost(&models.Collection.Host)
}

// Collect cpu collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		cpuCollect(&domain)
		return true
	})

	// collect host measurements
	cpuCollectHost(&models.Collection.Host)
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	hostFields := []string{
		"cpu_cores",
		"cpu_curfreq",
	}
	if config.Options.Verbose {
		hostFields = append(hostFields,
			"cpu_minfreq",
			"cpu_maxfreq",
		)
	}
	domainFields := []string{
		"cpu_cores",
		"cpu_total",
		"cpu_steal",
	}
	if config.Options.Verbose {
		domainFields = append(domainFields,
			"cpu_other_total",
			"cpu_other_steal",
		)
	}
	printable := models.Printable{
		HostFields:   hostFields,
		DomainFields: domainFields,
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = cpuPrint(&domain)
		return true
	})

	// lookup for host
	printable.HostValues = cpuPrintHost(&models.Collection.Host)

	return printable
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
