package memcollector

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

// Collector describes the memory collector
type Collector struct {
	models.Collector
}

const pagesize = 4096

// Lookup memory collector data
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

// Collect memory collector data
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
		"ram_Total",
		"ram_Free",
		"ram_Available",
	}
	domainFields := []string{
		"ram_total",
		"ram_used",
	}
	if config.Options.Verbose {
		hostFields = append(hostFields,
			"ram_Buffers",
			"ram_Cached",
			"ram_SwapCached",
			"ram_Active",
			"ram_Inactive",
			"ram_ActiveAanon",
			"ram_InactiveAanon",
			"ram_ActiveFile",
			"ram_InactiveFile",
			"ram_Unevictable",
			"ram_Mlocked",
			"ram_SwapTotal",
			"ram_SwapFree",
			"ram_Dirty",
			"ram_Writeback",
			"ram_AnonPages",
			"ram_Mapped",
			"ram_Shmem",
			"ram_Slab",
			"ram_SReclaimable",
			"ram_SUnreclaim",
			"ram_KernelStack",
			"ram_PageTables",
			"ram_NFSUnstable",
			"ram_Bounce",
			"ram_WritebackTmp",
			"ram_CommitLimit",
			"ram_CommittedAS",
			"ram_VmallocTotal",
			"ram_VmallocUsed",
			"ram_VmallocChunk",
			"ram_HardwareCorrupted",
			"ram_AnonHugePages",
			"ram_ShmemHugePages",
			"ram_ShmemPmdMapped",
			"ram_HugePagesTotal",
			"ram_HugePagesFree",
			"ram_HugePagesRsvd",
			"ram_HugePagesSurp",
			"ram_Hugepagesize",
			"ram_Hugetlb",
			"ram_DirectMap4k",
			"ram_DirectMap2M",
			"ram_DirectMap1G",
		)
		domainFields = append(domainFields,
			"ram_vsize",
			"ram_rss",
			"ram_minflt",
			"ram_cminflt",
			"ram_majflt",
			"ram_cmajflt",
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

// CreateCollector creates a new memory collector
func CreateCollector() Collector {
	return Collector{}
}
