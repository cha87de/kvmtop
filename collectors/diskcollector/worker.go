package diskcollector

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

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
	// ErrsSet            bool
	// Errs               int64
	Capacity   uint64
	Allocation uint64
	Physical   uint64
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
	disksources := ""
	for _, disk := range domcfg.Devices.Disks {
		if disk.Target == nil || disk.Target.Dev == "" {
			// skip if disk specs are invalid
			continue
		}
		dev := disk.Target.Dev
		ioStats, err := libvirtDomain.BlockStats(dev)

		if ioStats != nil && err == nil {
			// ioStats.ErrsSet - works only for xen
			/*if ioStats.ErrsSet {
				sums.ErrsSet = true
				sums.Errs += ioStats.Errs
			}*/
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

		sizeStats, err := libvirtDomain.GetBlockInfo(dev, 0)
		// sizes
		if sizeStats != nil && err == nil {
			sums.Capacity += sizeStats.Capacity
			sums.Allocation += sizeStats.Allocation
			sums.Physical += sizeStats.Physical
		}

		// find source path
		if disk.Source != nil && disk.Source.File != nil {
			// only consider file based disks
			sourcefile := disk.Source.File
			sourcedir := filepath.Dir(sourcefile.File)
			if !strings.Contains(disksources, sourcedir) {
				if disksources != "" {
					disksources += ","
				}
				disksources += sourcedir
			}
		}

	}

	// sizes
	domain.AddMetricMeasurement("disk_size_capacity", models.CreateMeasurement(uint64(sums.Capacity)))
	domain.AddMetricMeasurement("disk_size_allocation", models.CreateMeasurement(uint64(sums.Allocation)))
	domain.AddMetricMeasurement("disk_size_physical", models.CreateMeasurement(uint64(sums.Physical)))
	// IOs
	// domain.AddMetricMeasurement("disk_stats_errs", models.CreateMeasurement(uint64(sums.Errs)))
	domain.AddMetricMeasurement("disk_stats_flushreq", models.CreateMeasurement(uint64(sums.FlushReq)))
	domain.AddMetricMeasurement("disk_stats_flushtotaltimes", models.CreateMeasurement(uint64(sums.FlushTotalTimes)))
	domain.AddMetricMeasurement("disk_stats_rdbytes", models.CreateMeasurement(uint64(sums.RdBytes)))
	domain.AddMetricMeasurement("disk_stats_rdreq", models.CreateMeasurement(uint64(sums.RdReq)))
	domain.AddMetricMeasurement("disk_stats_rdtotaltimes", models.CreateMeasurement(uint64(sums.RdTotalTimes)))
	domain.AddMetricMeasurement("disk_stats_wrbytes", models.CreateMeasurement(uint64(sums.WrBytes)))
	domain.AddMetricMeasurement("disk_stats_wrreq", models.CreateMeasurement(uint64(sums.WrReq)))
	domain.AddMetricMeasurement("disk_stats_wrtotaltimes", models.CreateMeasurement(uint64(sums.WrTotalTimes)))
	// information
	domain.AddMetricMeasurement("disk_sources", models.CreateMeasurement(disksources))
}

func diskCollect(domain *models.Domain, host *models.Host) {
	pid := domain.PID
	stats := util.GetProcPIDStat(pid)
	domain.AddMetricMeasurement("disk_delayblkio", models.CreateMeasurement(uint64(stats.DelayacctBlkioTicks)))

	// calculate ioutil as estimation
	domainIOUtil := estimateIOUtil(domain, host)
	domain.AddMetricMeasurement("disk_ioutil", models.CreateMeasurement(domainIOUtil))
}

func diskPrint(domain *models.Domain) []string {
	capacity, _ := domain.GetMetricUint64("disk_size_capacity", 0)
	allocation, _ := domain.GetMetricUint64("disk_size_allocation", 0)
	physical, _ := domain.GetMetricUint64("disk_size_physical", 0)

	// errs := domain.GetMetricDiffUint64("disk_stats_errs", true)
	flushreq := domain.GetMetricDiffUint64("disk_stats_flushreq", true)
	flushtotaltimes := domain.GetMetricDiffUint64("disk_stats_flushtotaltimes", true)
	rdbytes := domain.GetMetricDiffUint64("disk_stats_rdbytes", true)
	rdreq := domain.GetMetricDiffUint64("disk_stats_rdreq", true)
	rdtotaltimes := domain.GetMetricDiffUint64("disk_stats_rdtotaltimes", true)
	wrbytes := domain.GetMetricDiffUint64("disk_stats_wrbytes", true)
	wrreq := domain.GetMetricDiffUint64("disk_stats_wrreq", true)
	wrtotaltimes := domain.GetMetricDiffUint64("disk_stats_wrtotaltimes", true)

	delayblkio := domain.GetMetricDiffUint64("disk_delayblkio", true)

	ioutil, _ := domain.GetMetricUint64("disk_ioutil", 0)

	result := append([]string{capacity}, allocation, ioutil)
	if config.Options.Verbose {
		result = append(result, physical, flushreq, flushtotaltimes, rdbytes, rdreq, rdtotaltimes, wrbytes, wrreq, wrtotaltimes, delayblkio)
	}
	return result
}

func estimateIOUtil(domain *models.Domain, host *models.Host) string {
	hostIOUtilstr := host.GetMetricString("disk_device_ioutil", 0)
	hostIOUtil, errc := strconv.Atoi(hostIOUtilstr)
	if errc != nil {
		return ""
	}
	hostReads := host.GetMetricDiffUint64AsFloat("disk_device_reads", true)
	hostWrites := host.GetMetricDiffUint64AsFloat("disk_device_writes", true)

	domainReads := domain.GetMetricDiffUint64AsFloat("disk_stats_rdreq", true)
	domainWrites := domain.GetMetricDiffUint64AsFloat("disk_stats_wrreq", true)

	hostLoad := hostReads + hostWrites
	domainLoad := domainReads + domainWrites

	ratio := domainLoad / hostLoad
	domainIOUtil := ratio * float64(hostIOUtil)

	return fmt.Sprintf("%.0f", domainIOUtil)
}
