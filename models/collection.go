package models

func init() {
	// initialize the collection variable
	Collection.Domains = *NewDomains()
	Collection.Collectors = *NewCollectors()
	Collection.Host = Host{
		Measurable: NewMeasurable(),
	}
}

// Collection of domains and other stuff
var Collection struct {
	Host           Host
	Domains        Domains
	Collectors     Collectors
	Printer        Printer
	LibvirtDomains LibvirtDomains
}
