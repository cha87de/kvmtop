package netcollector

import (
	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func hostPrint(host *models.Host) []string {
	receivedBytes := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedBytes", true)
	receivedPackets := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedPackets", true)
	receivedErrs := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedErrs", true)
	receivedDrop := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedDrop", true)
	receivedFifo := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedFifo", true)
	receivedFrame := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedFrame", true)
	receivedCompressed := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedCompressed", true)
	receivedMulticast := collectors.GetMetricDiffUint64(host.Measurable, "net_host_ReceivedMulticast", true)
	transmittedBytes := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedBytes", true)
	transmittedPackets := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedPackets", true)
	transmittedErrs := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedErrs", true)
	transmittedDrop := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedDrop", true)
	transmittedFifo := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedFifo", true)
	transmittedColls := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedColls", true)
	transmittedCarrier := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedCarrier", true)
	transmittedCompressed := collectors.GetMetricDiffUint64(host.Measurable, "net_host_TransmittedCompressed", true)

	speed := collectors.GetMetricUint64(host.Measurable, "net_host_speed", 0)

	result := append([]string{receivedBytes}, transmittedBytes, speed)
	if config.Options.Verbose {
		result = append(result, receivedPackets, receivedErrs, receivedDrop, receivedFifo, receivedFrame, receivedCompressed, receivedMulticast, transmittedPackets, transmittedErrs, transmittedDrop, transmittedFifo, transmittedColls, transmittedCarrier, transmittedCompressed)
	}

	return result
}
