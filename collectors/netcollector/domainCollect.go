package netcollector

import (
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
)

func domainCollect(domain *models.Domain) {
	/*
		// get stats from netstat
		stats := util.GetProcNetstat(domain.PID)
		domain.AddMetricMeasurement("net_ipextinoctets", models.CreateMeasurement(uint64(stats.IPExtInOctets)))
		domain.AddMetricMeasurement("net_ipextoutoctets", models.CreateMeasurement(uint64(stats.IPExtOutOctets)))
	*/

	// get stats from net/dev for domain interfaces
	ifs := domain.GetMetricStringArray("net_interfaces")
	statsSum := util.ProcPIDNetDev{}
	for _, devname := range ifs {
		devStats := util.GetProcPIDNetDev(domain.PID, devname)

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
	domain.AddMetricMeasurement("net_ReceivedBytes", models.CreateMeasurement(uint64(statsSum.ReceivedBytes)))
	domain.AddMetricMeasurement("net_ReceivedPackets", models.CreateMeasurement(uint64(statsSum.ReceivedPackets)))
	domain.AddMetricMeasurement("net_ReceivedErrs", models.CreateMeasurement(uint64(statsSum.ReceivedErrs)))
	domain.AddMetricMeasurement("net_ReceivedDrop", models.CreateMeasurement(uint64(statsSum.ReceivedDrop)))
	domain.AddMetricMeasurement("net_ReceivedFifo", models.CreateMeasurement(uint64(statsSum.ReceivedFifo)))
	domain.AddMetricMeasurement("net_ReceivedFrame", models.CreateMeasurement(uint64(statsSum.ReceivedFrame)))
	domain.AddMetricMeasurement("net_ReceivedCompressed", models.CreateMeasurement(uint64(statsSum.ReceivedCompressed)))
	domain.AddMetricMeasurement("net_ReceivedMulticast", models.CreateMeasurement(uint64(statsSum.ReceivedMulticast)))
	domain.AddMetricMeasurement("net_TransmittedBytes", models.CreateMeasurement(uint64(statsSum.TransmittedBytes)))
	domain.AddMetricMeasurement("net_TransmittedPackets", models.CreateMeasurement(uint64(statsSum.TransmittedPackets)))
	domain.AddMetricMeasurement("net_TransmittedErrs", models.CreateMeasurement(uint64(statsSum.TransmittedErrs)))
	domain.AddMetricMeasurement("net_TransmittedDrop", models.CreateMeasurement(uint64(statsSum.TransmittedDrop)))
	domain.AddMetricMeasurement("net_TransmittedFifo", models.CreateMeasurement(uint64(statsSum.TransmittedFifo)))
	domain.AddMetricMeasurement("net_TransmittedColls", models.CreateMeasurement(uint64(statsSum.TransmittedColls)))
	domain.AddMetricMeasurement("net_TransmittedCarrier", models.CreateMeasurement(uint64(statsSum.TransmittedCarrier)))
	domain.AddMetricMeasurement("net_TransmittedCompressed", models.CreateMeasurement(uint64(statsSum.TransmittedCompressed)))
}
