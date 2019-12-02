package memcollector

import (
	"kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

func domainLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	memStats, err := libvirtDomain.MemoryStats(uint32(libvirt.DOMAIN_MEMORY_STAT_NR), 0)
	if err != nil {
		return
	}
	var total, unused, used uint64
	for _, stat := range memStats {
		if stat.Tag == int32(libvirt.DOMAIN_MEMORY_STAT_UNUSED) {
			unused = stat.Val
		}
		if stat.Tag == int32(libvirt.DOMAIN_MEMORY_STAT_AVAILABLE) {
			total = stat.Val
		}
	}
	used = total - unused
	newMeasurementTotal := models.CreateMeasurement(total)
	domain.AddMetricMeasurement("ram_total", newMeasurementTotal)
	newMeasurementUsed := models.CreateMeasurement(used)
	domain.AddMetricMeasurement("ram_used", newMeasurementUsed)

}
