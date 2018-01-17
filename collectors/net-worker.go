package collectors

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

type stats struct {
	RxBytesSet   bool
	RxBytes      int64
	RxPacketsSet bool
	RxPackets    int64
	RxErrsSet    bool
	RxErrs       int64
	RxDropSet    bool
	RxDrop       int64
	TxBytesSet   bool
	TxBytes      int64
	TxPacketsSet bool
	TxPackets    int64
	TxErrsSet    bool
	TxErrs       int64
	TxDropSet    bool
	TxDrop       int64
}

func netLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// get list of network interfaces
	rawinterfaces, err := libvirtDomain.ListAllInterfaceAddresses(libvirt.DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE)
	if err != nil {
		return
	}
	var ifs []string
	for _, ifx := range rawinterfaces {
		ifs = append(ifs, ifx.Name)
	}
	newMeasurementInterfaces := models.CreateMeasurement(ifs)
	domain.AddMetricMeasurement("net_ifs", newMeasurementInterfaces)

	// query stats for each interface
	var sums stats
	for _, ifx := range ifs {
		stats, err := libvirtDomain.InterfaceStats(ifx)
		if err != nil {
			continue
		}
		// bytes ...
		if stats.RxBytesSet {
			sums.RxBytesSet = true
			sums.RxBytes += stats.RxBytes
		}
		if stats.TxBytesSet {
			sums.TxBytesSet = true
			sums.TxBytes += stats.TxBytes
		}
		// packets ...
		if stats.RxPacketsSet {
			sums.RxPacketsSet = true
			sums.RxPackets += stats.RxPackets
		}
		if stats.TxPacketsSet {
			sums.TxPacketsSet = true
			sums.TxPackets += stats.TxPackets
		}
		// drops ...
		if stats.RxDropSet {
			sums.RxDropSet = true
			sums.RxDrop += stats.RxDrop
		}
		if stats.TxDropSet {
			sums.TxDropSet = true
			sums.TxDrop += stats.TxDrop
		}
		// errors ...
		if stats.RxErrsSet {
			sums.RxErrsSet = true
			sums.RxErrs += stats.RxErrs
		}
		if stats.TxErrsSet {
			sums.TxErrsSet = true
			sums.TxErrs += stats.TxErrs
		}

	}
	newMeasurementBytes := models.CreateMeasurement(sums)
	domain.AddMetricMeasurement("net_stats", newMeasurementBytes)
}

func netCollect(domain *models.Domain) {
	// TODO
}

func netPrint(domain *models.Domain) []string {
	var thrpIn, thrpOut string
	if metric, ok := domain.GetMetric("net_stats"); ok {
		if len(metric.Values) > 1 {

			in1 := getIn(metric.Values[0].Value)
			in2 := getIn(metric.Values[1].Value)
			out1 := getOut(metric.Values[0].Value)
			out2 := getOut(metric.Values[1].Value)

			timeDiff := metric.Values[0].Timestamp.Sub(metric.Values[1].Timestamp).Seconds()
			timeConversionFactor := timeDiff

			fracIn := float64(in1-in2) / timeConversionFactor
			fracOut := float64(out1-out2) / timeConversionFactor

			thrpIn = fmt.Sprintf("%.2f", fracOut/1024/1024)
			thrpOut = fmt.Sprintf("%.2f", fracIn/1024/1024)
		}
	}

	result := append([]string{thrpOut}, thrpIn)
	return result
}

func getIn(byteValue []byte) int64 {
	reader := bytes.NewReader(byteValue)
	decoder := gob.NewDecoder(reader)
	var stats stats
	decoder.Decode(&stats)
	return stats.RxBytes
}

func getOut(byteValue []byte) int64 {
	reader := bytes.NewReader(byteValue)
	decoder := gob.NewDecoder(reader)
	var stats stats
	decoder.Decode(&stats)
	return stats.TxBytes
}
