package hostcollector

import (
	"os"

	"github.com/cha87de/kvmtop/models"
)

// Collector describes the host collector
type Collector struct {
	models.Collector
}

// Lookup host collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		hostLookup(&domain, libvirtDomain)
		return true
	})
}

// Collect host collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		hostCollect(&domain)
		return true
	})
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	printable := models.Printable{
		HostFields: []string{
			"host_name",
		},
		DomainFields: []string{
			"host_name",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = hostPrint(&domain)
		return true
	})

	// lookup for host
	// printable.HostValues = cpuPrintHost(host)
	hostname, _ := os.Hostname()
	printable.HostValues = append(printable.HostValues, hostname)

	return printable
}

// CreateCollector creates a new host collector
func CreateCollector() Collector {
	return Collector{}
}
