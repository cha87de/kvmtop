package netcollector

import (
	"strings"

	"kvmtop/config"
	"kvmtop/models"
)

func domainPrint(domain *models.Domain) []string {
	receivedBytes := domain.GetMetricDiffUint64("net_ReceivedBytes", true)
	receivedPackets := domain.GetMetricDiffUint64("net_ReceivedPackets", true)
	receivedErrs := domain.GetMetricDiffUint64("net_ReceivedErrs", true)
	receivedDrop := domain.GetMetricDiffUint64("net_ReceivedDrop", true)
	receivedFifo := domain.GetMetricDiffUint64("net_ReceivedFifo", true)
	receivedFrame := domain.GetMetricDiffUint64("net_ReceivedFrame", true)
	receivedCompressed := domain.GetMetricDiffUint64("net_ReceivedCompressed", true)
	receivedMulticast := domain.GetMetricDiffUint64("net_ReceivedMulticast", true)
	transmittedBytes := domain.GetMetricDiffUint64("net_TransmittedBytes", true)
	transmittedPackets := domain.GetMetricDiffUint64("net_TransmittedPackets", true)
	transmittedErrs := domain.GetMetricDiffUint64("net_TransmittedErrs", true)
	transmittedDrop := domain.GetMetricDiffUint64("net_TransmittedDrop", true)
	transmittedFifo := domain.GetMetricDiffUint64("net_TransmittedFifo", true)
	transmittedColls := domain.GetMetricDiffUint64("net_TransmittedColls", true)
	transmittedCarrier := domain.GetMetricDiffUint64("net_TransmittedCarrier", true)
	transmittedCompressed := domain.GetMetricDiffUint64("net_TransmittedCompressed", true)

	ifsRaw := domain.GetMetricStringArray("net_interfaces")
	interfaces := strings.Join(ifsRaw, ";")

	result := append([]string{receivedBytes}, transmittedBytes)
	if config.Options.Verbose {
		result = append(result, receivedPackets, receivedErrs, receivedDrop, receivedFifo, receivedFrame, receivedCompressed, receivedMulticast, transmittedPackets, transmittedErrs, transmittedDrop, transmittedFifo, transmittedColls, transmittedCarrier, transmittedCompressed, interfaces)
	}
	return result
}
