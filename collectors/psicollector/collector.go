package psicollector

import (
	"kvmtop/config"
	"kvmtop/models"
)

// Collector describes the Pressure Stall Information (PSI) collector
type Collector struct {
	models.Collector
}

// Lookup disk collector data
func (collector *Collector) Lookup() {
	hostLookup(&models.Collection.Host)
}

// Collect disk collector data
func (collector *Collector) Collect() {
	hostCollect(&models.Collection.Host)
}

// Print returns the collectors measurements in a Printable struct
func (collector *Collector) Print() models.Printable {
	hostFields := []string{
		"psi_some_cpu_avg60",

		"psi_some_io_avg60",
		"psi_full_io_avg60",

		"psi_some_mem_avg60",
		"psi_full_mem_avg60",
	}
	domainFields := []string{}

	if config.Options.Verbose {
		hostFields = []string{
			"psi_some_cpu_avg10",
			"psi_some_cpu_avg60",
			"psi_some_cpu_avg300",
			"psi_some_cpu_total",

			"psi_some_io_avg10",
			"psi_some_io_avg60",
			"psi_some_io_avg300",
			"psi_some_io_total",
			"psi_full_io_avg10",
			"psi_full_io_avg60",
			"psi_full_io_avg300",
			"psi_full_io_total",

			"psi_some_mem_avg10",
			"psi_some_mem_avg60",
			"psi_some_mem_avg300",
			"psi_some_mem_total",
			"psi_full_mem_avg10",
			"psi_full_mem_avg60",
			"psi_full_mem_avg300",
			"psi_full_mem_total",
		}
	}

	printable := models.Printable{
		HostFields:   hostFields,
		DomainFields: domainFields,
	}

	// lookup for host
	printable.HostValues = printHost(&models.Collection.Host)

	return printable
}

// CreateCollector creates a new Pressure Stall Information (PSI) collector
func CreateCollector() Collector {
	return Collector{}
}
