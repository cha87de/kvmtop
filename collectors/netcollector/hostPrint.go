package netcollector

import (
	"kvmtop/config"
	"kvmtop/models"
)

func hostPrint(host *models.Host) []string {
	receivedBytes := host.GetMetricDiffUint64("net_host_ReceivedBytes", true)
	receivedPackets := host.GetMetricDiffUint64("net_host_ReceivedPackets", true)
	receivedErrs := host.GetMetricDiffUint64("net_host_ReceivedErrs", true)
	receivedDrop := host.GetMetricDiffUint64("net_host_ReceivedDrop", true)
	receivedFifo := host.GetMetricDiffUint64("net_host_ReceivedFifo", true)
	receivedFrame := host.GetMetricDiffUint64("net_host_ReceivedFrame", true)
	receivedCompressed := host.GetMetricDiffUint64("net_host_ReceivedCompressed", true)
	receivedMulticast := host.GetMetricDiffUint64("net_host_ReceivedMulticast", true)
	transmittedBytes := host.GetMetricDiffUint64("net_host_TransmittedBytes", true)
	transmittedPackets := host.GetMetricDiffUint64("net_host_TransmittedPackets", true)
	transmittedErrs := host.GetMetricDiffUint64("net_host_TransmittedErrs", true)
	transmittedDrop := host.GetMetricDiffUint64("net_host_TransmittedDrop", true)
	transmittedFifo := host.GetMetricDiffUint64("net_host_TransmittedFifo", true)
	transmittedColls := host.GetMetricDiffUint64("net_host_TransmittedColls", true)
	transmittedCarrier := host.GetMetricDiffUint64("net_host_TransmittedCarrier", true)
	transmittedCompressed := host.GetMetricDiffUint64("net_host_TransmittedCompressed", true)

	speed, _ := host.GetMetricUint64("net_host_speed", 0)

	result := append([]string{receivedBytes}, transmittedBytes, speed)
	if config.Options.Verbose {
		result = append(result, receivedPackets, receivedErrs, receivedDrop, receivedFifo, receivedFrame, receivedCompressed, receivedMulticast, transmittedPackets, transmittedErrs, transmittedDrop, transmittedFifo, transmittedColls, transmittedCarrier, transmittedCompressed)
	}

	return result
}
