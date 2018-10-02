package netcollector

import (
	"strings"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/models"
)

func domainPrint(domain *models.Domain) []string {
	receivedBytes := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedBytes", true)
	receivedPackets := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedPackets", true)
	receivedErrs := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedErrs", true)
	receivedDrop := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedDrop", true)
	receivedFifo := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedFifo", true)
	receivedFrame := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedFrame", true)
	receivedCompressed := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedCompressed", true)
	receivedMulticast := collectors.GetMetricDiffUint64(domain.Measurable, "net_ReceivedMulticast", true)
	transmittedBytes := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedBytes", true)
	transmittedPackets := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedPackets", true)
	transmittedErrs := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedErrs", true)
	transmittedDrop := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedDrop", true)
	transmittedFifo := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedFifo", true)
	transmittedColls := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedColls", true)
	transmittedCarrier := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedCarrier", true)
	transmittedCompressed := collectors.GetMetricDiffUint64(domain.Measurable, "net_TransmittedCompressed", true)

	ifsRaw := domain.GetMetricStringArray("net_interfaces")
	interfaces := strings.Join(ifsRaw, ";")

	result := append([]string{receivedBytes}, receivedPackets, receivedErrs, receivedDrop, receivedFifo, receivedFrame, receivedCompressed, receivedMulticast, transmittedBytes, transmittedPackets, transmittedErrs, transmittedDrop, transmittedFifo, transmittedColls, transmittedCarrier, transmittedCompressed, interfaces)
	return result
}
