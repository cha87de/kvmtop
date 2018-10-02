package netcollector

import (
	"fmt"

	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func domainLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// get list of network interfaces
	/*rawinterfaces, err := libvirtDomain.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
	if err != nil {
		return
	}
	fmt.Printf("rawinterfaces: %+v\n", rawinterfaces)
	var ifs []string
	for _, ifx := range rawinterfaces {
		ifs = append(ifs, ifx.Name)
	}*/

	var ifs []string
	xmldoc, _ := libvirtDomain.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
	domcfg := &libvirtxml.Domain{}
	domcfg.Unmarshal(xmldoc)

	if domcfg.Devices == nil {
		fmt.Printf("devices for domain %s nil!\n", domain.UUID)
		return
	} else if domcfg.Devices.Interfaces == nil {
		fmt.Printf("device interfaces for domain %s nil!\n", domain.UUID)
		return
	} else {
		for _, devInterface := range domcfg.Devices.Interfaces {
			if devInterface.Target != nil {
				ifs = append(ifs, devInterface.Target.Dev)
			}
		}
	}

	newMeasurementInterfaces := models.CreateMeasurement(ifs)
	domain.AddMetricMeasurement("net_interfaces", newMeasurementInterfaces)
}
