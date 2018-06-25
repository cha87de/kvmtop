package hostcollector

import (
	"os"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

func hostLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	domain.AddMetricMeasurement("host_name", models.CreateMeasurement(name))
}

func hostCollect(domain *models.Domain) {
	// nothing to do at present
}

func hostPrint(domain *models.Domain) []string {
	host := collectors.GetMetricString(domain.Measurable, "host_name", 0)
	result := append([]string{host})
	return result
}
