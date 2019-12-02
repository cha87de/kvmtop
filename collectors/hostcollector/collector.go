package hostcollector

import (
	"kvmtop/config"
	"kvmtop/models"
)

// Collector describes the host collector
type Collector struct {
	models.Collector
}

// Lookup host collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		domainLookup(&domain, libvirtDomain)
		return true
	})
	hostLookup(&models.Collection.Host)
}

// Collect host collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		domainCollect(&domain)
		return true
	})
	hostCollect(&models.Collection.Host)
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	hostFields := []string{
		"host_name",
	}
	domainFields := []string{
		"host_name",
	}

	if config.Options.Verbose {
		hostFields = append(hostFields,
			"host_uuid",
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
		printable.DomainValues[uuid] = domainPrint(&domain)
		return true
	})

	// lookup for host
	printable.HostValues = hostPrint(&models.Collection.Host)

	return printable
}

// CreateCollector creates a new host collector
func CreateCollector() Collector {
	return Collector{}
}
