package netcollector

import (
	"github.com/cha87de/kvmtop/connector"
	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func hostLookup(host *models.Host) {
	bridges := getHostBridges()
	newMeasurementInterfaces := models.CreateMeasurement(bridges)
	host.AddMetricMeasurement("net_host_ifs", newMeasurementInterfaces)
}

func getHostBridges() []string {

	bridges := make(map[string]string)
	networks := make(map[string]string)

	models.Collection.LibvirtDomains.Map.Range(func(key, value interface{}) bool {
		libvirtDomain := value.(libvirt.Domain)

		xmldoc, _ := libvirtDomain.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
		domcfg := &libvirtxml.Domain{}
		domcfg.Unmarshal(xmldoc)

		for _, devInterface := range domcfg.Devices.Interfaces {
			if devInterface.Source.Network != nil {
				// lookup network bridge
				network := devInterface.Source.Network.Network
				networks[network] = network
			} else if devInterface.Source.Bridge != nil {
				bridge := devInterface.Source.Bridge.Bridge
				bridges[bridge] = bridge
			}
		}

		return true
	})

	// lookup bridges of networks
	for networkName := range networks {
		libvirtNetwork, _ := connector.Libvirt.Connection.LookupNetworkByName(networkName)
		bridge, _ := libvirtNetwork.GetBridgeName()
		bridges[bridge] = bridge
	}

	// build array of bridges
	bridgeArr := make([]string, 0, len(bridges))
	for k := range bridges {
		bridgeArr = append(bridgeArr, k)
	}
	return bridgeArr
}
