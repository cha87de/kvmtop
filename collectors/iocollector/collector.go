package iocollector

import (
	"github.com/cha87de/kvmtop/models"
)

// Collector describes the io collector
type Collector struct {
	models.Collector
}

// Lookup io collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		ioLookup(&domain, libvirtDomain)
		return true
	})
}

// Collect io collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		ioCollect(&domain)
		return true
	})
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	printable := models.Printable{
		HostFields: []string{},
		DomainFields: []string{
			"io_rchar",
			"io_wchar",
			"io_syscr",
			"io_syscw",
			"io_read_bytes",
			"io_write_bytes",
			"io_cancelled_write_bytes",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = ioPrint(&domain)
		return true
	})

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)

	return printable
}

// CreateCollector creates a new cpu collector
func CreateCollector() Collector {
	return Collector{}
}
