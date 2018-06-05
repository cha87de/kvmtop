package collectors

import (
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/util"
	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

type diskstats struct {
	RdBytesSet         bool
	RdBytes            int64
	RdReqSet           bool
	RdReq              int64
	RdTotalTimesSet    bool
	RdTotalTimes       int64
	WrBytesSet         bool
	WrBytes            int64
	WrReqSet           bool
	WrReq              int64
	WrTotalTimesSet    bool
	WrTotalTimes       int64
	FlushReqSet        bool
	FlushReq           int64
	FlushTotalTimesSet bool
	FlushTotalTimes    int64
	ErrsSet            bool
	Errs               int64
}

func diskLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {

	xmldoc, _ := libvirtDomain.GetXMLDesc(0)
	domcfg := &libvirtxml.Domain{}
	domcfg.Unmarshal(xmldoc)

	// generate list of virtual disks
	// var disks []string
	// for _, disk := range domcfg.Devices.Disks {
	//	disks = append(disks, disk.Target.Dev)
	// }
	// newMeasurementDisks := models.CreateMeasurement(disks)
	// domain.AddMetricMeasurement("disk_disks", newMeasurementDisks)

	// sum up stats from virtual disks
	var sums diskstats
	for _, disk := range domcfg.Devices.Disks {
		dev := disk.Target.Dev
		//sizeStats, _ := libvirtDomain.GetBlockInfo(dev, 0)
		ioStats, _ := libvirtDomain.BlockStats(dev)

		// ioStats.ErrsSet
		if ioStats.ErrsSet {
			sums.ErrsSet = true
			sums.Errs += ioStats.Errs
		}
		// ioStats.FlushReq
		if ioStats.FlushReqSet {
			sums.FlushReqSet = true
			sums.FlushReq += ioStats.FlushReq
		}
		// ioStats.FlushTotalTimes
		if ioStats.FlushTotalTimesSet {
			sums.FlushTotalTimesSet = true
			sums.FlushTotalTimes += ioStats.FlushTotalTimes
		}
		// ioStats.RdBytes
		if ioStats.RdBytesSet {
			sums.RdBytesSet = true
			sums.RdBytes += ioStats.RdBytes
		}
		// ioStats.RdReq
		if ioStats.RdReqSet {
			sums.RdReqSet = true
			sums.RdReq += ioStats.RdReq
		}
		// ioStats.RdTotalTimes
		if ioStats.RdTotalTimesSet {
			sums.RdTotalTimesSet = true
			sums.RdTotalTimes += ioStats.RdTotalTimes
		}
		// ioStats.WrBytes
		if ioStats.WrBytesSet {
			sums.WrBytesSet = true
			sums.WrBytes += ioStats.WrBytes
		}
		// ioStats.WrReq
		if ioStats.WrReqSet {
			sums.WrReqSet = true
			sums.WrReq += ioStats.WrReq
		}
		// ioStats.WrTotalTimes
		if ioStats.WrTotalTimesSet {
			sums.WrTotalTimesSet = true
			sums.WrTotalTimes += ioStats.WrTotalTimes
		}
	}
	domain.AddMetricMeasurement("disk_stats_errs", models.CreateMeasurement(uint64(sums.Errs)))
	domain.AddMetricMeasurement("disk_stats_flushreq", models.CreateMeasurement(uint64(sums.FlushReq)))
	domain.AddMetricMeasurement("disk_stats_flushtotaltimes", models.CreateMeasurement(uint64(sums.FlushTotalTimes)))
	domain.AddMetricMeasurement("disk_stats_rdbytes", models.CreateMeasurement(uint64(sums.RdBytes)))
	domain.AddMetricMeasurement("disk_stats_rdreq", models.CreateMeasurement(uint64(sums.RdReq)))
	domain.AddMetricMeasurement("disk_stats_rdtotaltimes", models.CreateMeasurement(uint64(sums.RdTotalTimes)))
	domain.AddMetricMeasurement("disk_stats_wrbytes", models.CreateMeasurement(uint64(sums.WrBytes)))
	domain.AddMetricMeasurement("disk_stats_wrreq", models.CreateMeasurement(uint64(sums.WrReq)))
	domain.AddMetricMeasurement("disk_stats_wrtotaltimes", models.CreateMeasurement(uint64(sums.WrTotalTimes)))
}

func diskCollect(domain *models.Domain) {
	pid := domain.PID
	stats := util.GetProcStat(pid)
	// fmt.Printf("\n%v\n", stats)
	domain.AddMetricMeasurement("disk_delayblkio", models.CreateMeasurement(uint64(stats.DelayacctBlkioTicks)))
}

func diskPrint(domain *models.Domain) []string {
	errs := getMetricDiffUint64(domain, "disk_stats_errs", true)
	flushreq := getMetricDiffUint64(domain, "disk_stats_flushreq", true)
	flushtotaltimes := getMetricDiffUint64(domain, "disk_stats_flushtotaltimes", true)
	rdbytes := getMetricDiffUint64(domain, "disk_stats_rdbytes", true)
	rdreq := getMetricDiffUint64(domain, "disk_stats_rdreq", true)
	rdtotaltimes := getMetricDiffUint64(domain, "disk_stats_rdtotaltimes", true)
	wrbytes := getMetricDiffUint64(domain, "disk_stats_wrbytes", true)
	wrreq := getMetricDiffUint64(domain, "disk_stats_wrreq", true)
	wrtotaltimes := getMetricDiffUint64(domain, "disk_stats_wrtotaltimes", true)

	delayblkio := getMetricDiffUint64(domain, "disk_delayblkio", true)

	result := append([]string{rdbytes}, wrbytes, delayblkio)
	if config.Options.Verbose {
		result = append([]string{errs}, flushreq, flushtotaltimes, rdbytes, rdreq, rdtotaltimes, wrbytes, wrreq, wrtotaltimes, delayblkio)
	}
	return result
}
