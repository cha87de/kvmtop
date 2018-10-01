package netcollector

import (
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

func domainLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// get list of network interfaces
	rawinterfaces, err := libvirtDomain.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
	if err != nil {
		return
	}
	var ifs []string
	for _, ifx := range rawinterfaces {
		ifs = append(ifs, ifx.Name)
	}
	newMeasurementInterfaces := models.CreateMeasurement(ifs)
	domain.AddMetricMeasurement("net_ifs", newMeasurementInterfaces)
}
