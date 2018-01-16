package collectors

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

func memLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	/*	memstats, err := libvirtDomain.MemoryStats(libvirt.DOMAIN_MEMORY_STAT_SWAP_IN, 0)
		if err != nil {
			return
		}
		cores := len(vcpus)
		newMeasurementCores := models.CreateMeasurement(cores)
		domain.AddMetricMeasurement("cpu_cores", newMeasurementCores)
	*/
}

func memCollect(domain *models.Domain) {

}

func memPrint(domain *models.Domain) []string {
	return nil
}
