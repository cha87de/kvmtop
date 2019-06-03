package netcollector

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
)

func hostCollect(host *models.Host) {
	// get stats from net/dev for host interfaces
	ifs := host.GetMetricStringArray("net_host_ifs")
	statsSum := util.ProcPIDNetDev{}
	for _, devname := range ifs {
		devStats := util.GetProcPIDNetDev(0, devname)

		statsSum.ReceivedBytes += devStats.ReceivedBytes
		statsSum.ReceivedPackets += devStats.ReceivedPackets
		statsSum.ReceivedErrs += devStats.ReceivedErrs
		statsSum.ReceivedDrop += devStats.ReceivedDrop
		statsSum.ReceivedFifo += devStats.ReceivedFifo
		statsSum.ReceivedFrame += devStats.ReceivedFrame
		statsSum.ReceivedCompressed += devStats.ReceivedCompressed
		statsSum.ReceivedMulticast += devStats.ReceivedMulticast

		statsSum.TransmittedBytes += devStats.TransmittedBytes
		statsSum.TransmittedPackets += devStats.TransmittedPackets
		statsSum.TransmittedErrs += devStats.TransmittedErrs
		statsSum.TransmittedDrop += devStats.TransmittedDrop
		statsSum.TransmittedFifo += devStats.TransmittedFifo
		statsSum.TransmittedColls += devStats.TransmittedColls
		statsSum.TransmittedCarrier += devStats.TransmittedCarrier
		statsSum.TransmittedCompressed += devStats.TransmittedCompressed
	}
	host.AddMetricMeasurement("net_host_ReceivedBytes", models.CreateMeasurement(uint64(statsSum.ReceivedBytes)))
	host.AddMetricMeasurement("net_host_ReceivedPackets", models.CreateMeasurement(uint64(statsSum.ReceivedPackets)))
	host.AddMetricMeasurement("net_host_ReceivedErrs", models.CreateMeasurement(uint64(statsSum.ReceivedErrs)))
	host.AddMetricMeasurement("net_host_ReceivedDrop", models.CreateMeasurement(uint64(statsSum.ReceivedDrop)))
	host.AddMetricMeasurement("net_host_ReceivedFifo", models.CreateMeasurement(uint64(statsSum.ReceivedFifo)))
	host.AddMetricMeasurement("net_host_ReceivedFrame", models.CreateMeasurement(uint64(statsSum.ReceivedFrame)))
	host.AddMetricMeasurement("net_host_ReceivedCompressed", models.CreateMeasurement(uint64(statsSum.ReceivedCompressed)))
	host.AddMetricMeasurement("net_host_ReceivedMulticast", models.CreateMeasurement(uint64(statsSum.ReceivedMulticast)))
	host.AddMetricMeasurement("net_host_TransmittedBytes", models.CreateMeasurement(uint64(statsSum.TransmittedBytes)))
	host.AddMetricMeasurement("net_host_TransmittedPackets", models.CreateMeasurement(uint64(statsSum.TransmittedPackets)))
	host.AddMetricMeasurement("net_host_TransmittedErrs", models.CreateMeasurement(uint64(statsSum.TransmittedErrs)))
	host.AddMetricMeasurement("net_host_TransmittedDrop", models.CreateMeasurement(uint64(statsSum.TransmittedDrop)))
	host.AddMetricMeasurement("net_host_TransmittedFifo", models.CreateMeasurement(uint64(statsSum.TransmittedFifo)))
	host.AddMetricMeasurement("net_host_TransmittedColls", models.CreateMeasurement(uint64(statsSum.TransmittedColls)))
	host.AddMetricMeasurement("net_host_TransmittedCarrier", models.CreateMeasurement(uint64(statsSum.TransmittedCarrier)))
	host.AddMetricMeasurement("net_host_TransmittedCompressed", models.CreateMeasurement(uint64(statsSum.TransmittedCompressed)))
}
