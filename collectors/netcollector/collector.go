package netcollector

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

// Collector describes the network collector
type Collector struct {
	models.Collector
}

// Lookup network collector data
func (collector *Collector) Lookup() {
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		libvirtDomain, _ := models.Collection.LibvirtDomains.Load(uuid)
		domainLookup(&domain, libvirtDomain)
		return true
	})

	hostLookup(models.Collection.Host)
}

// Collect network collector data
func (collector *Collector) Collect() {
	// lookup for each domain
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		// uuid := key.(string)
		domain := value.(models.Domain)
		domainCollect(&domain)
		return true
	})
	hostCollect(models.Collection.Host)
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	hostFields := []string{
		"net_host_receivedBytes",
		"net_host_transmittedBytes",
	}
	domainFields := []string{
		"net_receivedBytes",
		"net_transmittedBytes",
	}
	if config.Options.Verbose {
		hostFields = append(hostFields,
			"net_host_receivedPackets",
			"net_host_receivedErrs",
			"net_host_receivedDrop",
			"net_host_receivedFifo",
			"net_host_receivedFrame",
			"net_host_receivedCompressed",
			"net_host_receivedMulticast",
			"net_host_transmittedPackets",
			"net_host_transmittedErrs",
			"net_host_transmittedDrop",
			"net_host_transmittedFifo",
			"net_host_transmittedColls",
			"net_host_transmittedCarrier",
			"net_host_transmittedCompressed",
		)
		domainFields = append(domainFields,
			"net_receivedPackets",
			"net_receivedErrs",
			"net_receivedDrop",
			"net_receivedFifo",
			"net_receivedFrame",
			"net_receivedCompressed",
			"net_receivedMulticast",
			"net_transmittedPackets",
			"net_transmittedErrs",
			"net_transmittedDrop",
			"net_transmittedFifo",
			"net_transmittedColls",
			"net_transmittedCarrier",
			"net_transmittedCompressed",
			"net_interfaces",
		)
	}
	printable := models.Printable{
		HostFields:   hostFields,
		DomainFields: domainFields,
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = domainPrint(&domain)
		return true
	})

	// lookup for host
	printable.HostValues = hostPrint(models.Collection.Host)

	return printable
}

// CreateCollector creates a new network collector
func CreateCollector() Collector {
	return Collector{}
}
