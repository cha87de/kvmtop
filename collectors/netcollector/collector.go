package netcollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

// Collector describes the network collector
type Collector struct {
	models.Collector
}

// Lookup network collector data
func (collector *Collector) Lookup(host *models.Host, domains map[string]*models.Domain, libvirtDomains map[string]libvirt.Domain) {
	for uuid := range domains {
		domainLookup(domains[uuid], libvirtDomains[uuid])
	}
	hostLookup(host, libvirtDomains)
}

// Collect network collector data
func (collector *Collector) Collect(host *models.Host, domains map[string]*models.Domain) {
	// lookup for each domain
	for uuid := range domains {
		domainCollect(domains[uuid])
	}
	hostCollect(host)
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print(host *models.Host, domains map[string]*models.Domain) models.Printable {
	printable := models.Printable{
		HostFields: []string{
			"net_host_receivedBytes",
			"net_host_receivedPackets",
			"net_host_receivedErrs",
			"net_host_receivedDrop",
			"net_host_receivedFifo",
			"net_host_receivedFrame",
			"net_host_receivedCompressed",
			"net_host_receivedMulticast",
			"net_host_transmittedBytes",
			"net_host_transmittedPackets",
			"net_host_transmittedErrs",
			"net_host_transmittedDrop",
			"net_host_transmittedFifo",
			"net_host_transmittedColls",
			"net_host_transmittedCarrier",
			"net_host_transmittedCompressed",
		},
		DomainFields: []string{
			"net_receivedBytes",
			"net_receivedPackets",
			"net_receivedErrs",
			"net_receivedDrop",
			"net_receivedFifo",
			"net_receivedFrame",
			"net_receivedCompressed",
			"net_receivedMulticast",
			"net_transmittedBytes",
			"net_transmittedPackets",
			"net_transmittedErrs",
			"net_transmittedDrop",
			"net_transmittedFifo",
			"net_transmittedColls",
			"net_transmittedCarrier",
			"net_transmittedCompressed",
		},
	}

	// lookup for each domain
	printable.DomainValues = make(map[string][]string)
	for uuid := range domains {
		printable.DomainValues[uuid] = domainPrint(domains[uuid])
	}

	// lookup for host
	printable.HostValues = hostPrint(host)

	return printable
}

// CreateCollector creates a new network collector
func CreateCollector() Collector {
	return Collector{}
}
