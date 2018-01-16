package collectors

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

func memLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	memStats, err := libvirtDomain.MemoryStats(uint32(libvirt.DOMAIN_MEMORY_STAT_NR), 0)
	if err != nil {
		return
	}
	var total, unused, used uint64
	for _, stat := range memStats {
		if stat.Tag == int32(libvirt.DOMAIN_MEMORY_STAT_UNUSED) {
			unused = stat.Val
		}
		if stat.Tag == int32(libvirt.DOMAIN_MEMORY_STAT_AVAILABLE) {
			total = stat.Val
		}
	}
	used = total - unused
	newMeasurementTotal := models.CreateMeasurement(total)
	domain.AddMetricMeasurement("ram_total", newMeasurementTotal)
	newMeasurementUsed := models.CreateMeasurement(used)
	domain.AddMetricMeasurement("ram_used", newMeasurementUsed)
}

func memCollect(domain *models.Domain) {
	// when reading from libvirt, no collection possible
	// TODO read values from /proc/ fs instead.
}

func memPrint(domain *models.Domain) []string {
	var total, used string

	if metric, ok := domain.GetMetric("ram_total"); ok {
		if len(metric.Values) > 1 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var value uint64
			decoder.Decode(&value)
			total = fmt.Sprintf("%d", value/1024)
		}
	}

	if metric, ok := domain.GetMetric("ram_used"); ok {
		if len(metric.Values) > 1 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			decoder := gob.NewDecoder(reader)
			var value uint64
			decoder.Decode(&value)
			used = fmt.Sprintf("%d", value/1024)
		}
	}

	result := append([]string{total}, used)
	return result
}
