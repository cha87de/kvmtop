package collectors

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cha87de/kvmtop/models"
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

	var disks []string
	for _, disk := range domcfg.Devices.Disks {
		disks = append(disks, disk.Target.Dev)
	}
	newMeasurementDisks := models.CreateMeasurement(disks)
	domain.AddMetricMeasurement("disk_disks", newMeasurementDisks)

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
	newMeasurementStats := models.CreateMeasurement(sums)
	domain.AddMetricMeasurement("disk_stats", newMeasurementStats)
}

func diskCollect(domain *models.Domain) {
	// TODO
}

func diskPrint(domain *models.Domain) []string {
	var thrpR, thrpW string
	if metric, ok := domain.GetMetric("disk_stats"); ok {
		if len(metric.Values) > 1 {

			r1 := getRead(metric.Values[0].Value)
			r2 := getRead(metric.Values[1].Value)
			w1 := getWrite(metric.Values[0].Value)
			w2 := getWrite(metric.Values[1].Value)

			timeDiff := metric.Values[0].Timestamp.Sub(metric.Values[1].Timestamp).Seconds()
			timeConversionFactor := timeDiff

			fracR := float64(r1-r2) / timeConversionFactor
			fracW := float64(w1-w2) / timeConversionFactor

			thrpR = fmt.Sprintf("%.2f", fracR/1024/1024)
			thrpW = fmt.Sprintf("%.2f", fracW/1024/1024)
		}
	}

	result := append([]string{thrpR}, thrpW)
	return result
}

func getRead(byteValue []byte) int64 {
	reader := bytes.NewReader(byteValue)
	decoder := gob.NewDecoder(reader)
	var stats diskstats
	decoder.Decode(&stats)
	return stats.RdBytes
}

func getWrite(byteValue []byte) int64 {
	reader := bytes.NewReader(byteValue)
	decoder := gob.NewDecoder(reader)
	var stats diskstats
	decoder.Decode(&stats)
	return stats.WrBytes
}
